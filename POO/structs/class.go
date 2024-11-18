package main

import "fmt"

type Emmployee struct {
	id   int
	name string
}

func (e *Emmployee) SetId(id int) {
	e.id = id
}

func (e *Emmployee) SetName(name string) {
	e.name = name
}

func (e *Emmployee) GetId() int {
	return e.id
}
func (e *Emmployee) GetName() string {
	return e.name
}

func main() {
	e := Emmployee{}
	fmt.Printf("%v", e)
	e.id = 1
	e.name = "Name"
	e.SetId(5)
	e.SetName("Name 2")
	//fmt.Printf("%v", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

}
