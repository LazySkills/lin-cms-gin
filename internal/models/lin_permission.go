/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models

import (
	jtime "lin-cms-gin/internal/pkg/time"
)

// LinPermission [...]
type LinPermission struct {
	Model
	Name   string `gorm:"column:name;type:varchar(60);not null;comment:'权限名称，例如：访问首页'"`
	Module string `gorm:"column:module;type:varchar(50);not null;comment:'权限所属模块，例如：人员管理'"`
	Mount int `gorm:"column:mount;type:tinyint(2);default:1;comment:'0：关闭 1：开启'"`
	DeleteTime jtime.JSONTime  `gorm:"column:delete_time;type:datetime(3);default:NULL;comment:'删除时间'"`
}

func GetAllLinPermission() (permissions []LinPermission)  {
	db.Find(&permissions)
	return
}


func AddLinPermission(name string,module string,mount int) (per LinPermission) {
	db.Model(&per).Create(map[string]interface{}{"name": name, "module": module, "mount": mount})
	return
}

func (this *LinPermission) UpdatePermission(name string,module string,mount int) bool {
	db.Model(&this).Updates(map[string]interface{}{"name": name, "module": module, "mount": mount})
	return true
}

func GetLinPermissionByName(name string) (per LinPermission) {
	db.Where("name = ?", name).Find(&per)
	return
}
