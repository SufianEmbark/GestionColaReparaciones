# Milestones

Cada milestone es un producto mínimamente viable que se construye sobre el anterior. Los dos avanzan en el problema de la [HU001](historias-usuario.md#hu001-cuando-acepto-un-trabajo-no-sé-qué-día-y-a-qué-hora-darle-cita-al-cliente), que es la base de las otras dos historias, y la HU001 pasará de uno a otro según se avance. Los dos son internos; los siguientes, hacia un producto en la nube que Haquin pueda usar, se planificarán durante el desarrollo de estos.

## Milestone 0: representación de la HU001 en código

Se analiza la HU001 aplicando diseño dirigido por el dominio y lo que sale de ese análisis se programa en el lenguaje elegido para el proyecto.

**Validez:** quien revisa el pull request comprueba que el código recoge lo que plantea la HU001 sin añadir nada que no salga de ella, y que con él se puede empezar a programar la solución de la HU001.

**Historia de usuario asignada:** [HU001](historias-usuario.md#hu001-cuando-acepto-un-trabajo-no-sé-qué-día-y-a-qué-hora-darle-cita-al-cliente).

## Milestone 1: primera solución de la HU001 comprobada automáticamente

Sobre el código del milestone 0, la lógica mínima que empieza a resolver el problema de la HU001, con tests que se ejecutan automáticamente.

**Validez:** los tests se ejecutan automáticamente, pasan y comprueban que el reparto de citas respeta las condiciones de la HU001 y que, entre dos repartos posibles, se queda con el mejor según el criterio de la propia HU001.
