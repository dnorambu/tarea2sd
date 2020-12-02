# Tarea 2 Sistemas distribuidos

Tenemos que entregar una red de bibliotecas que comparten libros divididos en chunks. 
Lo más relevante es manejar adecuadamente la concurrencia, así que ojo con el 
acceso a zonas críticas. Aprender a usar *lock*!

### Links de interés
1. [Libros públicos](https://www.elejandria.com/coleccion/libros-llevados-al-cine)
2. [Chunkear un libro](https://www.socketloop.com/tutorials/golang-recombine-chunked-files-example)

### Bitácora
1. Cuando el clienteUploader manda chunks al DataNode, entonces los chunks quedan en memoria ram del
DN. Cuando la propuesta es aceptada, se mandan los chunks a los DN correspondientes y recién ahí, se 
almacenan en disco (físico).
2. Una vez que el clienteUploader termina de enviar los chunks al DataNode, este esperará a que termine
el proceso de creación de propuesta, su aceptación y posterior distribución de los chunks. (Esto es propio
de la implementación, no hay un criterio en la rúbrica (si es que el cliente debe esperar o no))
3. Ojo donde se estan corriendo los scripts. Para los DN tienes que estar dentro del directorio, o sea,
NO lo corras en la carpeta principal (tarea2sd), si no desde la carpeta *datanodex/*
4. El nombre de los libros tienen que ir sin espacios, sin guiones bajos y en formato .pdf
5. No me borre el log o exploto.
6. Una vez ejecutado el NameNode, este no debe botarse / caerse ¿por qué? Porque guarda en memoria RAM la lista
de libros
7. SUPUESTO: una vez que los libros se han subido, estos no se borrarán del sistema.
8. Usamos el orden de prioridad 3-2-1 (es decir, 3 es más prioritario que 2, y 2 es más prioritario que 1). No usamos
ID's como tal para los nodos, pero la ejecución del programa es consistente con el sistema de prioridad aquí
mencionado. (quizás esto varíe)
9. (Para distribuidos) la logica consiste en la siguiente: el DN que elabora la propuesta inicial, la envia a los otros 2 DataNode. Estos 2 datanode revisan independientemente si estan vivos. En caso de que ambos esten vivos, envian el OK al primer DN (que les mando la propuesta). En caso contrario, el Dn tendría que elaborar una nueva propuesta sin considerar al(los) nodo caido(s). Ej: DN1 hace una propuesta P, la envia a DN2 y DN3. DN2 verifica que DN3 está vivo si solo si P tiene chunks dirigidos a DN3; DN3 verifica que DN2 esta vivo si solo si P tiene chunks dirigidos a DN2. Si no hay problemas (todos los DN están vivos) entonces se envia el OK. En caso contrario, DN1 elabora la nueva propuesta, la cual no tendrá chunks dirigidos al (los) nodo caido(s). En el peor de los casos, DN1 se queda con los chunks.
10. Enviar chunks antes que escribir en LOG. Nuestra implementación considera la siguiente forma de pensar: es mejor mandar los chunks a los data nodes y luego escribir esta acción en el LOG, porque ¿qué pasaría si se escribe en el LOG pero no se mandan los chunks? (lo escribí con tuto, mejorar)
11. Ocupamos una chance 80% de aprobar la propuesta en el caso Distribuido
## Lista de pendientes
1. Cuando 2 o más clientes quieran mandar sus chunks al mismo DN, tenemos que manejar la concurrencia
para evitar que se sobreescriba el slice donde almacenamos en memoria los Chunks de cada cliente. (listo)
2. Dejamos pendiente el tema del time.Sleep para los ayudantes.