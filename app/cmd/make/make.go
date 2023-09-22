package make

import (
	"embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"gohub/pkg/console"
	"gohub/pkg/file"
	"gohub/pkg/str"
	"strings"
)

// Model 参数解释
// 单个单词，用户命令传参，以 User 模型为例:
// - user
// - User
// - users
// - Users
// 整理好的数据
//{
//	"TableName": "users",
//	"StructName": "User",
//	"StructNamePlural": "Users",
//	"VariableName": "user"
//	"VariableNamePlural": "users",
//	"PackageName": "user"
//}
// 两个词或者以上，用户命令传参，以 TopicComment 模型为例:
// - topic_comment
// - topic_comments
// - TopicComment
// - TopicComments
// 整理好的数据
//{
//	"TableName": "topic_comments",
//	"StructName": "TopicComment",
//	"StructNamePlural": "TopicComments",
//	"VariableName": "topicComment",
//	"VariableNamePlural": "topicComments",
//	"PackageName": "topic_comment"
//}

type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

// stubsFS 方便我们后面打包这些 .stub 为后缀名的文件

//go:embed stubs
var stubsFS embed.FS

// CmdMake 说明 cobra 命令
var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	// 注册 make 的自命令
	CmdMake.AddCommand(
		CmdMakeCmd,
		CmdMakeModel,
	)
}

// makeModelFromString 格式化用户输入的内容
func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(name))            // topic_comments > TopicComments > TopicComment
	model.StructNamePlural = str.Plural(model.StructName)             // TopicComment > TopicComments
	model.TableName = str.Snake(model.StructNamePlural)               // TopicComments > topic_comments
	model.VariableName = str.LowerCamel(model.StructName)             // TopicComment > topicComment
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural) // TopicComments > topicComments
	model.PackageName = str.Snake(model.StructName)                   // TopicComment > topic_comment
	return model
}

// createFileFromStub 读取 stub 文件并进行变量替换
// 最后一个选项可选，如若传参，应传 map[string]string 类型，作为附加的变量搜索替换
func createFileFromStub(filepath string, stubName string, model Model, variables ...interface{}) {

	// 实现最后一个参数可选
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	// 目标文件已存在
	if file.Exists(filepath) {
		console.Exit(filepath + " already exists!")
	}

	// 读取 stub 模版文件
	modelData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}
	modelStub := string(modelData)

	// 添加默认的替换变量
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName

	// 对模版内容做变量替换
	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	// 存储到目标文件中
	err = file.Put([]byte(modelStub), filepath)
	if err != nil {
		console.Exit(err.Error())
	}

	console.Success(fmt.Sprintf("[%s] created", filepath))
}
