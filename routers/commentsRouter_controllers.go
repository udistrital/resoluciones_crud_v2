package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ComponenteResolucionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:DisponibilidadVinculacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionResolucionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ModificacionVinculacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionEstadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:ResolucionVinculacionDocenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud_v2/controllers:VinculacionDocenteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
