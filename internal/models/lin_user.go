/** Created By wene<354007048@qq.com> . Date at 2020/6/2 */
package models

import (
	jtime "lin-cms-gin/pkg/time"
	"lin-cms-gin/pkg/tools"
)

// LinUser [...]
type LinUser struct {
	Model
	Username        string          `gorm:"unique_index:username_del;column:username;type:varchar(24);not null;comment:'用户名称'"`
	Nickname        string          `gorm:"column:nickname;type:varchar(24);comment:'昵称'"`
	Email           string          `gorm:"unique_index:email_del;column:email;type:varchar(100);not null;comment:'邮箱'"`
	Avatar          string          `gorm:"column:avatar;type:varchar(500);comment:'头像'"`
	DeleteTime      jtime.JSONTime  `gorm:"unique_index:username_del,email_del;column:delete_time;type:datetime(3);default:NULL;comment:'删除时间'"`
	linUserIdentity LinUserIdentity `gorm:"FOREIGNKEY:user_id"`
}

type LinUserJosn struct {
	Id string
	Nickname string
	Username string
}


func AddLinUser(username string, nickname string, email string, avatar string) bool {
	db.Create(&LinUser{
		Username : username,
		Nickname : nickname,
		Email : email,
		Avatar : avatar,
	})

	return true
}

func AddLinAuth(username string, password string) (b bool,err error) {
	var user = LinUser{
		Username : username,
		Nickname : username,
		Email : "",
		Avatar : "",
	}

	db.Create(&user)

	c,err := tools.PasswordEncode(password,"lin-cms",0)
	if err != nil {
		return false,err
	}

	AddIdentity(user.ID,"USERNAME_PASSWORD",user.Username,c)

	return true,err
}

func ExistLinUserByUsername(username string) bool {
	var user LinUser
	db.Where("username = ?", username).Find(&user)

	if user.ID > 0 {
		return  true
	}
	return false
}

func UpdateUserInfo(userID int, username string,nickname string,email string,avatar string) bool {
	user := GetLinUserByID(userID)

	if user.ID < 1 { return false }

	if user.Username != username {
		if ExistLinUserByUsername(username) {
			return false
		}else {
			ident := ExistIdentityByUserID(user.ID)
			maps := make(map[string]interface{})
			maps["identifier"] = username
			UpdateIdentityByID(ident.ID,maps)
		}
	}

	db.Model(&LinUser{}).Where("id = ?", userID).Updates(LinUser{
		Username: username,
		Nickname: nickname,
		Email: email,
		Avatar: avatar,
	})

	return true

}


func GetLinUserByID(id int) (user LinUser) {
	db.Where("id = ?", id).Find(&user)
	return
}

func GetLinUser(pageNum int, pageSize int, maps map[string]interface{}) (users []LinUser) {
	db.Where(maps).Offset(pageNum).Not([]int{1}).Limit(pageSize).Find(&users)

	return
}

func GetLinUserTotal(maps map[string]interface{}) (count int64){
	db.Model(&LinUser{}).Not([]int{1}).Distinct("id").Where(maps).Count(&count)

	return
}

func DeleteLinUser(id int) bool {
	db.Where("id = ?", id).Preload("linUserIdentity").Delete(&LinUser{})

	return true
}



func (this *LinUser)Migrate()  {
	m := db.Migrator()
	if !m.HasTable(&this) {
		m.AutoMigrate(&this)

	}
}
