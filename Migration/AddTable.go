package Migration

import "github.com/qiankaihua/ginDemo/Model"

func AddTable()  {
	InitMigration()
	AddMigration("test", &Model.Test{})
}
