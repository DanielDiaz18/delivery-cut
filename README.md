# rappi-cut

### concluidos:


1.	Generar una vista dentro de la aplicación que permita el registro de usuarios administradores del sitio (datos básicos como: nombre, apellidos, correo)
2.	Generar una vista que permita el registro de usuarios o clientes a la app (datos básicos).
3.	Generar una vista de LogIn para administradores.
4.	Generar una vista de LogIn para usuarios.
5.	Generar una vista que permita el registro de conductores o repartidores (por parte de un usuario administrador y nuevamente con datos básicos). Se podrá editar manualmente en la base de datos (o desde el FrontEnd según se vea conveniente) un campo en donde se guarde la posición geográfica del operador (latitud y longitud). 

---

### avanzados:

6.	Generar una vista en donde cada conductor sea visible dentro de un mapa en una posición geográfica específica (distribuirlos en el mapa, al menos 5 conductores), esta vista solo es visible para usuarios Administradores. 
7.	Cada conductor debe de tener un estatus (El estatus debe ser visible en el mapa, por ejemplo, con colores o con un popup en el pin del mapa en el conductor. En caso de tener un pedido, mostrarlo en el popup o en alguna descripción para ese conductor, esto puede ser mostrando un número de pedido):
    a.	Disponible: Disponible para recibir un pedido.
    b.	Ocupado: Ocupado y no debe recibir pedido.
    c.	Asignado: Se encuentra en trayecto a recibir y entregar el pedido (no recibe nuevos pedidos).

### avances extra:
    - Mapa mapbox con marcadores y rutas
    - websockets (se creo la conexion, faltan los eventos)  
    - configuración de redis

---

### faltantes:

8.	En esa misma vista, los administradores pueden dar por concluido un servicio/pedido (por ejemplo, al dar click en un conductor que está en estatus asignado, colocar un botón que permita dar por finalizado el servicio, simulando que el conductor ya lo concluyó).
9.	Generar una vista en donde usuarios de la aplicación puedan realizar pedidos a los conductores, un simple formulario en donde se pueda escribir lo que requiere el usuario, no es necesario tener tiendas ni objetos a elegir, pero es necesario contar con un apartado en donde puedan poner datos geográficos (número de calle, calle, colonia, estado, etc) de la posición en donde el conductor debe de ir a recoger el pedido (origen de pedido) y otro apartado para escribir la posición final de entrega del pedido (mismos campos que en la posición para recoger u origen del pedido), estos datos deben de convertirse a una posición geográfica (para esto existen apis de diferentes proveedores como Mapbox o Google). En dado caso de tener dificultades con la conversión, se puede sustituir por un mapa en donde el usuario mediante drag and drop posicione el origen y el destino de su pedido, de esta manera se ahorra la conversión a posición geográfica ya que se obtiene directamente desde el mapa y los pines drag and drop. Una vez el usuario haga click en enviar su pedido, se debe de elegir al mejor conductor tomando en cuenta el tiempo y la distancia (esto también se realiza con apis externas, ejemplo, ver api de Matrix de mapbox).
10.	Al finalizar la selección del operador, se debe de mostrar en un mapa al usuario el operador que lo va a atender, un mapa con la ruta trazada desde el origen al destino y un tiempo estimado de entrega (las apis regresan este tiempo, por ejemplo, la de Matrix de mapbox) además de los datos de su pedido. Para trazar la ruta también se requieren de Apis externas, estas regresan un arreglo de posiciones que se mandan como parámetro al mapa y este hace el trabajo de trazar la ruta.
11.	Generar una vista en donde se vea el historial de pedidos por parte del usuario, además del estatus en el que se encuentran (puede ser dentro de la misma vista para generar el pedido) y actualizar mediante una función de refrescado.
12.	Actualizar el estatus y la orden de los operadores según sea el caso (dependiendo de si está disponible, ocupado o si tiene un pedido), y mostrar eso en el mapa de los administradores (se puede utilizar una función de refrescado para que se actualice la información cada cierto tiempo).
