/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models

import (
	"lin-cms-gin/internal/pkg/logging"
	jtime "lin-cms-gin/internal/pkg/time"
	"lin-cms-gin/internal/pkg/tools"
)

// LinPoem [...]
type LinUserIdentity struct {
	Model
	UserId       int    `gorm:"column:user_id;type:int(11) unsigned;not null;comment:'用户id';"`
	IdentityType string `gorm:"column:identity_type;type:varchar(100);not null;comment:'加密类型';"`
	Identifier   string `gorm:"column:identifier;type:varchar(100);not null;comment:'检验人';"`
	Credential   string `gorm:"column:credential;type:varchar(100);not null;comment:'凭证';"`
	DeleteTime jtime.JSONTime  `gorm:"column:delete_time;type:datetime(3);default:NULL;comment:'删除时间'"`
}

func AddIdentity(userId int, identityType string, identifier string, credential string) bool {
	db.Create(&LinUserIdentity{
		UserId : userId,
		IdentityType : identityType,
		Identifier : identifier,
		Credential : credential,
	})

	return true
}

func ExistIdentityByName(name string) (identity LinUserIdentity) {
	db.Select("id").Where("identifier = ?", name).First(&identity)

	return identity
}


func ExistIdentityByUserID(userID int) (identity LinUserIdentity) {
	db.Select("id").Where("user_id = ?", userID).First(&identity)

	return identity
}

func UpdateIdentityByID(id int, data interface{}) bool {
	db.Model(&LinUserIdentity{}).Where("id = ?", id).Updates(data)
	return true
}

// 更新用户密码并检验旧密码
func UpdateIdentityPasswordByUserID(userId int, password string, newPassword string) bool {
	var identity = ExistIdentityByUserID(userId)
	if identity.ID > 0 {
		if !CheckIdentityPassword(identity,password) {
			logging.Warn(identity.Identifier+"用户密码错误")
		}else {
			db.Model(&identity).UpdateColumn("Credential",tools.CreatePbkdf2Hash256(newPassword))
			return true
		}
	}

	return false
}

func UpdateIdentityCredentialByUserID(userId int, password string) bool {
	var identity = ExistIdentityByUserID(userId)
	db.Model(&identity).UpdateColumn("Credential",tools.CreatePbkdf2Hash256(password))
	return true
}

func CheckIdentityPassword(identity LinUserIdentity, password string) bool {
	if _,err := tools.PasswordVerify(password, identity.Credential);err !=nil {
		logging.Warn(err.Error())
	}else {
		return true
	}
	return false
}
