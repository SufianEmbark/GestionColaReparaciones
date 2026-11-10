package reparacion

import "time"

const (
	PrimerDiaLaborable = time.Monday
	UltimoDiaLaborable = time.Friday
	InicioHorario      = 17 * time.Hour
	FinHorario         = 20*time.Hour + 30*time.Minute
)
