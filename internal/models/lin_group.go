/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models

import (
	jtime "lin-cms-gin/internal/pkg/time"
)

// LinGroup [...]
type LinGroup struct {
	Model
	Name string `gorm:"unique_index:name_del;column:name;type:varchar(60);not null;comment:'分组名称，例如：搬砖者'"`
	Info string `gorm:"column:info;type:varchar(255);comment:'分组信息：例如：搬砖的人'"`
	Level int `gorm:"column:level;type:tinyint(2);default:3;comment:'分组级别 1：root 2：guest 3：user  root（root、guest分组只能存在一个)'"`
	DeleteTime jtime.JSONTime  `gorm:"unique_index:name_del;column:delete_time;type:datetime(3);default:NULL;comment:'删除时间'"`
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

