package demo

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ SysConfigApi }

var sysConfigService = service.ServiceGroupApp.DemoServiceGroup.SysConfigService
