package main

import (
	"GoAnti-VirusX/Build/compileutils"
	"GoAnti-VirusX/log"
)

const (
	TmplDir        string = "Build/remotetmpl"  // 模板文件目录
	CompilationDir string = "Build/compilation" // 编译目录
)

func main() {
	err := compileutils.TmplBuilder(TmplDir, CompilationDir)
	if err != nil {
		log.Fatalf("编译错误: %s", err)
	}
	log.Info("Compile Susses")
}
