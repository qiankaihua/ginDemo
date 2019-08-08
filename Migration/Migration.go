package Migration

import (
	"fmt"
	"github.com/qiankaihua/ginDemo/Boot/Orm"
)

var Migration *migration

type migration struct {
	_mode map[string]interface{}
}

func InitMigration()  {
	Migration = new(migration)
	Migration._mode = make(map[string]interface{})
}

func AddMigration(tableName string, model interface{})  {
	Migration._mode[tableName] = model
}

func Fresh()  {
	db := Orm.GetDB()
	for key, value := range Migration._mode {
		if db.HasTable(key) {
			fmt.Println("table", key, "is existed")
		} else {
			db.Table(key).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(value)
			fmt.Println("table", key, "has been created successfully")
		}
	}
}

func Refresh()  {
	db := Orm.GetDB()
	for key, value := range Migration._mode {
		db.DropTableIfExists(key)
		fmt.Println("table", key, "has been dropped successfully")
		db.Table(key).Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(value)
		fmt.Println("table", key, "has been created successfully")
	}
}