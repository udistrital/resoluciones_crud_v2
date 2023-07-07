package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type ReporteFinanciera struct {
	Id                    int     `orm:"column(id);null"`
	Resolucion            string  `orm:"column(resolucion);null"`
	Cedula                int     `orm:"column(cedula);null"`
	Horas                 int     `orm:"column(horas);null"`
	Semanas               int     `orm:"column(semanas);null"`
	Total                 float64 `orm:"column(total);null"`
	Cdp                   int     `orm:"column(cdp);null"`
	SueldoBasico          float64 `orm:"column(sueldobasico);null"`
	PrimaNavidad          float64 `orm:"column(primanavidad);null"`
	Vacaciones            float64 `orm:"column(vacaciones);null"`
	PrimaVacaciones       float64 `orm:"column(primavacaciones);null"`
	Cesantias             float64 `orm:"column(cesantias);null"`
	InteresesCesantias    float64 `orm:"column(interesescesantias);null"`
	PrimaServicios        float64 `orm:"column(primaservicios);null"`
	BonificacionServicios float64 `orm:"column(bonificacionservicios);null"`
}

func ReporteFinancieraQuery(m *DatosReporte) (reporte []ReporteFinanciera, err error) {
	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return
	}

	fmt.Println("m ", m)

	query :=
		`SELECT r.id, r.numero_resolucion as resolucion,
			v.persona_id as cedula,
			SUM(v.numero_horas_semanales) filter (WHERE dv.rubro='SueldoBasico') as Horas,
			r.numero_semanas as Semanas,
			SUM(v.valor_contrato) filter (WHERE dv.rubro='SueldoBasico') as Total,
			SUM(DISTINCT dv.disponibilidad) as cdp,
			SUM(dv.valor) filter (WHERE dv.rubro='SueldoBasico') as sueldobasico,
			SUM(dv.valor) filter (WHERE dv.rubro='PrimaNavidad') as primanavidad,
			SUM(dv.valor) filter (WHERE dv.rubro='Vacaciones') as vacaciones,
			SUM(dv.valor) filter (WHERE dv.rubro='PrimaVacaciones') as primavacaciones,
			SUM(dv.valor) filter (WHERE dv.rubro='Cesantias') as cesantias,
			SUM(dv.valor) filter (WHERE dv.rubro='InteresesCesantias') as interesescesantias,
			SUM(dv.valor) filter (WHERE dv.rubro='PrimaServicios') as primaservicios,
			SUM(dv.valor) filter (WHERE dv.rubro='BonificacionServicios') as bonificacionservicios
		FROM resoluciones_new.resolucion r, resoluciones_new.resolucion_vinculacion_docente rv, resoluciones_new.resolucion_estado re, resoluciones_new.vinculacion_docente v, resoluciones_new.disponibilidad_vinculacion dv
		WHERE
			r.id = rv.id AND rv.id = v.resolucion_vinculacion_docente_id AND r.id = re.resolucion_id AND v.id=dv.vinculacion_docente_id
			AND r.dependencia_id=` + strconv.Itoa(m.Facultad) + `AND rv.nivel_academico='` + m.NivelAcademico + `'
			AND r.numero_resolucion='` + m.Resolucion + `' AND r.vigencia=` + strconv.Itoa(m.Vigencia) + `
			AND r.tipo_resolucion_id=583
			AND re.estado_resolucion_id=602
		GROUP BY r.id, r.numero_resolucion, v.persona_id, r.numero_semanas
		ORDER BY r.id DESC;`
	_, err = o.Raw(query).QueryRows(&reporte)
	return reporte, nil
}
