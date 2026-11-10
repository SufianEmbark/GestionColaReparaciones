package reparacion

import "errors"

var ErrCitasSolapadas = errors.New("Haquin no puede estar en dos trabajos a la vez")
var ErrUrgenciaInexistente = errors.New("la urgencia no es urgente, normal ni puede esperar")
var ErrFinAnteriorAlInicio = errors.New("la cita termina antes de empezar o a la vez")
var ErrCitaExcedeDuracionMaxima = errors.New("la cita dura más de las 3 horas seguidas que Haquin dedica a un trabajo")
var ErrCitaFueraDeHorario = errors.New("la cita queda fuera del horario de trabajo de Haquin")
var ErrCitaAntesDelMaterial = errors.New("la cita es anterior al día desde el que hay material")
var ErrCitaFueraDeDisponibilidad = errors.New("la cita queda fuera de la disponibilidad del cliente")
