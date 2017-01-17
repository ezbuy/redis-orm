package orm

import (
	"reflect"

	dlt "github.com/emirpasic/gods/lists/doublylinkedlist"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
	utils "github.com/emirpasic/gods/utils"
)

func PrimaryComparator(a, b interface{}) int {
	switch reflect.ValueOf(a).Kind() {
	case reflect.String:
		return utils.StringComparator(a, b)
	case reflect.Int:
		return utils.IntComparator(a, b)
	case reflect.Int8:
		return utils.Int8Comparator(a, b)
	case reflect.Int16:
		return utils.Int16Comparator(a, b)
	case reflect.Int32:
		return utils.Int32Comparator(a, b)
	case reflect.Int64:
		return utils.Int64Comparator(a, b)
	case reflect.Uint:
		return utils.UIntComparator(a, b)
	case reflect.Uint8:
		return utils.UInt8Comparator(a, b)
	case reflect.Uint16:
		return utils.UInt16Comparator(a, b)
	case reflect.Uint32:
		return utils.UInt32Comparator(a, b)
	case reflect.Uint64:
		return utils.UInt64Comparator(a, b)
	case reflect.Float32:
		return utils.Float32Comparator(a, b)
	case reflect.Float64:
		return utils.Float64Comparator(a, b)
	}
	return -1
}

type VSet struct {
	keys  *rbt.Tree
	sorts map[int]*dlt.List
}

func NewVSet() *VSet {
	set := new(VSet)
	set.keys = rbt.NewWith(PrimaryComparator)
	set.sorts = make(map[int]*dlt.List)
	return set
}

func (set *VSet) Add(items ...interface{}) {
	for _, item := range items {
		if v, ok := set.keys.Get(item); ok {
			set.keys.Put(item, v.(int)+1)
			continue
		}
		set.keys.Put(item, 1)
	}
}

func (set *VSet) SortAdd(times int, items ...interface{}) {
	for _, item := range items {
		if v, ok := set.keys.Get(item); ok {
			set.keys.Put(item, v.(int)+1)
			continue
		}
		set.keys.Put(item, 1)
	}
	if list, ok := set.sorts[times]; ok {
		list.Add(items...)
	} else {
		lt := dlt.New()
		lt.Append(items...)
		set.sorts[times] = lt
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

func (set *VSet) Unions(offset int, limit int) []interface{} {
	result := make([]interface{}, 0, limit)
	if limit == 0 {
		return result
	}
	//! normal
	ioffset := 0
	ilimit := 0
	it := set.keys.Iterator()
	for it.Next() {
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
	return result
}

func (set *VSet) Values(times int, offset int, limit int) []interface{} {
	result := []interface{}{}
	if limit == 0 {
		return result
	}

	//! pos
	ioffset := 0
	ilimit := 0

	//! sorted
	if list, ok := set.sorts[times]; ok {
		list.Each(func(index int, value interface{}) {
			if v, ok := set.keys.Get(value); ok {
				if v.(int) >= times {
					if ioffset < offset {
						ioffset++
						return
					}
					ioffset++
					if limit > 0 && ilimit >= limit {
						ilimit++
						return
					}
					ilimit++
					result = append(result, value)

				}
			}
		})
		return result
	}

	//! normal
	it := set.keys.Iterator()
	for it.Next() {
		v := it.Value().(int)
		if v >= times {
			if ioffset < offset {
				ioffset++
				continue
			}
			ioffset++
			if limit > 0 && ilimit >= limit {
				ilimit++
				break
			}
			ilimit++
			result = append(result, it.Key())
		}
	}
	return result
}
