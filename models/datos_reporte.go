package models

type DatosReporte struct {
	Resolucion     string
	NivelAcademico string
	Facultad       int
	Vigencia       int
}

type DatosReporteAll struct {
	NivelAcademico string
	Facultad       int
	Vigencia       int
}
