# redis-orm

![Go](https://github.com/ezbuy/redis-orm/workflows/Go/badge.svg?branch=master)

redis-orm fly orm up 

## Important NOTE

redis-orm is not maintained anymore , try to use the [ezorm's mysqlr](https://github.com/ezbuy/ezorm/blob/main/doc/schema/yaml_mysqlr.md) driver instead.


## features

## quick start

### generate command

```sh
go get github.com/ezbuy/redis-orm

redis-orm code -i example/yaml -o example/model

```

### read access usage

```go
import "github.com/ezbuy/redis-orm/example/model"

// mysql
model.MySQLSetup(cf)


db := model.MySQL()
// query (ids []string) by unique & index & range definitions
model.UserDBMgr(db).FindOne(unique)
model.UserDBMgr(db).Find(index)
model.UserDBMgr(db).Range(scope)
model.UserDBMgr(db).RangeRevert(scope)

// search usage
blogs, err := model.BlogDBMgr(db).Search("INNER JOIN users ON blog.user_id = users.user_id WHERE users.mail_box=?", "", "", "foo@bar.com")

// fetch object 
model.UserDBMgr(db).Fetch(pk PrimaryKey) (*User, error)
model.UserDBMgr(db).FetchByPrimaryKey(id int32) (*User, error)
model.UserDBMgr(db).FetchByPrimaryKeys(ids []int32) ([]*User, error)

// redis
model.RedisSetup(cf)

redis := model.Redis()
// query (ids []string) by unique & index & range definitions
model.UserRedisMgr(redis).FindOne(unique)
model.UserRedisMgr(redis).Find(index)
model.UserRedisMgr(redis).Range(scope)
model.UserRedisMgr(redis).RangeRevert(scope)

// fetch object 
model.UserRedisMgr(redis).Fetch(pk PrimaryKey) (*User, error)
model.UserRedisMgr(redis).FetchByPrimaryKeys(pks []PrimaryKey) ([]*User, error)


```

### write access usage

```go
import "github.com/ezbuy/redis-orm/example/model"

// mysql
model.MySQLSetup(cf)

db := model.MySQL()
tx, err := db.BeginTx()
defer tx.Close()

model.UserDBMgr(tx).Save(obj)
model.UserDBMgr(tx).Create(obj)
model.UserDBMgr(tx).Update(obj)
model.UserDBMgr(tx).Delete(obj)

model.UserDBMgr(tx).FindOne(unique)
model.UserDBMgr(tx).Find(index)
model.UserDBMgr(tx).Range(scope)
model.UserDBMgr(tx).RangeRevert(scope)

model.UserDBMgr(tx).Fetch(id string) (*User, error)
model.UserDBMgr(tx).FetchByPrimaryKey(id int32) (*User, error)
model.UserDBMgr(tx).FetchByPrimaryKeys(ids []int32) ([]*User, error)

// redis
model.RedisSetup(cf)

redis := model.Redis()
model.UserRedisMgr(redis).Save(obj)
model.UserRedisMgr(redis).Create(obj)
model.UserRedisMgr(redis).Update(obj)
model.UserRedisMgr(redis).Delete(obj)

```

### sync data

```go
import "github.com/ezbuy/redis-orm/example/model"

model.MySQLSetup(cf)
model.RedisSetup(cf)

db := model.MySQL()
redis := model.Redis()

model.UserRedisMgr(redis).Load(model.UserDBMgr(db))

```

## bench redis vs mysql

enviroment:
  
  mysql-server, redis-server, test client all in the same machine (mac air)

*redis-orm.redis.bench*
  
    Ran 1000 samples:
    unique.runtime:
      Fastest Time: 0.000s
      Slowest Time: 0.001s
      Average Time: 0.000s ± 0.000s
    index.runtime:
      Fastest Time: 0.000s
      Slowest Time: 0.000s
      Average Time: 0.000s ± 0.000s
    range.runtime:
      Fastest Time: 0.000s
      Slowest Time: 0.000s
      Average Time: 0.000s ± 0.000s
    range.revert.runtime:
      Fastest Time: 0.000s
      Slowest Time: 0.000s
      Average Time: 0.000s ± 0.000s
    fetch.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.004s
      Average Time: 0.002s ± 0.000s

*redis-orm.mysql.bench*
  
    Ran 1000 samples:
    unique.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.106s
      Average Time: 0.003s ± 0.005s
    index.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.106s
      Average Time: 0.003s ± 0.005s
    range.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.105s
      Average Time: 0.002s ± 0.005s
    range.revert.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.105s
      Average Time: 0.002s ± 0.006s
    fetch.runtime:
      Fastest Time: 0.004s
      Slowest Time: 0.150s
      Average Time: 0.006s ± 0.009s

