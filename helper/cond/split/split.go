// Package split provides `IN` cond help functions
package split

// Splittor defines in condition splittor
type Splittor interface {
	Size() int
}

// Split splits the slice into the slice of the slice
func Split(src []interface{}, splittor Splittor) (res [][]interface{}) {
	if splittor.Size() == 0 {
		return
	}
	tmp := make([]interface{}, 0, splittor.Size())
	for _, s := range src {
		tmp = append(tmp, s)
		if len(tmp) == splittor.Size() {
			res = append(res, tmp)
			tmp = make([]interface{}, 0, splittor.Size())
		}
	}
	if len(tmp) != 0 {
		res = append(res, tmp)
	}
	return
}

// DefaultSplittor defines the splittor
type DefaultSplittor struct {
	size int
}

// NewSplittor new the splittor
func NewSplittor(size int) DefaultSplittor {
	return DefaultSplittor{
		size: size,
	}
}

// Size impl the Splittor
func (ds DefaultSplittor) Size() int {
	return ds.size
}

// Split splits the src
func (ds DefaultSplittor) Split(src []interface{}) (res [][]interface{}) {
	return Split(src, ds)
}
