package reparacion

import "time"

type Trabajo struct {
	id             int
	urgencia       Urgencia
	llamada        time.Time
	materialDesde  time.Time
	duracion       time.Duration
	citas          []Cita
	disponibilidad []DisponibilidadDiaria
}
