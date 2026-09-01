# Exploración del Sistema de Microservicios

## Componentes del Sistema
- Originador de Créditos
- Motor Antifraude
- Core Bancario

## Restricciones Identificadas
- Throughput requerido: 1,500 solicitudes por segundo en hora pico.
- SLA: 99.9%

## Ambigüedades
- Método de comunicación entre servicios (REST vs gRPC).