package demo1

import "fmt"

type funcExample struct{}
type Person struct {
	name string
	age  uint
}

// 值传递 不会影响 原有初始化的结构体
// 此种方式类似于Person对象的实例方法
func (p Person) setName(name string) {
	p.name = name
}

// 指针地址传递 会影响原有的数据
// 此种方式类似于Person对象的实例指针对象方法
func (p *Person) setAge(age uint) {
	p.age = age
}

func sub(p, q int) (sub, sum int) {
	return q + p, p - q
}
func setValue(p *Person) {
	p.name = "测试"
	p.age = 17
}

func Example() {
	p := Person{
		"zhansan",
		18,
	}

	p.setName("lisi")
	v := &p
	v.setAge(20)

	var p1 Person
	setValue(&p1)
	fmt.Printf("p--->>> %v\n p1 ---->>>> %v \n", p, p1)
}
