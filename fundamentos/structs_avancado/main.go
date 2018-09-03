package main

import (
	"encoding/json"
	"fmt"

	"github.com/vinipis/aulas-jeff-go/fundamentos/structs_avancado/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = "18"
	casa.Y = "25"
	casa.Valor = "60000"
	//casa.SetValor(60000)
	fmt.Printf("Casa é: %+v\r\n", casa)
	//fmt.Printf("O valor da Casa é: %d\r\n", casa.GetValor())

	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON:", string(objJSON))
}
