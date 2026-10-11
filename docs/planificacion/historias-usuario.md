# Historias de usuario

Las historias de usuario expresan los problemas de [Haquin](personas.md) y salen de su [journey](user-journeys.md#journey-1-una-tarde-de-trabajo-de-haquin). En GitHub son issues con la etiqueta `user-stories`, y se irán resolviendo mediante otros issues agrupados en los milestones.

Haquin apunta los trabajos desde el móvil cuando le llaman, a lo largo de la mañana y a la hora de comer, y los consulta por la tarde mientras va de un trabajo a otro. Por eso la información tiene que estar disponible en todo momento desde cualquier lugar, y el proyecto se desplegará en la nube.

## [HU001] Cuando un trabajo se me bloquea, no sé qué otro trabajo pendiente puedo hacer en ese hueco

Como Haquin, cuando una visita se queda a medias porque me falta material o no me da tiempo, o cuando empiezo la tarde, no sé cuál de mis trabajos pendientes puedo hacer en ese momento: que el cliente pueda estar en casa a esa hora y que yo tenga lo que necesito. Por eso pierdo horas de la tarde sin trabajar teniendo trabajos que podría estar haciendo.

**Journey relacionado:** [Journey 1](user-journeys.md#journey-1-una-tarde-de-trabajo-de-haquin), pasos 3 y 6.

## [HU002] Cuando un trabajo se queda a medias, no sé cuándo volver sin solapar ni retrasar los demás

Como Haquin, cuando tengo que volver otro día a un trabajo porque espero una pieza o porque no me dio tiempo a terminarlo, no sé en qué hueco encajar la vuelta sin pisar otro trabajo ni retrasar todos los demás. Por eso acabo dejándolo para el final o reorganizándolo todo de cabeza.

**Journey relacionado:** [Journey 1](user-journeys.md#journey-1-una-tarde-de-trabajo-de-haquin), pasos 5 y 7.

## [HU003] Cuando me llama un cliente, no sé qué día y a qué hora podré ir

Como Haquin, cuando acepto un trabajo nuevo no sé cuándo podré atenderlo con los que ya tengo, así que no puedo darle al cliente una cita concreta. Por eso solo acepto trabajos que pueda empezar en una semana más o menos, y rechazo los demás.

**Journey relacionado:** [Journey 1](user-journeys.md#journey-1-una-tarde-de-trabajo-de-haquin), pasos 1 y 2.

## Conceptos

- **Trabajo:** una reparación que pide un cliente. Se conoce qué hay que hacer, cuándo llamó, una duración estimada y cuándo puede estar el cliente en casa.
- **Visita:** cada vez que Haquin va a un trabajo. Termina el trabajo o lo deja a medias.
- **Trabajo pendiente:** un trabajo solicitado que todavía no está terminado.
- **Trabajo bloqueado:** un trabajo pendiente que no se puede hacer hasta una fecha, porque falta un material que no estará disponible hasta entonces.
- **Disponibilidad del cliente:** las horas en que el cliente puede estar en casa. Si no dice nada, cualquier hora del horario.
- **Horario de trabajo:** de lunes a viernes, de 17:00 a 20:30.
- **Hueco:** tiempo dentro del horario de trabajo en el que Haquin no está haciendo ninguna visita.

Las fuentes de datos y la métrica están descritas en el [README](../../README.md).
