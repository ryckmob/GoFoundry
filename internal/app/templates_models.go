package app

import (
	"fmt"
	"strings"

	"github.com/ryckmob/GoFoundry/internal/common"
)

func modelTemplate(name string, fields []Field) string {
	var body strings.Builder

	body.WriteString("\tID uint\n")

	if len(fields) == 0 {
		body.WriteString("\tNome string\n")
		body.WriteString("\tAtivo bool\n")
	} else {
		for _, f := range fields {
			body.WriteString(fmt.Sprintf("\t%s %s\n", f.Name, f.Type))
		}
	}

	return fmt.Sprintf(`package %s

type %sModel struct {
%s}
`, name, common.Capitalize(name), body.String())
}
