
package argsUtil

import (
	"bytes"
	"encoding/gob"
	"errors"
)

type argUtil struct {
	srcArgs []interface{}
}

/*
args参数是要转换的不定参数
 */
func NewArgsUtil(args ...interface{}) *argUtil {
	return &argUtil{
		srcArgs: args,
	}
}

/*
转换到指定的类型

特别要注意，要传入变量的引用，否则无法给变量填充值

如果提示错误：gob: attempt to decode into a non-pointer 就表示你提供的参数中有的不是变量的引用
  */
func (me *argUtil) To(args ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			e, ok := r.(error)
			if ok {
				err = e
			}
		}
	}()

	if len(me.srcArgs) != len(args) {
		panic(errors.New("要转换的参数和目标参数个数不同"))
	}
	for i, v := range me.srcArgs {

		var bs []byte
		buf := bytes.NewBuffer(bs)
		en := gob.NewEncoder(buf)

		e := en.Encode(v)
		if e != nil {
			panic(e)
		}

		de := gob.NewDecoder(buf)
		e = de.Decode(args[i])
		if e != nil {
			panic(e)
		}
	}
	return
}
