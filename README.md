# 项目名称

gpool

# 作者

京城郭少

# 简介

Golang线程池

# Example

```go
package main
import (
	"github.com/GuoFlight/gpool"
	"fmt"
)
func Test1(a int)int{
    return a
}
func main() {
	gp := gpool.NewDefault()
	for i:=0;i<1000;i++{
		err := gp.AddGoroutine(Test1,i)
		if err!=nil{
			fmt.Println(err)
		}
	}
	gp.Run()
	for _,v1 := range gp.RetList{
		for _,v2 := range v1{
			fmt.Println(v2)
		}
	}
}
```

# 默认值

```go
DefaultGoroutineCountLimit = 10     //默认最大并发数量：10
DefaultIsWait = true                //执行Run方法默认会阻塞
```