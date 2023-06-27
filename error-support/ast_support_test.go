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

func Test_processText(t *testing.T) {
	type args struct {
		text          string
		targetComment string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				text:          "@ErrMsg(test)",
				targetComment: "@ErrMsg[(].*?[)]",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processText(tt.args.text, tt.args.targetComment); got != tt.want {
				t.Errorf("processText() = %v, want %v", got, tt.want)
			}
		})
	}
}
