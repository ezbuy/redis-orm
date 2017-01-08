package model

type _BlogMgr struct {
}

var BlogMgr *_BlogMgr

func (m *_BlogMgr) NewBlog() *Blog {
	return &Blog{}
}
func (m *_BlogMgr) MySQL() *ReferenceResult {
	return NewReferenceResult(BlogMySQLMgr())
}
