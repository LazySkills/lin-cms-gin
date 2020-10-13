/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models

import (
	"gorm.io/gorm"
	"lin-cms-gin/pkg/setting"
	jtime "lin-cms-gin/pkg/time"
)

// LinGroup [...]
type LinGroup struct {
	Model
	Name       string         `gorm:"unique_index:name_del;column:name;type:varchar(60);not null;comment:'分组名称，例如：搬砖者'"`
	Info       string         `gorm:"column:info;type:varchar(255);comment:'分组信息：例如：搬砖的人'"`
	Level      int            `gorm:"column:level;type:tinyint(2);default:3;comment:'分组级别 1：root 2：guest 3：user  root（root、guest分组只能存在一个)'"`
	DeleteTime jtime.JSONTime `gorm:"unique_index:name_del;column:delete_time;type:datetime(3);default:NULL;comment:'删除时间'"`
}

// LinGroup [...]
type LinGroupJson struct {
	ID  int  ` json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
}

func AddLinGroup(name string, info string, level int) (group LinGroup) {
	group = LinGroup{
		Name : name,
		Info : info,
		Level : level,
	}
	db.Create(&group)

	return
}

func (this *LinGroup) UpdateLinGroup(name string, info string) bool {
	result := db.Model(&this).Select("name","info","level").Updates(map[string]interface{}{"name": name, "info": info})

	if result.Error != nil {
		panic(result.Error)
	}

	return true
}

func GetLinGroupByID(id int) (linGroup LinGroup) {
	db.Where("id = ?", id).Find(&linGroup)
	return
}


func GetLinGroup(pageNum int, pageSize int, maps map[string]interface{}) (users []LinGroupJson) {
	db.Model(LinGroup{}).Where(maps).Offset(pageNum).Limit(pageSize).Select("id, name, info").Scan(&users)

	return
}

func GetLinGroupTotal(maps map[string]interface{}) (count int64){
	db.Model(&LinGroup{}).Where(maps).Count(&count)

	return
}

func DeleteLinGroup(id int) bool {

	linGroup := GetLinGroupByID(id)

	if linGroup.ID < 1 {
		return false
	}

	if linGroup.Level == setting.LinSetting.GroupLevelRoot {
		return false
	}else if linGroup.Level == setting.LinSetting.GroupLevelGuest  {
		return false
	}

	// 开启事务
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})

	tx.Transaction(func(txs *gorm.DB) error {
		if err := txs.Delete(&LinGroupPermission{GroupId: id}).Error; err != nil {
			return err
		}

		if err := txs.Delete(&LinUserGroup{GroupId: id}).Error; err != nil {
			return err
		}

		if err := txs.Delete(&linGroup).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

	return true
}

func DispatchPermission(groupID int, permissionID int) (bool, error)  {
	linGroup := GetLinGroupByID(groupID)

	if linGroup.ID < 1 {
		return false,nil
	}

	permission := GetLinPermissionById(permissionID)
	if permission.ID < 1 {
		return false,nil
	}

	isExist := ExistUserPermissionByPermissionId(permissionID,groupID)
	if isExist {
		return false,nil
	}

	AddLinGroupPermission(groupID,permissionID)

	return true,nil
}

func DispatchPermissions(groupID int, permissionIDS []int) (bool, error)  {
	linGroup := GetLinGroupByID(groupID)

	if linGroup.ID < 1 {
		return false,nil
	}
	// 开启事务
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})

	tx.Transaction(func(txs *gorm.DB) error {
		for permissionID, _ := range permissionIDS {
			permission := GetLinPermissionById(permissionID)
			if permission.ID < 1 {
				panic("分配权限不存在")
			}

			isExist := ExistUserPermissionByPermissionId(permissionID, groupID)
			if isExist {
				panic("分配权限不存在")
			}

			result := db.Create(&LinGroupPermission{
				GroupId: groupID,
				PermissionId: permissionID,
			})

			if result.Error != nil {
				panic(result.Error)
			}
		}
		return  nil
	})
	return true,nil
}

func RemovePermissions(groupID int, permissionIDS []int) bool  {
	linGroup := GetLinGroupByID(groupID)

	if linGroup.ID < 1 {
		return false
	}
	// 开启事务
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})

	tx.Transaction(func(txs *gorm.DB) error {
		for permissionID, _ := range permissionIDS {
			permission := GetLinPermissionById(permissionID)
			if permission.ID < 1 {
				panic("分配权限不存在")
			}

			isExist := ExistUserPermissionByPermissionId(permissionID, groupID)
			if isExist {
				panic("分配权限不存在")
			}

			txs.Where("group_id = ? AND permission_id = ?",groupID,permissionID).Delete(&LinGroupPermission{})
		}
		return  nil
	})
	return true
}