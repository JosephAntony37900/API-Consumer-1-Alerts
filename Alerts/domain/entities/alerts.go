package entities

import "time"

type Alerts struct {
	Id             int
	Id_Lectura     int
	Estado         string
	Fecha_Creacion time.Time
	IdRol int
	Codigo_Identificador string
	Tipo bool
}