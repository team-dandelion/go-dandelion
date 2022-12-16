package error_support

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

var routes []string

func walkFunc(route string, info os.FileInfo, err error) error {
	if info == nil {
		return nil
	}
	if info.IsDir() {
		return nil
	} else {
		if path.Ext(route) == ".go" {
			routes = append(routes, route)
		}
		return nil
	}
}

func showFileList(root string) {
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		fmt.Printf("filepath.Walk() error: %v\n", err)
	}
	return
}
