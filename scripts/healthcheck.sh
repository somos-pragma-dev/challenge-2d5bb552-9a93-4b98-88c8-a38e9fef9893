#!/bin/bash

set -e

# Verificar si el servicio está escuchando en el puerto especificado
nc -zv localhost $PORT

# Verificar si el servicio responde a una solicitud gRPC
grpcurl -plaintext localhost:$PORT describe