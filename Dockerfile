# Menggunakan image golang sebagai base image
FROM golang:1.16-alpine

# Menentukan working directory
WORKDIR /app

# Menyalin file-file proyek ke dalam container
COPY . .

# Menginstal dependensi dan build aplikasi
RUN go mod download
RUN go build -o main

# Menjalankan aplikasi saat container dijalankan
CMD ["./main"]
