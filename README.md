# Implementación de una API REST en Go con Gin y GORM

La empresa necesita una API REST para gestionar productos en un catálogo de comercio electrónico. Los productos tienen un nombre, precio, stock y categoría. La API debe registrar productos, listarlos, actualizarlos y eliminarlos. Se deben prohibir precios negativos y nombres duplicados. El sistema debe manejar la consistencia de los datos y la idempotencia de las operaciones. Los productos se almacenan en una base de datos PostgreSQL.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | Go Gin API |
| **Nivel** | junior-l2 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Go 1.21+, VS Code o GoLand.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Ejecuta `go build ./...`. Si no hay errores, estás listo.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Registro de productos

**Objetivo:** Implementar la funcionalidad para registrar productos en la API.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- La API debe aceptar solicitudes POST para registrar productos.
- Los productos deben tener un nombre, precio, stock y categoría.
- Prohibir precios negativos y nombres duplicados.
- La operación debe ser idempotente.

**Entregable:** API que acepta solicitudes POST para registrar productos con validación de precio y nombre.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo manejar la validación de datos en la entrada.
- Piensa en cómo garantizar la idempotencia de la operación.

</details>

### Fase 2: Listado de productos

**Objetivo:** Implementar la funcionalidad para listar productos en la API.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- La API debe aceptar solicitudes GET para listar productos.
- Los productos deben mostrarse con su nombre, precio, stock y categoría.

**Entregable:** API que acepta solicitudes GET para listar productos.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo paginar los resultados si hay muchos productos.
- Piensa en cómo manejar errores de la base de datos.

</details>

### Fase 3: Actualización de productos

**Objetivo:** Implementar la funcionalidad para actualizar productos en la API.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- La API debe aceptar solicitudes PUT para actualizar productos.
- Los productos deben tener un nombre, precio, stock y categoría.
- Prohibir precios negativos y nombres duplicados.
- La operación debe ser idempotente.

**Entregable:** API que acepta solicitudes PUT para actualizar productos con validación de precio y nombre.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo manejar la validación de datos en la entrada.
- Piensa en cómo garantizar la idempotencia de la operación.

</details>

### Fase 4: Eliminación de productos

**Objetivo:** Implementar la funcionalidad para eliminar productos en la API.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- La API debe aceptar solicitudes DELETE para eliminar productos.
- La operación debe ser idempotente.

**Entregable:** API que acepta solicitudes DELETE para eliminar productos.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo manejar la eliminación de productos en la base de datos.
- Piensa en cómo garantizar la idempotencia de la operación.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es una API REST y cuáles son sus principios fundamentales?
- **paraQueSirve**: ¿Para qué sirve una API REST en el contexto de un catálogo de comercio electrónico?
- **comoSeUsa**: ¿Cómo se usa Gin y GORM para implementar una API REST en Go?
- **erroresComunes**: ¿Cuáles son los errores comunes al implementar una API REST y cómo se pueden evitar?
- **queDecisionesImplica**: ¿Qué decisiones de diseño implica la implementación de una API REST con Gin y GORM?

## Criterios de Evaluacion

- Implementar una API REST que registre productos con validación de precio y nombre.
- Implementar una API REST que liste productos.
- Implementar una API REST que actualice productos con validación de precio y nombre.
- Implementar una API REST que elimine productos.

## Como trabajar con un asistente de IA

Hay dos caminos, elegi uno:

- **AGENTS.md** (recomendado) — instrucciones nativas del repo. Abri esta carpeta con tu agente local (Claude Code, Cursor, Codex, Copilot, Gemini) y las carga solo. Sabe que archivos faltan y con que comando se verifica, y completa el scaffold escribiendo en disco.
- **PROMPT_MEJORA.md** — para copiar y pegar en un chat (claude.ai, ChatGPT). Devuelve un ZIP con el proyecto. Sirve si no tenes un agente en el IDE.

Ninguno de los dos resuelve las fases del reto: eso es tu trabajo.

## Verificacion

El proyecto esta listo para trabajar cuando este comando corre sin errores:

```bash
el comando de build o arranque canonico del stack elegido
```

---

*Reto generado automaticamente por Challenge Generator - Pragma*
