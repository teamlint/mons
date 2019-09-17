package timex

import "time"

var (
	// DateStart DateBegin 函数别名
	DateStart = DateBegin
)

// DateBegin 返回指定日期时间最开始的值
func DateBegin(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// DateEnd 返回指定日期时间结束的值
func DateEnd(t time.Time) time.Time {
	ed := DateBegin(t)
	return ed.AddDate(0, 0, 1).Add(-time.Nanosecond * 1)
}

// IsSameDate 判断两个时间是否是同一日期
func IsSameDate(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return (y1 == y2 && m1 == m2 && d1 == d2)
}

// IsBeforeDate 判断第一个日期是否在第二个日期之前
func IsBeforeDate(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	nd1 := time.Date(y1, m1, d1, 0, 0, 0, 0, t1.Location())
	y2, m2, d2 := t2.Date()
	nd2 := time.Date(y2, m2, d2, 0, 0, 0, 0, t2.Location())

	return nd1.Before(nd2)
}

// DaysBetween 返回开始时间和结束时间之间的天数
func DaysBetween(fromDate time.Time, toDate time.Time) int {
	return int(toDate.Sub(fromDate) / (24 * time.Hour))
}

// StartOfMonth 返回指定时间月份的第一天
func StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

// EndOfMonth 返回指定时间月份的最后一天
func EndOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, date.Location())
}
