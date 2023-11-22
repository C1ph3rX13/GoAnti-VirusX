package selectfunctmpl

import (
	"go/ast"    // Go语言抽象语法树包
	"go/parser" // Go语言解析器包
	"go/token"  // Go语言标记包
	"os"
	"path/filepath"
	"strings"
)

// GetExportedFunctionsFromFolder 获取指定文件夹中所有Go文件，并提取每个文件中的导出函数
func GetExportedFunctionsFromFolder(folderPath string) ([]string, error) {
	fset := token.NewFileSet() // 创建一个新的文件集合，用于保存解析产生的信息
	pkgs, err := parser.ParseDir(fset, folderPath, func(info os.FileInfo) bool {
		// 只解析Go源代码文件
		return !info.IsDir() && filepath.Ext(info.Name()) == ".go"
	}, parser.ParseComments) // 解析指定文件夹中的Go代码，并保留注释信息
	if err != nil {
		return nil, err
	}

	exportedFunctions := make([]string, 0)
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			functions := getExportedFunctionsFromFile(file)             // 获取当前文件中的导出函数
			exportedFunctions = append(exportedFunctions, functions...) // 将导出函数添加到结果列表中
		}
	}

	return exportedFunctions, nil
}

// getExportedFunctionsFromFile 获取指定Go源文件中的导出函数
func getExportedFunctionsFromFile(file *ast.File) []string {
	functions := make([]string, 0)

	for _, decl := range file.Decls {
		// 只提取函数声明
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// 只提取导出函数
		if funcDecl.Name.IsExported() {
			if strings.Contains(funcDecl.Name.Name, "Decrypt") {
				functions = append(functions, funcDecl.Name.Name)
			}
			continue
		}
	}

	return functions
}
