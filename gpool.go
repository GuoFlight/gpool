package gpool

import (
	"errors"
	"reflect"
	"sync"
)

type GPool struct {
	WaitGroup sync.WaitGroup
	IsWait bool
	GoroutineCountLimit chan int
	funcList []Function
	RetList [][]reflect.Value
	lockRetList sync.Mutex
}
type Function struct {
	FunctionAddr interface{}
	Params []interface{}
}

func New(goroutineCountLimit int,isWait bool) GPool {
	var gpool GPool
	gpool.GoroutineCountLimit = make(chan int,goroutineCountLimit)
	gpool.IsWait = isWait
	return gpool
}
func NewDefault() GPool {
	var gpool GPool
	gpool.GoroutineCountLimit = make(chan int,DefaultGoroutineCountLimit)
	gpool.IsWait = DefaultIsWait
	return gpool
}

func (this *GPool)AddGoroutine(functionName interface{},params... interface{})error{
	//校验参数
	valueFunction := reflect.ValueOf(functionName)
	if valueFunction.Kind() != reflect.Func {
		return errors.New("第一个参数不是函数 the first Parameter are not a function")
	}
	if valueFunction.Type().NumIn() != len(params){
		return errors.New("参数数量与函数类型不匹配")
	}
	//添加到函数列表
	functionAndParams := Function{
		FunctionAddr: functionName,
		Params: params,
	}
	this.funcList = append(this.funcList,functionAndParams)
	this.WaitGroup.Add(1)
	return nil
}

func (this *GPool)Run(){
	for _,curFunction := range this.funcList{
		//获取函数
		valueFunction := reflect.ValueOf(curFunction.FunctionAddr)
		//获取参数
		valueParams := make([]reflect.Value, len(curFunction.Params))
		for j:=0;j<len(curFunction.Params);j++ {
			valueParams[j] = reflect.ValueOf(curFunction.Params[j])
		}
		//调用函数
		go func(valueFunction reflect.Value,valueArgs []reflect.Value) {
			this.GoroutineCountLimit <- 1
			valueRets := valueFunction.Call(valueArgs)	//调用函数
			this.lockRetList.Lock()
			this.RetList = append(this.RetList, valueRets)
			this.lockRetList.Unlock()
			this.WaitGroup.Done()
			<- this.GoroutineCountLimit
		}(valueFunction,valueParams)
	}
	if this.IsWait{
		this.WaitGroup.Wait()
	}
}