package util

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func Test_FormatFileSize(t *testing.T) {
	fmt.Println(FormatFileSize(1))
	fmt.Println(FormatFileSize(1000))
	fmt.Println(FormatFileSize(10000))
	fmt.Println(FormatFileSize(100000))
	fmt.Println(FormatFileSize(1000000))
	fmt.Println(FormatFileSize(10000000))
	fmt.Println(FormatFileSize(100000000))
	fmt.Println(FormatFileSize(1000000000))
	fmt.Println(FormatFileSize(10000000000))
	fmt.Println(FormatFileSize(100000000000))
	fmt.Println(FormatFileSize(1000000000000))
	fmt.Println(FormatFileSize(10000000000000))
	fmt.Println(FormatFileSize(100000000000000))
	fmt.Println(FormatFileSize(1000000000000000))
	fmt.Println(FormatFileSize(10000000000000000))
	fmt.Println(FormatFileSize(10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000))
}

func TestExecuteTimeTotal_Auto(t *testing.T) {
	NewExecuteTimeTotal().Auto(func() {
		for i := 0; i < 999999999; i++ {

		}
	}, func(total int64) {
		fmt.Println(total)
	})
}

func BenchmarkSliceIsIn(b *testing.B) {
	SliceInValue(uint64(23), []uint64{123, 23, 23, 2, 32, 32, 3, 12, 3, 123, 1, 23, 12, 3, 2323232323233232, 123, 232323, 123, 2323, 23, 2, 3, 2, 3, 23, 2, 32, 3, 2, 3, 23, 2, 3, 2, 3, 2, 23333333333333})
}

func TestIsNumeric(t *testing.T) {
	fmt.Println(IsNumeric("000000123111100"))
}

func Test_SliceUnique(t *testing.T) {
	//data := []string{"wewewe",
	//	"wewewewe",
	//	"wewewewewewewewe",
	//	"wewewewewewewewe",
	//	"wewewewewewewewe",
	//	"wewewewewewewewe",
	//	"wewewewew23ewewewe",
	//	"wewewewewewewewe",
	//	"wewewewew23ewewewe",
	//	"wewewewewewewewe",
	//	"wewewewewewewewe",
	//	"wewewewewe23wewewe",
	//	"wewewewewewewewe",
	//	"wewewew321ewewewewe",
	//	"wewewewewewewewe",
	//	"wewewewew23ewewewe",
	//	"wewewewewewewewe",
	//	"wewewewewewewewe",
	//	"wewewewewewewewe",
	//	"wewewewewewewewe",
	//	"wewewewewewewewe"}
	data := []uint64{1, 4, 3, 4, 4}
	//type ass struct {
	//	A string
	//	B uint64
	//}
	//
	//data := make([]ass, 0)
	//tt := ass{"wew", 1}
	//for i := 1; i < 1000; i++ {
	//	data = append(data, tt)
	//}
	new := SliceUnique(data)
	for _, v := range new {
		fmt.Println(v)
	}
}

func BenchmarkSliceUnique(b *testing.B) {
	data := []uint64{1, 2, 4, 5, 4}
	SliceUnique(data)
}

func TestHttpClient_NewHttpClient(t *testing.T) {
	cli := NewHttpClient()
	//_ = cli.SetProxy("tcp", "127.0.0.1:1080")
	cli.timeOut = 3 * time.Second
	data := cli.Params(map[string]int{"name": 1}).GET("https://www.google.com").ToString()
	fmt.Println(data, "cli err", cli.GetLastErr())
}

func TestHttpClient_NewHttpClient2(t *testing.T) {
	cli := NewHttpClient()
	cli.timeOut = 3 * time.Second
	data := cli.JsonParams(map[string]string{"name": "quqiang"}).POST("http://127.0.0.1:8182/?a=1&b=2").ToString()
	fmt.Println(data, "cli err", cli.GetLastErr())
}

func TestHttpClient_NewHttpClient3(t *testing.T) {
	cli := NewHttpClient()
	cli.timeOut = 3 * time.Second
	cli.SetProxy("tcp", "127.0.0.1:1080")
	data := cli.Params(map[string]string{"name": "quqiang", "data[0]": "1", "data[1]": "1"}).POST("http://www.google.com").ToString()
	fmt.Println(data, "cli err", cli.GetLastErr())
}

func TestHttpClient_NewHttpClient4(t *testing.T) {
	cli := NewHttpClient()
	cli.timeOut = 300 * time.Second
	type par struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	a := &par{
		Name: "22",
		Age:  33,
	}
	cli.SetRequestOption(func(r *http.Request) {
		r.Header.Add("xxxx", "xxx")
	})
	cli.SetFileByFileName("/Users/quqiang/Desktop/request.php", "file", "xxx.php")
	data := cli.Params(a).POST("http://127.0.0.1:8182").ToString()
	fmt.Println(data, "cli err", cli.GetLastErr())
}

func TestHttpClient_POST(t *testing.T) {
	//a := make(map[string]interface{})
	//a["startsessiondatetime"] = "2021-07-21 00:00:00"
	//a["endsessiondatetime"] = "2021-07-21 23:59:59"
	type Params struct {
		Start string `json:"startsessiondatetime"`
		End  string    `json:"endsessiondatetime"`
	}
	a := &Params{
		Start: "2021-07-21 00:00:00",
		End:  "2021-07-21 23:59:59",
	}
	cli := NewHttpClient()
	cli.timeOut = 300 * time.Second
	cli.jsonParams = true
	cli.SetRequestOption(func(r *http.Request) {
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("VistaAccessToken", "4cjtMTaNldmeixhw3gUbQIqGk9FCpOun")
	})
	resp := cli.Params(a).POST("http://svc.b.1.jycinema.com:6155/api/query/runQueryGetDataTable?queryId=e83ae344-a9af-4df6-9ba7-469adb13b3af&version=v2&mode=all&DBLabelType=2&dbLabel=JinYi").ToString()
	fmt.Println(  resp)
}

func TestGetDistance(t *testing.T) {
	lat1, lng1 := 39.922425132446875, 116.39216423034668
	lat2, lng2 := 39.91355467158645, 116.40218496322632
	distance := GetDistance(lat1, lng1, lat2, lng2)
	fmt.Printf("%f m", distance*1000)

	fmt.Println()
	lat, lng := GetMid(lat1, lng1, lat2, lng2)
	fmt.Println(lat, lng)
}

func TestFloat(t *testing.T)  {
	a := 0.1
	b := 0.2
	fmt.Printf("%+v", a+b)
}
