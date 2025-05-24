package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysConfigRouter struct {}

// InitSysConfigRouter 初始化 sysConfig表 路由信息
func (s *SysConfigRouter) InitSysConfigRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	sysConfigRouter := Router.Group("sysConfig").Use(middleware.OperationRecord())
	sysConfigRouterWithoutRecord := Router.Group("sysConfig")
	sysConfigRouterWithoutAuth := PublicRouter.Group("sysConfig")
	{
		sysConfigRouter.POST("createSysConfig", sysConfigApi.CreateSysConfig)   // 新建sysConfig表
		sysConfigRouter.DELETE("deleteSysConfig", sysConfigApi.DeleteSysConfig) // 删除sysConfig表
		sysConfigRouter.DELETE("deleteSysConfigByIds", sysConfigApi.DeleteSysConfigByIds) // 批量删除sysConfig表
		sysConfigRouter.PUT("updateSysConfig", sysConfigApi.UpdateSysConfig)    // 更新sysConfig表
	}
	{
		sysConfigRouterWithoutRecord.GET("findSysConfig", sysConfigApi.FindSysConfig)        // 根据ID获取sysConfig表
		sysConfigRouterWithoutRecord.GET("getSysConfigList", sysConfigApi.GetSysConfigList)  // 获取sysConfig表列表
	}
	{
	    sysConfigRouterWithoutAuth.GET("getSysConfigPublic", sysConfigApi.GetSysConfigPublic)  // 获取sysConfig表列表
	}
}
