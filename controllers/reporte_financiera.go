package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/resoluciones_crud_v2/models"
)

// Reporte_financieraController operations for Gestion_vinculaciones
type ReporteFinancieraController struct {
	beego.Controller
}

// URLMapping ...
func (c *ReporteFinancieraController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("PostAll", c.PostAll)
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
			fmt.Println("REPORTE ", reporte)
			fmt.Println("REPORTE ", reporte[0])
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

// Post ...
// @Title Reporte Financiera
// @Description Crea el reporte de financiera para todas las facultades y tipos de vinculación
// @Param	body		body 	models.DatosReporte	true		"body for DatosReporte content"
// @Success 201 {object} []models.ReporteFinanciera
// @Failure 400 bad request
// @Failure 500 Internal server error
// @router /all [post]
func (c *ReporteFinancieraController) PostAll() {
	var v models.DatosReporteAll
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if reporte, err := models.ReporteFinancieraV2Query(&v); err == nil {
			fmt.Println("REPORTE ", reporte)
			fmt.Println("REPORTE ", reporte[0])
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
