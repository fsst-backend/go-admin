package swagger

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/spf13/cobra"

	"go-admin/app/admin/models"
	"go-admin/common/database"

	filewrap "go-admin/config/filewarp"

	"github.com/go-admin-team/go-admin-core/sdk/config"
)

var (
	cfgFile  string
	jsonFile string
)

var Cmd = &cobra.Command{
	Use:   "update-sysapi",
	Short: "Update SysApi table from Swagger JSON file",
	Long:  `Parse Swagger JSON file and update SysApi table, update if path and action match, insert if not exist`,
	Run:   run,
}

func init() {
	Cmd.PersistentFlags().StringVar(&cfgFile, "config", "/app/settings.yaml", "config file")
	Cmd.PersistentFlags().StringVar(&jsonFile, "json", "/app/docs/violet/backend-api.json", "swagger json file")
}

func run(cmd *cobra.Command, args []string) {
	// 加载配置
	config.Setup(
		filewrap.NewFileWrap(cfgFile),
		database.Setup,
	)

	db := sdk.Runtime.GetDbByKey("*")
	if db == nil {
		fmt.Println("数据库连接不存在")
		return
	}

	data, err := os.ReadFile(jsonFile)

	if err != nil {
		fmt.Printf("读取文件失败: %s\n", err.Error())
		return
	}

	// 解析 JSON
	var swaggerData map[string]interface{}
	err = json.Unmarshal(data, &swaggerData)
	if err != nil {
		fmt.Printf("解析 JSON 失败: %s\n", err.Error())
		return
	}

	// 提取 paths
	paths, ok := swaggerData["paths"].(map[string]interface{})
	if !ok {
		fmt.Println("JSON 中没有找到 paths 字段")
		return
	}

	// 遍历 paths
	for path, methods := range paths {
		methodsMap, ok := methods.(map[string]interface{})
		if !ok {
			continue
		}

		// 遍历 HTTP 方法
		for method, operation := range methodsMap {
			operationMap, ok := operation.(map[string]interface{})
			if !ok {
				continue
			}

			// 获取标题和标签
			title := ""
			if titleVal, exists := operationMap["summary"]; exists {
				if titleStr, ok := titleVal.(string); ok {
					title = titleStr
				}
			}

			// 获取标签
			tag := ""
			if tagsVal, exists := operationMap["tags"]; exists {
				if tagsArray, ok := tagsVal.([]interface{}); ok && len(tagsArray) > 0 {
					// 取第一个标签
					if tagStr, ok := tagsArray[0].(string); ok {
						tag = tagStr
					}
				}
			}

			// 构建 SysApi 对象
			sysApi := models.SysApi{
				Path:   path,
				Action: strings.ToUpper(method),
				Title:  title,
				Tag:    tag,
				Handle: fmt.Sprintf("%s %s", strings.ToUpper(method), path),
			}

			// 检查是否已存在相同的 path 和 action
			var existingApi models.SysApi
			result := db.Where("path = ? AND action = ?", path, strings.ToUpper(method)).First(&existingApi)

			if result.Error == nil && result.RowsAffected > 0 {
				// 更新现有记录
				sysApi.Id = existingApi.Id
				err := db.Model(&existingApi).Updates(&sysApi).Error
				if err != nil {
					fmt.Printf("更新记录失败: path=%s, action=%s, error=%s\n", path, method, err.Error())
				} else {
					fmt.Printf("更新记录成功: path=%s, action=%s, title=%s\n", path, method, title)
				}
			} else {
				// 插入新记录
				err := db.Create(&sysApi).Error
				if err != nil {
					fmt.Printf("插入记录失败: path=%s, action=%s, error=%s\n", path, method, err.Error())
				} else {
					fmt.Printf("插入记录成功: path=%s, action=%s, title=%s\n", path, method, title)
				}
			}
		}
	}

	fmt.Println("Swagger JSON 更新 SysApi 表完成")
}
