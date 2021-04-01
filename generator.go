package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"strings"
)

type TemplateParams struct {
	LimbCount int
	LimbBits int
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
		n := end - start
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = start + i
		}
		return result
	},
	"sub": func(val, v2 int) int {
		return val - v2
	},
	"mul": func(val, v2 int) int {
		return val * v2
	},
	// returns "[x]uint64 {0, 0, 0, ......, 0}"
	"makeZeroedLimbs": func(numLimbs int) string {
		result := fmt.Sprintf("[%d]uint64 {", numLimbs)
		return result + strings.Repeat(" 0,", numLimbs - 1) + " 0}"
	},
}

func buildTemplate(dest_path, template_path string, params *TemplateParams) {
	f, err := os.Create(dest_path)
	if err != nil {
		log.Println("create file: ", err)
		panic("fuck1")
	}

	template_content := loadTextFile(template_path)

	tmpl := template.Must(template.New("").Funcs(funcs).Parse(template_content))

	if err := tmpl.Execute(f, *params); err != nil {
		log.Fatal(err)
		panic("fuck2")
	}

	f.Close()
}

func generateLimbFuncList(maxLimbs int) {
	params := TemplateParams{maxLimbs, 64}
	buildTemplate("build/entrypoint.go", "templates/entrypoint.go.template", &params)
}

func generateLimbImpl(maxLimbs int) {
	params := TemplateParams{maxLimbs, 64}

	buildTemplate(fmt.Sprintf("build/arith_%dlimbs.go", maxLimbs), "templates/arith.go.template", &params)
	buildTemplate(fmt.Sprintf("build/arith_%dlimbs_test.go", maxLimbs), "templates/arith_test.go.template", &params)
}

func main() {
	var min_limbs int = 4
	//var max_limbs int = 12
	var max_limbs int = 20

	generateLimbFuncList(max_limbs)

	fmt.Println("generating limb impls")

	for i := min_limbs; i <= max_limbs; i++ {
		fmt.Println(fmt.Sprintf("generating limb impl for %dx64-bit limbs", i))
		generateLimbImpl(i)
	}

}
