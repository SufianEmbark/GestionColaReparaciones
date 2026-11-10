package reparacion

import "time"

const DuracionMaximaCita = 3 * time.Hour

type Cita struct {
	inicio time.Time
	fin    time.Time
}
