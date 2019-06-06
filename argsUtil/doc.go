/*

功能


方便不定参数的传出和取出的。 前提是要确定传入的参数个数和类型。

使用这个工具之前的代码

	func test(args ...interface{}){
		// 想在这里取出args中存取的各种类型，需要写一堆类型断言语句，非常麻烦
	}

	test(true,1,2.3,"test")


使用后


	func test(args ...interface{}){
		var b bool
		var n int
		var f float
		var s string
		new argsUtil(args).To(&b,&n,&f,&s)
		// 这里就可以使用 b  n f s 这些变量了，并且已经传入了对应数据
	}

	test(true,1,2.3,"test")


 */
package argsUtil

