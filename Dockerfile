# Usa una imagen base de Go
FROM golang:1.20

# Establece el directorio de trabajo en /app
WORKDIR /app

# Copia los archivos de la aplicación al contenedor
COPY . .

# Compila la aplicación
RUN go build -o main exercise5.go

# Expone el puerto 8000
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["./main"]


