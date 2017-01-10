package model_test

// import (
// 	. "github.com/ezbuy/redis-orm/example/model"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("MySQL", func() {
// 	user := UserMgr.NewUser()

// 	BeforeEach(func() {
// 		MySQLSetup(&MySQLConfig{
// 			Host:     "localhost",
// 			Port:     3306,
// 			UserName: "ezorm_user",
// 			Password: "ezorm_pass",
// 			Database: "ezorm",
// 		})
// 	})

// 	Describe("CREATE", func() {
// 		tx, err := UserMySQLMgr().BeginTx()
// 		defer tx.Close()
// 		Ω(err).ShouldNot(HaveOccurred())
// 		Ω(tx.Save(user)).ShouldNot(HaveOccurred())
// 	})
// })
