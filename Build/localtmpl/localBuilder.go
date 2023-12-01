package main

import (
	"GoAnti-VirusX/Build/compileutils"
	"GoAnti-VirusX/Utils"
)

const (
	TmplDir        string = "Build/localtmpl"   // 模板文件目录
	CompilationDir string = "Build/compilation" // 编译目录
)

func main() {
	err := compileutils.TmplBuilder(TmplDir, CompilationDir)
	if err != nil {
		Utils.GetLogger().Fatalf("Compilate error: %s", err)
	}
}
