package model_test

import (
	"fmt"
	"time"

	. "github.com/ezbuy/redis-orm/example/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("redis-orm.mysql", func() {
	BeforeEach(func() {
		MySQLSetup(&MySQLConfig{
			Host:     "localhost",
			Port:     3306,
			UserName: "ezorm_user",
			Password: "ezorm_pass",
			Database: "ezorm",
		})
	})

	Describe("CRUD", func() {
		It("user crud test", func() {
			user := UserMgr.NewUser()
			user.Name = "user01"
			user.Mailbox = "user01@sss.fff"
			user.HeadUrl = "aaaa.png"
			user.Password = "123456"
			user.CreatedAt = time.Now()
			user.UpdatedAt = user.CreatedAt
			user.Longitude = 103.754
			user.Latitude = 1.3282
			tx, err := UserMySQLMgr().BeginTx()
			Ω(err).ShouldNot(HaveOccurred())
			defer tx.Close()
			//! debug sql
			// tx.Debug(true)

			//! create
			Ω(tx.Create(user)).ShouldNot(HaveOccurred())

			//! update
			user.HeadUrl = "bbbb.png"
			user.UpdatedAt = time.Now()
			Ω(tx.Update(user)).ShouldNot(HaveOccurred())

			//! fetch check
			obj, err := tx.Fetch(user.Id)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(obj.HeadUrl).To(Equal(user.HeadUrl))

			//! delete
			Ω(tx.Delete(obj)).ShouldNot(HaveOccurred())

			//! fetch check
			_, err = tx.Fetch(user.Id)
			Ω(err).Should(HaveOccurred())

			//! save
			Ω(tx.Save(user)).ShouldNot(HaveOccurred())
			Ω(tx.Save(user)).Should(HaveOccurred())
			Ω(tx.Delete(user)).ShouldNot(HaveOccurred())
		})
	})

	Describe("Finder", func() {
		ids := []string{}
		BeforeEach(func() {
			tx, err := UserMySQLMgr().BeginTx()
			Ω(err).ShouldNot(HaveOccurred())
			defer tx.Close()

			for i := 0; i < 100; i++ {
				user := UserMgr.NewUser()
				user.Name = fmt.Sprintf("name%d", i)
				user.Mailbox = fmt.Sprintf("name%d@ezbuy.com", i)
				user.HeadUrl = fmt.Sprintf("name%d.png", i)
				user.Password = fmt.Sprintf("pwd%d", i)
				if i%2 == 0 {
					user.Sex = true
				} else {
					user.Sex = false
				}
				user.Age = int32(i)
				user.CreatedAt = time.Now()
				user.UpdatedAt = user.CreatedAt
				user.Longitude = 103.754
				user.Latitude = 1.3282
				Ω(tx.Save(user)).ShouldNot(HaveOccurred())
				ids = append(ids, fmt.Sprint(user.Id))
			}
			//! debug sql
			UserMySQLMgr().Debug(true)
		})
		AfterEach(func() {
			tx, err := UserMySQLMgr().BeginTx()
			Ω(err).ShouldNot(HaveOccurred())
			defer tx.Close()
			Ω(tx.DeleteByIds(ids)).ShouldNot(HaveOccurred())
		})

		It("unique", func() {
			unique := &MailboxPasswordOfUserUK{
				Mailbox:  "name20@ezbuy.com",
				Password: "pwd20",
			}
			obj, err := UserMySQLMgr().FindOne(unique)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(obj).ShouldNot(BeNil())
		})

		It("index", func() {
			sexIdx := &SexOfUserIDX{
				Sex: false,
			}
			us, err := UserMySQLMgr().Find(sexIdx)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(50))
		})

		It("range", func() {
			scope := &AgeOfUserRNG{
				AgeBegin: 10,
				AgeEnd:   35,
			}
			us, err := UserMySQLMgr().Range(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(24))
			fmt.Println("range us =>", us)
			Ω(us[1] > us[0]).To(Equal(true))
		})

		It("range.revert", func() {
			scope := &AgeOfUserRNG{
				AgeBegin: 10,
				AgeEnd:   35,
			}
			us, err := UserMySQLMgr().RevertRange(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(24))
			fmt.Println("range.revert us =>", us)
			Ω(us[1] > us[0]).To(Equal(false))
		})
	})

})

var _ = Describe("redis-orm.redis", func() {
	BeforeEach(func() {
		RedisSetUp(&RedisConfig{
			Host:     "localhost",
			Port:     6379,
			Password: "",
		})
	})

	Describe("create", func() {
		It("should be a novel", func() {
			var err error
			err = nil

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("update", func() {
		It("should be a novel", func() {
			var err error
			err = nil

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("delete", func() {
		It("should be a novel", func() {
			var err error
			err = nil

			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("save", func() {
		It("should be a novel", func() {
			var err error
			err = nil

			Ω(err).ShouldNot(HaveOccurred())
		})
	})
})
