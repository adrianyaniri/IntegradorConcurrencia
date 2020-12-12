/* ejemplo de lectura y escritura
en buffer implementacion en lenguaje go
*/
/* ejemplo de patron productor consumidor */
/* channels son equivalentes a los buffer */
package IntegradorConcurrencia

import "fmt"

func main() {

	direccionCanal()

}

// indica la direccion de lectura o escritura en el buffer

func direccionCanal() {
	sender := make(chan string, 1) //creacion del channel y el tipo de dato que puede resibir
	// segundo parametro indica el tamaño del channel( buffer)
	reciver := make(chan string, 1)

	send(sender, "enviando mensaje") // envia el mensaje al channel pasado en el parametro
	reciveAndSend(sender, reciver)

	mesaje := <-reciver // lee lo que recibe del channel

	fmt.Println(mesaje) // escribe por pantalla el mensaje
}

// envia mensaje al channel

func send(sender chan<- string, mensaje string) {
	sender <- mensaje

}

// funcion que resibe el mensaje y envia el mensaje resibido

func reciveAndSend(sender <-chan string, receive chan<- string) {
	mensaje := <-sender //envia el mensaje
	receive <- mensaje  // chanell que recibe el mensaje

}
