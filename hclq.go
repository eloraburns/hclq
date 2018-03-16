package main

import (
	"fmt"
	"github.com/hashicorp/hcl"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: hclq hcl_filename 'literal go template'\n")
		os.Exit(1)
	}

	hcl_filename := os.Args[1]
	template_text := os.Args[2]

	tmpl, err := template.New("hclq").Parse(template_text)
	if err != nil {
		fmt.Printf("Template parse err: %s\n", err)
		os.Exit(2)
	}

	hcl_file, err := ioutil.ReadFile(hcl_filename)
	if err != nil {
		fmt.Printf("File read err: %s\n", err)
		os.Exit(3)
	}

	var out interface{}
	err = hcl.Decode(&out, string(hcl_file))
	if err != nil {
		fmt.Printf("Input: %s\n\nError: %s\n", hcl_filename, err)
		os.Exit(4)
	}

	err = tmpl.Execute(os.Stdout, out)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err)
		os.Exit(5)
	}
	fmt.Printf("\n")

	os.Exit(0)
}
