# GestionColaReparaciones

## Contexto

Mi padre tiene un trabajo de mañana y por la tarde trabaja de autónomo como técnico de reparaciones y mantenimiento del hogar. Apunta los trabajos pendientes en notas y WhatsApp, y lleva el orden de memoria, atendiéndolos por orden de llegada.

Cada trabajo empieza con una visita en la que intenta arreglarlo en el momento. A veces no puede terminarlo: le falta un material que no tiene en ese momento (la tienda está cerrada, no le da tiempo a ir y volver...) o el trabajo es más largo de lo que se puede hacer en una tarde. Entonces el trabajo queda bloqueado hasta otro día. Mientras tanto, al seguir el orden de llegada, deja huecos de su tarde sin trabajar aunque tenga otros trabajos pendientes que sí podría hacer.

## Problema

**Asignar a cada trabajo pendiente un día y una hora de visita, dentro del horario disponible, de forma que no queden horas sin trabajar mientras haya trabajos que se podrían estar haciendo.**

### Qué se sabe

- **Al recibir la llamada:** el tipo de trabajo, qué hay que hacer, cuándo llamó el cliente, su urgencia (urgente, normal o puede esperar), una duración estimada del trabajo y la disponibilidad del cliente (si no dice nada, puede a cualquier hora del horario). Si ya sabe que necesitará un material que no tiene, también sabe desde qué fecha podrá tenerlo.

- **Tras cada visita:** si el trabajo quedó terminado o no. Si no, por qué (falta material o falta tiempo), cuánto tiempo le queda y, si falta material, desde qué fecha podrá tenerlo.

La fecha del material siempre se conoce en el momento (en la llamada o en la visita), así que no hay incertidumbre sobre cuándo llega. Que un trabajo necesite material no es un problema por sí mismo: solo importa si en el momento en que tocaría hacerlo todavía no lo tiene.

### Qué hay que obtener

Para cada trabajo pendiente, **la fecha y hora de su próxima visita**.

### Restricciones

- **Horario:** de lunes a viernes, de 17:00 a 20:30. Es el horario que se planifica y el único que cuenta para la métrica. A veces una visita se alarga más allá de las 20:30, y los fines de semana no se planifican porque son muy irregulares; en ambos casos la visita se registra igual.
- **Duración de una visita:** habitualmente no más de 3 horas. Un trabajo más largo se reparte en varias visitas en días distintos.
- **Material:** un trabajo que necesita material no puede visitarse antes de la fecha en que el material está disponible.
- **Cliente:** la visita tiene que caer dentro de la disponibilidad que ha indicado el cliente.
- **Sin solapamientos:** dos visitas no pueden coincidir en el tiempo.
- **Prioridad:** los trabajos se atienden por orden de llamada, pero los urgentes van antes que los normales, y estos antes que los que pueden esperar.

Al planificar con fecha y hora ya no es necesario limitarse a aceptar solo trabajos que se puedan empezar en una semana. Ese límite existía para poder llevarlo todo de cabeza, así que **los rechazos dejan de formar parte del problema**.

## Métrica

Se divide el horario (L-V, 17:00 a 20:30) en bloques de media hora. Un bloque es **tiempo muerto aprovechable** si se cumplen las dos condiciones:

1. No se estaba haciendo ninguna visita.
2. Había al menos un trabajo **que se podía hacer**: ya solicitado, sin terminar, con el material ya disponible si lo necesita y con el cliente disponible en ese bloque.

**La métrica son las horas muertas aprovechables por semana (cada bloque cuenta 0,5 h), y se busca minimizarlas.** Una solución es mejor que otra si deja menos horas muertas aprovechables con los mismos trabajos.

Las horas muertas en las que no había nada que se pudiera hacer (no hay trabajos, todos están bloqueados o los clientes no pueden) no cuentan, porque no dependen de la planificación.

La situación actual, siguiendo el orden de llegada, es la línea base con la que se comparará.

El ideal es 0: ninguna hora muerta mientras haya algún trabajo que se pueda hacer.

## Datos

Los datos los apunta mi padre (o yo con la información que me da) en una hoja de Google Sheets con dos hojas.

https://docs.google.com/spreadsheets/d/1TZXDSEyXSjovz1fJbzHfa1Fm0bqNKjzyctNGtZC7k9I/edit?usp=sharing

### Trabajos (una fila por trabajo, al recibir la llamada)

| ID | Tipo | Qué hay que hacer | Llamó (fecha y hora) | Urgencia | Duración estimada (h) | Disponibilidad del cliente | Material disponible desde |
| -- | ---- | ----------------- | -------------------- | -------- | --------------------- | -------------------------- | ------------------------- |

