package main

import (
	"encoding/json"
	"fmt"

	"github.com/marcovargas74/cursoDeGo/ERRO/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Lucy"
	casa.X = 17
	casa.Y = 29
	if err := casa.SetValor(100); err != nil {
		fmt.Println("[main] Houve um erro na atribuicao de valor:", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor, por favor verigfique sua avaliação")
		}
		return
	}

	fmt.Printf("A Casa é: %+v\r\n", casa)
	fmt.Printf("O Valor da Casa é: %d\r\n", casa.GetValor())

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geracaodo objeto JSON:", err.Error())
		return
	}
	fmt.Printf("A Casa em JSON: %s\r\n", string(objJSON))
}
