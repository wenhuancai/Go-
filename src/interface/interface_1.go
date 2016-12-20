package main

import (
	"fmt"
)

type Tester struct {
	s interface {
		String() string
	}
}


type Test struct {
	 ( s func *)() int
}


type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func main() {
	p := &User{1, "Tom"}
	fmt.Printf("p is %p\n", p)
	t := Tester{p}
	fmt.Println("t is ", t)
	fmt.Println(t.s.String)
	fmt.Println(p.String)
}
