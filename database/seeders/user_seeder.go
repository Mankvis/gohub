package seeders

import (
	"fmt"
	"gohub/database/fatcories"
	"gohub/pkg/console"
	"gohub/pkg/logger"
	"gohub/pkg/seed"
	"gorm.io/gorm"
)

func init() {

	// 添加 Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		// 创建 10 个用户对象
		users := fatcories.MakeUsers(10)

		// 批量创建用户（注意批量创建不会调用模型钩子）
		result := db.Table("users").Create(&users)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Tbale [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}