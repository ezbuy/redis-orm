# redis-orm
redis-orm fly orm up 

## standard yaml definition

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
    - storetype: pair
    - storetype: int 
    - storetype: ReferenceModelName
  importSQL: 'select fields... from table'

````

please check the example yaml.


## Generate code from yaml

````
$: go get github.com/ezbuy/redis-orm

$: redis-orm code -i example/yamls -o example/model

````

## How to use codes from project


### How to use MySQL orm

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

	

