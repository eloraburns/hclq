package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: hclq hcl_filename 'literal go template'\n")
		os.Exit(1)
	}

	hcl_filename := os.Args[1]
	template_text := os.Args[2]

	hcl_text, err := ioutil.ReadFile(hcl_filename)
	if err != nil {
		fmt.Printf("File read err: %s\n", err)
		os.Exit(2)
	}

	out, err := ParseAndFormat(hcl_text, template_text)
	if err != nil {
		fmt.Printf("Error parsing/formatting: %s\n", err)
		os.Exit(3)
	}

	fmt.Printf("%s\n", out)
	os.Exit(0)
}
