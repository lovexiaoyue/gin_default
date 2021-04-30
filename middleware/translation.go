package middleware

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/lovexiaoyue/gin-default/app/validate"
	"reflect"
)

func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置语言支持
		en := en2.New()
		zh := zh2.New()

		//设置国际化翻译器
		uni := ut.New(zh,zh,en)
		val := validator.New()

		//根据参数取翻译器实例
		loacle := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(loacle)

		//翻译器注册到 validator
		switch loacle {
		case "en":
			en_translations.RegisterDefaultTranslations(val,trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			zh_translations.RegisterDefaultTranslations(val,trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			val.RegisterValidation("is-validuser", func(fl validator.FieldLevel) bool {
				return fl.Field().String() == "admin"
			})

			//自定义验证器
			val.RegisterTranslation("is-validuser", trans, func(ut ut.Translator) error {
				return ut.Add("is-validuser","{0} 填写不正确",true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t,_  := ut.T("is-validuser",fe.Field())
				return t
			})
			break
		}
		c.Set(validate.TranslatorKey, trans)
		c.Set(validate.ValidatorKey, val)
		c.Next()
	}
}
