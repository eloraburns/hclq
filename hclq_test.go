package main

import "testing"

func TestThings(t *testing.T) {
	cases := []struct {
		hcl, tmpl, want string
	}{
		{"foo = \"bar\"", "{{ .foo }}", "bar"},
		{"foo = 2", "{{ .foo }}", "2"},
		{"foo = 3", "{{ . }}", "map[foo:3]"},
		// Would you believe that HCL wraps all your objects in lists? Well…believe it.
		{"foo = { bar = 4 }", "{{ (index .foo 0).bar }}", "4"},
		{"foo \"bar\" { baz = 5 }", "{{ (index (index .foo 0).bar 0).baz }}", "5"},
	}
	for _, c := range cases {
		got, err := ParseAndFormat([]byte(c.hcl), c.tmpl)
		if err != nil || got != c.want {
			t.Errorf("ERROR: %s\n\nHCL: %s\n\nTemplate: %s\n\nGot: %s\n\nWant: %s\n", err, c.hcl, c.tmpl, got, c.want)
		}
	}
}
