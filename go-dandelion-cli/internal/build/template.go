package build

import (
	"os"
	"text/template"
)

func CreateTemplateFile(filePath, tmplPath string, data any) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	file, oErr := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
	if oErr != nil {
		return oErr
	}
	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}
