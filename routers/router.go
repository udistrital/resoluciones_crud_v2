// @APIVersion 1.0.0
// @Title Resoluciones CRUD Version 2
// @Description Nueva versión de API CRUD para el sistema de Resoluciones
// @Contact computo@udistrital.edu.co
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/resoluciones_crud_v2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/resolucion_estado",
			beego.NSInclude(
				&controllers.ResolucionEstadoController{},
			),
		),

		beego.NSNamespace("/modificacion_resolucion",
			beego.NSInclude(
				&controllers.ModificacionResolucionController{},
			),
		),

		beego.NSNamespace("/modificacion_vinculacion",
			beego.NSInclude(
				&controllers.ModificacionVinculacionController{},
			),
		),

		beego.NSNamespace("/componente_resolucion",
			beego.NSInclude(
				&controllers.ComponenteResolucionController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion_docente",
			beego.NSInclude(
				&controllers.ResolucionVinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/vinculacion_docente",
			beego.NSInclude(
				&controllers.VinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/disponibilidad_vinculacion",
			beego.NSInclude(
				&controllers.DisponibilidadVinculacionController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
