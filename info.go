// 最好是小写字母，并且最好不要带下划线，见名知义
package logprint

import (
	"fmt"
	"time"
)

func info(msg interface{}) { //接受传入任何类型的值，方法为小写字母开头是私有包，不能被导入
	t := time.Now()
	fmt.Printf("[info] %s: %s \n", t.Format("2006-01-02 15:04:05.000"), msg)
}

func Info(msg interface{}) { //接受传入任何类型的值，方法为大写字母开头是公开包，可以被其它包导入
	t := time.Now()
	fmt.Printf("[info] %s: %s \n", t.Format("2006-01-02 15:04:05.000"), msg)
}
