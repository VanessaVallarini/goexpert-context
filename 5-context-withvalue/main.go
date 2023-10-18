package main

import (
	"context"
	"fmt"
)

//utilizando o contexto para envio de dados
//normalmente são metadados
//não é uma boa prática, não sabemos se os dados realmente foram passados, se existem no contexto
//o ideal é usar parâmetros via funções
//mas, vamos supor que temos uma request e ela possui dados importantes, podemos repassar esses valores assim

func main() {
	ctx := context.Background()

	//passando um chave valor do meu token para fazer uma reserva de hotel
	ctx = context.WithValue(ctx, "token", "senha")

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
