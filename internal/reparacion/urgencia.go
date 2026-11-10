package reparacion

type Urgencia int

const (
	PuedeEsperar Urgencia = iota
	Normal
	Urgente
)
