# Tarea 2 Sistemas distribuidos
## Grupo Go Diego Go
Diego Gutierrez 201521066-7
Diego Norambuena 201704002-5

## Instrucciones para Namenode
Ejecutar el NameNode
* make nn

## Instrucciones para Data Node 3
Ejecutar el Data Node 3
* make dn3

## Instrucciones para Data Node 2
Ejecutar el Data Node 2
* make dn2

## Instrucciones para Data Node 1
Ejecutar el Data Node 1
* make dn1

## Ejecutar el cliente
Ejecutar el cliente
* make cliente

## Bitacora
* Borrar el archivo "log.txt" llevará a una falla en la ejecución del sistema. **No lo haga**. Si quiere una ejecución limpia solo elimine su contenido, junto con el contenido de "libros.txt", y los chunks almacenados en los diferentes DataNodes.
* El nombre de los libros debe ir sin guiones bajos ni espacios (se dejan algunos en la carpeta books/ para que realicen las pruebas necesarias) Por defecto, solo se buscaran archivos que sean en formato PDF, por lo que solo deben ingresar el nombre del libro, sin extension.
* Cuando el clienteUploader manda chunks al DataNode aleatorio, este los almacena en memoria ram (slicedeChunks). Cuando la propuesta es aceptada, se mandan los chunks a los DN correspondientes y recién ahí, se almacenan en disco (físico).
* El formato del log es el siguiente: 
    NombreLibro NumPartes
    nombreLibro_parte_1 ipMaq
    . . .
    nombreLibro_parte_n ipMaq
    NombreLibro2 NumPartes
    nombreLibro2_parte_1 ipMaq
    . . .
    nombreLibro2_parte_n ipMaq