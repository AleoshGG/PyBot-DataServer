FROM golang:1.23-alpine AS builder

# Install build tools and git for go mod
RUN apk add --no-cache git build-base

# Install migrate CLI
# Using v4.18.3 as specified in your go.mod
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.3

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos de dependencias
COPY go.mod go.sum ./

# Descargamos las dependencias (Docker guardará esto en caché)
RUN go mod download

# Copiamos todo el código fuente de Go
COPY . .

# Compilamos la aplicación
# - CGO_ENABLED=0: Crea un binario estático (crucial para Alpine)
# - o api-sensors: El nombre de nuestro archivo ejecutable final
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/api-sensors .

# --- Etapa 2: Imagen Final (La "Runner") ---
# Empezamos desde una imagen base de Alpine, que es súper ligera
FROM alpine:latest

# (Opcional) Instala certificados raíz, necesario si tu app hace llamadas HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copiamos SOLAMENTE el binario compilado desde la etapa "builder"
COPY --from=builder /app/api-sensors .
# Copiamos la carpeta de migraciones
COPY ./database/migrations ./database/migrations

# --- ¡IMPORTANTE! ---
# Expone el puerto en el que tu servidor de Go escucha.
# En nuestra configuración de Nginx, dijimos que era el 8080.
# Asegúrate de que tu app de Go escuche en este puerto.
EXPOSE 8081

# Comando para ejecutar el servidor de api-sensors
CMD ["./api-sensors"]