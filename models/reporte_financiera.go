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
	ProyectoCurricular    int     `orm:"column(proyectocurricular);null"`
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
			v.proyecto_curricular_id as proyectocurricular,
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
			AND (r.tipo_resolucion_id=663 OR r.tipo_resolucion_id=664)
			AND (re.estado_resolucion_id=671 AND re.activo = true)
			AND v.numero_contrato is not null
		GROUP BY r.id, r.numero_resolucion, v.id
		ORDER BY r.id DESC;`
	fmt.Println("QUERY ", query)
	_, err = o.Raw(query).QueryRows(&reporte)
	fmt.Println(reporte)
	return reporte, nil
}

type ReporteResolucion struct {
	Id                    int     `db:"id"`
	Resolucion            string  `db:"resolucion"`
	Vigencia              int     `db:"vigencia"`
	Periodo               int     `db:"periodo"`
	NivelAcademico        string  `db:"nivel_academico"`
	TipoVinculacion       string  `db:"tipo_vinculacion"`
	DocumentoDocente      int     `db:"documento_docente"`
	Horas                 float64 `db:"horas"`
	Semanas               int     `db:"semanas"`
	Total                 float64 `db:"total"`
	Cdp                   int     `db:"cdp"`
	Rp                    int     `db:"rp"`
	Proyectocurricular    int     `db:"proyectocurricular"`
	TipoResolucion        string  `db:"tipo_resolucion"`
	Sueldobasico          float64 `db:"sueldobasico"`
	Primanavidad          float64 `db:"primanavidad"`
	Vacaciones            float64 `db:"vacaciones"`
	Primavacaciones       float64 `db:"primavacaciones"`
	Cesantias             float64 `db:"cesantias"`
	Interesescesantias    float64 `db:"interesescesantias"`
	Primaservicios        float64 `db:"primaservicios"`
	Bonificacionservicios float64 `db:"bonificacionservicios"`
}

func ReporteFinancieraV2Query(m *DatosReporteAll) (reporte []ReporteResolucion, err error) {
	o := orm.NewOrm()
	err = o.Begin()

	if err != nil {
		return
	}

	fmt.Println("m ", m)

	query :=
		`SELECT r.id, r.numero_resolucion as resolucion,
 			r.vigencia,
			r.periodo,
			case 
				when r.dependencia_id = 17 then 'FACULTAD DE CIENCIAS Y EDUCACION'
				when r.dependencia_id = 65 then 'FACULTAD DE MEDIO AMBIENTE'
				when r.dependencia_id = 66 then 'FACULTAD TECNOLOGICA'
				when r.dependencia_id = 8 then 'VICERRECTORIA ACADEMICA'
				when r.dependencia_id = 14 then 'FACULTAD DE INGENIERIA'
				when r.dependencia_id = 35 then 'FACULTAD DE ARTES - ASAB'
				when r.dependencia_id = 232 then 'FACULTAD DE CIENCIAS MATEMATICAS Y NATURALES'
 			end as facultad,	

			rv.nivel_academico, 
			rv.dedicacion tipo_vinculacion,
			v.persona_id as documento_docente,
			SUM(v.numero_horas_semanales) filter (WHERE dv.rubro='SueldoBasico') as Horas,
			r.numero_semanas as Semanas,
			SUM(v.valor_contrato) filter (WHERE dv.rubro='SueldoBasico') as Total,
			SUM(DISTINCT dv.disponibilidad) as cdp,
			v.numero_rp rp,
			v.proyecto_curricular_id as proyectocurricular,
			case 
				when r.tipo_resolucion_id = 663 then 'vinculación'
				when r.tipo_resolucion_id = 664 then 'adición'
				when r.tipo_resolucion_id = 665 then 'reducción'
				when r.tipo_resolucion_id = 666 then 'cancelación'
			end as tipo_resolucion,
			SUM(dv.valor) filter (WHERE dv.rubro='SueldoBasico') as sueldobasico,
			SUM(dv.valor) filter (WHERE dv.rubro='PrimaNavidad') as primanavidad,
			SUM(dv.valor) filter (WHERE dv.rubro='Vacaciones') as vacaciones,
			SUM(dv.valor) filter (WHERE dv.rubro='PrimaVacaciones') as primavacaciones,
			SUM(dv.valor) filter (WHERE dv.rubro='Cesantias') as cesantias,
			SUM(dv.valor) filter (WHERE dv.rubro='InteresesCesantias') as interesescesantias,
			SUM(dv.valor) filter (WHERE dv.rubro='PrimaServicios') as primaservicios,
			SUM(dv.valor) filter (WHERE dv.rubro='BonificacionServicios') as bonificacionservicios
		FROM resoluciones_new.resolucion r, 
			resoluciones_new.resolucion_vinculacion_docente rv, 
			resoluciones_new.resolucion_estado re, 
			resoluciones_new.vinculacion_docente v, 
			resoluciones_new.disponibilidad_vinculacion dv
			WHERE
				r.id = rv.id AND rv.id = v.resolucion_vinculacion_docente_id AND r.id = re.resolucion_id AND v.id=dv.vinculacion_docente_id
				AND r.dependencia_id= ` + strconv.Itoa(m.Facultad) + `
				AND r.vigencia=` + strconv.Itoa(m.Vigencia) + `
				AND rv.nivel_academico='` + m.NivelAcademico + `'
				AND (r.tipo_resolucion_id=663 OR r.tipo_resolucion_id=664 OR r.tipo_resolucion_id=665 OR r.tipo_resolucion_id=666)
				AND (re.estado_resolucion_id=671 AND re.activo = true)
				AND v.numero_contrato is not null
		GROUP BY r.id, r.numero_resolucion, v.id, rv.nivel_academico, rv.dedicacion 
		ORDER BY r.id DESC;`
	fmt.Println("QUERY ", query)
	_, err = o.Raw(query).QueryRows(&reporte)
	fmt.Printf("Data: %+v\n", reporte)
	return reporte, nil
}
