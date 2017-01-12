package orm

import rbt "github.com/emirpasic/gods/trees/redblacktree"

type VSet struct {
	keys *rbt.Tree
}

func NewVSet() *VSet {
	set := new(VSet)
	set.keys = rbt.NewWithStringComparator()
	return set
}

func (set *VSet) Add(val int, items ...string) {
	for _, item := range items {
		if v, ok := set.keys.Get(item); ok {
			set.keys.Put(item, v.(int)+val)
			continue
		}
		set.keys.Put(item, val)
	}
}

func (set *VSet) SortAdd(val int, items ...string) {
	for _, item := range items {
		if v, ok := set.keys.Get(item); ok {
			set.keys.Put(item, v.(int)+val)
			continue
		}
		set.keys.Put(item, val)
	}
}

func (set *VSet) Remove(items ...string) {
	for _, item := range items {
		set.keys.Remove(item)
	}
}

func (set *VSet) Clear() {
	set.keys.Clear()
}

func (set *VSet) Values(start int, offset int, limit int) []string {
	result := []string{}
	if limit == 0 {
		return result
	}

	ioffset := 0
	ilimit := 0
	it := set.keys.Iterator()
	for it.Next() {
		v := it.Value().(int)
		if v >= start {
			if ioffset < offset {
				continue
			}
			ioffset++
			if limit > 0 && ilimit >= limit {
				break
			}
			ilimit++
			result = append(result, it.Key().(string))
		}
	}
	return result
}

func (set *VSet) SortValues(start int, offset int, limit int) []string {
	result := []string{}
	if limit == 0 {
		return result
	}

	ioffset := 0
	ilimit := 0
	it := set.keys.Iterator()
	for it.Next() {
		v := it.Value().(int)
		if v >= start {
			if ioffset < offset {
				continue
			}
			ioffset++
			if limit > 0 && ilimit >= limit {
				break
			}
			ilimit++
			result = append(result, it.Key().(string))
		}
	}
	return result
}
