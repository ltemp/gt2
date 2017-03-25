package controllers

import (
	"gett2/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type DriverController struct {
	beego.Controller
}

// @Title GetDriver
// @Description get driver by id
// @Param	id		path 	int	true		"driver id"
// @Success 200 {object} models.Driver
// @Failure 403 :id is not int or empty
// @router /:id [get]
func (u *DriverController) Get() {
	id, err := u.GetInt(":id")
	if err != nil {
		u.Abort("400")
	}
	if err == nil {
		driver, err := models.LoadDriver(id)
		if err != nil {
			u.Abort("404")
		} else {
			u.Data["json"] = driver
		}
	}

	u.ServeJSON()
}

// @Title CreateDriver
// @Description create driver
// @Param	body		body 	models.Driver	true		"body with driver info"
// @Failure 500 error adding driver
// @Failure 400 bad request
// @router / [post]
func (u *DriverController) Post() {
	var driver models.Driver
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &driver)
	if err != nil {
		u.Abort("400")
	}
	err = models.InsertDriver(driver)
	if err != nil {
		u.Abort("500")
	}
	u.Data["json"] = "added succesfully"
	u.ServeJSON()
}

// @Title UpdateDriver
// @Description update driver
// @Param	body		body 	models.Driver	true		"body with driver info"
// @Success 200
// @Failure 400 bad request
// @Failure 500 error updating driver
// @router / [put]
func (u *DriverController) Put() {
	var driver models.Driver
	var err = json.Unmarshal(u.Ctx.Input.RequestBody, &driver)
	if err != nil {
		u.Abort("400")
	}
	err, updated := models.UpdateDriver(driver)
	if err != nil {
		u.Abort("500")
	}
	if !updated {
		u.Abort("404")
	}
	u.Data["json"] = "updated succesfully"
	u.ServeJSON()
}

// @Title Delete
// @Description delete driver by id
// @Param	id		path 	int	true		"driver id"
// @Success 200 {object} models.Driver
// @Failure 403 :id is not int or empty
// @router /:id [delete]
func (u *DriverController) Delete() {
	id, err := u.GetInt(":id")
	if err != nil {
		u.Abort("400")
	}
	if err == nil {
		err, updated := models.DeleteDriver(id)
		if err != nil {
			u.Abort("500")
		}
		if !updated {
			u.Abort("404")
		}
	}
	u.Data["json"] = "updated succesfully"
	u.ServeJSON()
}
