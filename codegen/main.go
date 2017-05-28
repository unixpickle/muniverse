// Command codegen generates a file called spec.go with
// specifications for each available environment.
package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/unixpickle/essentials"
)

func main() {
	spec, err := ioutil.ReadFile("games/spec.json")
	if err != nil {
		essentials.Die(err)
	}
	var specObj interface{}
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

func writeSpec(f *os.File, specObj interface{}) error {
	t := template.New("template")
	t, err := t.Parse(TemplateSource)
	if err != nil {
		return err
	}
	return t.Execute(f, specObj)
}

var TemplateSource = `package muniverse

// An EnvSpec contains meta-data about an environment.
type EnvSpec struct {
	Name    string
	BaseURL string
	Width   int
	Height  int

	KeyWhitelist []string
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
