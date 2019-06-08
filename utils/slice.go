package utils

import (
	"reflect"
	"fmt"
)

func SliceContains(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("1-st parameter given a non-slice type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

// SliceJoin convert all values to []string
func SliceToStr(slice interface{}) []string {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("parameter given a non-slice type")
	}

	st := []string{}
	for i := 0; i < s.Len(); i++ {
		st = append(st, fmt.Sprintf("%v", s.Index(i).Interface()))
	}
	return st
}

func SliceIndex(slice interface{}, item interface{}) int {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("1-st parameter given a non-slice type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return i
		}
	}
	return -1
}

func SliceDelete(slicePtr interface{}, item interface{}) {
	ptr := reflect.ValueOf(slicePtr)

	if ptr.Kind() != reflect.Ptr {
		panic("1-st parameter muct be ptr on slice type")
	}
	val := ptr.Elem()

	i := SliceIndex(val.Interface(), item)
	if i == -1 {
		return
	}
	val.Set(reflect.AppendSlice(val.Slice(0, i), val.Slice(i+1, val.Len())))
	return
}

func SliceDeleteIndex(slicePtr interface{}, index int) {
	ptr := reflect.ValueOf(slicePtr)

	if ptr.Kind() != reflect.Ptr {
		panic("1-st parameter muct be ptr on slice type")
	}
	val := ptr.Elem()

	len := val.Len()
	if index < 0 || index >= len {
		panic(fmt.Sprintf("index %d out of range [0:%d]", index, len-1))
	}

	val.Set(reflect.AppendSlice(val.Slice(0, index), val.Slice(index+1, len)))
	return
}

