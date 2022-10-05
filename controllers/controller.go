package controller

import (
	"crud-user-mvc/helpers"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type Controller struct {
	db     *gorm.DB
	helper helper.Helper
}

func Init(db *gorm.DB, hp helper.Helper) *Controller {
	return &Controller{db: db, helper: hp}
}

func (c *Controller) BindParam(ctx *gin.Context, param interface{}) error {
	if err := ctx.ShouldBindUri(param); err != nil {
		return err
	}

	return ctx.ShouldBindWith(param, binding.Query)
}

func (c *Controller) BindBody(ctx *gin.Context, body interface{}) error {
	return ctx.ShouldBindWith(body, binding.Default(ctx.Request.Method, ctx.ContentType()))
}
