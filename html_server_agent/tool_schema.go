package server

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type ArgsHtml struct {
	HtmlCode string `json:"htmlcode"`
}

// Tool function definition for editting the HTML.
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

// Tool function definition for running the server
var (
	RunServerFuncDef = openai.FunctionDefinition{
		Name:        "run_server_func",
		Description: "Run the server",
	}

	RunTheServer = openai.Tool{
		Type:     openai.ToolType("function"),
		Function: &RunServerFuncDef,
	}
)
