package controllers

import (
	"gett2/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type MetricController struct {
	beego.Controller
}


/**
NOTE- DUE TO A BUG IN BEEGO - this was added as PUT request.
When changing to GET - the request body isn't retrieved.
 */

// @Title GetMetrics
// @Description get metrics
// @Param	body		body 	models.MetricQueryInfo	true
// @Success 200 {object} models.Metric
// @router / [put]
func (u *MetricController) Get() {
	mi := models.MetricQueryInfo{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &mi)
	err, metrics := models.LoadMetrics(mi)
	if err != nil {
		u.Abort("404")
	}

	u.Data["json"] = metrics
	u.ServeJSON()
}

// @Title Delete
// @Description delete metrics by the given MetricQueryInfo
// @Param	body	body 	models.MetricQueryInfo	true
// @Success 200
// @Failure 500 error deleting
// @Failure 400 bad request
// @router / [delete]
func (u *MetricController) Delete() {
	mi := models.MetricQueryInfo{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &mi)
	if err != nil {
		u.Abort("400")
	}
	err, deleted := models.DeleteMetrics(mi)
	if err != nil {
		u.Abort("500")
	}
	if !deleted {
		u.Abort("404")
	}

	u.Data["json"] = "deleted succesfully"
	u.ServeJSON()
}

// @Title CreateMetric
// @Description creates metric
// @Param	body		body 	models.Metric	true	"body with metric info"
// @Success 200
// @Failure 500 error adding metric
// @Failure 400 bad request
// @router / [post]
func (u *MetricController) Post() {
	var metric models.Metric
	var err = json.Unmarshal(u.Ctx.Input.RequestBody, &metric)
	if err != nil {
		u.Abort("400")
	}
	err = models.InsertMetric(metric)
	if err != nil {
		u.Abort("500")
	}
	u.Data["json"] = "added succesfully"
	u.ServeJSON()
}
