package main

import (
	"log"
	"os"
	"text/template"
	"io/ioutil"
)

type TemplateParams struct {
    LimbCount int
}

func loadTextFile(file_name string) string {
    content, err := ioutil.ReadFile(file_name)
    if err != nil {
		panic(err)
    }

    // Convert []byte to string and print to screen
    return string(content)
}

var funcs = template.FuncMap{
    "intRange": func(start, end int) []int {
        n := end - start + 1
        result := make([]int, n)
        for i := 0; i < n; i++ {
            result[i] = start + i
        }
        return result
    },
    "sub": func(val, v2 int) int {
        return val - v2
    },
}

func main() {
    params := TemplateParams{3}
	template_content := loadTextFile("arith.go.template")

	tmpl := template.Must(template.New("").Funcs(funcs).Parse(template_content))

	if err := tmpl.Execute(os.Stdout, params); err != nil {
		log.Fatal(err)
	}
}
