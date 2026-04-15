package models

type DatosReporte struct {
	Resolucion                  string
	NivelAcademico              string
	Facultad                    int
	Vigencia                    int
	TipoResolucionVinculacionId int
	TipoResolucionAdicionId     int
	TipoResolucionReduccionId   int
	TipoResolucionCancelacionId int
	EstadoResolucionExpedidaId  int
}

type DatosReporteAll struct {
	NivelAcademico              string
	Facultad                    int
	Vigencia                    int
	TipoResolucionVinculacionId int
	TipoResolucionAdicionId     int
	TipoResolucionReduccionId   int
	TipoResolucionCancelacionId int
	EstadoResolucionExpedidaId  int
}
