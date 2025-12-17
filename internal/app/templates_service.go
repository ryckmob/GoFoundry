package app

import (
	"fmt"
)

func serviceTemplate(name string) string {
	return fmt.Sprintf(`package %s

type Carro struct {
	ID    uint   `+"`json:\"id\"`"+`
	Nome  string `+"`json:\"nome\"`"+`
	Marca string `+"`json:\"marca\"`"+`
}

// Service contém as regras do app
type Service struct{}

// Top5Carros retorna os 5 carros mais populares
func (s *Service) Top5Carros() []Carro {
	return []Carro{
		{ID: 1, Nome: "Generated", Marca: "Generated"},
	}
}

`, name)
}
