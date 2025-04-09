package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"math"
	"os"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type defFile struct {
	Fields           [][]any          `json:"FIELDS"`
	Types            map[string]int16 `json:"TYPES"`
	TransactionTypes map[string]int32 `json:"TRANSACTION_TYPES"`
}

type field struct {
	Name           string
	IsSerialized   bool   `json:"isSerialized"`
	IsSigningField bool   `json:"isSigningField"`
	IsVLEncoded    bool   `json:"isVLEncoded"`
	Nth            int16  `json:"nth"`
	Type           string `json:"type"`
}

func parseField(r []any) (field, error) {
	name, ok := r[0].(string)
	if !ok {
		return field{}, fmt.Errorf("reading name %v", r[0])
	}

	values, ok := r[1].(map[string]any)

	if !ok {
		return field{}, fmt.Errorf("reading value %v", r[1])
	}

	return field{
		Name:           name,
		IsSerialized:   values["isSerialized"].(bool),
		IsSigningField: values["isSigningField"].(bool),
		IsVLEncoded:    values["isVLEncoded"].(bool),
		Nth:            int16(math.Round(values["nth"].(float64))),
		Type:           values["type"].(string),
	}, nil

}

func generateType(f field) string {
	return fmt.Sprintf(`"%s": {
        IsSerialized: %t,
	    IsSigningField: %t,
       	IsVLEncoded: %t,
        Nth: %d,
        Type: %s,
    },
`,
		f.Name,
		f.IsSerialized,
		f.IsSigningField,
		f.IsVLEncoded,
		f.Nth,
		f.Type,
	)
}

func main() {
	file, err := os.Open("./pkg/xrpl/defs/main/definitions.json")

	if err != nil {
		fmt.Printf("building xrpl definitions: %v", err)
		return
	}

	dc := json.NewDecoder(file)

	var df defFile

	err = dc.Decode(&df)
	if err != nil {
		logger.Errorf("reading definitions: %v", err)
		return
	}

	generatedFile := []byte{}
	generatedFile = fmt.Appendf(generatedFile, "// Code generated - DO NOT EDIT. \n\n")

	generatedFile = fmt.Appendf(generatedFile, "package defs\n\n")

	// Types
	generatedFile = fmt.Appendf(generatedFile, "const (\n")
	for k, v := range df.Types {
		generatedFile = fmt.Appendf(generatedFile, "%s XType = %d\n", k, v)
	}
	generatedFile = fmt.Appendf(generatedFile, ")\n\n")

	// Transaction Types
	generatedFile = fmt.Appendf(generatedFile, "var TxTypeToValue = map[string]int32{\n")
	for k, v := range df.TransactionTypes {
		generatedFile = fmt.Appendf(generatedFile, "\"%s\":%d,\n", k, v)
	}
	generatedFile = fmt.Appendf(generatedFile, "}\n\n")

	// Fields
	generatedFile = fmt.Appendf(generatedFile, "var NameToField = map[string]Field{\n")

	for j := range df.Fields {
		f, err := parseField(df.Fields[j])
		if err != nil {
			fmt.Printf("parsing xrpl definitions: %v", err)
			return
		}
		generatedFile = fmt.Append(generatedFile, generateType(f))
	}
	generatedFile = fmt.Appendf(generatedFile, "}")

	auto, err := os.OpenFile("./pkg/xrpl/defs/autogen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("opening write file: %v", err)
		return
	}

	// format
	generatedFile, err = format.Source(generatedFile)
	if err != nil {
		fmt.Printf("formatting: %v", err)
		return
	}

	_, err = auto.Write(generatedFile)
	if err != nil {
		logger.Errorf("writing to definitions.go file: %v", err)
		return
	}

	defer auto.Close()

}
