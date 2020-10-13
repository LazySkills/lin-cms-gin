/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models

import jtime "lin-cms-gin/pkg/time"

// LinLog [...]
type LinLog struct {
	Model
	Message    string         `gorm:"column:message;type:varchar(450);not null;comment:'消息内容'"`
	UserID     int            `gorm:"column:user_id;type:int(11);not null;comment:'用户ID'"`
	Username   string         `gorm:"column:username;type:varchar(24);not null;comment:'用户昵称'"`
	StatusCode int            `gorm:"column:status_code;type:int(11);not null;comment:'状态码'"`
	Method     string         `gorm:"column:method;type:varchar(20);not null;comment:'方法'"`
	Path       string         `gorm:"column:path;type:varchar(50);not null;comment:'路径'"`
	Permission string         `gorm:"column:permission;type:varchar(100);default:NULL;comment:'许可'"`
	DeleteTime jtime.JSONTime `gorm:"column:delete_time;type:datetime(3);default:NULL;comment:'删除时间'"`
}

type LinLogUserJson struct {
	Names string
}

func GetLinLog(pageNum int, pageSize int, maps map[string]interface{}) (linLog []LinLog) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&linLog)

	return
}


func GetLinLogTotal(maps map[string]interface{}) (count int64){
	db.Model(&LinUser{}).Distinct("id").Where(maps).Count(&count)

	return
}

func GetLinLogUsers(pageNum int, pageSize int) (users []LinLogUserJson) {
	db.Select("username as names").Group("username").Having("COUNT(lin_log.username)>0").Offset(pageNum).Limit(pageSize).Find(&LinLog{}).Scan(&users)

	return
}


func GetLinLogUsersTotal() (count int64){
	db.Model(&LinUser{}).Distinct("id").Group("username").Having("COUNT(lin_log.username)>0").Count(&count)

	return
}

func AddLinLog(maps interface{}) bool {
	if res := db.Model(&LinLog{}).Create(maps); res.Error != nil{
		panic(res.Error)
	}
	return true
}
