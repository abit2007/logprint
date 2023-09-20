package vegetable

import "packagestudy/logprint"

func MakeDish(name string) {
	msg := "做的素菜：" + name
	logprint.Info(msg) //这里调用其它包，不调也行自己写方法
}
