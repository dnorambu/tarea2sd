# Tarea 2 Sistemas distribuidos

Tenemos que entregar una red de bibliotecas que comparten libros divididos en chunks. 
Lo más relevante es manejar adecuadamente la concurrencia, así que ojo con el 
acceso a zonas críticas. Aprender a usar *lock*!

### Links de interés
1. [Libros públicos](https://www.elejandria.com/coleccion/libros-llevados-al-cine)
2. [Chunkear un libro](https://www.socketloop.com/tutorials/golang-recombine-chunked-files-example)

### Discusion 26/11/2020
1. Cuando el clienteUploader manda chunks al DataNode, entonces los chunks quedan en memoria ram del
DN. Cuando la propuesta es aceptada, se mandan los chunks a los DN correspondientes y recién ahí, se 
almacenan en disco (físico).

## Lista de pendientes
1. Cuando 2 o más clientes quieran mandar sus chunks al mismo DN, tenemos que manejar la concurrencia
para evitar que se sobreescriba el slice donde almacenamos en memoria los Chunks de cada cliente.
