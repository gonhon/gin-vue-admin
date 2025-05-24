package demo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/demo"
	demoReq "github.com/flipped-aurora/gin-vue-admin/server/model/demo/request"
)

type SysConfigService struct{}

// CreateSysConfig 创建sysConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) CreateSysConfig(sysConfig *demo.SysConfig) (err error) {
	err = global.GVA_DB.Create(sysConfig).Error
	return err
}

// DeleteSysConfig 删除sysConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) DeleteSysConfig(id string) (err error) {
	err = global.GVA_DB.Delete(&demo.SysConfig{}, "id = ?", id).Error
	return err
}

// DeleteSysConfigByIds 批量删除sysConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) DeleteSysConfigByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]demo.SysConfig{}, "id in ?", ids).Error
	return err
}

// UpdateSysConfig 更新sysConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) UpdateSysConfig(sysConfig demo.SysConfig) (err error) {
	err = global.GVA_DB.Model(&demo.SysConfig{}).Where("id = ?", sysConfig.Id).Updates(&sysConfig).Error
	return err
}

// GetSysConfig 根据id获取sysConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) GetSysConfig(id string) (sysConfig demo.SysConfig, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysConfig).Error
	return
}

// GetSysConfigInfoList 分页获取sysConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) GetSysConfigInfoList(info demoReq.SysConfigSearch) (list []demo.SysConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&demo.SysConfig{})
	var sysConfigs []demo.SysConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysConfigs).Error
	return sysConfigs, total, err
}
