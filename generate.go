package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed _templates/*
var templateFs embed.FS

func replaceDirWithValues(dir string, opts options) error {
	outDir := opts.appName
	fs.WalkDir(templateFs, dir, func(path string, d fs.DirEntry, err error) error {
		// fmt.Println(dir, path)
		if err != nil {
			return err
		}
		
		parts := []string{}
		parts = append(parts, outDir)
		parts = append(parts, strings.Split(path, "/")[2:]...)
		// fmt.Println(parts)
		outPath := filepath.Join(parts...)
		if d.IsDir() {
			err := os.Mkdir(outPath, 0777)
			if err != nil {
				return err
			}
		} else {
			tmpl, err := template.ParseFS(templateFs, path)
			if err != nil {
				return err
			}
			if strings.HasSuffix(outPath, ".tmpl") {
				outPath = strings.Replace(outPath, ".tmpl", "", -1)
			}
			file, err := os.Create(outPath)
			if err != nil {
				return err
			}
			err = tmpl.Execute(file, struct{ Name string }{Name: "andy"})
			if err != nil {
				return err
			}
		}

		return nil
	})
	return nil

}

func generateTemplatedAPI(opts options) error {

	var dir string
	switch opts.appTemplate {
	case string(ConnectCloudRun):
		dir = "_templates/connectrpc-cloudrun"
	case string(HumaFlyIO):
		dir = "_templates/huma-fly"
	default:
		return errors.New(fmt.Sprintf("unknown template type", opts.appTemplate))
	} 

	return replaceDirWithValues(dir, opts)
}