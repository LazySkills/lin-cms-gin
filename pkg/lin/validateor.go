/** Created By wene<354007048@qq.com> . Date at 2020/6/21 */
package lin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器T
var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
			case "en":
				err = enTranslations.RegisterDefaultTranslations(v, trans)
			case "zh":
				err = zhTranslations.RegisterDefaultTranslations(v, trans)
			default:
				err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		if err != nil {
			return err
		}
		return
	}
	return
}

// 参数验证器
func Validator(c *gin.Context,form interface{}) (error string) {
	if err := c.ShouldBind(form); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			return err.Error()
		}
		// validator.ValidationErrors类型错误则进行翻译
		for _,value := range errs.Translate(trans){
			error += " "+value
		}
		return error
	}
	return
}