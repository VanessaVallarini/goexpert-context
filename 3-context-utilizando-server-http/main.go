package main

//como executar esse projeto:
//para saída Request processada com sucesso: curl localhost:8080 e aguarda 5 segundos
//para saída Request cancelada pelo cliente: curl localhost:8080 e em menos de 5 segundos executa no terminal ctrl+C

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

//vamos ter 5 segundos para executar a requisição
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		//imprime no command line stdout
		log.Println("Request processada com sucesso")
		//imprime no browser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		//imprime no command line stdout
		log.Println("Request cancelada pelo cliente")
	}
}
