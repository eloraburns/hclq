package main

import (
	"errors"
	"fmt"
	"github.com/hashicorp/hcl"
	"strings"
	"text/template"
)

func ParseAndFormat(hcl_text []byte, template_text string) (string, error) {
	tmpl, err := template.New("hclq").Parse(template_text)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Template parse error: %s", err))
	}

	var out interface{}
	err = hcl.Decode(&out, string(hcl_text))
	if err != nil {
		return "", errors.New(fmt.Sprintf("HCL parse error: %s", err))
	}

	var out_buf strings.Builder
	err = tmpl.Execute(&out_buf, out)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error executing template: %s", err))
	}

	return out_buf.String(), nil
}
