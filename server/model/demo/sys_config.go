// 自动生成模板SysConfig
package demo

import (
	"time"
)

// sysConfig表 结构体  SysConfig
type SysConfig struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:主键编码;size:19;"`  //主键编码 
    ConfigName  string `json:"configName" form:"configName" gorm:"column:config_name;comment:ConfigName;size:128;"`  //ConfigName 
    ConfigKey  string `json:"configKey" form:"configKey" gorm:"column:config_key;comment:ConfigKey;size:128;"`  //ConfigKey 
    ConfigValue  string `json:"configValue" form:"configValue" gorm:"column:config_value;comment:ConfigValue;size:255;"`  //ConfigValue 
    ConfigType  string `json:"configType" form:"configType" gorm:"column:config_type;comment:ConfigType;size:64;"`  //ConfigType 
    IsFrontend  string `json:"isFrontend" form:"isFrontend" gorm:"column:is_frontend;comment:是否前台;size:64;"`  //是否前台 
    Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:Remark;size:128;"`  //Remark 
    CreateBy  *int `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者;size:19;"`  //创建者 
    UpdateBy  *int `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:更新者;size:19;"`  //更新者 
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"`  //创建时间 
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:最后更新时间;"`  //最后更新时间 
    DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:删除时间;"`  //删除时间 
}


// TableName sysConfig表 SysConfig自定义表名 sys_config
func (SysConfig) TableName() string {
    return "sys_config"
}

