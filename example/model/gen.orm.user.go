package model

type _UserMgr struct {
}

var UserMgr *_UserMgr

func (m *_UserMgr) NewUser() *User {
	return &User{}
}
func (m *_UserMgr) MySQL() *ReferenceResult {
	return NewReferenceResult(UserMySQLMgr())
}
func (m *_UserMgr) Redis() *ReferenceResult {
	return NewReferenceResult(UserRedisMgr())
}
