package main

import (
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type Args1 struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
	Num3 int `json:"num3"`
}

type Args2 struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
	Num3 int `json:"num3"`
	Num4 int `json:"num4"`
}

var (
	jsonSchema1 = jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"num1": {
				Type:        jsonschema.Integer,
				Description: "The first number",
			},
			"num2": {
				Type:        jsonschema.Integer,
				Description: "The second number",
			},
			"num3": {
				Type:        jsonschema.Integer,
				Description: "The third number",
			},
		},
	}

	myFuncDef1 = openai.FunctionDefinition{
		Name:        "three_number_multiplier",
		Description: "Multiplies three numbers together",
		Parameters:  &jsonSchema1,
	}

	MyTool1 = openai.Tool{
		Type:     openai.ToolType("function"),
		Function: &myFuncDef1,
	}

	jsonSchema2 = jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"num1": {
				Type:        jsonschema.Integer,
				Description: "The first number",
			},
			"num2": {
				Type:        jsonschema.Integer,
				Description: "The second number",
			},
			"num3": {
				Type:        jsonschema.Integer,
				Description: "The third number",
			},
			"num4": {
				Type:        jsonschema.Integer,
				Description: "The third number",
			},
		},
	}

	myFuncDef2 = openai.FunctionDefinition{
		Name:        "four_number_multiplier",
		Description: "Multiplies four numbers together",
		Parameters:  &jsonSchema2,
	}

	MyTool2 = openai.Tool{
		Type:     openai.ToolType("function"),
		Function: &myFuncDef2,
	}
)
