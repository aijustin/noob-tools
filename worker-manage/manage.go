/**
 * @Author: quqiang
 * @Email: 77347042@qq.com
 * @Version: 1.0.0
 * @Date: 2022/4/3 6:06 PM
 */
package worker_manage

import (
	"context"
	"fmt"
	"log"
	"noob-tools/util"
	"time"
)

type WorkerConfig struct {
	WorkerNum     int
	WorkerName    string
	Log           *log.Logger
	WorkerTimeOut time.Duration
}

type Worker struct {
	Config      *WorkerConfig
	workerQueue chan WorkerHandle //队列
	workerCtx   context.Context
	cancel      context.CancelFunc
}

type WorkerHandle func(worker *Worker)

func New(config *WorkerConfig) *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		Config:      config,
		workerCtx:   ctx,
		cancel:      cancel,
		workerQueue: make(chan WorkerHandle, 0),
	}
}

func (qw *Worker) Run() *Worker {
	for i := 0; i < qw.Config.WorkerNum; i++ {
		tid := i
		qw.workerGo(fmt.Sprintf("workerId:%d", tid))
	}
	return qw
}

func (qw *Worker) Wait() {
	for {
		time.Sleep(1 * time.Second)
	}
}

func (qw *Worker) SetQueue(handle WorkerHandle) *Worker {
	qw.workerQueue <- handle
	return qw
}

func (qw *Worker) CloseQueue() {
	close(qw.workerQueue)
}

func (qw *Worker) Cancel() {
	qw.cancel()
}

func (qw *Worker) workerGo(id string) {
	qw.Config.Log.Println(id, "Worker Start:", qw.Config.WorkerName)
	go func() {
		for {
			select {
			case handle := <-qw.workerQueue:
				if handle == nil {
					qw.Config.Log.Println(id, "Worker Queue handle Is nil:", qw.Config.WorkerName)
					break
				}
				qw.handleRun(id, handle)
			case <-qw.workerCtx.Done():
				qw.Config.Log.Println(id, "Worker Exit:", qw.Config.WorkerName)
				return
			}
		}
	}()
}

func (qw *Worker) handleRun(id string, handle WorkerHandle) {
	util.NewExecuteTimeTotal().Auto(func() {
		defer func() {
			e := recover()
			if e != nil {
				qw.Config.Log.Println(id, "Worker Panic:", qw.Config.WorkerName, e)
			}
		}()
		handle(qw)
	}, func(total int64) {
		qw.Config.Log.Println(id, "Worker End:", qw.Config.WorkerName, " Time", total, "ms")
	})
}
