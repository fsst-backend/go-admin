package apis

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysApi struct {
	api.Api
}

// GetPage 获取接口管理列表
// @Summary 获取接口管理列表
// @Description 获取接口管理列表
// @Tags 接口管理
// @Param name query string false "名称"
// @Param title query string false "标题"
// @Param path query string false "地址"
// @Param action query string false "类型"
// @Param limit query int false "页条数"
// @Param offset query int false "页码"
// @Success 200 {object} response.Response{message=response.Page{list=[]models.SysApi}} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-api [get]
// @Security Bearer
func (e SysApi) GetPage(c *gin.Context) {
	s := service.SysApi{}
	req := dto.SysApiGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysApi, 0)
	var count int64
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.PageOK(list, int(count), req.GetOffset(), req.GetLimit(), "查询成功")
}

// Get 获取接口管理
// @Summary 获取接口管理
// @Description 获取接口管理
// @Tags 接口管理
// @Success 200 {object} response.Response{message=models.SysApi} "{"code": 0, "message": [...]}"
// @Router /lotus/api/v1/sys-api/get [get]
// @Security Bearer
func (e SysApi) Get(c *gin.Context) {
	req := dto.SysApiGetReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	p := actions.GetPermissionFromContext(c)
	var object models.SysApi
	err = s.Get(&req, p, &object).Error
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Insert 新增接口管理
// @Summary 新增接口管理
// @Description 新增接口管理
// @Tags 接口管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiInsertReq true "body"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": "创建成功"}"
// @Router /lotus/api/v1/sys-api [post]
// @Security Bearer
func (e SysApi) Insert(c *gin.Context) {
	req := dto.SysApiInsertReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, "创建失败")
		return
	}
	e.OK(req.GetId(), "创建成功")
}

// Update 修改接口管理
// @Summary 修改接口管理
// @Description 修改接口管理
// @Tags 接口管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiUpdateReq true "body"
// @Success 200 {object} response.Response{message=string}	"{"code": 0, "message": "修改成功"}"
// @Router /lotus/api/v1/sys-api [put]
// @Security Bearer
func (e SysApi) Update(c *gin.Context) {
	req := dto.SysApiUpdateReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)
	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, "更新失败")
		return
	}
	e.OK(req.GetId(), "更新成功")
}

// DeleteSysApi 删除接口管理
// @Summary 删除接口管理
// @Description 删除接口管理
// @Tags 接口管理
// @Param data body dto.SysApiDeleteReq true "body"
// @Success 200 {object} response.Response{message=string}	"{"code": 0, "message": "删除成功"}"
// @Router /lotus/api/v1/sys-api [delete]
// @Security Bearer
func (e SysApi) DeleteSysApi(c *gin.Context) {
	req := dto.SysApiDeleteReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	p := actions.GetPermissionFromContext(c)
	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// GenerateFromSwagger 从 Swagger 文件自动生成 API
// @Summary 从 Swagger 文件自动生成 API
// @Description 扫描 /app/doc 目录下的 swagger 文件，自动生成或更新 API
// @Tags 接口管理
// @Success 200 {object} response.Response{message=map[string]interface{}} "{"code": 0, "message": "生成成功"}"
// @Router /lotus/api/v1/sys-api/generate-from-swagger [post]
// @Security Bearer
func (e SysApi) GenerateFromSwagger(c *gin.Context) {
	err := e.MakeContext(c).
		MakeOrm().
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	docDir, err := ResolveSwaggerDocDir()
	if err != nil {
		e.Error(500, err, "Swagger 目录不存在")
		return
	}

	stats, err := syncSysApiFromSwaggerDir(e.Orm, e.Logger, docDir)
	if err != nil {
		e.Error(500, err, "扫描目录失败")
		return
	}

	e.OK(stats, "Swagger 文件处理完成")
}

// ResolveSwaggerDocDir 返回 Swagger JSON 所在目录（优先 /app/doc，否则项目下 docs），与 GenerateFromSwagger 行为一致。
func ResolveSwaggerDocDir() (string, error) {
	docDir := "/app/doc"
	if _, err := os.Stat(docDir); os.IsNotExist(err) {
		docDir = "docs"
		if _, err := os.Stat(docDir); os.IsNotExist(err) {
			return "", fmt.Errorf("目录不存在: %s", docDir)
		}
	}
	return docDir, nil
}

// SyncSysApiFromSwagger 扫描 Swagger 目录并将 path/action 写入 sys_api，供服务启动等无 HTTP 上下文的场景调用。
// 若目录不存在则返回错误（由调用方决定是否忽略）。
func SyncSysApiFromSwagger(db *gorm.DB, lg *logger.Helper) (map[string]interface{}, error) {
	if db == nil {
		return nil, fmt.Errorf("数据库连接为空")
	}
	if lg == nil {
		return nil, fmt.Errorf("logger 为空")
	}
	docDir, err := ResolveSwaggerDocDir()
	if err != nil {
		return nil, err
	}
	return syncSysApiFromSwaggerDir(db, lg, docDir)
}

func syncSysApiFromSwaggerDir(db *gorm.DB, lg *logger.Helper, docDir string) (map[string]interface{}, error) {
	var totalFiles, processed, inserted, updated, errCount int
	var errorDetails []string

	err := filepath.Walk(docDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".json") {
			totalFiles++
			fileStats, err := processSwaggerFileDB(db, lg, path)
			if err != nil {
				errCount++
				errorDetails = append(errorDetails, fmt.Sprintf("处理文件 %s 失败: %v", path, err))
				lg.Error(err)
			} else {
				processed++
				if insertedVal, ok := fileStats["inserted"].(int); ok {
					inserted += insertedVal
				}
				if updatedVal, ok := fileStats["updated"].(int); ok {
					updated += updatedVal
				}
			}
		}
		return nil
	})

	stats := map[string]interface{}{
		"totalFiles":   totalFiles,
		"processed":    processed,
		"inserted":     inserted,
		"updated":      updated,
		"errors":       errCount,
		"errorDetails": errorDetails,
	}
	if err != nil {
		return stats, err
	}
	return stats, nil
}

