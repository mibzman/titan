package titan

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"text/template"
)

func TemplateIze(w io.Writer, file string, data interface{}) {
	parsedTemplate, error := template.ParseFiles(generateComponentList(file)...)
	if error != nil {
		w.Write(generateError(error))
		log.Fatal("template error ", error)
	}
	error = parsedTemplate.Execute(w, data)
	if error != nil {
		w.Write(generateError(error))
	}
}

func generateComponentList(file string) []string {
	files, err := ioutil.ReadDir("components/")
	if err != nil {
		log.Fatal(err)
	}

	fileNames := []string{"pages/" + file}
	for _, fileName := range files {
		fileNames = append(fileNames, "components/"+fileName.Name())
	}
	return fileNames
}

func generateError(err error) []byte {
	return []byte(fmt.Sprintf("\n\n## an error occurred:\n```\n" + err.Error() + "\n```\n"))
}
