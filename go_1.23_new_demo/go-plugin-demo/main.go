package main

import "plugin"

func main() {
	// 将编译好的插件 .so 文件加载并获取其中的方法作为备用,可以实现动态加载所需要的功能
	// 该种插件方式  针对Linux支持较好，Mac有少量问题但支持，Windows不支持，不同的Go版本也有问题
	p, err := plugin.Open("./plugins/plugin.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}

	*v.(*int) = 999
	f.(func())()
}
