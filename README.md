## Contenido

- [Requerimientos](#requerimientos)
- [Estructura de proyecto](#estructura-de-proyecto)
- [Configuracion](#configuracion)
- [Explicacion y mejoras](#explicacion-y-mejoras)


## Requerimientos

- golang 1.23.10
- Docker y Docker Compose

## Estructura de proyecto

```plaintext
root/
├── api/
│   ├── handler/  
│   │   ├── auth.go  # controlador de autenticación
│   │   └── user.go  # controlador de usuarios   
│   ├── middleware/                
│   │  └── auth.go   # middleware para proteccon de rutas          
│   ├── models/
│   │   └── user.go  # modelo de usuario y validaciones
│   ├── repository/
│   │    ├── postgres_repository.go  # interfaz de repositorio para postgres
│   │    └── user_repository.go  # interfaz de repositorio de usuarios
│   ├── router/
│   │   └── router.go  # configuración de rutas
│   └── service/
│       ├── auth_service.go  # servicio de autenticacion
│       └── user_service.go  # servicio de usuarios
├── docs/ # documentacion de swagger
├── docker-compose.yaml       # orquestacion de contenedores
├── Dockerfile                # imagen de contenedor
├── go.mod                    # dependencias de golang
├── go.sum                    # dependencias de golang
├── init.sql                  # script de inicialización de la base de datos y creacion del usuario admin
├── main.go                   # punto de entrada de la aplicacion
└── README.md                 # documentacion   

```


## Configuracion

1. Clonar repositorio.
   ```bash
   git clone https://github.com/Estebanfonseca/prueba_tecnica_BRM.git
   ```
2. Build imagen de docker.
   ```bash
   docker-compose up -d  --build
   ```
3. la aplicacion estara disponible
   ```bash
    http:localhost:8000/
    ```
4. una ves dentro vera la documentacion de swagger

## Explicacion y mejoras

*** Explicacion ***
- el proyecto esta estructurado en capas, cada capa tiene una responsabilidad bien definida.
- el controlador maneja las peticiones entrantes y las envia a los servicios.
- los servicios son responsables de la logica de negocio.
- los modelos manejan la validacion de datos.
- los repositorios manejan la comunicacion con la base de datos.
- el middleware maneja la autenticacion.
- la documentacion de swagger esta disponible en el endpoint /
- la aplicacion utiliza postgres como base de datos.
- la aplicacion utiliza docker para la orquestacion de contenedores.
- la aplicacion utiliza golang como lenguaje de programacion.
- la aplicacion usa el patron de diseño de repositorio para la comunicacion con la base de datos lo que permite cambiar de base de datos sin afectar la aplicacion 
- decidi usar este patron de diseño para la aplicacion ya que permite escalar mas facil por su modularidad ya que he trabajado con nestjs me siento comodo al usar este patron de diseño.

*** Mejoras ***
- Agregar roles a los usuarios para tener mas control de las acciones sensibles como la creacion  y eliminacion de usuarios
- Agregar Oath2 para tener mas seguridad en la autenticacion
- Agregar un sistema de logs mas robustos para tener un mejor control de los errores
