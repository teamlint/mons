// Copyright 2018 The Teamlint Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// you can obtain one at https://github.com/teamlint/go.

// Package timex 实现time常用操作
// 所有time类型使用local时区转化处理
package timex

import "time"

const (
	// DatetimeFormat 默认日期时间格式
	DatetimeFormat = "2006-01-02 15:04:05"
	// DateFormat 默认日期格式
	DateFormat = "2006-01-02"
	// TimeFormat 默认时间格式
	TimeFormat = "15:04:05"
)

var (
	// TimeFormats 时间格式列表
	TimeFormats = []string{
		DateFormat,
		DatetimeFormat,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}
)
