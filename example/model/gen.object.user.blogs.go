package model



import (
	"fmt"
	"time"
	"strings"
	"database/sql"
	"github.com/ezbuy/redis-orm/orm"
)
var (
	_ sql.DB
	_ time.Time
	_ fmt.Formatter
	_ strings.Reader
	_ orm.VSet
)

	
	
		type UserBlogs struct {
			UserId  int32 `db:"user_id"`
			BlogId  int32 `db:"blog_id"`
		}
	

	

	

	

	

	type _UserBlogsMgr struct {
	}
	var UserBlogsMgr *_UserBlogsMgr

	func (m *_UserBlogsMgr) NewUserBlogs() *UserBlogs {
		return &UserBlogs{}
	}

	//! object function
	


func (obj *UserBlogs) GetNameSpace() string {
	return "model"
}

func (obj *UserBlogs) GetClassName() string {
	return "UserBlogs"
}

func (obj *UserBlogs) GetTableName() string {
	return "user_blogs"	
}

func (obj *UserBlogs) GetColumns() []string {
	columns := []string{
		"`user_id`",
		"`blog_id`",
	}
	return columns
}

	//! uniques

	//! indexes

	//! ranges
	func (m *_UserBlogsMgr) MySQL() *ReferenceResult {
		return NewReferenceResult(UserBlogsMySQLMgr())
	}
	



type _UserBlogsMySQLMgr struct {
	*orm.MySQLStore
}

func UserBlogsMySQLMgr() *_UserBlogsMySQLMgr {
	return &_UserBlogsMySQLMgr{_mysql_store}
}

func NewUserBlogsMySQLMgr(cf *MySQLConfig) (*_UserBlogsMySQLMgr, error) {
	store, err := orm.NewMySQLStore(cf.Host, cf.Port, cf.Database, cf.UserName, cf.Password)
	if err != nil {
		return nil, err
	}
	return &_UserBlogsMySQLMgr{store}, nil
}


func (m *_UserBlogsMySQLMgr) Search(where string, args ...interface{}) ([]*UserBlogs, error) {
	obj := UserBlogsMgr.NewUserBlogs()
	if where != "" {
	 	where = " WHERE " + where
	}
	query := fmt.Sprintf("SELECT %s FROM `user_blogs` %s", strings.Join(obj.GetColumns(), ","), where)
	objs, err := m.FetchBySQL(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]*UserBlogs, 0, len(objs))
	for _, obj := range objs {
		results = append(results, obj.(*UserBlogs))
	}
	return results, nil
}

func (m *_UserBlogsMySQLMgr) SearchCount(where string, args ...interface{}) (int64, error){ 
	if where != "" {
		where = " WHERE " + where
	}
	return m.queryCount(where, args...)
}

func (m *_UserBlogsMySQLMgr) FetchBySQL(q string, args ... interface{}) (results []interface{}, err error) {
	rows, err := m.Query(q, args...)	
	if err != nil {
		return nil, fmt.Errorf("UserBlogs fetch error: %v", err)
	}
	defer rows.Close()

	

	for rows.Next() {
		var result UserBlogs
		err = rows.Scan(&(result.UserId),&(result.BlogId),)
		if err != nil {
			return nil, err
		}

		
		
		

		results = append(results, &result)
	}
	if err = rows.Err() ;err != nil {
		return nil, fmt.Errorf("UserBlogs fetch result error: %v", err)
	}
	return
}
func (m *_UserBlogsMySQLMgr) Fetch(