package controllers

import (
	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-3/services"
	"github.com/santiagoh1997/weather-api/version-3/utils"
)

// JobController connects the request with the model
type JobController struct {
	beego.Controller
	Service *services.JobService
}

// NewJobController returns a new JobController
func NewJobController(js *services.JobService) *JobController {
	return &JobController{
		Service: js,
	}
}

// Schedule schedules a new job to be performed hourly
func (jc *JobController) Schedule() {
	var req utils.Request
	req.City = jc.GetString("city")
	req.Country = jc.GetString("country")
	if err := req.ValidateRequest(); err != nil {
		jc.Ctx.Output.SetStatus(err.StatusCode)
		jc.Data["json"] = err
		jc.ServeJSON()
		return
	}
	if _, apiErr := jc.Service.NewJob(req.City, req.Country); apiErr != nil {
		jc.Ctx.Output.SetStatus(apiErr.StatusCode)
		jc.Data["json"] = apiErr
		return
	}
	response := struct {
		Message string `json:"message"`
	}{
		"Job scheduled",
	}
	jc.Data["json"] = response
	jc.Ctx.Output.SetStatus(202)
	jc.ServeJSON()
}
