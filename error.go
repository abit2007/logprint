package logprint

import (
	"fmt"
	"time"
)

func Error(msg interface{}) { //这里用error类型较好些 func Error(err error) ，想不起如何填了
	t := time.Now()
	fmt.Printf("[error] %s: %s \n", t.Format("2006-01-02 15:04:05.000"), msg)
}
