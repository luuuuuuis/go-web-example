package dao

import (
    "fmt"
    "gjob-admin/pkg/model"
    "gjob-admin/pkg/utils"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接和自动迁移，并设置全局的 DB 变量
func InitDB() {
    var err error
    DB, err = gorm.Open(mysql.Open(utils.Config.GetString("mysql.base.dsn")), &gorm.Config{})
    if err!= nil {
        fmt.Printf("Error initializing database: %v\n", err)
        os.Exit(1)
    }

    // 执行自动迁移
    if err := autoMigrate(); err!= nil {
        fmt.Printf("Auto migration error: %v\n", err)
        os.Exit(1)
    }
}

// autoMigrate 执行数据库模型的自动迁移
func autoMigrate() error {
    return DB.AutoMigrate(&model.Node{}, &model.SqlJob{}, &model.ScriptJob{})
}