package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

/*
Написать функцию, которая принимает на вход имя файла и название функции.
Необходимо подсчитать в этой функции количество вызовов асинхронных функций.
Результат работы должен возвращать количество вызовов int и ошибку error.
Разрешается использовать только go/parser, go/ast и go/token.
 */

/*
Комментарий:
Сумел вывести список функций и подфункций,
буду очень признателен, если подскажете как отфильтровать только go функции
 */

const srcFileName = "/home/user/go/go2hw/gb_go2_hw/hw7/app2/src.go"

func main() {

	// ast парсит и собирает Функции и подфункции

	funcs, funcsx, err := getFuncs(srcFileName)
	if err != nil {
		log.Fatal(err)
	}

	for i, d := range funcs {
		fmt.Println("found func:", d)
		for _, dx := range funcsx[i]{
			fmt.Println("sub_func:", dx)
		}
	}

}

func extractFuncs(funcs []*ast.FuncDecl) [][]string {
	var out [][]string
	var row []string
	for _, fun := range funcs {
		row = nil
		ast.Inspect(fun, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.CallExpr:
				b1 := types.ExprString(n.Fun)
				row = append(row, b1)
			}
			return true
		})
		out = append(out, row)
	}

	return out
}

func getFuncs(srcFileName string) ([]string, [][]string, error) {
	fset := token.NewFileSet()
	// парсим файл, чтобы получить AST
	astFile, err := parser.ParseFile(fset, srcFileName, nil, 0)
	if err != nil {
		return nil, nil, err
	}

	var out []string

	//funcs is slice of functions in file
	funcs := []*ast.FuncDecl{}
	for _, decl := range astFile.Decls {
		if fn, isFn := decl.(*ast.FuncDecl); isFn {
			funcs = append(funcs, fn)
			out = append(out, fn.Name.String())
		}
	}
	outx := extractFuncs(funcs)

	return out, outx, nil
}
