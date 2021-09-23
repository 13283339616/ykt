package main

import "github.com/panjf2000/ants/v2"

/**
在程序启动的时候就启动一个动态的协程池
*/
var poolObj *ants.Pool

func poolInit() {
	//默认给10000的动态容量
	poolObj, _ = ants.NewPool(10000)
}

//提交任务
func poolSubmit(fun func()) {
	ants.Submit(fun)
}

//池释放
func poolRelease() {
	poolObj.Release()
}

//池重启
func poolReboot() {
	poolObj.Reboot()
}

//池动态调整大小
func poolTune(size int) {
	poolObj.Tune(size)
}

//返回池还有多少空闲的协程
func poolGetFreeSize() int {
	return poolObj.Free()
}

//返回池中总共有多少协程
func poolGetCapSize() int {
	return poolObj.Cap()
}

//返回池中总共有多少运行的协程
func poolGetRunningSize() int {
	return poolObj.Running()
}
