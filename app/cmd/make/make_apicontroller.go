package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/console"
	"strings"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller, example: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1),
}

func runMakeAPIController(cmd *cobra.Command, args []string) {

	// 处理参数，要求附带 API 版本（v1 或者 v2）
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}

	// apiVersion 用来拼接目标路径
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)

	// 组建目标目录
	//dirPath := fmt.Sprintf("app/http/controllers/api/%s/", apiVersion)
	//err := os.MkdirAll(dirPath, os.ModePerm)
	//if err != nil {
	//	console.ExitIf(err)
	//}

	//filePath := fmt.Sprintf("%s_controller.go", model.TableName)

	filePath := fmt.Sprintf("app/http/controllers/api/%s/%s_controller.go", apiVersion, model.TableName)

	// 基于模版创建文件（做好变量替换）
	createFileFromStub(filePath, "apicontroller", model, map[string]string{"{{version}}": apiVersion})
}
