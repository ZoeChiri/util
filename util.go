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