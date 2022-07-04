package demo1

import (
	"errors"
	"fmt"
)

/**
interface 常见的使用介绍
*/
func init() {
	fmt.Println("init func .....")
}

func TestFunc() {
	fmt.Println("test  func ....")
}

type stringer interface {
	string() string
}
type notifier interface {
	notify()
	test(args int) (returnValue string)
	stringer
}

type user struct {
	name string
	role int
}

func (u *user) notify() {
	fmt.Printf("user struct call this interface %s\n", u.name)
}
func (u *user) test(i int) string {
	fmt.Printf("user struct call this interface  -- test method %s\n", u.name)
	return "" + u.name + "test method"
}

func (u *user) string() string {
	return "interface 嵌入 其他的 interface --- user struct"
}

type tester interface {
	toString()
}
type integer int

func (i integer) toString() {
	fmt.Printf("test inferface --- int to string --->>> %v\n", i)
}

type admin struct {
	name string
	role int
}

func (a *admin) notify() {
	fmt.Printf("admin struct call this interface %s\n", a.name)
}

func (a *admin) test(i int) string {
	return "admin struct test interface"
}

func (a *admin) string() string {
	return "interface 嵌入 其他的 interface --- user struct"
}
func InterfaceDemo() {
	var t1, t2 interface{}
	fmt.Printf("t1 %d t2\n", t2 == t1)
	t1 = 100
	t2 = 200
	fmt.Printf("t1 %d t2\n", t2 == t1)

	u1 := user{
		"zhansan",
		11,
	}
	sendNotification(&u1)
	fmt.Printf("%s\n", u1.string())
	a1 := admin{
		"admin",
		10,
	}
	sendNotification(&a1)
	var a integer = 10
	a.toString()

	var err error = errors.New("123")
	fmt.Printf("自定义err 信息 --->>>%v\n", err.Error())
}

func sendNotification(n notifier) {
	n.notify()
	res := n.test(11)
	fmt.Printf("res -->>>%s\n", res)
}
