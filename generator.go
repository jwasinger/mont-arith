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

func generateLimbImpl(maxLimbs int) {
	f, err := os.Create(fmt.Sprintf("build/arith_impl%d.go", maxLimbs))
	if err != nil {
		log.Println("create file: ", err)
		panic("fuck1")
	}

    params := TemplateParams{maxLimbs}
	template_content := loadTextFile("arith.go.template")

	tmpl := template.Must(template.New("").Funcs(funcs).Parse(template_content))

	if err := tmpl.Execute(f, params); err != nil {
		log.Fatal(err)
		panic("fuck2")
	}

	f.Close()
}

func main() {
    var max_limbs int = 3

    for i := 0; i < max_limbs; i++ {
		generateLimbImpl(i)
    }

}
