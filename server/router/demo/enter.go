package demo

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ SysConfigRouter }

var sysConfigApi = api.ApiGroupApp.DemoApiGroup.SysConfigApi
