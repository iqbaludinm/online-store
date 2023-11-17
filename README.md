# Online Store API

## Table of Contents

- [Tentang](#tentang)
- [Panduan Menjalankan](#panduan-menjalankan)
- [Struktur Folder](#struktur-folder)
- [Feature](#feature)
- [Dokumentasi API](#dokumentasi-api)

---

### Tentang

Online Store API merupakan sebuah project RESTful API Online Shopp. Project ini dibuat menggunakan bahasa pemrograman Go dan menggunakan framework Fiber. Dalam implementasi pemprograman dan layout menggunakan Clean Architecture yang dipopulerkan oleh Uncle Bob. Project ini dibuat dengan menggunakan database PostgreSQL dan dideploy menggunakan Docker dan AWS.

----
### Panduan Menjalankan

Proses menjalankan project dapat dilakukan dengan 3 cara yaitu menjalankan file main secara langsung maupun menggunakan Makefile. Berikut merupakan penjelasan cara menjalankan proyek ini:
+ Cara Pertama
Untuk menjalankan secara langsung file main.go dapat menggunakan command sebagai berikut:
    ```bash
    go run ./app/main.go
    ```
+ Cara Kedua
Untuk menjalankan menggunakan Makefile dapat menggunakan command sebagai berikut:
    ```bash
    Make run
    ```
+ Cara Ketiga
Untuk menjalankan menggunakan nodemon dapat menggunakan command sebagai berikut:
    ```bash
    nodemon --exec go run ./app/main.go
    ```
    Atau dapat menggunakan Makefile yang telah menyediakan command nodemon:

    ```bash
    Make run-nodemon
    ```
    Untuk detail lengkap cara menjalankan menggunakan Makefile dapat dilihat difile Makefile.

----
### Struktur Folder


---

### Feature

- [x] Database ORM
- [x] Database Relational
- [x] Json Validation
- [x] JWT Security
- [x] Database migration
- [x] Docker Support
- [x] Open API / Swagger
- [x] Http Client
- [x] Error Handling
- [x] Logging
- [x] Cache

----

### Dokumentasi API