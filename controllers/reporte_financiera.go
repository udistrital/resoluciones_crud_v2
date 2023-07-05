package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_crud_v2/models"
)

// Reporte_financieraController operations for Gestion_vinculaciones
type ReporteFinancieraController struct {
	beego.Controller
}

// Post ...
// @Title Reporte Financiera
// @Description Crea el reporte de financiera
// @Param	body		body 	models.DatosReporte	true		"body for DatosReporte content"
// @Success 201 {object} []models.ReporteFinanciera
// @Failure 400 bad request
// @Failure 500 Internal server error
// @router / [post]
func (c *ReporteFinancieraController) Post() {
	var v models.DatosReporte
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if reporte, err := models.ReporteFinancieraQuery(&v); err == nil {
			c.Ctx.Output.SetStatus(201)

			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "201", "Message": "Registration successful", "Data": reporte}
		} else {
			logs.Error(err)
			c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Data["mesaage"] = "Error service POST: The request contains an incorrect data type or an invalid parameter"
		c.Abort("400")
	}
	c.ServeJSON()
}
