package orm

import rbt "github.com/emirpasic/gods/trees/redblacktree"
import dlt "github.com/emirpasic/gods/lists/doublylinkedlist"

type VSet struct {
	keys  *rbt.Tree
	sorts map[int]*dlt.List
}

func NewVSet() *VSet {
	set := new(VSet)
	set.keys = rbt.NewWithIntComparator()
	set.sorts = make(map[int]*dlt.List)
	return set
}

func (set *VSet) Add(val int, items ...interface{}) {
	for _, item := range items {
		if v, ok := set.keys.Get(item); ok {
			set.keys.Put(item, v.(int)+val)
			continue
		}
		set.keys.Put(item, val)
	}
}

func (set *VSet) SortAdd(val int, items ...interface{}) {
	for _, item := range items {
		if v, ok := set.keys.Get(item); ok {
			set.keys.Put(item, v.(int)+val)
			continue
		}
		set.keys.Put(item, val)
	}
	if list, ok := set.sorts[val]; ok {
		list.Add(items...)
	} else {
		lt := dlt.New()
		lt.Append(items...)
		set.sorts[val] = lt
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

func (set *VSet) Values(start int, offset int, limit int) []interface{} {
	result := make([]interface{}, 0, limit)
	if limit == 0 {
		return result
	}

	//! sorted
	if list, ok := set.sorts[start]; ok {
		list.Each(func(index int, value interface{}) {
			if v, ok := set.keys.Get(value); ok {
				if v.(int) >= start {
					result = append(result, value)
				}
			}
		})
		return result
	}

	//! normal
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
			result = append(result, it.Key())
		}
	}
	return result
}
