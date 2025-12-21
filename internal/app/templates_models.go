package app

import (
	"fmt"

	"github.com/ryckmob/GoFoundry/internal/common"
)

func modelTemplate(name string) string {
	return fmt.Sprintf(`package %s

type %sModel struct {
	ID     uint
	Teste1 string
	Teste2 int
	Teste3 float64
	//altere os campos
}
`, name, common.Capitalize(name))
}
