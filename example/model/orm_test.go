package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestPeopleObject(t *testing.T) {
	MySQLSetup(&MySQLConfig{
		Host:     "localhost",
		Port:     3306,
		UserName: "ezorm_user",
		Password: "ezorm_pass",
		Database: "ezorm",
	})

	now := time.Now()
	user1 := UserMgr.NewUser()
	user1.Name = "user01"
	user1.Mailbox = "user01@sss.fff"
	user1.HeadUrl = "aaaa.png"
	user1.Password = "123456"
	user1.CreatedAt = now
	user1.UpdatedAt = now
	user1.Longitude = 103.754
	user1.Latitude = 1.3282

	tx, err := UserMySQLMgr().BeginTx()
	assert.Equal(t, nil, err)
	defer tx.Close()

	assert.Equal(t, tx.Save(user1), err)
	user1.HeadUrl = "bbb.png"
	assert.Equal(t, tx.Save(user1), err)

	unique := &MailboxPasswordOfUserUnique{
		Mailbox:  "user01@sss.fff",
		Password: "123456",
	}

	id, err := UserMySQLMgr().FindOne(unique)
	assert.Equal(t, nil, err)

	user, err := UserMySQLMgr().Fetch(id)
	assert.Equal(t, nil, err)
	fmt.Println("user => ", user)

	blogTx, err := BlogMySQLMgr().BeginTx()
	assert.Equal(t, nil, err)

	blog11 := Blog{
		UserId:    user1.Id,
		Title:     "BlogTitle1",
		Content:   "hello! everybody!!!",
		Status:    1,
		Readed:    10,
		CreatedAt: now,
		UpdatedAt: now,
	}
	assert.Equal(t, blogTx.Save(&blog11), err)

	blog12 := Blog{
		UserId:    user1.Id,
		Title:     "BlogTitle1222",
		Content:   "hello! everybody!!!",
		Status:    1,
		Readed:    10,
		CreatedAt: now,
		UpdatedAt: now,
	}
	assert.Equal(t, blogTx.Save(&blog12), err)

	blogTx.Close()

	index := &UserIdOfBlogIndex{
		UserId: user1.Id,
	}

	vals, err := BlogMgr.MySQL().Find(index).Result()
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(vals))
}
