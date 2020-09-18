/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models

// LinPermission [...]
type LinGroupPermission struct {
	ID         int           `gorm:"primary_key;column:id;type:int(11) unsigned auto_increment;not null;comment:'ID';" json:"id"`
	GroupId 	int 		`gorm:"index:group_id_permission_id;column:group_id;type:int(11);unsigned;comment:'分组id'"`
	PermissionId int 		`gorm:"index:group_id_permission_id;column:permission_id;type:int(11)unsigned;autoIncrement:false;not null;comment:'权限ID'"`
}

type JsonLinGroupPermission struct {
	ID         int
	Name   string
	Module string
	Mount int
}

func GetLinGroupPermissionByGroupId(groupID int) (permissions []JsonLinGroupPermission)  {
	db.Model(&LinGroupPermission{}).Select("lin_group_permission.id,lin_permission.name,lin_permission.module,lin_permission.mount").Where("group_id = ?", groupID).Joins("left join lin_permission on lin_permission.id = lin_group_permission.permission_id").Find(&permissions)
	return
}

func ExistUserPermissionByPermissionId(permissionId int, groupId int) bool {
	var permission LinGroupPermission
	db.Where("permission_id = ? and group_id = ?", permissionId , groupId).Find(&permission)

	if permission.ID > 0 {
		return true
	}
	return false
}

