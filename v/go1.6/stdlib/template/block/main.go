package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// master := `{{ define "T1" }} T1 {{ end }}{{ template "T1" . }}`
	master := `{{ block "T1" . }} T1 {{ end }}`
	masterTmpl, err := template.New("master").Parse(master)
	if err != nil {
		log.Fatal(err)
	}

	// if err := masterTmpl.Execute(os.Stdout, nil); err != nil {
	// 	log.Fatal(err)
	// }

	overlay := `{{ define "T1" }} redefinition T1 {{ end }}`
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err)
	}
}
