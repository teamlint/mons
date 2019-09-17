// Copyright 2018 The Teamlint Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// You can obtain one at https://github.com/teamlint/go.

// Copyright 2017 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// Package binaryx 二进制及byte操作
package binaryx

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

// Bit 二进制位(0|1)
type Bit int8

// Encode 针对基本类型进行二进制打包，支持的基本数据类型包括:int/8/16/32/64、uint/8/16/32/64、float32/64、bool、string、[]byte
// 其他未知类型使用 fmt.Sprintf("%v", value) 转换为字符串之后处理
func Encode(vs ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(vs); i++ {
		switch value := vs[i].(type) {
		case int:
			buf.Write(EncodeInt(value))
		case int8:
			buf.Write(EncodeInt8(value))
		case int16:
			buf.Write(EncodeInt16(value))
		case int32:
			buf.Write(EncodeInt32(value))
		case int64:
			buf.Write(EncodeInt64(value))
		case uint:
			buf.Write(EncodeUint(value))
		case uint8:
			buf.Write(EncodeUint8(value))
		case uint16:
			buf.Write(EncodeUint16(value))
		case uint32:
			buf.Write(EncodeUint32(value))
		case uint64:
			buf.Write(EncodeUint64(value))
		case bool:
			buf.Write(EncodeBool(value))
		case string:
			buf.Write(EncodeString(value))
		case []byte:
			buf.Write(value)
		case float32:
			buf.Write(EncodeFloat32(value))
		case float64:
			buf.Write(EncodeFloat64(value))
		default:
			if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
				buf.Write(EncodeString(fmt.Sprintf("%v", value)))
			}
		}
	}
	return buf.Bytes()
}

// Decode 整形二进制解包
// 注意第二个及其后参数为字长确定的整形变量的指针地址，以便确定解析的[]byte长度，
// 例如：int8/16/32/64、uint8/16/32/64、float32/64等等
func Decode(b []byte, vs ...interface{}) error {
	buf := bytes.NewBuffer(b)
	for i := 0; i < len(vs); i++ {
		err := binary.Read(buf, binary.LittleEndian, vs[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// EncodeString 编码字符串
func EncodeString(s string) []byte {
	return []byte(s)
}

// DecodeToString 字符串解码
func DecodeToString(b []byte) string {
	return string(b)
}

// EncodeBool 编码bool类型
func EncodeBool(b bool) []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}

// EncodeInt 自动识别int类型长度，转换为[]byte
func EncodeInt(i int) []byte {
	if i <= math.MaxInt8 {
		return EncodeInt8(int8(i))
	} else if i <= math.MaxInt16 {
		return EncodeInt16(int16(i))
	} else if i <= math.MaxInt32 {
		return EncodeInt32(int32(i))
	}
	return EncodeInt64(int64(i))
}

// EncodeUint 自动识别uint类型长度，转换为[]byte
func EncodeUint(i uint) []byte {
	if i <= math.MaxUint8 {
		return EncodeUint8(uint8(i))
	} else if i <= math.MaxUint16 {
		return EncodeUint16(uint16(i))
	} else if i <= math.MaxUint32 {
		return EncodeUint32(uint32(i))
	}
	return EncodeUint64(uint64(i))
}

// EncodeInt8 编码int8
func EncodeInt8(i int8) []byte {
	return []byte{byte(i)}
}

// EncodeUint8 编码uint8
func EncodeUint8(i uint8) []byte {
	return []byte{byte(i)}
}

// EncodeInt16 编码int16
func EncodeInt16(i int16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, uint16(i))
	return bytes
}

// EncodeUint16 编码uint16
func EncodeUint16(i uint16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, i)
	return bytes
}

// EncodeInt32 编码int32
func EncodeInt32(i int32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(i))
	return bytes
}

// EncodeUint32 编码uint32
func EncodeUint32(i uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, i)
	return bytes
}

// EncodeInt64 编码int64
func EncodeInt64(i int64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(i))
	return bytes
}

// EncodeUint64 编码uint64
func EncodeUint64(i uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, i)
	return bytes
}

