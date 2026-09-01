# Diseño y Evaluación de Microservicios en Go

En el contexto de una fintech, debes diseñar y evaluar un conjunto de microservicios que interactúan para gestionar transacciones financieras. Los servicios deben manejar solicitudes de múltiples canales (API REST, gRPC), asegurar la idempotencia de las transacciones y manejar errores específicos del dominio. Los actores involucrados son el 'originador de créditos', el'motor antifraude' y el 'core bancario'. El sistema debe procesar al menos 1 500 solicitudes por segundo en hora pico, con un SLA del 99.9%.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | diseñar microservicios en Go |
| **Nivel** | junior-l2 |
| **Tipo** | mixed |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Un IDE o editor de código.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Verifica que el proyecto arranca sin errores.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Exploración del Sistema

**Objetivo:** identificar las restricciones y ambigüedades del sistema de microservicios

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Analiza el sistema existente para identificar sus componentes y responsabilidades.
- Enumera las restricciones del sistema, como el throughput requerido y el SLA.
- Identifica ambigüedades o áreas que requieren clarificación.

**Entregable:** Documento que describe los componentes del sistema, sus responsabilidades y las restricciones identificadas.

<details>
<summary>Pistas de conocimiento</summary>

- Considera el impacto de las restricciones en el diseño del sistema.
- Piensa en cómo las ambigüedades podrían afectar la implementación.

</details>

### Fase 2: Evaluación de Decisiones de Diseño

**Objetivo:** evaluar una decisión controversial en el diseño del sistema

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Elige una decisión controversial en el diseño del sistema, como la elección entre REST y gRPC para la comunicación entre servicios.
- Evalúa los pros y contras de la decisión.
- Documenta la decisión, incluyendo contexto, fuerzas, opciones, decisión tomada y consecuencias.

**Entregable:** Registro de decisiones que documenta la evaluación de una decisión controversial en el diseño del sistema.

<details>
<summary>Pistas de conocimiento</summary>

- Considera el impacto de la decisión en la escalabilidad y la latencia del sistema.
- Piensa en cómo la decisión afecta la complejidad del código y la facilidad de mantenimiento.

</details>

### Fase 3: Comunicación de Decisiones

**Objetivo:** comunicar las decisiones tomadas a diferentes audiencias

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Prepara una presentación que comunique las decisiones tomadas en las fases anteriores a dos audiencias diferentes: el equipo técnico y la gerencia de la fintech.
- Asegúrate de que cada audiencia comprenda las decisiones y sus consecuencias.
- Valida que la presentación sea clara y efectiva para cada audiencia.

**Entregable:** Presentación que comunica las decisiones tomadas a diferentes audiencias.

<details>
<summary>Pistas de conocimiento</summary>

- Considera el lenguaje y los detalles técnicos apropiados para cada audiencia.
- Piensa en cómo presentar los beneficios y riesgos de las decisiones de manera clara y concisa.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es un microservicio y por qué se utilizan en sistemas distribuidos?
- **paraQueSirve**: ¿Para qué sirve la idempotencia en las transacciones financieras y cómo se implementa?
- **comoSeUsa**: ¿Cómo se usan los servicios de gRPC y REST en la comunicación entre microservicios?
- **erroresComunes**: ¿Cuáles son los errores comunes al diseñar microservicios y cómo se pueden evitar?
- **queDecisionesImplica**: ¿Qué decisiones implica la elección entre REST y gRPC para la comunicación entre servicios?

## Criterios de Evaluacion

- Identificación y documentación clara de las restricciones y ambigüedades del sistema.
- Evaluación exhaustiva de una decisión controversial en el diseño del sistema.
- Comunicación efectiva de las decisiones tomadas a diferentes audiencias.

---

*Reto generado automaticamente por Challenge Generator - Pragma*
