package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"math"
	"os"
	"sort"
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
	if len(r) < 2 {
		return field{}, fmt.Errorf("field entry needs >=2 elements, got %d", len(r))
	}
	name, ok := r[0].(string)
	if !ok {
		return field{}, fmt.Errorf("reading name %v", r[0])
	}

	values, ok := r[1].(map[string]any)

	if !ok {
		return field{}, fmt.Errorf("reading value %v", r[1])
	}

	// Validate Nth against int16 range before narrowing. An out-of-range
	// value here would silently wrap and put the wrong field tag into the
	// generated autogen.go — producing wrong XRPL serialization on every
	// signed transaction with no compile- or run-time error.
	nthRounded := math.Round(values["nth"].(float64)) //nolint:forcetypeassert
	if nthRounded < math.MinInt16 || nthRounded > math.MaxInt16 {
		return field{}, fmt.Errorf("nth %v out of int16 range for field %q", nthRounded, name)
	}

	//nolint:forcetypeassert
	return field{
		Name:           name,
		IsSerialized:   values["isSerialized"].(bool),
		IsSigningField: values["isSigningField"].(bool),
		IsVLEncoded:    values["isVLEncoded"].(bool),
		Nth:            int16(nthRounded),
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
		panic(fmt.Sprintf("building xrpl definitions: %v", err))
	}
	defer file.Close() //nolint:errcheck

	dc := json.NewDecoder(file)

	var df defFile

	err = dc.Decode(&df)
	if err != nil {
		panic(fmt.Sprintf("reading definitions: %v", err))
	}

	generatedFile := []byte{}
	generatedFile = fmt.Appendf(generatedFile, "// Code generated - DO NOT EDIT. \n\n")

	generatedFile = fmt.Appendf(generatedFile, "package defs\n\n")

	// Sort keys for deterministic output.
	generatedFile = fmt.Appendf(generatedFile, "const (\n")
	typeKeys := sortedKeys(df.Types)
	for _, k := range typeKeys {
		generatedFile = fmt.Appendf(generatedFile, "%s XType = %d\n", k, df.Types[k])
	}
	generatedFile = fmt.Appendf(generatedFile, ")\n\n")

	// Transaction Types
	generatedFile = fmt.Appendf(generatedFile, "var TxTypeToValue = map[string]int32{\n")
	txKeys := sortedKeys(df.TransactionTypes)
	for _, k := range txKeys {
		generatedFile = fmt.Appendf(generatedFile, "\"%s\":%d,\n", k, df.TransactionTypes[k])
	}
	generatedFile = fmt.Appendf(generatedFile, "}\n\n")

	// Fields
	generatedFile = fmt.Appendf(generatedFile, "var NameToField = map[string]Field{\n")

	for j := range df.Fields {
		f, err := parseField(df.Fields[j])
		if err != nil {
			panic(fmt.Sprintf("parsing xrpl definitions: %v", err))
		}
		generatedFile = fmt.Append(generatedFile, generateType(f))
	}
	generatedFile = fmt.Appendf(generatedFile, "}")

	// Format BEFORE opening the destination with O_TRUNC. The original
	// order zeroed autogen.go first and only then ran format.Source; any
	// formatter error left the committed file at zero bytes, breaking
	// every downstream build until the generator was successfully rerun.
	generatedFile, err = format.Source(generatedFile)
	if err != nil {
		panic(fmt.Sprintf("formatting: %v", err))
	}

	auto, err := os.OpenFile("./pkg/xrpl/defs/autogen.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(fmt.Sprintf("opening write file: %v", err))
	}
	defer auto.Close() //nolint:errcheck

	_, err = auto.Write(generatedFile)
	if err != nil {
		panic(fmt.Sprintf("writing to definitions.go file: %v", err))
	}
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
