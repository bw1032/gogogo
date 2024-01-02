package main

import (
	"fmt"
	"reflect"
)

func E() {

	//reflect.TypeOf()和reflect.ValueOf()

	//interface{}类型变量其具体类型使用reflect.Tpye来表示，具体值则使用reflect.Value来表示。 而reflect.Type和reflect.Value分别提供reflect.TypeOf()和reflect.ValueOf()来获取interface{}的具体类型及具体值
	var a interface{}
	a = 1
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a)) //int 1
	a = "a"
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a)) //string a

	//reflect.Kind
	var b struct{ a int `json:"bbb"` }
	fmt.Println(reflect.TypeOf(b))        //struct { a int }
	fmt.Println(reflect.TypeOf(b).Kind()) //struct

	//NumField() 和 Field()
	num := reflect.TypeOf(b).NumField()           //NumField()方法获取一个struct所有的fields的数量
	fmt.Println(reflect.TypeOf(b).Field(num - 1)) //{a gogo/interview int  0 [0] false} //Field(i int)获取指定第i个field,
	fmt.Println(reflect.TypeOf(b).Field(num - 1).Tag.Get("json")) //bbb

    //Int() 和String()
    var c = "a"
	fmt.Println(reflect.ValueOf(c).Int()) //panic: reflect: call of reflect.Value.Int on string Value


}
