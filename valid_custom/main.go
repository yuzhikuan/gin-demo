package main

import "github.com/gin-gonic/gin"

import "time"

import "net/http"

import "gopkg.in/go-playground/validator.v9"

import "github.com/gin-gonic/gin/binding"

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// 自定义校验函数
func bookableDate(fl validator.FieldLevel) bool {
	filed := fl.Field()
	if date, ok := filed.Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

func main() {
	r := gin.Default()

	// 注册验证的tag
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/bookable", func(c *gin.Context) {
		var b Booking
		if err := c.ShouldBind(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok!",
			"booking": b,
		})
	})

	r.Run()
}
