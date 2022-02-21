# Manual Tecnico
 
 Esta practica esta compuesta por un API en golang que es consumida por una app en REACT.

### Docker Compose

Servicios: db, Backend, Frontend

### Puertos en que corre cada contenedor

Base de datos Mongo: 21017.\
Backend en golang: 8081.\
Frontend en React: 80.\

### Volumen

Para la persistencia de datos se manejaron volumen que ayudara a que cada vez que se reiniciara el contenedor la informacion no se perdiera.

### REGISTRY

Las imagenes tanto del backen como del frontend fueron subidas a dockerHub.\
Nombre de las Imagenes .\
edson2021/backend_p1_201701029.\
edson2021/frontend_p1_201701029.

### Comandos

Para ejecutar el proyecto se deben ejecutar el siguiente comando.\
docker-compose up -d --build. 

# Manual de Usuario

Una breve descripcion de como funciona la aplicacion.

## Pasos a seguir

1. Ingresar a la direccion localhost:80/.\
2. Una vez estando se mostrara la siguiente imagen.
![Test](https://user-images.githubusercontent.com/14056462/154874348-5cb161cd-797f-45b2-9745-9f26cb22a5cd.png)
.\
3. Para ingresar una operacion, presionar el boton que aparece en la imagen. Ingresar los datos correspondientes.
