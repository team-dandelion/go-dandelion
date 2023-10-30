package build

import (
	"github.com/team-dandelion/go-dandelion/go-dandelion-cli/internal/asset"
	"os"
	"text/template"
)

func CreateTemplateFile(filePath, tmplPath string, data any) error {
	txtData, err := asset.Asset(tmplPath)
	if err != nil {
		return err
	}
	tmpl, err := template.New(tmplPath).Parse(string(txtData))
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
