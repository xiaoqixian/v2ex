// Date:   Thu Jul 10 11:18:53 2025
// Mail:   lunar_ubuntu@qq.com
// Author: https://github.com/xiaoqixian

package util

import (
	"log"
	"reflect"
)

func MustCast[T any](in any) T {
	to, ok := in.(T)
	if !ok {
		var tmp T
		log.Fatalf("Convert to type '%s' failed, actual type is '%s'", 
			reflect.TypeOf(tmp).Name(), reflect.TypeOf(in).Name())
	}
	return to
}
