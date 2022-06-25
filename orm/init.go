package orm

import (
	"database/sql/driver"
	"fmt"
	paginator "github.com/dmitryburov/gorm-paginator"
	"github.com/gomodule/redigo/redis"
	"github.com/ipipdotnet/ipdb-go"
	"github.com/spf13/viper"

	"go-hash/a"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	// 定义常量
	Db *gorm.DB
	// 定义常量
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_DB    string
	// IP 地址
	Ipdata *ipdb.City
)

type Lists struct {
	Items      interface{}
	Pagination *paginator.Pagination
}

var ChYH map[string]chan string
var ChCount map[string]chan int

// SpotMap 现货委托
var SpotMap sync.Map

const TimeFormat = "2006-01-02 15:04:05"

var GO_DUIJIAN string

//= "go_btc_job"

//.Set("gorm:query_option", "FOR UPDATE")
func Start() {
	// 加载IP地址库
	Ipdata, _ = ipdb.NewCity(a.Path + "/ip.ipdb")
	GO_DUIJIAN = "go_hash_job" + viper.GetString("database.database")
	var err error
	ChYH = make(map[string]chan string)
	ChCount = make(map[string]chan int)

	database := viper.GetString("database.database")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	//conn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local&sql_mode=`STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION`", username, password, host, database)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&sql_mode=`STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION,ALLOW_INVALID_DATES`", username, password, host, port, database)

	fmt.Println("orm init conn", conn)
	//Db, err = gorm.Open("mysql", conn)
	Db, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		QueryFields:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		//DisableDatetimePrecision:                 true,
	})
	if err != nil {
		fmt.Println("err", err)
	}
	d, _ := Db.DB()
	//Db.DB().Ping()
	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(100)

	//Db.AutoMigrate(&AdminUser{})
	//Db.AutoMigrate(&Lottery{})
	//new(Lottery).First()
	//Db.AutoMigrate(&PlayKind{})
	//Db.AutoMigrate(&User{})
	//Db.AutoMigrate(&UserLower{})
	//Db.AutoMigrate(&UserDebt{})
	//Db.AutoMigrate(&UserCaptial{})
	//Db.AutoMigrate(&PayManage{})
	//Db.AutoMigrate(&LotteryResult{})
	//Db.AutoMigrate(&LotteryResultDate{})
	//Db.AutoMigrate(&SysStatistics{})
	//Db.AutoMigrate(&PayLog{})
	//Db.AutoMigrate(&UserBet{})
	//Db.AutoMigrate(&Setting{})
	//Db.AutoMigrate(&BackwaterRule{})
	//Db.AutoMigrate(&Withdraw{})
	//Db.AutoMigrate(&Chat{})
	//new(LotteryResultDate).GameAddTime()
	//new(Setting).Init()

	// 设置初始化
	//new(Setting).Init()
	// 启用Logger，显示详细日志
	if viper.GetBool("database.debug") == true {
		//Db.LogMode(true)
		//Db.Logger.LogMode(4)
		//Db.Logger.LogMode(1)
	}

	//以下连接redis
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = viper.GetString("redis.host")
	REDIS_DB = viper.GetString("redis.db")
	port = viper.GetString("redis.port")
	fmt.Println("REDIS_HOST", REDIS_HOST, REDIS_DB)
	MaxIdle := viper.GetInt("redis.maxidle")
	MaxActive := viper.GetInt("redis.maxactive")
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: 180 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST+":"+port)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
	//RabbitClientConnect()
}

type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("15:04:05"))
	if formatted == `"00:00:00"` {
		formatted = fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	} else {
		formatted = fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	}
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t JSONTime) String() string {
	return t.Time.Format("2006-01-02 15:04:05")
}

type JSONDate struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONDate{Time: value}
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t JSONDate) String() string {
	return t.Time.Format("2006-01-02")
}
