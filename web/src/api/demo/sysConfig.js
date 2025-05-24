import service from '@/utils/request'

// @Tags SysConfig
// @Summary 创建sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "创建sysConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysConfig/createSysConfig [post]
export const createSysConfig = (data) => {
  return service({
    url: '/sysConfig/createSysConfig',
    method: 'post',
    data
  })
}

// @Tags SysConfig
// @Summary 删除sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "删除sysConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysConfig/deleteSysConfig [delete]
export const deleteSysConfig = (params) => {
  return service({
    url: '/sysConfig/deleteSysConfig',
    method: 'delete',
    params
  })
}

// @Tags SysConfig
// @Summary 批量删除sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除sysConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysConfig/deleteSysConfig [delete]
export const deleteSysConfigByIds = (params) => {
  return service({
    url: '/sysConfig/deleteSysConfigByIds',
    method: 'delete',
    params
  })
}

// @Tags SysConfig
// @Summary 更新sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysConfig true "更新sysConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysConfig/updateSysConfig [put]
export const updateSysConfig = (data) => {
  return service({
    url: '/sysConfig/updateSysConfig',
    method: 'put',
    data
  })
}

// @Tags SysConfig
// @Summary 用id查询sysConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SysConfig true "用id查询sysConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysConfig/findSysConfig [get]
export const findSysConfig = (params) => {
  return service({
    url: '/sysConfig/findSysConfig',
    method: 'get',
    params
  })
}

// @Tags SysConfig
// @Summary 分页获取sysConfig表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取sysConfig表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysConfig/getSysConfigList [get]
export const getSysConfigList = (params) => {
  return service({
    url: '/sysConfig/getSysConfigList',
    method: 'get',
    params
  })
}
