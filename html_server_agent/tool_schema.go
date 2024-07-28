package server

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type ArgsHtml struct {
	HtmlCode string `json:"htmlcode"`
}

var (
	HtmljsonSchema = jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"htmlcode": {
				Type:        jsonschema.String,
				Description: "The new HTML code of the website",
			},
		},
	}

	HtmlEditFuncDef = openai.FunctionDefinition{
		Name:        "html_edit_func",
		Description: "Replaces the HTML code of the website",
		Parameters:  &HtmljsonSchema,
	}

	HtmlEdit = openai.Tool{
		Type:     openai.ToolType("function"),
		Function: &HtmlEditFuncDef,
	}
)
