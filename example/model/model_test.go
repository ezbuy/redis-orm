package model_test

import (
	"fmt"
	"time"

	"github.com/ezbuy/redis-orm/orm"

	. "github.com/ezbuy/redis-orm/example/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("manager", func() {
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

	It("vset", func() {
		vset := orm.NewVSet()
		Ω(vset).ShouldNot(BeNil())

		vset.Add(101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116)
		Ω(vset.Values(1, 0, 6)).Should(Equal([]interface{}{101, 102, 103, 104, 105, 106}))
		Ω(vset.Values(1, 6, 6)).Should(Equal([]interface{}{107, 108, 109, 110, 111, 112}))
		Ω(vset.Values(1, 12, -1)).Should(Equal([]interface{}{113, 114, 115, 116}))
		vset.SortAdd(2, 103, 102, 105, 108, 113, 112)
		Ω(vset.Values(2, 0, 2)).Should(Equal([]interface{}{103, 102}))
		Ω(vset.Values(2, 2, 2)).Should(Equal([]interface{}{105, 108}))
		Ω(vset.Values(2, 4, -1)).Should(Equal([]interface{}{113, 112}))
	})

	It("rr", func() {
		scope1 := &IdOfUserRNG{}
		us1, err := UserMgr.MySQL().Range(scope1).Result()
		Ω(err).ShouldNot(HaveOccurred())
		Ω(len(us1)).To(Equal(100))

		sexIdx := &SexOfUserIDX{
			Sex: false,
		}
		scope2 := &AgeOfUserRNG{
			AgeBegin: 10,
			AgeEnd:   35,
		}
		us2, err := UserMgr.Redis().Find(sexIdx).Range(scope2).Result()
		Ω(err).ShouldNot(HaveOccurred())
		Ω(len(us2)).To(Equal(12))
	})
})

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

		Measure("mysql.bench", func(b Benchmarker) {
			b.Time("crud.runtime", func() {
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
		}, 1)
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
			Ω(us[1].(int32) > us[0].(int32)).To(Equal(true))
		})

		It("range.revert", func() {
			scope := &AgeOfUserRNG{}
			us, err := UserMySQLMgr().RangeRevert(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(100))
			Ω(us[1].(int32) > us[0].(int32)).To(Equal(false))
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
			objs2, err := UserMySQLMgr().RangeFetch(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(objs2)).To(Equal(100))
			for i, obj := range objs2 {
				Ω(obj.Name).To(Equal(fmt.Sprintf("name%d", i)))
			}
		})

		It("search", func() {
			us, err := UserMySQLMgr().Search("age < 50 and sex = 1")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(25))

			cnt, err := UserMySQLMgr().SearchCount("age < 50 and sex = 1")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(cnt).To(Equal(int64(25)))
		})

		Measure("mysql.bench", func(b Benchmarker) {
			b.Time("unique.runtime", func() {
				unique := &MailboxPasswordOfUserUK{
					Mailbox:  "name20@ezbuy.com",
					Password: "pwd20",
				}
				obj, err := UserMySQLMgr().FindOne(unique)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(obj).ShouldNot(BeNil())
			})
			b.Time("index.runtime", func() {
				sexIdx := &SexOfUserIDX{
					Sex: false,
				}
				us, err := UserMySQLMgr().Find(sexIdx)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(len(us)).To(Equal(50))
			})
			b.Time("range.runtime", func() {
				scope := &AgeOfUserRNG{
					AgeBegin: 10,
					AgeEnd:   35,
				}
				us, err := UserMySQLMgr().Range(scope)
				Ω(err).ShouldNot(HaveOccurred())
				UserMySQLMgr().Debug(true)
				count, err := UserMySQLMgr().RangeCount(scope)
				fmt.Println("err=>", err)
				Ω(len(us)).To(Equal(int(count)))
				Ω(us[1].(int32) > us[0].(int32)).To(Equal(true))
			})
			b.Time("range.revert.runtime", func() {
				scope := &AgeOfUserRNG{}
				us, err := UserMySQLMgr().RangeRevert(scope)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(len(us)).To(Equal(100))
				Ω(us[1].(int32) > us[0].(int32)).To(Equal(false))
			})
			b.Time("fetch.runtime", func() {
				scope := &IdOfUserRNG{}
				us, err := UserMySQLMgr().Range(scope)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(len(us)).To(Equal(100))
				objs, err := UserMySQLMgr().FetchByIds(us)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(len(objs)).To(Equal(100))
			})
		}, 1)
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
			Ω(UserRedisMgr().Clear()).ShouldNot(HaveOccurred())
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

		Measure("redis.bench", func(b Benchmarker) {
			b.Time("crud.runtime", func() {
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

				user.Age = int32(40)
				Ω(UserRedisMgr().Update(user)).ShouldNot(HaveOccurred())
				obj, err = UserRedisMgr().Fetch(fmt.Sprint(user.Id))
				Ω(err).ShouldNot(HaveOccurred())
				Ω(obj.Age).To(Equal(int32(40)))

				Ω(UserRedisMgr().Delete(user)).ShouldNot(HaveOccurred())
				_, err = UserRedisMgr().Fetch(fmt.Sprint(user.Id))
				Ω(err).Should(HaveOccurred())
			})
		}, 1)
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
			count, err := UserRedisMgr().FindCount(sexIdx)
			Ω(len(us)).To(Equal(int(count)))
			Ω(len(us)).To(Equal(50))
		})
		It("range", func() {
			scope := &AgeOfUserRNG{
				AgeBegin: 10,
				AgeEnd:   35,
			}
			us, err := UserRedisMgr().Range(scope)
			Ω(err).ShouldNot(HaveOccurred())
			count, err := UserRedisMgr().RangeCount(scope)
			Ω(len(us)).To(Equal(int(count)))
			Ω(us[1].(int32) > us[0].(int32)).To(Equal(true))
		})
		It("range.revert", func() {
			scope := &AgeOfUserRNG{}
			us, err := UserRedisMgr().RangeRevert(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(100))
			Ω(us[1].(int32) > us[0].(int32)).To(Equal(false))
		})

		It("fetch", func() {
			scope := &IdOfUserRNG{}
			us, err := UserRedisMgr().Range(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(us)).To(Equal(100))
			objs, err := UserRedisMgr().FetchByIds(us)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(objs)).To(Equal(100))
			for i, obj := range objs {
				Ω(obj.Name).To(Equal(fmt.Sprintf("name%d", i)))
			}
			objs2, err := UserRedisMgr().RangeFetch(scope)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(len(objs2)).To(Equal(100))
			for i, obj := range objs2 {
				Ω(obj.Name).To(Equal(fmt.Sprintf("name%d", i)))
			}
		})

		Measure("redis.bench", func(b Benchmarker) {
			b.Time("unique.runtime", func() {
				unique := &MailboxPasswordOfUserUK{
					Mailbox:  "name20@ezbuy.com",
					Password: "pwd20",
				}
				obj, err := UserRedisMgr().FindOne(unique)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(obj).ShouldNot(BeNil())
			})
			b.Time("index.runtime", func() {
				sexIdx := &SexOfUserIDX{
					Sex: false,
				}
				us, err := UserRedisMgr().Find(sexIdx)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(len(us)).To(Equal(50))
			})
			b.Time("range.runtime", func() {
				scope := &AgeOfUserRNG{
					AgeBegin: 10,
					AgeEnd:   35,
				}
				us, err := UserRedisMgr().Range(scope)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(len(us)).To(Equal(24))
				Ω(us[1].(int32) > us[0].(int32)).To(Equal(true))
			})
			b.Time("range.revert.runtime", func() {
				scope := &AgeOfUserRNG{}
				us, err := UserRedisMgr().RangeRevert(scope)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(len(us)).To(Equal(100))
				Ω(us[1].(int32) > us[0].(int32)).To(Equal(false))
			})
			b.Time("fetch.runtime", func() {
				scope := &AgeOfUserRNG{}
				us, err := UserRedisMgr().RangeRevert(scope)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(len(us)).To(Equal(100))
				objs, err := UserRedisMgr().FetchByIds(us)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(len(objs)).To(Equal(100))
			})
		}, 1)
	})
})
