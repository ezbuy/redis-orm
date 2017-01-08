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

func (set *VSet) SortAdd(start int, items ...string) {
	for i, item := range items {
		if v, ok := set.keys.Get(item); ok {
			set.keys.Put(item, v.(int)+start+i)
			continue
		}
		set.keys.Put(item, start+i)
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

func (set *VSet) Values(start int) []string {
	result := []string{}
	it := set.keys.Iterator()
	for it.Next() {
		v := it.Value().(int)
		if v >= start {
			result = append(result, it.Key().(string))
		}
	}
	return result
}
