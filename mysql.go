package m7sdb

import (
	"strconv"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "m7s.live/engine/v4"
)
var once sync.Once
type MysqlConfig struct {
	Host string
	Port int
	Username string
	Password string
	DBName string
}

var MysqlPlugConfig =  &MysqlConfig{}

var plugin = InstallPlugin(MysqlPlugConfig)


func (conf *MysqlConfig) OnEvent(event any){
 switch event.(type) {
    case FirstConfig: //插件初始化逻辑
    case *Stream://按需拉流逻辑
    case SEwaitPublish://由于发布者掉线等待发布者
    case SEpublish://首次进入发布状态
    case SErepublish://再次进入发布状态
    case SEwaitClose://由于最后一个订阅者离开等待关闭流
    case SEclose://关闭流
    case UnsubscribeEvent://订阅者离开
  }
}
var db *gorm.DB

func DB() *gorm.DB{
	 once.Do(func() {
       	dsn := MysqlPlugConfig.Username + ":" + MysqlPlugConfig.Password + "@tcp(" + MysqlPlugConfig.Host + ":" + strconv.Itoa( MysqlPlugConfig.Port )+ ")/" + MysqlPlugConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
				db, _  =  gorm.Open(mysql.Open(dsn), &gorm.Config{}) // database: *gorm.DB
    })
    // 核心代码,初始化后,在返回实例
    return db
}
