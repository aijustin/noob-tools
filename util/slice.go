/**
 * @Author: quqiang
 * @Email: 77347042@qq.com
 * @Version: 1.0.0
 * @Date: 2021/1/21 3:02 下午
 */

package util

import (
	"reflect"
)

func SliceInValue(val interface{}, items interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(items).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(items)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func SliceUnique(items interface{}) []interface{} {
	maps := make(map[interface{}]interface{})
	switch reflect.TypeOf(items).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(items)
		newSlice := make([]interface{}, 0, s.Len())
		for i := 0; i < s.Len(); i++ {
			v := s.Index(i).Interface()
			if _, ok := maps[v]; !ok {
				maps[v] = nil
				newSlice = append(newSlice, v)
			}

		}
		return newSlice
	}
	return nil
}

func SliceDiff() {

}
