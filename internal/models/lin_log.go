/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models

import jtime "lin-cms-gin/pkg/time"

// LinLog [...]
type LinLog struct {
	Model
	Message    string         `gorm:"column:message;type:varchar(450);not null;comment:'消息内容'"`
	UserID     int            `gorm:"column:user_id;type:int(11);not null;comment:'用户ID'"`
	UserName   string         `gorm:"column:user_name;type:varchar(24);not null;comment:'用户昵称'"`
	StatusCode int            `gorm:"column:status_code;type:int(11);not null;comment:'状态码'"`
	Method     string         `gorm:"column:method;type:varchar(20);not null;comment:'方法'"`
	Path       string         `gorm:"column:path;type:varchar(50);not null;comment:'路径'"`
	Permission string         `gorm:"column:permission;type:varchar(100);default:NULL;comment:'许可'"`
	DeleteTime jtime.JSONTime `gorm:"column:delete_time;type:datetime(3);default:NULL;comment:'删除时间'"`
}

