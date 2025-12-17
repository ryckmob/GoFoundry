package app

import (
	"fmt"

	"github.com/ryckmob/GoFoundry/internal/common"
)

func modelTemplate(name string) string {
	return fmt.Sprintf(`package %s

type %s struct {
	ID uint `+"`gorm:\"primaryKey\"`"+`
}
`, name, common.Capitalize(name))
}
