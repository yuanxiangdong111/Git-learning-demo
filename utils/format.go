package utils

import (
	"bytes"
	"text/template"
)

func FormatTplByMap(f string, m map[string]interface{}) string {
	var tpl bytes.Buffer
	t := template.Must(template.New("").Parse(f))
	if err := t.Execute(&tpl, m); err != nil {
		return ""
	}
	return tpl.String()
}
