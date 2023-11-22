package main

import (
	"GoShellCodeLoader/Build/compileutils"
	"GoShellCodeLoader/Utils"
)

const (
	TmplDir        string = "Build/remotetmpl"  // 模板文件目录
	CompilationDir string = "Build/compilation" // 编译目录
)

func main() {
	err := compileutils.TmplBuilder(TmplDir, CompilationDir)
	if err != nil {
		Utils.GetLogger().Fatalf("编译错误: %s", err)
	}
}
