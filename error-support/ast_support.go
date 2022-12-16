package error_support

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strconv"
	"strings"
)

const comment_prefix = string("//")

type ErrorCode struct {
	Package string
	errors  map[int]*Error
}

func (ec *ErrorCode) Errors() map[int]*Error {
	return ec.errors
}

// print go file ast detail
func PrintAstInfo(fileName, src string, mode parser.Mode) {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, fileName, src, mode)
	if err != nil {
		panic(err)
	}
	ast.Print(fSet, f)
}

// find func and method in go file by target comment
// @see github.com\astaxie\beego\parser.go parserPkg
// @see github.com\astaxie\beego\parser.go parserComments
func scanFuncDeclByComment(fileName, src, targetComment string) *ErrorCode {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, fileName, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	result := &ErrorCode{
		errors: make(map[int]*Error),
	}
	result.Package = f.Name.String()
	var isAuto bool
	var lastCode int
	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.GenDecl:
			if len(decl.Specs) > 0 {
				for _, spec := range decl.Specs {
					switch value := spec.(type) {
					case *ast.ValueSpec:
						code, auto, err, bad := analysisComment(value, targetComment, lastCode, isAuto)
						if !bad {
							result.errors[code] = err
							isAuto = auto
							lastCode = code
						}
					}
				}
			}
		}

	}
	return result
}

func analysisComment(valueSpec *ast.ValueSpec, comment string, lastCode int, isAuto bool) (code int, auto bool, err *Error, bad bool) {
	if valueSpec.Doc == nil {
		bad = true
		return
	}
	if len(valueSpec.Names) != len(valueSpec.Doc.List) {
		bad = true
		return
	}

	if !isContainComment(valueSpec.Doc.List, comment) {
		bad = true
		return
	}

	err = &Error{}
	for _, name := range valueSpec.Names {
		err.Name = name.Name
	}

	if len(valueSpec.Values) == 0 && isAuto {
		if lastCode >= 0 {
			err.Code = lastCode + 1
		} else {
			err.Code = lastCode - 1
		}
		auto = isAuto
	} else if len(valueSpec.Values) == 0 {
		err.Code = lastCode
	}

	for _, value := range valueSpec.Values {
		switch value.(type) {
		case *ast.BasicLit:
			data, vErr := strconv.Atoi(value.(*ast.BasicLit).Value)
			if vErr != nil {
				panic(data)
			}
			err.Code = data
			auto = false
		case *ast.BinaryExpr:
			if value.(*ast.BinaryExpr).X.(*ast.Ident).Name == "iota" {
				auto = true
			}

			data, vErr := strconv.Atoi(value.(*ast.BinaryExpr).Y.(*ast.BasicLit).Value)
			if vErr != nil {
				panic(data)
			}
			if value.(*ast.BinaryExpr).Op.String() == "-" {
				err.Code = -data
			} else {
				err.Code = data
			}
		}
	}

	for _, doc := range valueSpec.Doc.List {
		err.Msg = processText(doc.Text, comment)
	}

	return err.Code, auto, err, false
}

func isContainComment(lines []*ast.Comment, targetComment string) bool {
	for _, l := range lines {
		c := strings.TrimSpace(strings.TrimLeft(l.Text, comment_prefix))

		result, _ := regexp.MatchString(targetComment, c)
		return result
	}
	return false
}

func processText(text, targetComment string) string {
	compileRegex := regexp.MustCompile(targetComment)
	matchArr := compileRegex.FindStringSubmatch(text)
	if len(matchArr) <= 0 {
		return text
	}

	return matchArr[0][8 : len(matchArr[0])-1]
}
