package server

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type ArgsAppJS struct {
	AppJSCode string `json:"appjscode"`
}

// Tool function definition for editting the HTML.
var (
	AppJSjsonSchema = jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"appjscode": {
				Type:        jsonschema.String,
				Description: "The new App.js code of the website",
			},
		},
	}

	AppJSEditFuncDef = openai.FunctionDefinition{
		Name:        "app_js_edit_func",
		Description: "Replaces the App.js code of the react website with that desired",
		Parameters:  &AppJSjsonSchema,
	}

	AppJSEdit = openai.Tool{
		Type:     openai.ToolType("function"),
		Function: &AppJSEditFuncDef,
	}
)