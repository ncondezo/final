# Clinica API

Esta API es la encargada de manera administrar turnos medicos a pacientes.

## Environment Variables

Para correr este proyecto deberás agregar las siguientes variables de entorno de la Api en un archivo .env en la raiz del proyecto.

`MYSQL_ROOT_PASSWORD`: Password root para base de datos MYSQL.

`MYSQL_DATABASE`: Nombre de la base de datos de la aplicacion.

`MYSQL_USER`: Username para acceder a la base de datos.

`MYSQL_PASSWORD`: Password de base de datos.

`TOKEN_SECRET_KEY`: Secret Key para funcionamiento del token (cualquier valor en formato String por ej: "MysecretKey")


## Seteo de ambiente

Clonar el proyecto

```bash
  git clone git@github.com:ncondezo/final.git
```

Moverse a la carpeta root

```bash
  cd final
```

Abrir en visual studio el proyecto
```bash
  code .
```

**Ejecución con visual studio**

    1. Abrir la solución y en el proyecto FINAL abrir crear un archivo .env en la raiz del proyecto agregando las variables de entorno descriptas arriba y cambiando sus valores de preferencia.
  
    2. Correr el siguiente comando para iniciar docker compose con mysql

```bash
  docker compose up 
```

**Ejecución de la api en Go desde terminal**

    Posicionarse en la raiz de la aplicacion y ejecutar el siguiente comando para correr la Api.

```bash
  go run ./cmd/server/main.go
```





