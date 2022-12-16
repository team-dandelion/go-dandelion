package error_support

import (
	"fmt"
	"go/parser"
	"io/ioutil"
	"testing"
)

type Transact struct {
}

func TestPrintAstInfo(t *testing.T) {
	bt, _ := ioutil.ReadFile(`ast_text.go`)
	src := string(bt)
	PrintAstInfo(``, src, parser.ParseComments)
}

func TestScanFuncDeclByComment(t *testing.T) {
	bt, _ := ioutil.ReadFile(`ast_text.go`)
	src := string(bt)
	result := scanFuncDeclByComment(``, src, "@ErrMsg(.*?)")
	fmt.Println(fmt.Sprintf("%v", result))
}
