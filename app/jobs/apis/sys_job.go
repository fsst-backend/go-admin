package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/jobs/service"
	"go-admin/common/dto"
)

type SysJob struct {
	api.Api
}

// @Summary 删除定时任务
// @Description 删除定时任务
// @Tags 定时任务
// @Accept application/json
// @Product application/json
// @Param data body dto.GeneralDelDto true "删除数据"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": "删除成功"}"
// @Router /lotus/api/v1/job/remove [post]
// @Security Bearer
func (e SysJob) RemoveJobForService(c *gin.Context) {
	v := dto.GeneralDelDto{}
	s := service.SysJob{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&v, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	s.Cron = sdk.Runtime.GetCrontabKey(c.Request.Host)
	err = s.RemoveJob(&v)
	if err != nil {
		e.Logger.Errorf("RemoveJob error, %s", err.Error())
		e.Error(500, err, "")
		return
	}
	e.OK(nil, s.Msg)
}

// @Summary 启动定时任务
// @Description 启动定时任务
// @Tags 定时任务
// @Accept application/json
// @Product application/json
// @Param data body dto.GeneralGetDto true "启动数据"
// @Success 200 {object} response.Response{message=string} "{"code": 0, "message": "启动成功"}
// @Router /lotus/api/v1/job/start [post]
// @Security Bearer
func (e SysJob) StartJobForService(c *gin.Context) {
	e.MakeContext(c)
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}
	var v dto.GeneralGetDto
	err = c.ShouldBind(&v)
	if err != nil {
		log.Warnf("参数验证错误, error: %s", err)
		e.Error(http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	s := service.SysJob{}
	s.Orm = db
	s.Log = log
	s.Cron = sdk.Runtime.GetCrontabKey(c.Request.Host)
	err = s.StartJob(&v)
	if err != nil {
		log.Errorf("GetCrontabKey error, %s", err.Error())
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, s.Msg)
}
