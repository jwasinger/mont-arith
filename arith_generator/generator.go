package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"bytes"
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
	"add": func(val, v2 int) int {
		return val + v2
	},
	"sub": func(val, v2 int) int {
		return val - v2
	},
	"max": func(val, val2 int) int {
		if val > val2 {
			return val
		} else {
			return val2
		}
	},
	"gte": func(v1, v2 int) bool {
		return v1 >= v2
	},
	"mul": func(val, v2 int) int {
		return val * v2
	},
	// returns "[x]uint64 {0, 0, 0, ......, 0}"
/*
	"makeZeroedLimbs": func(numLimbs int) string {
		result := fmt.Sprintf("[%d]uint64 {", numLimbs)
		return result + strings.Repeat(" 0,", numLimbs - 1) + " 0}"
	},
*/
}

func buildTemplate(dest_path, template_path string, params *TemplateParams) {
	f, err := os.Create(dest_path)
	if err != nil {
		log.Println("create file: ", err)
		panic("")
	}

	template_content := loadTextFile(template_path)

	tmpl := template.Must(template.New("").Funcs(funcs).Parse(template_content))

	if err := tmpl.Execute(f, *params); err != nil {
		log.Fatal(err)
		panic("")
	}

	f.Close()
}

func genAddModUnrolled(minLimbs, maxLimbs int) {
	headerTemplateContent := loadTextFile("templates/header.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	addModUnrolledTemplateContent := loadTextFile("templates/addmod_unrolled.go.template")
	addModUnrolledTemplate := template.Must(template.New("").Funcs(funcs).Parse(addModUnrolledTemplateContent))

	params := TemplateParams{0, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create("addmod_unrolled.go")
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	for i := minLimbs; i <= maxLimbs; i++ {
		params = TemplateParams{i, 64}
		if err := addModUnrolledTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}
}

func genAddMod(minLimbs, maxLimbs int) {
	headerTemplateContent := loadTextFile("templates/addmodsubmodheader.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	params := TemplateParams{0, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create("addmod_nonunrolled.go")
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	addModNonUnrolledTemplateContent := loadTextFile("templates/addmod_nonunrolled.go.template")
	addModNonUnrolledTemplate := template.Must(template.New("").Funcs(funcs).Parse(addModNonUnrolledTemplateContent))

	for i := minLimbs; i <= maxLimbs; i++ {
		params = TemplateParams{i, 64}
		if err := addModNonUnrolledTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}
}


func genSubModUnrolled(minLimbs, maxLimbs int) {
	headerTemplateContent := loadTextFile("templates/header.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	addModUnrolledTemplateContent := loadTextFile("templates/submod_unrolled.go.template")
	addModUnrolledTemplate := template.Must(template.New("").Funcs(funcs).Parse(addModUnrolledTemplateContent))

	params := TemplateParams{0, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create("submod_unrolled.go")
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	for i := minLimbs; i <= maxLimbs; i++ {
		params = TemplateParams{i, 64}
		if err := addModUnrolledTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}
}

func genSubMod(minLimbs, maxLimbs int) {
	headerTemplateContent := loadTextFile("templates/header.go.template")
	headerTemplate := template.Must(template.New("").Funcs(funcs).Parse(headerTemplateContent))

	params := TemplateParams{0, 64}
	buf := new(bytes.Buffer)

	f, err := os.Create("submod.go")
	if err != nil {
		log.Fatal(err)
		panic("")
	}

	if err := headerTemplate.Execute(buf, params); err != nil {
		log.Fatal(err)
		panic("")
	}

	subModNonUnrolledTemplateContent := loadTextFile("templates/submod.go.template")
	subModNonUnrolledTemplate := template.Must(template.New("").Funcs(funcs).Parse(subModNonUnrolledTemplateContent))

	for i := minLimbs; i <= maxLimbs; i++ {
		params = TemplateParams{i, 64}
		if err := subModNonUnrolledTemplate.Execute(buf, params); err != nil {
			log.Fatal(err)
			panic("")
		}
	}

	if n, err := f.Write(buf.Bytes()); err != nil || n != len(buf.Bytes()) {
		panic(err)
	}
}

func genPresets(maxLimbs int) {
	params := TemplateParams{maxLimbs, 64}
	buildTemplate("presets.go", "templates/presets.go.template", &params)
}

func generateMulModMontImpl(limbCount int) {
	params := TemplateParams{limbCount, 64}
	buildTemplate(fmt.Sprintf("mulmodmont%d.go", limbCount * 64), "templates/mulmodmont.go.template", &params)
}

func genMulModMont() {
	var min_limbs int = 2
	//var max_limbs int = 11
	var max_limbs int = 11

	for i := min_limbs; i <= max_limbs; i++ {
		generateMulModMontImpl(i)
	}
}

func main() {
	genPresets(64)

	genMulModMont()
	genAddMod(2, 64)
	genSubMod(2, 64)

	genAddModUnrolled(2, 64)
	genSubModUnrolled(2, 64)
}
