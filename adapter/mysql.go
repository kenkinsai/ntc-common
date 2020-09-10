package adapter

import (
	"fmt"
	"log"
	"sync"
	"time"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

// MYSQL ..
type MYSQL struct {
	Name    string `mapstructure:"name"`
	Server  string `mapstructure:"server"`
	Port    int    `mapstructure:"port"`
	DBName  string `mapstructure:"dbname"`
	User    string `mapstructure:"user"`
	Pass    string `mapstructure:"password"`
	Session *xorm.Engine
}


// MYSQLs ..
type MYSQLs map[string]*MYSQL

// Get MYSQL
func (adapters MYSQLs) Get(name string) (result *MYSQL) {
	if adapter, ok := adapters[name]; ok {
		result = adapter
	} else {
		panic("Không tìm thấy config MYSQL " + name)
	}
	return
}


var (
	onceMYSQL      map[string]*sync.Once
	onceMYSQLMutex = sync.RWMutex{}
)

func init() {
	onceMYSQL = make(map[string]*sync.Once)
}

func (config MYSQL) isExist() bool {
	return config.Name != ""
}

// Init ..
func (config *MYSQL) Init() {
	if onceMYSQL[config.Name] == nil {
		onceMYSQL[config.Name] = &sync.Once{}
	}
	onceMYSQL[config.Name].Do(func() {
		onceMYSQLMutex.Lock()
		log.Printf("[%s][%s] MYSQL [connecting]\n", config.Name, config.Server)
		connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Pass, config.Server, config.Port, config.DBName)
		xormEngine, err := xorm.NewEngine("mysql", connString)
		if err != nil {
			log.Printf("[%s][%s] MYSQL [error]: %s", config.Name, config.Server, err)
			time.Sleep(1 * time.Second)
			onceMYSQL[config.Name] = &sync.Once{}
			onceMYSQLMutex.Unlock()
			config.Init()
		} else {
			config.Session = xormEngine
			config.Session.SetMaxOpenConns(100)
			config.Session.SetMaxIdleConns(5)
			//connect thử
			errPing := config.Session.Ping()
			if errPing != nil {
				log.Printf("[%s][%s] MYSQL ping [error]: %s", config.Name, config.Server, errPing)
				time.Sleep(1 * time.Second)
				onceMYSQL[config.Name] = &sync.Once{}
				onceMYSQLMutex.Unlock()
				config.Init()
				return
			}
			log.Printf("[%s][%s] MYSQL [connected]\n", config.Name, config.Server)
			onceMYSQLMutex.Unlock()
		}
	})
}

// GetMYSQLSession func
func (config MYSQL) GetMYSQLSession() (session *xorm.Session) {
	if config.Session != nil {
		session = config.Session.NewSession()
	} else {
		panic(fmt.Errorf("[%s][%s] Chưa init MYSQL", config.Name, config.Server))
	}
	return
}