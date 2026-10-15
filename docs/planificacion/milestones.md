# Milestones

Cada milestone entrega un producto mínimamente viable que se construye sobre el anterior. Los milestones marcan hasta qué nivel se resuelve el problema en cada fase, no el cierre de una historia de usuario concreta: las [historias de usuario](historias-usuario.md) se trabajan mediante issues asignados a cada milestone.

Los dos primeros son internos, para el propio desarrollo. Los siguientes, que se planificarán durante el desarrollo de estos, irán hacia un producto externo desplegado en la nube que Haquin pueda usar.

## Milestone 0: representación del problema en código

**Para qué:** que el problema descrito en el [README](../../README.md) y en las historias de usuario quede expresado en el código del proyecto, de forma que el siguiente milestone pueda programar la lógica sobre él sin tener que volver a interpretar el problema.

**Qué se entrega:**
- Código en el repositorio, en el lenguaje elegido para el proyecto y siguiendo sus buenas prácticas, que represente los conceptos del problema y sus relaciones, sin lógica de negocio.
- Un documento en `docs/` que explique las decisiones tomadas al representar cada concepto.

**Cómo se entrega:** mediante un pull request a la rama principal, enlazado a los issues que lo componen.

**Qué lo hace válido:**
- Se ha seguido una metodología de diseño dirigido por el dominio, partiendo de las historias de usuario y de su sección de conceptos.
- Cada decisión del documento se puede relacionar con una historia de usuario o con un concepto.
- El código se puede comprobar con las herramientas de calidad del lenguaje sin que dé errores.

**Historia de usuario asignada:** [HU001](historias-usuario.md#hu001-cuando-un-trabajo-se-me-bloquea-no-sé-qué-otro-trabajo-pendiente-puedo-hacer-en-ese-hueco).

Se prioriza la HU001 porque es el problema que mide la métrica del proyecto (horas muertas aprovechables) y porque, para representarla, hacen falta todos los conceptos del problema: trabajos pendientes, bloqueos por material, disponibilidad del cliente, duración restante y huecos. Las otras dos historias se apoyan en esos mismos conceptos.

## Milestone 1: primera lógica de negocio comprobada automáticamente

**Para qué:** que sobre lo entregado en el milestone 0 exista la lógica mínima que empiece a resolver los problemas de las historias de usuario.

**Qué se entrega:**
- El código de esa lógica de negocio en el repositorio.
- Tests que comprueben su comportamiento, ejecutables de forma automática con el gestor de tareas del proyecto.

**Cómo se entrega:** mediante un pull request a la rama principal, enlazado a los issues que lo componen.

**Qué lo hace válido:**
- Los tests se ejecutan automáticamente y pasan.
- Los tests comprueban el comportamiento esperado según las historias de usuario que se trabajen en el milestone, respetando las restricciones descritas en el README.
