# hclq
A simple utility to parse HashiCorp HCL files and return values

Given:
```
# some.hcl
foo = "bar"
bar = {
  quux = "zounds"
}
```

You can extract simple values:
```
$ ./hclq some.hcl '{{.foo}}'
bar
```

Nested objects are wrapped in lists, because HCL:
```
$ ./hclq some.hcl '{{ (index .bar 0).quux }}'
zounds
```

## Build
`make`
