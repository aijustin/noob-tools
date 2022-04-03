/**
 * @Author: quqiang
 * @Email: 77347042@qq.com
 * @Version: 1.0.0
 * @Date: 2020/12/21 1:14 下午
 */
package util

import (
	"math"
	"time"
)

const (
	TIME_TEMPLATE_YYYYMMDDHHIISS_NORMAL = "2006-01-02 15:04:05"
	TIME_TEMPLATE_YYYYMMDD              = "20060102"
	TIME_TEMPLATE_YYYY_MM_DD            = "2006-01-02"
	TIME_TEMPLATE_YYYYMMDDHHIISS        = "20060102150405"
	ZERO_TEMPLATE_DEFAULT               = "0000-00-00 00:00:00"
)

// 获取当前时间时间戳
func GetNowSeconds() int64 {
	return time.Now().Unix()
}

// 时间戳转换字符串时间
func UnixTime2Str(timestamp int64, TimeTemplate string) string {
	_time := time.Unix(timestamp, 0)
	if TimeIsZero(_time) {
		return ZERO_TEMPLATE_DEFAULT
	}
	return _time.Format(TimeTemplate)
}

// 时间搓是否为0，time直接判断isZero是不对的，需要转换成utc时区
func TimeIsZero(t time.Time) bool {
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	return t.IsZero()
}

// 获取当前时间开始的时间点
func GetTimeStart(_time time.Time) time.Time {
	year, month, day := _time.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// 获取当前时间的0点的s数
func GetTimeStartSec(_time time.Time) int64 {
	return GetTimeStart(_time).Unix()
}

func TimeStr2TimeSec(_time string) int64 {
	parse, err := time.Parse(TIME_TEMPLATE_YYYYMMDDHHIISS_NORMAL, _time)
	if err != nil {
		return 0
	}
	return parse.Unix()
}

// 获取明天零点时间
func GetTomorrow(_time time.Time) time.Time {
	year, month, day := _time.Date()
	data := time.Date(year, month, day+1, 0, 0, 0, 0, time.Local)
	return data
}

func TimeNewByYearMonthDay(year int, month uint8, day int) int64 {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Unix()
}

func GetDurationNatureDays(startTime int64, endTime int64) []string {
	var result []string
	nowTime := time.Now().Unix()
	if endTime == 0 || endTime > nowTime {
		endTime = nowTime
	}
	if startTime == 0 || startTime > nowTime {
		startTime = nowTime
	}

	start := time.Unix(startTime, 0)
	startZeroTime := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())

	end := time.Unix(endTime, 0)
	endZeroTime := time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, end.Location())

	var days float64
	if start.Unix() > end.Unix() {
		days = 1
	} else {
		days = float64(float64(endZeroTime.Unix()-startZeroTime.Unix()) / (24 * 3600))
	}

	daysInt := int(math.Ceil(days))

	for i := 0; i < daysInt; i++ {
		result = append(result, startZeroTime.AddDate(0, 0, i).Format("2006-01-02"))
	}

	return result
}

func init() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = location
}
