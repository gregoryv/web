package files

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)

func MustLoadFunc(filename, funcName string) string {
	b, err := LoadFunc(filename, funcName)
	if err != nil {
		panic(err)
	}
	return b
}

// LoadFunc parses the filename for a given function name and returns
// it in its current form as a string.
func LoadFunc(filename, funcName string) (string, error) {
	funcAST, fset := parseFunc(filename, funcName)
	if funcAST == nil {
		return "", fmt.Errorf("func %s not found in %s", funcName, filename)
	}
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, funcAST)
	b := buf.String()
	// Remove the tab after signature
	b = strings.ReplaceAll(b, ")\t{", ") {")
	return b, nil
}

func parseFunc(filename, funcName string) (fun *ast.FuncDecl, fset *token.FileSet) {
	fset = token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return nil, nil
	}
	for _, d := range file.Decls {
		f, ok := d.(*ast.FuncDecl)
		if !ok || f.Name.Name != funcName {
			continue
		}
		fun = f
		return
	}
	return nil, nil
}
