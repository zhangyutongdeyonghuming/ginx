package cmd

import (
	"os"
	"path"
	"runtime"
	"text/template"

	"github.com/ginx/templates"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [projectName]",
	Short: "Create a new project",
	Long:  `This command allows you to create a new project with a specified name.`,
	Args:  cobra.ExactArgs(1), // 确保只有一个参数
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0] // 读取第一个参数作为项目名称
		initProjce(moduleName)
	},
}

// 初始化
func initProjce(moduleName string) {
	// 获取当前 Go 版本数字部分
	goVersion := runtime.Version()[2:]
	err := os.Mkdir(moduleName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 定义数据
	data := map[string]string{
		"moduleName": moduleName,
		"goVersion":  goVersion,
	}
	ts := templates.Templates
	for _, t := range ts {
		fullPath := path.Join(moduleName, t.Path)
		// 创建多级目录
		err := os.MkdirAll(fullPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		// 创建模板并解析模板字符串
		tmpl, err := template.New("example").Parse(t.Template)
		if err != nil {
			panic(err)
		}

		// 创建输出文件
		outputFile, err := os.Create(path.Join(fullPath, t.Name))
		if err != nil {
			panic(err)
		}
		defer outputFile.Close()

		// 将模板执行结果写入文件
		err = tmpl.Execute(outputFile, data)
		if err != nil {
			panic(err)
		}
	}
}
