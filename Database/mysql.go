package Database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init() {

    // Log := logrus.New()
    // fileName := "runtime.log"
    // 写入文件
    // src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    // if err != nil {
    // 	fmt.Println("err", err)
    // }
    //
    // Log.Out = src
    // newLogger := logger.New(
    // 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
    // 	logger.Config{
    // 		SlowThreshold:             time.Second,   // 慢 SQL 阈值
    // 		LogLevel:                  logger.Silent, // 日志级别
    // 		IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
    // 		Colorful:                  false,         // 禁用彩色打印
    // 	},
    // )

    // dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
    // 	conf.DbUser,
    // 	conf.DbPassWord,
    // 	conf.DbHost,
    // 	conf.DbPort,
    // 	conf.DbName,
    // )
    dsn := ""
    // logrus.Fatal(dsn)
    gConf := &gorm.Config{
        NamingStrategy: schema.NamingStrategy{
            // TablePrefix: "",   // 表名前缀，`User`表为`t_users`
            SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
        },
        // Logger:            newLogger,
        AllowGlobalUpdate: true,
    }
    db, err := gorm.Open(mysql.New(mysql.Config{
        DSN:                       dsn,
        DefaultStringSize:         256,   // string 类型字段的默认长度
        DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
        DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
        DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
        SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
    }), gConf)

    if err != nil {
        logrus.Fatal("连接数据库错误:", err)
    }
    DB = db
}
