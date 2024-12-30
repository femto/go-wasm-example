//package go_wasm_example

package main

import (
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return nil
	}
	return args[0].Int() + args[1].Int()
}

func fibonacci(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return nil
	}
	n := args[0].Int()
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func greet(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return nil
	}
	return "Hello, " + args[0].String() + "!"
}

func main() {
	c := make(chan struct{}, 0)

	// 注册函数到JavaScript
	js.Global().Set("wasmAdd", js.FuncOf(add))
	js.Global().Set("wasmFibonacci", js.FuncOf(fibonacci))
	js.Global().Set("wasmGreet", js.FuncOf(greet))

	<-c // 保持程序运行
}
