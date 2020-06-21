// Code generated by 'goexports hash/crc32'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"go/constant"
	"go/token"
	"hash/crc32"
	"reflect"
)

func init() {
	Symbols["hash/crc32"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Castagnoli":   reflect.ValueOf(constant.MakeFromLiteral("2197175160", token.INT, 0)),
		"Checksum":     reflect.ValueOf(crc32.Checksum),
		"ChecksumIEEE": reflect.ValueOf(crc32.ChecksumIEEE),
		"IEEE":         reflect.ValueOf(constant.MakeFromLiteral("3988292384", token.INT, 0)),
		"IEEETable":    reflect.ValueOf(&crc32.IEEETable).Elem(),
		"Koopman":      reflect.ValueOf(constant.MakeFromLiteral("3945912366", token.INT, 0)),
		"MakeTable":    reflect.ValueOf(crc32.MakeTable),
		"New":          reflect.ValueOf(crc32.New),
		"NewIEEE":      reflect.ValueOf(crc32.NewIEEE),
		"Size":         reflect.ValueOf(constant.MakeFromLiteral("4", token.INT, 0)),
		"Update":       reflect.ValueOf(crc32.Update),

		// type definitions
		"Table": reflect.ValueOf((*crc32.Table)(nil)),
	}
}
