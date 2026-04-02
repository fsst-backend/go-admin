package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cockroachdb/cmux"
	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	adminGrpc "go-admin/app/admin/grpc"
	"go-admin/app/admin/models"
	"go-admin/app/admin/router"
	"go-admin/app/jobs"
	jobsModels "go-admin/app/jobs/models"
	otherModels "go-admin/app/other/models/tools"
	"go-admin/common/database"
	"go-admin/common/global"
	common "go-admin/common/middleware"
	"go-admin/common/middleware/handler"
	"go-admin/common/storage"
	ext "go-admin/config"
	filewrap "go-admin/config/filewarp"
	"go-admin/grpc/pb"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "go-admin server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

var AppRouters = make([]func(), 0)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck, "api", "a", false, "Start server with check api data")

	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	AppRouters = append(AppRouters, router.InitRouter)
}

func setup() {
	// 注入配置扩展项
	config.ExtendConfig = &ext.ExtConfig
	//1. 读取配置
	config.Setup(
		filewrap.NewFileWrap(configYml),
		// file.NewSource(file.WithPath(configYml)), // 原生支持重环境变量中替换
		database.Setup,
		storage.Setup,
	)

	//2. 执行数据库自动迁移
	autoMigrate()

	//注册监听函数
	queue := sdk.Runtime.GetMemoryQueue("")
	queue.Register(global.LoginLog, models.SaveLoginLog)
	queue.Register(global.OperateLog, models.SaveOperaLog)
	// queue.Register(global.ApiCheck, models.SaveSysApi)
	go queue.Run()

	// 打印 ExtConfig 配置参数
	log.Infof("ExtConfig - AMap.Key: %s", ext.ExtConfig.AMap.Key)
	log.Infof("ExtConfig - Violet.TargetURL: %s", ext.ExtConfig.Violet.TargetURL)
	log.Infof("ExtConfig - Violet.DomainID: %d", ext.ExtConfig.Violet.DomainID)
	log.Infof("ExtConfig - LinkForty.TargetURL: %s", ext.ExtConfig.LinkForty.TargetURL)
	log.Infof("ExtConfig - Telemarketing.TargetURL: %s", ext.ExtConfig.Telemarketing.TargetURL)
	log.Infof("ExtConfig - SMS.TargetURL: %s", ext.ExtConfig.SMS.TargetURL)
	log.Infof("ExtConfig - ChatAdminAPI.TargetURL: %s", ext.ExtConfig.ChatAdminAPI.TargetURL)
	log.Infof("ExtConfig - Upload.AppKey: %s", ext.ExtConfig.Upload.AppKey)
	log.Infof("ExtConfig - Upload.Secret: %s", ext.ExtConfig.Upload.Secret)
	log.Infof("ExtConfig - Frontend.BaseSiteURL: %s", ext.ExtConfig.Frontend.BaseSiteURL)
	log.Infof("ExtConfig - Frontend.BaseAPIURL: %s", ext.ExtConfig.Frontend.BaseAPIURL)
	log.Infof("ExtConfig - Frontend.BaseH5URL: %s", ext.ExtConfig.Frontend.BaseH5URL)
	log.Infof("ExtConfig - Frontend.BaseUploadURL: %s", ext.ExtConfig.Frontend.BaseUploadURL)

	usageStr := `starting api server...`
	log.Info(usageStr)
}

// autoMigrate 执行数据库自动迁移
func autoMigrate() {
	db := sdk.Runtime.GetDbByKey("*")
	if db == nil {
		log.Warn("数据库连接不存在,跳过自动迁移")
		return
	}

	log.Info("开始执行数据库自动迁移...")

	// 设置表选项
	if config.DatabaseConfig.Driver == "mysql" || config.DatabaseConfig.Driver == "tidb" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
	}

	// 执行 AutoMigrate
	err := db.AutoMigrate(
		// admin 模块模型
		new(models.SysDept),
		new(models.SysRoleDept),
		new(models.SysConfig),
		new(models.SysMenu),
		new(models.SysRoleMenu),
		new(models.SysLoginLog),
		new(models.SysOperaLog),
		new(models.SysUserRole),
		new(models.SysRolePermission),
		new(models.SysUser),
		new(models.SysRole),
		new(models.SysPost),
		new(models.SysDictData),
		new(models.SysDictType),
		new(models.SysApi),
		new(models.SysPermission),
		new(models.SysPermissionApi),
		new(models.CasbinRule),
		// jobs 模块模型
		new(jobsModels.SysJob),
		// other 模块模型
		new(otherModels.SysTables),
		new(otherModels.SysColumns),
	)

	if err != nil {
		log.Errorf("数据库自动迁移失败: %v", err)
		// 不中断程序启动,只记录错误
	} else {
		log.Info("数据库自动迁移完成")
	}
}

