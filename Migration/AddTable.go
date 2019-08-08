package Migration

import "github.com/qiankaihua/ginDemo/Model"

func AddTable()  {
	InitMigration()
	test := &Model.Test{}
	AddMigration("test", test)
}
