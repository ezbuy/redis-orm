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
			tx.Debug(false)

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
			user.HeadUrl = "ccc.png"
			Ω(tx.Save(user)).ShouldNot(HaveOccurred())
			Ω(tx.Delete(user)).ShouldNot(HaveOccurred())
		})
	})

	Describe("Finder", func() {

		BeforeEach(func() {
			tx, err := UserMySQLMgr().BeginTx()
			Ω(err).ShouldNot(HaveOccurred())
			defer tx.Close()
			users := []*User{}
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
				users = append(users, user)
			}
			//! debug sql
			UserMySQLMgr().Debug(false)
			Ω(tx.BatchCreate(users)).ShouldNot(HaveOccurred())

		})
		AfterEach(func() {
			tx, err := UserMySQLMgr().BeginTx()
			Ω(err).ShouldNot(HaveOccurred())
			defer tx.Close()

			scope := &IdOfUserRNG{}
			us, err := UserMySQLMgr().Range(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(tx.DeleteByIds(us)).ShouldNot(HaveOccurred())
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
			Ω(us[1] > us[0]).To(Equal(true))
		})

		It("range.revert", func() {
			scope := &AgeOfUserRNG{}
			us, err := UserMySQLMgr().RevertRange(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(100))
			Ω(us[1] > us[0]).To(Equal(false))
		})

		It("fetch", func() {
			scope := &IdOfUserRNG{}
			us, err := UserMySQLMgr().Range(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(100))
			objs, err := UserMySQLMgr().FetchByIds(us)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(objs)).To(Equal(100))
			for i, obj := range objs {
				Ω(obj.Name).To(Equal(fmt.Sprintf("name%d", i)))
			}
		})

	})

})

var _ = Describe("redis-orm.redis", func() {

	MySQLSetup(&MySQLConfig{
		Host:     "localhost",
		Port:     3306,
		UserName: "ezorm_user",
		Password: "ezorm_pass",
		Database: "ezorm",
	})

	RedisSetUp(&RedisConfig{
		Host:     "localhost",
		Port:     6379,
		Password: "",
	})

	BeforeEach(func() {
		tx, err := UserMySQLMgr().BeginTx()
		Ω(err).ShouldNot(HaveOccurred())
		defer tx.Close()
		users := []*User{}
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
			users = append(users, user)
		}
		//! debug sql
		UserMySQLMgr().Debug(false)
		Ω(tx.BatchCreate(users)).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		tx, err := UserMySQLMgr().BeginTx()
		Ω(err).ShouldNot(HaveOccurred())
		defer tx.Close()

		scope := &IdOfUserRNG{}
		us, err := UserMySQLMgr().Range(scope)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(tx.DeleteByIds(us)).ShouldNot(HaveOccurred())
	})

	Describe("load", func() {
		It("mysql => redis", func() {
			Ω(UserRedisMgr().Load(UserMySQLMgr())).ShouldNot(HaveOccurred())
		})
	})

	Describe("crud", func() {
		var user *User
		It("create", func() {
			user = UserMgr.NewUser()
			user.Id = 101
			user.Name = fmt.Sprintf("name%d", 101)
			user.Mailbox = fmt.Sprintf("name%d@ezbuy.com", 101)
			user.HeadUrl = fmt.Sprintf("name%d.png", 101)
			user.Password = fmt.Sprintf("pwd%d", 101)
			user.Sex = true
			user.Age = int32(32)
			user.CreatedAt = time.Now()
			user.UpdatedAt = user.CreatedAt
			user.Longitude = 103.754
			user.Latitude = 1.3282
			Ω(UserRedisMgr().Create(user)).ShouldNot(HaveOccurred())

			obj, err := UserRedisMgr().Fetch(fmt.Sprint(user.Id))
			Ω(err).ShouldNot(HaveOccurred())
			Ω(obj.Name).To(Equal(fmt.Sprintf("name%d", 101)))
		})
		It("update", func() {
			user.Age = int32(40)
			Ω(UserRedisMgr().Update(user)).ShouldNot(HaveOccurred())
			obj, err := UserRedisMgr().Fetch(fmt.Sprint(user.Id))
			Ω(err).ShouldNot(HaveOccurred())
			Ω(obj.Age).To(Equal(int32(40)))
		})
		It("delete", func() {
			Ω(UserRedisMgr().Delete(user)).ShouldNot(HaveOccurred())
			_, err := UserRedisMgr().Fetch(fmt.Sprint(user.Id))
			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("finder", func() {
		It("unique", func() {
			unique := &MailboxPasswordOfUserUK{
				Mailbox:  "name20@ezbuy.com",
				Password: "pwd20",
			}
			obj, err := UserRedisMgr().FindOne(unique)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(obj).ShouldNot(BeNil())
		})
		It("index", func() {
			sexIdx := &SexOfUserIDX{
				Sex: false,
			}
			us, err := UserRedisMgr().Find(sexIdx)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(50))
		})
		It("range", func() {
			scope := &AgeOfUserRNG{
				AgeBegin: 10,
				AgeEnd:   35,
			}
			us, err := UserRedisMgr().Range(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(24))
			Ω(us[1] > us[0]).To(Equal(true))
		})
		It("range.revert", func() {
			scope := &AgeOfUserRNG{}
			us, err := UserRedisMgr().RevertRange(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(100))
			Ω(us[1] > us[0]).To(Equal(false))
		})
	})
})
