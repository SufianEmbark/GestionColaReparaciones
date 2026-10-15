# Milestones

Cada milestone es un producto mínimamente viable que se construye sobre el anterior. Las [historias de usuario](historias-usuario.md) se trabajan mediante issues asignados a cada milestone. Los dos primeros son internos; los siguientes, hacia un producto en la nube que Haquin pueda usar, se planificarán durante el desarrollo de estos.

## Milestone 0: representación del problema en código

Código que representa los conceptos del problema y sus relaciones, sin lógica de negocio.

**Validez:** sigue una metodología de diseño dirigido por el dominio a partir de las historias de usuario.

**Historia de usuario asignada:** [HU001](historias-usuario.md#hu001-cuando-un-trabajo-se-me-bloquea-no-sé-qué-otro-trabajo-pendiente-puedo-hacer-en-ese-hueco), porque es el problema que mide la métrica y para representarla hacen falta todos los conceptos del problema.

## Milestone 1: primera lógica de negocio comprobada automáticamente

Lógica mínima sobre el milestone 0 que empieza a resolver los problemas de las historias de usuario, con tests automáticos.

**Validez:** los tests se ejecutan automáticamente, pasan y comprueban el comportamiento esperado según las historias de usuario, respetando las restricciones del README.
