package main

import (
	"fmt"
	"github.com/GuoFlight/gpool"
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
