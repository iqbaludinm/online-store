# base image
# golang:1:20.1-alpine3.17
FROM golang@sha256:18da4399cedd9e383beb6b104d43aa1d48bd41167e312bb5306d72c51bd11548 as builder


# Buat direktori kerja di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum terlebih dahulu untuk mengoptimalkan caching dependencies
COPY go.mod .
COPY go.sum .

# Download dan install dependencies
RUN go mod download

# Salin seluruh kode sumber
COPY . .

# Build aplikasi
RUN go build -o main .

# Eksekusi aplikasi
CMD ["./main"]

EXPOSE 8080