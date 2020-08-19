package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := Person{"william", 18}
	p := &people
	fmt.Printf("in main: %p,%p\n", p, &p) //0xc000096440,0xc0000ca018
	change(p)
	fmt.Printf("in main: %p,%p\n", p, &p) //0xc000096440,0xc0000ca018
	fmt.Println("in main", *p)            //in main {william 18}
}

func change(p *Person) {
	fmt.Printf("in func: %p,%p\n", p, &p) //0xc000096440,0xc0000ca028  &p is different from &p in main
	otherPeople := Person{"alice", 18}
	p = &otherPeople
	fmt.Printf("in func: %p,%p\n", p, &p) //0xc000096480,0xc0000ca028
	fmt.Println("in func", *p)            //{alice 18}
}