// EncodeFloat32 编码float32
func EncodeFloat32(f float32) []byte {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

// EncodeFloat64 编码float64
func EncodeFloat64(f float64) []byte {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

// fillUpSize 当b位数不够时，进行高位补0
func fillUpSize(b []byte, l int) []byte {
	c := make([]byte, 0)
	c = append(c, b...)
	for i := 0; i <= l-len(b); i++ {
		c = append(c, 0x00)
	}
	return c
}

// DecodeToInt 将二进制解析为int类型，根据[]byte的长度进行自动转换
func DecodeToInt(b []byte) int {
	if len(b) < 2 {
		return int(DecodeToInt8(b))
	} else if len(b) < 3 {
		return int(DecodeToInt16(b))
	} else if len(b) < 5 {
		return int(DecodeToInt32(b))
	}
	return int(DecodeToInt64(b))
}

// DecodeToUint 将二进制解析为uint类型，根据[]byte的长度进行自动转换
func DecodeToUint(b []byte) uint {
	if len(b) < 2 {
		return uint(DecodeToUint8(b))
	} else if len(b) < 3 {
		return uint(DecodeToUint16(b))
	} else if len(b) < 5 {
		return uint(DecodeToUint32(b))
	}
	return uint(DecodeToUint64(b))
}

// DecodeToBool 将二进制解析为bool类型，识别标准是判断二进制中数值是否都为0，或者为空
func DecodeToBool(b []byte) bool {
	if len(b) == 0 {
		return false
	}
	if bytes.Equal(b, make([]byte, len(b))) {
		return false
	}
	return true
}

// DecodeToInt8 将二进制解析为int8类型
func DecodeToInt8(b []byte) int8 {
	return int8(b[0])
}

// DecodeToUint8 将二进制解析为uint8类型
func DecodeToUint8(b []byte) uint8 {
	return uint8(b[0])
}

// DecodeToInt16 将二进制解析为int16类型
func DecodeToInt16(b []byte) int16 {
	return int16(binary.LittleEndian.Uint16(fillUpSize(b, 2)))
}

// DecodeToUint16 将二进制解析为uint16类型
func DecodeToUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(fillUpSize(b, 2))
}

// DecodeToInt32 将二进制解析为int32类型
func DecodeToInt32(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(fillUpSize(b, 4)))
}

// DecodeToUint32 将二进制解析为uint32类型
func DecodeToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(fillUpSize(b, 4))
}

// DecodeToInt64 将二进制解析为int64类型
func DecodeToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(fillUpSize(b, 8)))
}

// DecodeToUint64 将二进制解析为uint64类型
func DecodeToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(fillUpSize(b, 8))
}

// DecodeToFloat32 将二进制解析为float32类型
func DecodeToFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(fillUpSize(b, 4)))
}

// DecodeToFloat64 将二进制解析为float64类型
func DecodeToFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(fillUpSize(b, 8)))
}

// EncodeBits 默认编码
func EncodeBits(bits []Bit, i int, l int) []Bit {
	return EncodeBitsWithUint(bits, uint(i), l)
}

// EncodeBitsWithUint 将ui按位合并到bits数组中，并占length长度位(注意：uis数组中存放的是二进制的0|1数字)
func EncodeBitsWithUint(bits []Bit, ui uint, l int) []Bit {
	a := make([]Bit, l)
	for i := l - 1; i >= 0; i-- {
		a[i] = Bit(ui & 1)
		ui >>= 1
	}
	if bits != nil {
		return append(bits, a...)
	}
	return a
}

// EncodeBitsToBytes 将bits转换为[]byte，从左至右进行编码，不足1 byte按0往末尾补充
func EncodeBitsToBytes(bits []Bit) []byte {
	if len(bits)%8 != 0 {
		for i := 0; i < len(bits)%8; i++ {
			bits = append(bits, 0)
		}
	}
	b := make([]byte, 0)
	for i := 0; i < len(bits); i += 8 {
		b = append(b, byte(DecodeBitsToUint(bits[i:i+8])))
	}
	return b
}

// DecodeBits 解析为int
func DecodeBits(bits []Bit) int {
	v := int(0)
	for _, i := range bits {
		v = v<<1 | int(i)
	}
	return v
}

// DecodeBitsToUint 解析为uint
func DecodeBitsToUint(bits []Bit) uint {
	v := uint(0)
	for _, i := range bits {
		v = v<<1 | uint(i)
	}
	return v
}

// DecodeBytesToBits 解析[]byte为字位数组[]uint8
func DecodeBytesToBits(bs []byte) []Bit {
	bits := make([]Bit, 0)
	for _, b := range bs {
		bits = EncodeBitsWithUint(bits, uint(b), 8)
	}
	return bits
}
