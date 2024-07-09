package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 数据源，类似MySQL中的数据
var users = []User{
	{ID: 1, Name: "张三"},
	{ID: 2, Name: "李四"},
	{ID: 3, Name: "王五"},
}

func main() {
	//创建一个gin的引擎实例
	r := gin.Default()
	//绑定引擎实例的url对应的处理函数
	r.GET("/users", listUser)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.DELETE("/users/:id", deleteUser)
	r.PATCH("/users/:id", updateUserName)
	r.Run(":8080") //引擎实例在8080端口运行
}

// 列举用户的函数
func listUser(c *gin.Context) {
	c.JSON(200, users) //返回用户列表,并且返回200的ok状态码
}

// 根据请求的id获取用户信息
func getUser(c *gin.Context) {
	id := c.Param("id") //获取请求的id的数据
	var user User       //创建一个用户对象
	found := false      //初始化查找的布尔值
	//类似于数据库的SQL查询
	for _, u := range users { //在数据库中逐一的对比id值
		if strings.EqualFold(id, strconv.Itoa(u.ID)) { //如果id值相等的情况下
			user = u     //将对象赋给用户
			found = true //设置找到的结果为真
			break        //结束循环
		}
	}

	if found { //如果是找到的情况
		c.JSON(200, user) //将正确的信息放入到json中
	} else {
		c.JSON(404, gin.H{ //否则返回自创的错误json信息
			"message": "用户不存在",
		})
	}
}

// 创建用户的处理函数
func createUser(c *gin.Context) {
	name := c.DefaultPostForm("name", "") //通过post中的带有name标签获得name这一数据
	if name != "" {                       //如果名字不为空
		u := User{ID: len(users) + 1, Name: name} //根据需求创建符合要求的用户对象
		users = append(users, u)                  //将对象插入数据库
		c.JSON(http.StatusCreated, u)             //返回正确插入用户数据的信息
	} else { //否则说明名字为空直接返回没有用户名的错误信息
		c.JSON(http.StatusOK, gin.H{
			"message": "请输入用户名称",
		})
	}
}

// 删除用户的操作
func deleteUser(c *gin.Context) {
	id := c.Param("id") //获取删除的id值
	i := -1             //暂存待删除元素索引值的变量
	//类似于数据库的SQL查询
	for index, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) { //如果查找到了该数据
			i = index //将index暂存到i中
			break     //结束搜索循环
		}
	}

	if i >= 0 { //如果当前的i值合法直接执行数据库删除操作
		users = append(users[:i], users[i+1:]...)
		c.JSON(http.StatusNoContent, "") //返回正常信息
	} else { //否则不合法返回错误信息
		c.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
	}
}

// 更新用户信息的函数
func updateUserName(c *gin.Context) {
	id := c.Param("id") //获取到用户id
	i := -1             //继续设置用于暂存用户信息的变量
	//类似于数据库的SQL查询
	for index, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) { //对比如果找到了index直接暂存
			i = index
			break
		}
	}

	if i >= 0 { //如果当前找到的是有意义的用户
		users[i].Name = c.DefaultPostForm("name", users[i].Name) //更新当前对象的name字段
		c.JSON(http.StatusOK, users[i])                          //返回正常更新信息
	} else { //否则返回错误信息
		c.JSON(http.StatusNotFound, gin.H{
			"message": "用户不存在",
		})
	}
}

// 用于处理用户的请求的含税
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method { //根据r引擎对象获得的方法
	case "GET": //如果是Get方法
		users, err := json.Marshal(users) //将当前的用户对象序列化
		if err != nil {                   //如果出错的情况
			w.WriteHeader(http.StatusInternalServerError)       //写入出错信息
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}") //将信息打出来
		} else { //否则输出正常的信息
			w.WriteHeader(http.StatusOK)
			w.Write(users) //再写入用户的信息
		}
	default: //如果用其他方式访问直接输出方法未找到的信息
		w.WriteHeader(http.StatusNotFound)            //将未找到的状态写入
		fmt.Fprint(w, "{\"message\": \"not found\"}") //将未找到的信息打出来
	}
}

// 用户
type User struct {
	ID   int
	Name string
}
