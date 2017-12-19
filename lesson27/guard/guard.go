package main

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}

func main() {
	var nobody *person
	nobody.birthday()
}
