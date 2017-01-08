package model

type _UserBaseInfoMgr struct {
}

var UserBaseInfoMgr *_UserBaseInfoMgr

func (m *_UserBaseInfoMgr) NewUserBaseInfo() *UserBaseInfo {
	return &UserBaseInfo{}
}
func (m *_UserBaseInfoMgr) MySQL() *ReferenceResult {
	return NewReferenceResult(UserBaseInfoMySQLMgr())
}
