# LIBCORE

## **Introduccion**

Libreria que contiene funciones e estructura para utilidades en cualquier proyecto go para instalacion por el paquete **database** y es windows ir a este [GoOracleWindow](https://github.com/rafael180496/tips/blob/main/GoOracleWindow.md)

### **TEST:**

Carpeta donde contiene ejemplos de las librerias creadas.

### **Utility:**

Paquete que contiene funciones e estructuras variadas como utilidad de cualquier proyecto(cast,archivo,encriptacion,comando,etc..)

### **Database:**

Paquete contiene funciones e estructura para manipulacion de base de datos

### **Server:**

Paquete para crea un api rest con ECHO mas practico un ejemplo seguir este [challengeBatchApi](https://github.com/rafael180496/challengeBatchApi).

### **Documentacion**

* [Introduccion]([https://medium.com/mindorks/create-projects-independent-of-gopath-using-go-modules-802260cdfb51])

* [GoLibModule]([https://blog.golang.org/using-go-modules])

  **Inicializar**

```bash
    go mod init github.com/{your_username}/{repo_name}
    or go test -v #para inicializar las dependencias
    go build
    ./gomod
```

### **Librerias Externas:**

**Toda libreria externa se debe de instalar para que
funcione correctamente el proyecto.**

* [fatih/color]([https://github.com/fatih/color])

* [go-ora]([https://github.com/sijms/go-ora])

* [lib/pq]([https://github.com/lib/pq])

* [go-sql-driver/mysql]([https://github.com/go-sql-driver/mysql])

* [denisenkom/go-mssqldb]([https://github.com/denisenkom/go-mssqldb])

* [jmoiron/sqlx]([github.com/jmoiron/sqlx])

* [go-ini/ini]([https://github.com/go-ini/ini])

* [labstack/echo]([https://github.com/labstack/echo])

* [satori/go.uuid]([https://github.com/satori/go.uuid])

**Envio de correo**
[Soporte de Cuente en gmail](https://www.hostinger.com.ar/tutoriales/como-usar-el-servidor-smtp-gmail-gratuito/)
**Ejemplo de Envio**

```go
    func TestEmail(t *testing.T) {
    t.Logf("Email:%s\n", "Envio")
    email := utl.StEmailAdmin{}
    email.AddUser("...@gmail.com",
        "...*", utl.SMTPURL["gmail"], utl.SMTPPORT["gmail1"])
    email.AddDest("...@thedataage.com")
    email.AddBody(utl.TypeContent["html"], "Prueba",    "Prueba")
    err := email.SendMail()
    if err != nil {
        t.Logf("Email:%s\n", err.Error())
    } else {
        t.Logf("Email:%s\n", "Terminado")
    }

}
```
