# IntegradorConcurrencia


Concurrencia en Sistemas Operativos:

El diseño de sistemas operativos modernos se basan fundamentamente en la gestion de procesos.

-Multiprogramacion: Es la gestion de varios procesos en un sistema mono-procesador
-Multiprocesamiento: Es la gestion de varios procesos dentro de un sistema multi-procesador
-Procesamiento Distribuido: Es la gestion de varios procesos que se ejecuntan en un sistema de computadores multiples y distribuidos

En este diseño de sistemas operativos, la concurrencia es una parte fundamental. La concurrencia abarca varios aspectos, la comunicacion
entre distintos procesos, la competencia o comparticipacion de recursos, la reserva del CPU para varios procesos.
La ejecucion de varios procesos en concurrente permite optimizar el uso del CPU, y evitar que este tenga tiempo ocioso.
En un sistema operativo existen tres formas en la que se pueden ejecutar procesos concurrentes:

-Multiprogramcion con un unico procesador: 
    Los procesos concurrente se ejecutan en un solo procesador. El sistema operativo se encarga de gestionar los tiempo del uso del CPU
    entre los procesos, intercalando la ejecucion de los mismo para dar la apariencia de ejecucion simultanea

- Multiprocesador: Los procesos concurrente no solo intercambia su ejecucion, sino tambien se superponen. En este caso si existe una simultanedad de los proceso. Puede haber tantos procesos ejecutando simultaneamente como procesadores tenga el sistema

-Multicomputadora : Es una maquina que tiene memoria distribuidad. Esta formada por varias computadoras completas (CPU, memoria,etc) denomidados nodos. Cada nodo esta conectado entre si por una red de interconexion. El sistema operativo se encarga de la comunicacion entre ellos mediante el paso de mensajes. En este sistema tambien puede haber procesos que se este ejecutando en simultaneo.


Ventajas:


Desventajas:




Bibliografia consultada: 

