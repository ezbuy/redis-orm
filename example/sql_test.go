package example

import (
	"testing"

	"github.com/ezbuy/redis-orm/example/model"
)

func init() {
	model.MySQL().Debug(true)
	model.MySQLSetup(&model.MySQLConfig{
		Host:            "127.0.0.1",
		Port:            3306,
		UserName:        "u",
		Password:        "pwd",
		Database:        "db",
		PoolSize:        8,
		ConnMaxLifeTime: 5000000000,
	})
}

func TestSql(t *testing.T) {
	model.BlogDBMgr(model.MySQL()).FindAllByStatus(1)

	model.BlogDBMgr(model.MySQL()).FindAllByUserIdTitle(101, "Title101")

	model.BlogDBMgr(model.MySQL()).DeleteByPrimaryKey(101, 1001)

	blog := &model.Blog{
		Id:     102,
		UserId: 1002,
	}
	model.BlogDBMgr(model.MySQL()).Delete(blog)
}
