package tfvars

import (
	"bytes"
	"fmt"

	"github.com/ollliegits/terraform-docs/internal/pkg/doc"
	"github.com/ollliegits/terraform-docs/internal/pkg/print"
	"github.com/ollliegits/terraform-docs/internal/pkg/settings"
)

// Print prints a terraform.tfvars template for a variables.tfvars file.
func Print(document *doc.Doc, settings settings.Settings) (string, error) {
	var buffer bytes.Buffer

	if document.HasInputs() {
		if settings.Has(print.WithSortByName) {
			if settings.Has(print.WithSortInputsByRequired) {
				doc.SortInputsByRequired(document.Inputs)
			} else {
				doc.SortInputsByName(document.Inputs)
			}
		}

		printInputs(&buffer, document.Inputs, settings)
	}

	return buffer.String(), nil
}

func getInputDefaultValue(input *doc.Input, settings settings.Settings) string {
	var result = ""

	if input.HasDefault() {
		result = print.GetPrintableValue(input.Default, settings, false)
	}

	return result
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, settings settings.Settings) {
	buffer.WriteString("\n")

	for _, input := range inputs {
		var format = "%s=\"%s\"\n# %s\n\n"
		if input.HasDefault() {
			format = "# %s=%s\n# %s\n\n"
		}
		buffer.WriteString(
			fmt.Sprintf(
				format,
				input.Name,
				getInputDefaultValue(&input, settings),
				input.Description))
	}

	buffer.WriteString("\n")
}
