"# utils" 


## argsUtil例子

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/noxue/utils/argsUtil"
)

type Data struct {
	Id       int               `json:"id"`
	Name     string            `json:"name"`
	User     User              `json:"user"`
	Articles map[string]string `json:"articles"`
	Admin    bool              `json:"admin"`
}
type User struct {
	Age     int    `json:"age"`
	Address string `json:"address"`
}

// 实现Stringer接口，格式化输出，方便展示
func (me *Data) String() string {
	bs, e := json.MarshalIndent(me,"","\t")
	if e != nil {
		panic(e)
	}
	return string(bs)
}

func main() {

	sd := Data{
		Id:    100,
		Name:  "不学网",
		Admin: true,
		User: User{
			Age:     20,
			Address: "中国哈哈哈",
		},
		Articles: map[string]string{
			"第一篇文章": "第一篇文章内容第一篇文章内容第一篇文章内容第一篇文章内容",
			"第二篇文章": "第二篇文章内容第二篇文章内容第二篇文章内容第二篇文章内容",
		},
	}
	var d Data
	var i int
	var b bool
	var nums []int
	e := argsUtil.NewArgsUtil(sd, 1,true,[]int{1,2,3,4}).To(&d, &i,&b,&nums)
	if e != nil {
		panic(e)
		return
	}
	fmt.Println("int:",i)
	fmt.Println("bool:",b)
	fmt.Print("array:")
	for _,v:=range nums{
		fmt.Print(v,",")
	}
	fmt.Println()
	fmt.Println("data:","\n",&d)
}
```