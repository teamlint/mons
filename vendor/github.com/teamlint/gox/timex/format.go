package timex

import (
	"math"
	"strconv"
	"time"
)

// Lang 语言
type Lang string

const (
	// Unix epoch (or Unix time or POSIX time or Unix timestamp)  1 year (365.24 days)
	infinity float64 = 31556926 * 1000
	// CHS 简体中文
	CHS Lang = "CHS"
	// CHT 繁体中文
	CHT Lang = "CHT"
	// ENG 英语
	ENG Lang = "ENG"
)

// Handler 基于最小秒数单位显示友好字符串格式函数
// timeIntervalThreshold 最小秒数,timePeriod 时间单位, message 友好字符串, processPlural 处理单位复数

func handler(timeIntervalThreshold float64, timePeriod, message string, processPlural bool) func(float64) string {
	return func(difference float64) string {
		n := difference / timeIntervalThreshold
		nStr := strconv.FormatFloat(n, 'f', 0, 64)
		if !processPlural {
			return nStr + timePeriod + message
		}
		if int(n) > 1 {
			return nStr + " " + timePeriod + "s " + message
		}
		return nStr + " " + timePeriod + " " + message
	}
}

// timeLapse condition struct
type timeLapse struct {
	// Time stamp threshold to handle the time lap condition
	Threshold float64
	// Handler function which determines the time lapse based on the condition
	Handler func(float64) string
}

var (
	timeLapses = []timeLapse{
		{Threshold: -31535999, Handler: handler(-31536000, "year", "from now", true)},
		{Threshold: -2591999, Handler: handler(-2592000, "month", "from now", true)},
		{Threshold: -604799, Handler: handler(-604800, "week", "from now", true)},
		{Threshold: -172799, Handler: handler(-86400, "day", "from now", true)},
		{Threshold: -86399, Handler: func(diff float64) string { return "tomorrow" }},
		{Threshold: -3599, Handler: handler(-3600, "hour", "from now", true)},
		{Threshold: -59, Handler: handler(-60, "minute", "from now", true)},
		{Threshold: -0.9999, Handler: handler(-1, "second", "from now", true)},
		{Threshold: 1, Handler: func(diff float64) string { return "just now" }},
		{Threshold: 60, Handler: handler(1, "second", "ago", true)},
		{Threshold: 3600, Handler: handler(60, "minute", "ago", true)},
		{Threshold: 86400, Handler: handler(3600, "hour", "ago", true)},
		{Threshold: 172800, Handler: func(diff float64) string { return "yesterday" }},
		{Threshold: 604800, Handler: handler(86400, "day", "ago", true)},
		{Threshold: 2592000, Handler: handler(604800, "week", "ago", true)},
		{Threshold: 31536000, Handler: handler(2592000, "month", "ago", true)},
		{Threshold: infinity, Handler: handler(31536000, "year", "ago", true)},
	}
	timeLapsesCHS = []timeLapse{
		{Threshold: -172799, Handler: handler(-86400, "天", "后", false)},
		{Threshold: -86399, Handler: func(diff float64) string { return "明天" }},
		{Threshold: -3599, Handler: handler(-3600, "小时", "后", false)},
		{Threshold: -59, Handler: handler(-60, "分钟", "后", false)},
		{Threshold: -0.9999, Handler: handler(-1, "秒", "后", false)},
		{Threshold: 1, Handler: func(diff float64) string {
			return "刚刚"
		}},
		{Threshold: 60, Handler: handler(1, "秒", "前", false)},
		{Threshold: 3600, Handler: handler(60, "分钟", "前", false)},
		{Threshold: 86400, Handler: handler(3600, "小时", "前", false)},
		{Threshold: 172800, Handler: func(diff float64) string { return "昨天" }},
		{Threshold: 604800, Handler: handler(86400, "天", "前", false)},
	}
)

// Pretty 返回指定时间消逝时间字符串
func Pretty(t time.Time, overtimeFormatter ...string) string {
	return PrettyByLang(t, CHS, overtimeFormatter...)
}

// PrettyByLang 根据指定语言返回指定时间消逝时间字符串
func PrettyByLang(t time.Time, lang Lang, overtimeFormatter ...string) (timeSince string) {
	timestamp := t.Unix()
	now := time.Now().Unix()

	if timestamp > now || timestamp <= 0 {
		timeSince = ""
	}

	timeElapsed := float64(now - timestamp)

	matched := false
	var timeFormats []timeLapse
	switch lang {
	case ENG:
		timeFormats = timeLapses
	default:
		timeFormats = timeLapsesCHS
	}
	defaultFormatter := DatetimeFormat
	if len(overtimeFormatter) > 0 {
		defaultFormatter = overtimeFormatter[0]
	}

	for _, formatter := range timeFormats {
		if timeElapsed < 0 {
			if math.Abs(timeElapsed) >= 604800 {
				return Format(t, defaultFormatter)
			}
		}
		if timeElapsed < formatter.Threshold {
			timeSince = formatter.Handler(timeElapsed)
			matched = true
			break
		}
	}
	if !matched {
		return Format(t, defaultFormatter)
	}
	return timeSince
}

// Format 日期时间格式化
// 可包含三个参数 1、时间对象 2、格式化字符串(string) 3、空值时返回值字符串(string)
func Format(t time.Time, f ...string) string {
	if t.IsZero() {
		if len(f) > 1 {
			return f[1]
		}
		return ""
	}
	format := DatetimeFormat
	if len(f) > 0 {
		format = f[0]
	}

	return t.Format(format)
}
