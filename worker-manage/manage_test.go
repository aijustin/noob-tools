/**
 * @Author: quqiang
 * @Email: 77347042@qq.com
 * @Version: 1.0.0
 * @Date: 2022/4/3 7:13 PM
 */
package worker_manage

import (
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestWorker_Run(t *testing.T) {
	wokerManage := New(&WorkerConfig{
		WorkerNum:     100,
		WorkerName:    "",
		Log:           log.New(os.Stdout, "", 1),
		WorkerTimeOut: 5,
	})

	wokerManage.Run()

	for i := 0; i < 10; i++ {
		wokerManage.SetQueue(func(worker *Worker) {
			panic("1")
		})
	}

	for i := 0; i < 10; i++ {
		wokerManage.SetQueue(func(worker *Worker) {
			time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
		})
	}
	go func() {
		time.Sleep(5 * time.Second)
		wokerManage.Cancel()
	}()
	wokerManage.Wait()
}
