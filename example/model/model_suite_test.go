package model_test

import (
	"testing"

	. "github.com/ezbuy/redis-orm/example/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	MySQLSetup(&MySQLConfig{
		Host:     "localhost",
		Port:     3306,
		UserName: "ezorm_user",
		Password: "ezorm_pass",
		Database: "ezorm",
	})

	MySQL().Debug(true)
}

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "redis-orm")
}

/*
SELECT blogs.`id`,blogs.`user_id`,blogs.`title`,blogs.`content`,blogs.`status`,blogs.`readed`,blogs.`created_at`,blogs.`updated_at` FROM blogs INNER JOIN users ON blog.user_id = users.user_id WHERE users.mail_box=?   [test@ezbuy.com]
*/
func TestSearch(t *testing.T) {
	blogs, err := BlogDBMgr(MySQL()).Search("INNER JOIN users ON blog.user_id = users.user_id WHERE users.mail_box=?", "", "", "test@ezbuy.com")
	_, _ = blogs, err
}