func run() error {
	if config.ApplicationConfig.Mode == pkg.ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}
	initRouter()

	for _, f := range AppRouters {
		f()
	}

	go func() {
		jobs.InitJob()
		jobs.Setup(sdk.Runtime.GetDb())
	}()

	if apiCheck {
		var routers = sdk.Runtime.GetRouter()
		q := sdk.Runtime.GetMemoryQueue("")
		mp := make(map[string]interface{})
		mp["List"] = routers
		message, err := sdk.Runtime.GetStreamMessage("", global.ApiCheck, mp)
		if err != nil {
			log.Infof("GetStreamMessage error, %s \n", err.Error())
			//日志报错错误，不中断请求
		} else {
			err = q.Append(message)
			if err != nil {
				log.Infof("Append message error, %s \n", err.Error())
			}
		}
	}

	plainl, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ApplicationConfig.Port))
	if err != nil {
		log.Fatalf("tcp conn:%v", err)
	}
	m := cmux.New(plainl)

	// gRPC 必须优先
	grpcL := m.Match(cmux.HTTP2())

	httpL := m.Match(cmux.HTTP1Fast())
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler:      sdk.Runtime.GetEngine(),
		ReadTimeout:  time.Duration(config.ApplicationConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.ApplicationConfig.WriterTimeout) * time.Second,
	}

	grpcServer := NewGrpcServer()

	// 初始化 AdminService 并注入日志、ORM 和配置
	adminService := &adminGrpc.AdminService{
		Logger: log.NewHelper(sdk.Runtime.GetLogger()), // 注入日志
		Config: &ext.ExtConfig,                         // 注入扩展配置（环境变量）
	}
	adminService.MakeOrm() // 注入 ORM（从 Runtime 获取数据库连接）
	pb.RegisterAdminUserServiceServer(grpcServer, adminService)

	go func() {
		// 服务连接
		log.Infof("http server start on port: %d", config.ApplicationConfig.Port)
		if err := srv.Serve(httpL); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("listen: ", err)
		}
		log.Infof("http server stop")
	}()

	go func() {
		log.Infof("grpc server start on port: %d", config.ApplicationConfig.Port)
		if err := grpcServer.Serve(grpcL); err != nil {
			log.Fatal("grpc serve:", err)
		}
		log.Infof("grpc server stop")
	}()

	go func() {
		log.Infof("cmux server start on port: %d", config.ApplicationConfig.Port)
		if err := m.Serve(); err != nil {
			log.Fatal("cmux serve:", err)
		}
		log.Infof("cmux server stop")
	}()

	fmt.Println(pkg.Red(string(global.LogoContent)))
	tip()
	fmt.Println(pkg.Green("Server run at:"))
	fmt.Printf("-  Local:   %s://localhost:%d/ \r\n", "http", config.ApplicationConfig.Port)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", pkg.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Info("Shutdown Server ... ")

	grpcServer.GracefulStop()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Info("Server exiting")

	return nil
}

//var Router runtime.Router

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`go-admin `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}

func initRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		h = gin.New()
		sdk.Runtime.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		//os.Exit(-1)
	}
	if config.SslConfig.Enable {
		r.Use(handler.TlsHandler())
	}
	//r.Use(middleware.Metrics())
	r.Use(common.Sentinel()).
		Use(common.RequestId(pkg.TrafficKey)).
		Use(api.SetRequestLogger)

	common.InitMiddleware(r)
}
