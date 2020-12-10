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

-Paralelismo real: En arquitectura de multinucleos, se podrian asignar una tarea para cada nucleo y asi podrian trabajar las tareas en
simultaneo sin problemas. En arquitectura mononucleos,se podria tambien obtener concurrencia utilizando multiples procesadores. El sistema opretivo
se encargar que los distintos procesos que se esten ejecutando, puedan colarborar entre si y comunicarse entre ellos.

-Paralelismo virtual: Utiliza la tecnica de multiplexacion que permite que ejecucion de varios procesos en un unico procesador, de manera que parece
que se estan ejecutando todos los procesos a la vez, aunque en el procesador solo de este ejecutando un proceso en cada instante de tiempo.
Aunque haya solo un solo procesador esto no es niguna retrincion para que se puedan ejecutar programas concurrentes

-Permiten compartir recursos fisicos y logicos del sistema

-Permiten acelerar los calculos: divide las tareas en subtareas y las ejecuta en paralelo

_ Modulalidad : Se puede construir un sistema modular, dividiendo las funciones del sistema en procesos separados





Desventajas:
-Condicion de carrera: Esto se da cuando uno o varios procesos quieren acceder a un mismo recurso compartido sin control.

-Aplazamiento: Consiste en que uno o varios procesos no tenga suficiente uso del CPU o de otro recurso, para terminar su tarea

-Condicion de espera circula: Esto ocurre cuando un proceso necesita de un recurso que esta siendo utilizado por otro y este otro necesite
que el primero libre un recurso, creado asi una cola de espera

-Condicion de espera ocupada: Cuando un proceso pide la utilizacion de un recurso que esta siendo utilizado por otro, consumiendo asi su tiempo 
chequeando si el recuro solicitado fue liberado

-Exclusion mutua: La condicion de exclusion mutua establece que solo un proceso puede acceder a la seccion critica ( variable, espacio de memoria, recurso
compartido). Para evitar esto se utilizan alguna tecnica para evitar la entrar en la seccion como la implementacion de semaforos, monitores, mutex




Bibliografia consultada: "Sitemas Operativos" de William Stanllings

