package config

import (
	"io/ioutil"
	"os"
	"path"
)

var appConfigPath string

func lookingConfig(filepath ...string) {
	var fp string
	var err error
	if len(filepath) > 0 {
		fp = filepath[0]
	} else {
		fp, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}
	if len(fp) < 2 {
		panic("can't find config")
	}
	files, err := ioutil.ReadDir(fp)
	if err != nil {
		panic(err)
	}
	for i := range files {
		if "config" == files[i].Name() {
			appConfigPath = fp + "/config/"
			return
		}
	}
	lookingConfig(path.Dir(fp))
}
