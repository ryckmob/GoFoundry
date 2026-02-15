package app

import (
	"fmt"
	"strings"

	"github.com/ryckmob/GoFoundry/internal/common"
)

func modelTemplate(name string, fields []Field) string {
	var body strings.Builder
	var imports []string
	usesTime := false

	for _, f := range fields {
		if f.Type == "time.Time" {
			usesTime = true
			break
		}
	}

	if usesTime {
		imports = append(imports, `"time"`)
	}

	body.WriteString("\tID uint\n")

	if len(fields) == 0 {
		body.WriteString("\tNome string\n")
		body.WriteString("\tAtivo bool\n")
	} else {
		for _, f := range fields {
			body.WriteString(fmt.Sprintf("\t%s %s\n", f.Name, f.Type))
		}
	}

	var importBlock string
	if len(imports) > 0 {
		importBlock = "import (\n"
		for _, imp := range imports {
			importBlock += "\t" + imp + "\n"
		}
		importBlock += ")\n\n"
	}

	return fmt.Sprintf(`package %s

%s
type %sModel struct {
%s}
`, name, importBlock, common.Capitalize(name), body.String())
}
