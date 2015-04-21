package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

// some defines...
// var defines = `
// {{define "func"}}func(){}{{end}}
//
// {{define "main"}}func main(){
// 	(.arg1)
// 	}{{end}}
// `

// run the application.
func main() {
	// variable to store the source string.
	src := ""

	// parse the commandline args and read file.
	if len(os.Args) == 1 {
		println("missing source...")
		println("try $ tplate source.tpl")
	} else {
		filepath := os.Args[1] + ".tplate"
		dat, err := ioutil.ReadFile(filepath)
		if err != nil {
			panic(err)
		}
		src = string(dat)
	}

	/*funcMap := template.FuncMap{
	  "Get": func() string {
	    return "--> get..."
	  },
	}*/
	tmpl, errTmpl := template.New("tpl"). /*.Funcs(funcMap)*/ Parse( /*defines +*/ src + "\n")
	if errTmpl != nil {
		panic(errTmpl)
	}

	tmplData := TemplateData{Name: "Mary"}
	err := tmpl.Execute(os.Stdout, tmplData)
	if err != nil {
		log.Printf("execution failed: %s", err)
	}
}

type TemplateData struct {
	Name string
}

func (t TemplateData) Yo(name string) string {
	return "Yo " + name
}
