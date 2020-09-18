/** Created By wene<354007048@qq.com> . Date at 2020/6/2 */
package tools

import (
	"lin-cms-gin/internal/pkg/setting"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)



type Claims struct {
	UniqueId int `json:"uniqueId"`
	jwt.StandardClaims
}

// 根据类型创建授权token
//
// int uniqueId 唯一ID
// string userInfo 用户信息
// bool isRefresh 是否为刷新token
func GenerateToken(uniqueId int, isRefresh bool) (string, error) {
	nowTime := time.Now()
	jwtSecret,expire := getJwtConfigByType(isRefresh)
	expireTime := nowTime.Add(expire)
	claims := Claims{
		uniqueId,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}


// 根据传入的token值和类型获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string, isRefresh bool)(*Claims,error){

	jwtSecret,_ := getJwtConfigByType(isRefresh)

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims!=nil{
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims,ok:=tokenClaims.Claims.(*Claims);ok&&tokenClaims.Valid{
			return claims,nil
		}
	}
	return nil,err

}

// 根据类型获取Jwt配置
func getJwtConfigByType(isRefresh bool) (jwtSecret []byte, expireTime time.Duration) {
	if isRefresh {
		return []byte(setting.JwtSetting.JwtRefreshSecret),setting.JwtSetting.JwtRefreshExpireTime
	}else {
		return []byte(setting.JwtSetting.JwtSecret),setting.JwtSetting.JwtExpireTime
	}
}