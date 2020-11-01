package docs

import "testing"

func Test_generate_pages(t *testing.T) {
	NewIndexPage().SaveAs("index.html")
}
