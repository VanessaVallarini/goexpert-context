package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	//iniciando o contexto com regras de timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //se esse tempo for menor que o server a requisição será cancelada
	defer cancel()

	//preparando a request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	//executando a request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
