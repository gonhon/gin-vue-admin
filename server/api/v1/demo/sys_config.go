package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/demo"
    demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysConfigApi struct {}

// CreateSysConfig 创建sysConfig表
// @Tags SysConfig
// @Summary 创建sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body demo.SysConfig true "创建sysConfig表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /sysConfig/createSysConfig [post]
func (sysConfigApi *SysConfigApi) CreateSysConfig(c *gin.Context) {
	var sysConfig demo.SysConfig
	err := c.ShouldBindJSON(&sysConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysConfigService.CreateSysConfig(&sysConfig)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSysConfig 删除sysConfig表
// @Tags SysConfig
// @Summary 删除sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body demo.SysConfig true "删除sysConfig表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /sysConfig/deleteSysConfig [delete]
func (sysConfigApi *SysConfigApi) DeleteSysConfig(c *gin.Context) {
	id := c.Query("id")
	err := sysConfigService.DeleteSysConfig(id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysConfigByIds 批量删除sysConfig表
// @Tags SysConfig
// @Summary 批量删除sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /sysConfig/deleteSysConfigByIds [delete]
func (sysConfigApi *SysConfigApi) DeleteSysConfigByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	err := sysConfigService.DeleteSysConfigByIds(ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysConfig 更新sysConfig表
// @Tags SysConfig
// @Summary 更新sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body demo.SysConfig true "更新sysConfig表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /sysConfig/updateSysConfig [put]
func (sysConfigApi *SysConfigApi) UpdateSysConfig(c *gin.Context) {
	var sysConfig demo.SysConfig
	err := c.ShouldBindJSON(&sysConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysConfigService.UpdateSysConfig(sysConfig)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysConfig 用id查询sysConfig表
// @Tags SysConfig
// @Summary 用id查询sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query demo.SysConfig true "用id查询sysConfig表"
// @Success 200 {object} response.Response{data=demo.SysConfig,msg=string} "查询成功"
// @Router /sysConfig/findSysConfig [get]
func (sysConfigApi *SysConfigApi) FindSysConfig(c *gin.Context) {
	id := c.Query("id")
	resysConfig, err := sysConfigService.GetSysConfig(id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(resysConfig, c)
}

// GetSysConfigList 分页获取sysConfig表列表
// @Tags SysConfig
// @Summary 分页获取sysConfig表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query demoReq.SysConfigSearch true "分页获取sysConfig表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /sysConfig/getSysConfigList [get]
func (sysConfigApi *SysConfigApi) GetSysConfigList(c *gin.Context) {
	var pageInfo demoReq.SysConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysConfigService.GetSysConfigInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetSysConfigPublic 不需要鉴权的sysConfig表接口
// @Tags SysConfig
// @Summary 不需要鉴权的sysConfig表接口
// @accept application/json
// @Produce application/json
// @Param data query demoReq.SysConfigSearch true "分页获取sysConfig表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /sysConfig/getSysConfigPublic [get]
func (sysConfigApi *SysConfigApi) GetSysConfigPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的sysConfig表接口信息",
    }, "获取成功", c)
}
