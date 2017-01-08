# redis-orm
redis-orm fly orm up 

*注意* 版本开发中，请勿在生产环境中使用

## 标准Yaml格式定义

````yaml

ModelName:
  dbs: [mysql, mssql, mongo, redis, elastic]
  dbname: DBName
  dbtable: TableName
  dbview: ViewName
  fields:
    - FieldName1:
      flags: [primary, autoinc, noinc, nullable, unique, index, range, order, fulltext]
      attrs: []
    - FieldName2:
      flags: [primary, autoinc, noinc, nullable, unique, index, range, order, fulltext]
      attrs: []	
  uniques: [[FieldName1, ..., FieldNameN],[FieldName1, ..., FieldNameM]]
  indexes: [[FieldName1, ..., FieldNameN],[FieldName1, ..., FieldNameM]]
  ranges: [[FieldName1, ..., RangeFieldName],[FieldName1, ..., RangeFieldName]]
  orders: [[FieldName1, ..., FieldNameN],[FieldName1, ..., FieldNameM]]
  relation:
    - storetype: pair | set | zset | geo | list
    - valuetype: int32 
    - modeltype: ReferenceModelName
  importSQL: 'select fields... from table'

````
具体使用参见[样例](example/yaml/user.yaml)

## 代码生成

````
$: go get github.com/ezbuy/redis-orm

$: redis-orm code -i example/yaml -o example/model

````

## 快速开始


### MySQL ORM的使用

````
import "github.com/ezbuy/redis-orm/example/model"

model.MySQLSetup(cf)

//! read access

//! query (ids []string) by unique & index & range & order definitions
model.UserMySQLMgr().FindOne(unique)
model.UserMySQLMgr().Find(index)
model.UserMySQLMgr().Range(scope)
model.UserMySQLMgr().OrderBy(sort)

//! fetch object 
model.UserMySQLMgr().Fetch(id string) (*User, error)
model.UserMySQLMgr().FetchByIds(ids []string) ([]*User, error)

//! write access
tx, _ := model.UserMySQLMgr().BeginTx()
tx.Create(obj)
tx.Update(obj)
tx.Delete(obj)
tx.Close()

//! high level manager op
model.UserMgr.MySQL().FindOne(unique)
model.UserMgr.MySQL().Find(index)
model.UserMgr.MySQL().Range(scope)
model.UserMgr.MySQL().OrderBy(sort)

````

	