- **Tipo:** aires acondicionados, fontanería, electricidad, cámaras frigoríficas, pintura, comunidades, máquinas u otros.
- **Urgencia:** urgente, normal o puede esperar.
- **Material disponible desde:** solo si en la llamada ya sabe que necesitará un material que no tiene. Vacío si no lo necesita o ya lo tiene.

### Visitas (una fila cada vez que va a un trabajo)

| Trabajo | Inicio (fecha y hora) | Fin (fecha y hora) | Resultado | Duración restante (h) | Material disponible desde |
| ------- | --------------------- | ------------------ | --------- | --------------------- | ------------------------- |

- **Resultado:** *Terminado*, *Falta material* o *Falta tiempo*.
- **Duración restante:** solo cuando el resultado no es *Terminado*; cuánto le queda al trabajo.
- **Material disponible desde:** solo cuando el resultado es *Falta material*.

Un trabajo que se bloquea tiene varias visitas: la primera con *Falta material* o *Falta tiempo* y la última con *Terminado*. A partir de estas dos hojas se deduce todo lo demás: cuándo empezó y terminó cada trabajo, si se aplazó, si necesitó material, en qué momentos había trabajos disponibles y, por tanto, la métrica.

## Planificación

La planificación del proyecto está en los siguientes documentos:

- [User journeys](docs/planificacion/user-journeys.md)
- [Personas](docs/planificacion/personas.md)
- [Historias de usuario](docs/planificacion/historias-usuario.md)
- [Milestones](docs/planificacion/milestones.md)

## Configuración del repositorio: [Ver documentación de configuración](docs/configuracion/configuracion.md)

## Instrucciones para los revisores

El objetivo 0 se aprobó con condiciones. Este es el mensaje del profesor:

> Te voy a ser honesto: si hubieras enviado esto antes de ayer, lo habría rechazado. Creo que le falta todavía concretar exactamente cuál es el problema y como se sabe que está solucionado. Pero para ser también honesto, no nos va a dar tiempo ni a ti ni a mi para solucionarlo y corregirlo para mañana.
>
> Creo que todo lo que hay se puede trabajar en el objetivo 1, pero *tienes que trabajarlo en serio* y solucionar todos estos problemas o te arriesgas a no pasar este objetivo. Estaré atento a lo que hagan los revisores, y te agradecería que incluyeras en algún sitio como instrucciones para los mismos este mensaje y estos cambios.
>
> Por lo demás, lo doy por aprobado.

Cambios realizados a partir de sus comentarios:

| Comentario | Cambio |
| ---------- | ------ |
| Concretar exactamente cuál es el problema y cómo se sabe que está solucionado | Nueva sección **Problema**, con qué se sabe, qué hay que obtener (fecha y hora de cada visita) y las restricciones. En **Métrica** se indica que el ideal es 0 horas muertas aprovechables y que se compara con la situación actual (orden de llegada). |
| Métrica explícita, de lo específico a lo general | Nueva sección **Métrica**: horas muertas aprovechables por semana, a minimizar. No se usan los trabajos terminados por semana porque dependen de cuántos clientes llamen, ni el dinero porque depende de si llegan trabajos caros o baratos; ninguna de las dos mide lo bien que se planifica. |
| Evitar datos que se puedan deducir unos de otros | Una tabla por trabajo pasa a ser dos: **Trabajos** y **Visitas**. Se eliminan «¿Necesita material?», «Fecha que empezó/terminó» y «¿Se aplazó a mitad?», porque se deducen de las visitas, y «Fecha prevista/que tuvo el material» se unen en una. |
| ¿El material se puede deducir de lo que hay que hacer? | No: depende de lo que encuentre (un aire puede fallar por la electrónica, el compresor o el gas). Por eso se apunta cuando se sabe, en la llamada o en la visita. |
| ¿Cuándo se puede saber la fecha prevista del material? ¿Hay que introducir incertidumbre? | Se sabe en el momento en que se descubre que hace falta, así que no hay incertidumbre. La fecha prevista y la real se unen en «Material disponible desde». |
| ¿Hay que minimizar la diferencia entre la llamada y el inicio? | No directamente: se minimizan las horas muertas aprovechables, y que el cliente espere menos es una consecuencia. |
| ¿Inicio y fin deben ser el mismo día si se tiene el material? | Lo ideal es terminar en una visita, pero si se bloquea no se puede evitar. Lo que la planificación sí evita es el tiempo perdido mientras tanto, y para planificar la siguiente visita se apunta la duración restante. |
| ¿El aplazamiento no está implícito en las fechas? | Sí: se elimina la columna y se deduce de que un trabajo tenga más de una visita. |
| ¿Qué aportan los trabajos rechazados? | Nada al problema: se elimina la tabla de rechazados. |
| ¿Cómo no rechazar si la semana está llena? Reservar más allá de una semana | El límite de una semana era para llevarlo de memoria; al planificar con fecha y hora se puede reservar más allá, así que los rechazos dejan de ser parte del problema. |
