package util

import "sync"

type UT struct {
	Function func(string) map[string]string
	Mymap map[string]string
	Name string
	sync.Mutex
}

func NewUT() *UT{
	ut := &UT{}
	ut.Mymap = map[string]string{}
	ut.Name = "GreyGrayPuff"
	ut.Function = func(s string) map[string]string{
		m:=map[string]string{}
		m[s]="Mouse"
		return m
	}
	return ut
}
func (ut *UT)DoSomething(s string){
	ut.Lock()
	defer ut.Unlock()

	ut.Mymap[s] = "Boseman"

}

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}
func (L *List) Insert(key interface{}){
	list := &Node{
		next: L.head,
		key: key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}
func main(){
	link := List{}
	link.Insert("Nuts")
	link.Insert("Vegan Ice Cream")
	link.Insert("Spinach")
}
