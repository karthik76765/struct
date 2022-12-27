package
import "fmt"
type person struct {
	name string
	age int
}
func newperson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Printlnln(person{name:"alice", age: 30})
	fmt.Println(person{name: "fred"})
	fmt.Println("&person{name: "ann",age: 40})
	fmt.println("newperson("jon"))
	s := person{name: "sean", age: 50}
fmt.Println("s.name")
sp := &s
fmt.Println(sp.age)
sp.age = 51
fmt.Println(sp.age)
}