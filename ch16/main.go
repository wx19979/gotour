// 反射是计算机语言中程序检视自身结构的一种方法,属于元编程的一种形式
// 反射三大定律:(1)任何接口值interface{}都可反射出反射对象
// (2)反射对象也可还原为interface{}变量(3)要修改反射对象,该值必须可设置,可寻址
// 字符串和结构体之间的转换是通过运行时反射实现的
// Go语言反射定义中,任何接口都由两部分组成:
// (1)接口的具体类型(2)具体类型对应的值
// (1)reflect.Value(本质是一个结构体) (2)reflect.Type(本质是一个接口)
// 其中interface{}是空接口,可以表示任何类型，可以把任何类型
// 转换为空接口,常用于反射、类型断言以减少代码重复,简化编程
// reflect.Value对应的函数类别:(1)获取或修改对应的值(2)和struct类型字段
// 有关用于获取对应字段(3)和类型上方法集有关用于获取对应方法
// 修改对应的值的步骤:(1)传递struct结构体指针对应的值reflect.Value
// (2)通过Elem方法获取指针指向的值(3)通过Field方法获取要修改的字段
// (4)通过Set系列方法修改成对应的值
// reflect.Type几个特殊的函数:(1)Implements方法用于判断是否实现了接口u
// (2)AssignableTo方法用于判断是否可以赋值给类型u(3)ConertibleTo方法用于
// 判读是否可以转换为类型u(4)Comparable方法用于判断该类型是否是可比较的
// 字符串和结构体互转场景最多的就是JSON和struct的互转
// struct tag是添加在struct字段上的标记,用它进行辅助,可以完成一些额外操作
// 如JSON和struct互转
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func main() {
	//测试int类型i的反射函数
	i := 3
	iv := reflect.ValueOf(i)  //获取i的值
	it := reflect.TypeOf(i)   //获取i的类型
	fmt.Println(iv.Int(), it) //3,int
	//转换iv的类型为int
	i1 := iv.Interface().(int)
	fmt.Println(i1)
	//根据地址获取到i的指针值
	ipv := reflect.ValueOf(&i)
	ipv.Elem().SetInt(4) //根据该指针修改当前的值
	fmt.Println(i)

	p := person{Name: "飞雪无情", Age: 20}  //创建一个人的对象
	ppv := reflect.ValueOf(&p)          //获取到人的指针
	ppv.Elem().Field(0).SetString("张三") //设置这个人的指针对应第一个参数的值为张三
	fmt.Println(p)
	fmt.Println("ppv ", ppv.Kind()) //输出ppv的类型为指针ppv.Kind()是直接输出其对应的底层类型

	pv := reflect.ValueOf(p) //输出当前反射的值
	fmt.Println(pv.Kind())   //输出当前pv的类型

	pt := reflect.TypeOf(p) //获取p的类型
	//遍历person的字段
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段：", pt.Field(i).Name)
	}
	//遍历person的方法
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法：", pt.Method(i).Name)
	}
	//创建一个stringerType和一个writerType
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer：", pt.Implements(stringerType)) //判断是否实现了该接口
	fmt.Println("是否实现了io.Writer：", pt.Implements(writerType))

	//struct to json
	jsonB, err := json.Marshal(p) //结构体转换为字符串
	if err == nil {
		fmt.Println(string(jsonB))
	}

	//json to struct
	respJSON := "{\"name\":\"李四\",\"age\":40}"
	json.Unmarshal([]byte(respJSON), &p) //字符串转换为结构体
	fmt.Println(p)

	//遍历person字段中key为json、bson的tag
	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n", sf.Name, sf.Tag.Get("json"))
		fmt.Printf("字段%s上,bson tag为%s\n", sf.Name, sf.Tag.Get("bson"))
	}

	jsonBuilder := strings.Builder{} //初始化jsonBuilder
	jsonBuilder.WriteString("{")     //向该对象写入"{"
	num := pt.NumField()             //获取pt对象字段个数
	for i := 0; i < num; i++ {       //依次遍历字段
		jsonTag := pt.Field(i).Tag.Get("json")                      //通过json这个标签判断是不是要找的字段
		jsonBuilder.WriteString("\"" + jsonTag + "\"")              //根据获取到的数据前后加上反斜杠
		jsonBuilder.WriteString(":")                                //然后再写入冒号
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i))) //最后将字段对应的数据写入
		if i < num-1 {                                              //最后再结束循环
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")      //然后再在尾部写上"}"
	fmt.Println(jsonBuilder.String()) //输出最后加工好的结果

	//反射调用person的Print方法
	mPrint := pv.MethodByName("Print")             //获取方法
	args := []reflect.Value{reflect.ValueOf("登录")} //获取参数
	mPrint.Call(args)                              //最后调用方法
}

// 定义人结构体
type person struct {
	Name string `json:"name" bson:"b_name"` //名字字段,上面标识json和bjson
	Age  int    `json:"age" bson:"b_name"`  //年龄字段,上面标识json和bjson
}

// 设置字符串的函数
func (p person) String() string {
	return fmt.Sprintf("Name is %s,Age is %d", p.Name, p.Age)
}

// 输出字符串的函数
func (p person) Print(prefix string) {
	fmt.Printf("%s:Name is %s,Age is %d\n", prefix, p.Name, p.Age)
}
