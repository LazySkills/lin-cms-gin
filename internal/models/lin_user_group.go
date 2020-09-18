/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models


// LinGroup [...]
type LinUserGroup struct {
	ID         int           `gorm:"primary_key;column:id;type:int(11) unsigned auto_increment;not null;comment:'ID';" json:"id"`
	UserId int `gorm:"index:user_id_group_id;column:user_id;type:int(11);unsigned;not null;comment:'用户id'"`
	GroupId int `gorm:"index:user_id_group_id;column:group_id;type:int(11);unsigned;comment:'分组id'"`
	//Permissions []LinGroupPermission `gorm:"ForeignKey:group_id;" json:"Permissions"`
}

func AddLinUserGroup(userId int, groupId int) bool {
	db.Create(&LinUserGroup{
		UserId : userId,
		GroupId : groupId,
	})

	return true
}

func GetLinUserGroupUserIds(groupID int, maps interface{}) (linUserGroup []LinUserGroup)  {
	db.Select("user_id").Where(maps).Where("group_id = ?", groupID).Find(&linUserGroup)
	return
}

func GetLinUserGroupByUserID(userID int) (linUserGroup LinUserGroup)  {
	db.Where("user_id = ?", userID).Find(&linUserGroup)
	return
}
