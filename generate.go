package main

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed _templates/*
var templateFs embed.FS

func replaceDirWithValues(dir string, opts options) {
	outDir := opts.appName
	// err := os.Mkdir(outDir, 0777)
	// if err != nil {
	// 	log.Fatalln("directory already exists here")
	// }
	
	// dirs, err := fs.ReadDir(templateFs, dir)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(dirs)

	fs.WalkDir(templateFs, dir, func(path string, d fs.DirEntry, err error) error {
		// fmt.Println(dir, path)
		if err != nil {
			panic(err)
		}
		
		parts := []string{}
		parts = append(parts, outDir)
		parts = append(parts, strings.Split(path, "/")[2:]...)
		// fmt.Println(parts)
		outPath := filepath.Join(parts...)
		if d.IsDir() {
			err := os.Mkdir(outPath, 0777)
			if err != nil {
				panic(err)
			}
		} else {
			tmpl, err := template.ParseFS(templateFs, path)
			if err != nil {
				panic(err)
			}
			if strings.HasSuffix(outPath, ".tmpl") {
				outPath = strings.Replace(outPath, ".tmpl", "", -1)
			}
			file, err := os.Create(outPath)
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(file, struct{ Name string }{Name: "andy"})
			if err != nil {
				panic(err)
			}
		}

		return nil
	})

}

func generateTemplatedAPI(opts options) {
	dir := "_templates/connectrpc-cloudrun"
	replaceDirWithValues(dir, opts)
}


func generateTemplatedMain(opts options) {
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	templateDir := "_templates"
	outDir := "out"

	fs.WalkDir(templateFs, templateDir, func(path string, d fs.DirEntry, err error) error {
		// fmt.Println(path)
		parts := strings.Split(path, "/")
		parts[0] = outDir
		outPath := filepath.Join(parts...)
		if d.IsDir() {
			os.Mkdir(outPath, 0777)
		} else {
			tmpl, err := template.ParseFS(templateFs, path)
			if err != nil {
				panic(err)
			}
			file, err := os.Create(outPath)
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(file, struct{ Name string }{Name: "andy"})
			if err != nil {
				panic(err)
			}
		}

		return nil
	})
}
