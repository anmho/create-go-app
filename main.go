package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templateFs embed.FS

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println("pwd", pwd)
	templateDir := "templates"
	outDir := "out"
	fs.WalkDir(templateFs, templateDir, func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
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
            // file, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0666)
            if err != nil {
                panic(err)
            }
            err = tmpl.Execute(file, struct{Name string}{Name: "andy"})
            if err != nil {
                panic(err)
            }
        }

		return nil
	})

}
