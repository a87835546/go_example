package hprose_example

import (
	"errors"
	"github.com/hprose/hprose-go/hprose"
	"log"
	"net/http"
	"time"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

type Person struct {
	Name     string
	Age      int
	Sex      bool
	birthday int64
}
type myService struct{}

func (myService) Swap(a int, b int) (int, int) {
	return b, a
}

func (myService) Sum(args ...int) (int, error) {
	if len(args) < 2 {
		return 0, errors.New("Requires at least two parameters")
	}
	a := args[0]
	for i := 1; i < len(args); i++ {
		a += args[i]
	}
	return a, nil
}

func (myService) AddTest(p Person) (Person, error) {
	p.Name = "测试===%v" + p.Name
	p.Age = 100 + p.Age
	p.Sex = true
	log.Printf("person -->>> %v", p)
	return p, nil
}

func (myService) GetPersons() []Person {
	var persons []Person
	p1 := Person{
		"lisi",
		18,
		true,
		time.Now().Unix(),
	}
	p2 := Person{
		"zhangsan",
		19,
		true,
		time.Now().Unix(),
	}
	persons = append(persons, p1)
	persons = append(persons, p2)
	return persons
}
func init() {
	log.Println("init----->>>>>>>>>>")
	service := hprose.NewHttpService()
	service.AddFunction("hello", hello)
	service.AddMethods(myService{})
	http.ListenAndServe(":8080", service)
}
