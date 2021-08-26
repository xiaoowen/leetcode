package main

import (
	"strings"
)

func entityParser(text string) string {
	return strings.NewReplacer(`&quot;`, `"`, `&apos;`, `'`, `&amp;`, `&`, `&gt;`, `>`, `&lt;`, `<`, `&frasl;`, `/`).Replace(text)
}
