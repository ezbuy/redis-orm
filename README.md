# redis-orm
redis-orm fly orm up 

## features

## quick start

### generate command

````
$: go get github.com/ezbuy/redis-orm

$: redis-orm code -i example/yaml -o example/model

````

### read access usage

````
import "github.com/ezbuy/redis-orm/example/model"

# mysql

//! query (ids []string) by unique & index & range definitions
model.UserMySQLMgr().FindOne(unique)
model.UserMySQLMgr().Find(index)
model.UserMySQLMgr().Range(scope)
model.UserMySQLMgr().RevertRange(scope)

//! fetch object 
model.UserMySQLMgr().Fetch(id string) (*User, error)
model.UserMySQLMgr().FetchByIds(ids []string) ([]*User, error)

# redis

//! query (ids []string) by unique & index & range definitions
model.UserRedisMgr().FindOne(unique)
model.UserRedisMgr().Find(index)
model.UserRedisMgr().Range(scope)
model.UserRedisMgr().RevertRange(scope)

//! fetch object 
model.UserRedisMgr().Fetch(id string) (*User, error)
model.UserRedisMgr().FetchByIds(ids []string) ([]*User, error)

````

### write access usage

````
import "github.com/ezbuy/redis-orm/example/model"

# mysql

tx, _ := model.UserMySQLMgr().BeginTx()
tx.Save(obj)
tx.Create(obj)
tx.Update(obj)
tx.Delete(obj)
tx.Close()

# redis

model.UserRedisMgr().Save(obj)
model.UserRedisMgr().Create(obj)
model.UserRedisMgr().Update(obj)
model.UserRedisMgr().Delete(obj)

````

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
      Fastest Time: 0.013s
      Slowest Time: 0.046s
      Average Time: 0.018s ± 0.002s

*redis-orm.mysql.bench*
  
    Ran 1000 samples:
    unique.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.138s
      Average Time: 0.003s ± 0.009s
    index.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.017s
      Average Time: 0.002s ± 0.001s
    range.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.142s
      Average Time: 0.003s ± 0.008s
    range.revert.runtime:
      Fastest Time: 0.002s
      Slowest Time: 0.105s
      Average Time: 0.003s ± 0.007s
    fetch.runtime:
      Fastest Time: 0.004s
      Slowest Time: 0.157s
      Average Time: 0.006s ± 0.010s	

