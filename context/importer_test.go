package context

import (
	"bytes"
	"testing"
)

func TestSassImport_single(t *testing.T) {
	in := bytes.NewBufferString(`@import "a";`)

	var out bytes.Buffer
	ctx := Context{}
	ctx.AddImport("a", "a { color: blue; }")
	err := ctx.Compile(in, &out)
	if err != nil {
		t.Fatal(err)
	}
	e := `a {
  color: blue; }
`
	if e != out.String() {
		t.Fatalf("got:\n%s\nwanted:\n%s", out.String(), e)
	}

}

func TestSassImport_multi(t *testing.T) {
	in := bytes.NewBufferString(`@import "a";
@import "b";`)

	var out bytes.Buffer
	ctx := Context{}
	ctx.AddImport("a", "a { color: blue; }")
	ctx.AddImport("b", "b { font-weight: bold; }")
	err := ctx.Compile(in, &out)
	if err != nil {
		t.Fatal(err)
	}
	e := `a {
  color: blue; }

b {
  font-weight: bold; }
`

	if e != out.String() {
		t.Fatalf("got:\n%s\nwanted:\n%s", out.String(), e)
	}

}
