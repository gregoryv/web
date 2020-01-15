/*
   Package web provides html writing capabilities.

   In most cases when writing interactive webapplications you
   should stick to html/template package.
   This package however can be suitable when writing documentation
   or generating websites.


      article := Article(
         Class("fancy"),
         H1("Title of my article"),
      )
      page := NewPage("", Html(Body(article)))
      page.WriteTo(os.Stdout)

   By default the page is written as html, expected.

      <!DOCTYPE html>

      <html>
      <body>
      <article class="fancy">
      <h1>Title of my article</h1>

      </article>
      </body>
      </html>

*/
package web

var (
	_ = 1
)
