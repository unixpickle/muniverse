// Command codegen generates a file called spec.go with
// specifications for each available environment.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"text/template"

	"github.com/unixpickle/essentials"
)

func main() {
	spec, err := ioutil.ReadFile("games/spec.json")
	if err != nil {
		essentials.Die(err)
	}
	var specObj []map[string]interface{}
	if err := json.Unmarshal(spec, &specObj); err != nil {
		essentials.Die(err)
	}

	out, err := os.Create("spec.go")
	if err != nil {
		essentials.Die(err)
	}
	defer out.Close()

	if err := writeSpec(out, specObj); err != nil {
		out.Close()
		os.Remove("spec.go")
		essentials.Die(err)
	}
}

func writeSpec(f *os.File, specObj []map[string]interface{}) error {
	t := template.New("template")
	t, err := t.Parse(TemplateSource)
	if err != nil {
		return err
	}
	if err := fillVariantFields(specObj); err != nil {
		return err
	}
	return t.Execute(f, specObj)
}

func fillVariantFields(spec []map[string]interface{}) error {
	byName := map[string]map[string]interface{}{}
	for _, item := range spec {
		name, ok := item["name"].(string)
		if !ok {
			return errors.New("missing or invalid name attribute")
		}
		byName[name] = item
	}
	for _, item := range spec {
		if parent, ok := item["variant_of"].(string); ok {
			ref, ok := byName[parent]
			if !ok {
				return fmt.Errorf("make variant of %s: spec not found", parent)
			}

			// Inherit missing values.
			for key, val := range ref {
				if _, ok := item[key]; !ok {
					item[key] = val
				}
			}

			opts, ok := item["variant_opts"]
			if !ok {
				opts = map[string]interface{}{}
			}
			data, err := json.Marshal(opts)
			if err != nil {
				return essentials.AddCtx("marshal variant opts", err)
			}
			item["variant_opts"] = strconv.Quote(string(data))
		}
	}
	return nil
}

var TemplateSource = `package muniverse

// An EnvSpec contains meta-data about an environment.
type EnvSpec struct {
	Name    string
	BaseURL string
	Width   int
	Height  int

	KeyWhitelist []string

	VariantOf   string
	VariantOpts string
}

var EnvSpecs = []*EnvSpec{ {{- range .}}
	{
		Name:    "{{.name}}",
		BaseURL: "{{.base}}",
		Width:   {{.width}},
		Height:  {{.height}},
		{{- $length := len .key_whitelist -}}
		{{if gt $length 0}}

		KeyWhitelist: []string{
			{{- range .key_whitelist}}
			"{{.}}", {{- end}}
		}, {{- end}}
		{{- if .variant_of}}

		VariantOf:   "{{.variant_of}}",
		VariantOpts: {{.variant_opts}},
		{{- end}}
	},
{{end -}} }

// SpecForName finds the first entry in EnvSpecs with the
// given name.
// If no entry is found, nil is returned.
func SpecForName(name string) *EnvSpec {
	for _, s := range EnvSpecs {
		if s.Name == name {
			return s
		}
	}
	return nil
}
`
