/**
 * @Author: quqiang
 * @Email: 77347042@qq.com
 * @Version: 1.0.0
 * @Date: 2021/3/3 1:01 下午
 */
package util

import "time"

type ExecuteTimeTotal struct {
	ExecuteTimeTotalDetail
}

type ExecuteTimeTotalDetail struct {
	StartTime time.Time
	EndTime   time.Time
}

func NewExecuteTimeTotal() *ExecuteTimeTotal {
	obj := new(ExecuteTimeTotal)
	obj.Start()
	return obj
}

func (obj *ExecuteTimeTotal) Start() *ExecuteTimeTotal {
	obj.StartTime = time.Now()
	return obj
}

func (obj *ExecuteTimeTotal) End() *ExecuteTimeTotal {
	obj.EndTime = time.Now()
	return obj
}

func (obj *ExecuteTimeTotal) Total() int64 {
	return time.Now().Sub(obj.StartTime).Milliseconds()
}

func (obj *ExecuteTimeTotal) Auto(f func(), calback func(total int64)) {
	obj.Start()
	f()
	obj.End()
	calback(obj.Total())
}
