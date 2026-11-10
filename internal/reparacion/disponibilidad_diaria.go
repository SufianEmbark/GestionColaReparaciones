package reparacion

import "time"

type DisponibilidadDiaria struct {
	dia    time.Weekday
	inicio time.Duration
	fin    time.Duration
}
