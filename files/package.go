/* Package files provides file loading utils.

 */
package files

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

// MustLoad returns the content of filename or panics.
func MustLoad(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// MustLoadLines returns the content of filename within a range. If to is -1
// the file is read until EOF. Lines start at 1.
func MustLoadLines(filename string, from, to int) string {
	var buf bytes.Buffer
	fh, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fh)
	for i := from; i > 1; i-- {
		scanner.Scan()
		to--
	}
	for scanner.Scan() {
		to--
		buf.WriteString(scanner.Text() + "\n")
		if to == 0 {
			break
		}
	}
	return buf.String()
}

// ----------------------------------------

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
