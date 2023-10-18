package main

import (
	"context"
	"fmt"
	"time"
)

//Ex: vamos tentar fazer uma reserva em um hotel e se bater um determinado tempo cancelamos a tentativa

func main() {
	//contexto em branco -> iniciamos ele
	ctx := context.Background()

	//inserindo regra de cancelamento no contexto -> 3 segundos
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)

	//boa prática
	defer cancel()
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	//o select é extremamente útil, é quase um case, porém, trabalha de forma assíncrona
	//ele fica aguardando o resultado e quando o resultado chega ele toma a ação
	select {
	case <-ctx.Done(): //se o contexto estiver finalizado a tentativa de reserva será cancelada
		fmt.Println("Hotel booking cancelled. Tomeout reached.")
		return
	case <-time.After(1 * time.Second): //caso passe 1 segundo e nosso contexto não foi cancelado nossa reserva será concluída
		//se essa operação ultrapassar os 3 segundos do contexto a reserva nunca será concluída
		fmt.Println("Hotel bocked.")
	}
}
