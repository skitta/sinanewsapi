package controllers

import (
	"skitta.com/goweb/sinanewsAPI/models"

	"github.com/astaxie/beego"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title Get
// @Description find object by publishID
// @Param	publishID	path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :publish ID is empty
// @router /:publishID [get]
func (o *ObjectController) Get() {
	publishID := o.Ctx.Input.Param(":publishID")
	if publishID != "" {
		ob, err := models.GetOne(publishID)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ObjectController) GetAll() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	obs, err := models.GetAll()
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = obs
	}
	o.ServeJSON()
}

