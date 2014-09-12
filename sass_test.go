package sprite_sass_test

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"testing"

	. "github.com/drewwells/sprite_sass"
)

func fileString(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func fileReader(path string) io.Reader {
	reader, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return reader
}

func TestRun(t *testing.T) {

	var scanned []byte
	ipath := "test/_var.scss"
	exp, err := ioutil.ReadFile("test/var.css")
	if err != nil {
		panic(err)
	}

	ctx := Context{
		OutputStyle:  NESTED_STYLE,
		IncludePaths: make([]string, 0),
		Out:          "",
	}
	r, w := io.Pipe()
	go ctx.Run(fileReader(ipath), w, "test")

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		scanned = append(scanned, scanner.Bytes()...)
	}

	scanned = rerandom.ReplaceAll(scanned, []byte(""))
	if string(scanned) != string(exp) {
		t.Errorf("Processor file did not match was: "+
			"\n~%s~\n exp:\n~%s~", string(scanned), string(exp))
	}

}

/*func TestCompile(t *testing.T) {
	ctx := Context{
		OutputStyle:  NESTED_STYLE,
		IncludePaths: make([]string, 0),
		Src:          fileString("test/file1.scss"),
		Out:          "",
	}
	ctx.Compile()
	fpath := "test/file1.css"
	bytes, _ := ioutil.ReadFile(fpath)
	exp := string(bytes)

	if ctx.Out != exp {
		t.Errorf("%s string mismatch found: \n%s, expected \n%s",
			fpath, ctx.Out, exp)
	}
	fpath = "test/file1a.scss"
	ctx.Src = fileString(fpath)
	ctx.Compile()

	if ctx.Out != exp {
		t.Errorf("%s string mismatch found: \n%s, expected \n%s",
			fpath, ctx.Out, exp)
	}
}

func TestExport(t *testing.T) {
	ctx := Context{
		OutputStyle:  NESTED_STYLE,
		IncludePaths: make([]string, 0),
		Src:          fileString("test/sheet1.scss"),
		Out:          "",
	}
	ctx.Compile()
}*/
