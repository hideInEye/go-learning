package main

import (
	"fmt"
)

type Foo struct {
	name string `json:"name"`
}

func PointerMethod(f *Foo) {
	fmt.Println("pointer method on", f.name)
}

func ValueMethod(f Foo) {
	fmt.Println("value method on", f.name)
}

func NewFoo() Foo {
	return Foo{name: "right value struct"}
}

func f1() int {
	x := 5
	defer func() { x++ }()
	return x
}

func f2() (x int) {
	defer func() { x++ }()
	return 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	// 数组字符串储存
	//name := [...]string{"吴佩","alex"}
	//fmt.Printf("数组的内存地址:%p \n",&name)
	//fmt.Printf("数组的第一个内存地址：%p \n",&name[0])
	//fmt.Printf("数组的第二个内存地址：%p \n",&name[1])
	//f1:= Foo{name:"value struct"}
	//ValueMethod(f1)
	//f3 :=NewFoo()
	//ValueMethod(f3)
	//f2 := &Foo{name:"pointer struct"}
	//PointerMethod(f2)

	// 在切片不执行扩容的时候
	//user := make([]int,3,7)
	//fmt.Println(&user[1])
	//fmt.Println(&user[2])
	//v1 := make([]int,1,3)
	//v2 :=append(v1,66)
	//// append 执行时他们的内存地址时相同的
	//v1[0]=999
	//fmt.Println(v2)
	//fmt.Println(v1)
	// 另一种赋值
	//v1 := make([]int,1,3)
	//v1 = append(v1,99)
	//fmt.Println(v1)

	// 切片执行扩容的时候
	//v1 :=[]int{1,2,3}
	//v1:= make([]int,3,3)
	//fmt.Println(v1,len(v1),cap(v1))
	//
	//v2:= append(v1,44)
	//fmt.Println(v1,len(v1),cap(v1))
	//fmt.Println(v2,len(v2),cap(v2))
	//v1[0]=5555
	//fmt.Println(v1)
	//fmt.Println(v2)

	// 注意：通过切片切出来的数据和原切片的内存地址是相同的
	//v1 := []int{1,2,3,4,5}
	//v2 := v1[1:3]
	//v3 := v1[1:]
	//v4 := v1[:3]

	// 删除 注意：使用切片是不太会使用删除 【链表】
	//v1 := []int {11,22,33,44,55,66}
	//dataIndex :=2
	//v2 := append(v1[:dataIndex],v1[dataIndex+1:]...)
	//fmt.Println(v2)

	// 插入
	//v1 :=[]int{1,2,3,4,5,6}
	//dataIndex :=3
	//result := make([]int,0, len(v1)+1)
	//result = append(result,v1[:dataIndex]...)
	//result = append(result,99)
	//result = append(result,v1[dataIndex:]...)
	//fmt.Println(result)

	// 循环
	//v1 :=[]int{1,2,3,4,5,6}
	//
	//for i:=0;i< len(v1);i++{
	//	fmt.Println(i,v1[i])
	//}
	//for index,value := range v1{
	//	fmt.Println(index,value)
	//}

	// map 理解
	// 初始化
	// userInfo :=map[string]string{}
	//userinfo := map[string]string{"name": "sad", "age": "22"}
	//info := Foo{
	//	name: "1123",
	//}
	//userInfo := []Foo{info}
	//fmt.Println(userInfo)
	//userinfo["age"]="20"
	//userinfo["email"] = "123@123.com"
	//fmt.Println(userinfo)
	//data := make(map[string]int,10)
	//data["100"] = 100
	//data["200"] = 200
	//fmt.Println(len(data))
	//var row map[string]int
	////row["name"] =1 会出现编译报错
	//fmt.Println(row)
	//
	//// value 是个指针
	//value := new(map[string]int)
	////value["key"]=1 会出现编译报错
	//fmt.Println(&value)
	//value = &data
	//fmt.Println(value)
	//v1 := make(map[[2]int] float32)
	//v1[[2]int{1,1}] = 1.2
	// 打印结果
	//sad
	//map[age:20 email:123@123.com name:sad]
	//2
	//map[]
	//0xc00000e038
	//&map[100:100 200:200]
	// 长度
	//data := map[string]string{"n1":"11"}
	//val := len(data)
	//fmt.Println(val)
	//
	//// 根据容量的值，计算出合适的容量
	//// 一个map中会包含很多容器，每个容器中可以存在8个键值对
	//info := make(map[string]string,10)
	//info["n1"] = "11"
	//v1 :=len(info)
	//fmt.Println(v1)
	//// 2
	////v2 :=cap(info)  会报错不支持map计算

	// crud
	//data := map[string]string{"n1":"111"}
	//data["n3"] = "222"
	//data["n1"]="1"
	//delete(data,"n2")
	//fmt.Println(data["n1"])
	//for _,value := range data{
	//	fmt.Println(value)
	//}

	//number := decimal.NewFromFloat(2021.556368).Round(2)
	//fmt.Printf("number:%s", number)

	//number := 5
	//for number < 10 {
	//	defer fmt.Printf("defer:%d\n", number)
	//	fmt.Printf(" number: %d\n", number)
	//	number++
	//}

}
