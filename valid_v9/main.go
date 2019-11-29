package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

// Person 定义
type Person struct {
	Age     int    `form:"age" validate:"required,gt=10"`
	Name    string `form:"name" validate:"required"`
	Address string `form:"address" validate:"required"`
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

// 验证信息多语言化
func main() {
	Validate := validator.New()

	zh := zh2.New()
	en := en2.New()
	uni = ut.New(zh, en)

	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {

		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(Validate, trans)
		default:
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		}

		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusBadRequest, "%v", err)
			c.Abort()
			return
		}
		if err := Validate.Struct(person); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceErrs := []string{}
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			c.String(http.StatusBadRequest, "%v", sliceErrs)
			c.Abort()
			return
		}
		c.String(http.StatusOK, "%v", person)
	})

	r.Run()
}
