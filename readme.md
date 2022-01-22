# ¿Como dockerizar una app en golang?

Esta docu esta basado en este [tutorial](https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e) respecto a la parte de montar una imagen de golang (la construccion del docker-file)

Luego viene la parte del `docker-compose.yml` el cual utiliza el dockerfile creado para hacer la build de la app y tambien se agrega el uso de una base de datos postgres

## Como correr

Manda un 
`docker-compose up --build`
con ello se inicia buildeando desde cero, luego puedes omitir el ``-build``

## Como probar

Dale al endpoint /pingDB y si te devuelve un 200 es por que pudo establecer conexion exitosa con la db :)