// processSwaggerFile 处理单个 swagger 文件（HTTP 请求上下文）
func (e SysApi) processSwaggerFile(filePath string) (map[string]interface{}, error) {
	return processSwaggerFileDB(e.Orm, e.Logger, filePath)
}

func processSwaggerFileDB(db *gorm.DB, lg *logger.Helper, filePath string) (map[string]interface{}, error) {
	stats := map[string]interface{}{
		"inserted": 0,
		"updated":  0,
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return stats, fmt.Errorf("读取文件失败: %w", err)
	}

	var swaggerData map[string]interface{}
	err = json.Unmarshal(data, &swaggerData)
	if err != nil {
		return stats, fmt.Errorf("解析 JSON 失败: %w", err)
	}

	paths, ok := swaggerData["paths"].(map[string]interface{})
	if !ok {
		return stats, fmt.Errorf("JSON 中没有找到 paths 字段")
	}

	for path, methods := range paths {
		methodsMap, ok := methods.(map[string]interface{})
		if !ok {
			continue
		}

		for method, operation := range methodsMap {
			operationMap, ok := operation.(map[string]interface{})
			if !ok {
				continue
			}

			title := ""
			if titleVal, exists := operationMap["summary"]; exists {
				if titleStr, ok := titleVal.(string); ok {
					title = titleStr
				}
			}

			tag := ""
			if tagsVal, exists := operationMap["tags"]; exists {
				if tagsArray, ok := tagsVal.([]interface{}); ok && len(tagsArray) > 0 {
					if tagStr, ok := tagsArray[0].(string); ok {
						tag = tagStr
					}
				}
			}

			sysApi := models.SysApi{
				Path:   path,
				Action: strings.ToUpper(method),
				Title:  title,
				Tag:    tag,
				Handle: fmt.Sprintf("%s %s", strings.ToUpper(method), path),
			}

			var existingApi models.SysApi
			result := db.Where("path = ? AND action = ?", path, strings.ToUpper(method)).First(&existingApi)

			if result.Error == nil && result.RowsAffected > 0 {
				sysApi.Id = existingApi.Id
				err := db.Model(&existingApi).Updates(&sysApi).Error
				if err != nil {
					lg.Errorf("更新记录失败: path=%s, action=%s, error=%s", path, method, err.Error())
				} else {
					stats["updated"] = stats["updated"].(int) + 1
				}
			} else {
				err := db.Create(&sysApi).Error
				if err != nil {
					lg.Errorf("插入记录失败: path=%s, action=%s, error=%s", path, method, err.Error())
				} else {
					stats["inserted"] = stats["inserted"].(int) + 1
				}
			}
		}
	}

	return stats, nil
}
