package render

import (
	"bytes"
	"html/template"
)

type PageData struct {
	Title  string
	Header string
	Items  []string
}

func RenderTemplate(movies []string) string {
	// Load the template file
	htmlFile, err := template.ParseFiles("example.html")
	if err != nil {
		panic(err)
	}

	// // Prepare the data to insert into the template
	data := PageData{
		Title:  "My Page",
		Header: "Movies",
		Items:  movies,
	}

	var buf bytes.Buffer

	// Render the template with the data
	err = htmlFile.Execute(&buf, data)
	if err != nil {
		panic(err)
	}

	output := buf.String()

	return output
}
