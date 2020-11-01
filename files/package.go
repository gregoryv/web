/* Package files provides file loading utils.

 */
package files

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
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
