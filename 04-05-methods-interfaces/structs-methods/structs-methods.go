package main

import "fmt"

type T0 struct {
	x  int  // -> 0
	xp *int // -> nil
}

func (*T0) M0() int { return 1 }

type T1 struct {
	y int
}
func (T1) M1() int     { return 1 }
func (t1 *T1) M3() int { return t1.y }

type PT1 *T1

type I interface{
	M4() string
}

type IImpl struct{}
func (IImpl) M4() string {
	return "M4() result"
}

type T2 struct {
	z int
	T1
	*T0
	I
}

func (*T2) M2() int { return 1 }

type Q *T2


var t T2 = T2{z: 5, T1: T1{y: 6}, T0: &T0{x: 4}, I:IImpl{}} // with t.T0 != nil
var xpv = 42
var p *T2 = &t // with p != nil and (*p).T0 != nil
var q Q = p


func main() {
	//t0 := T0 {}
	//fmt.Println(t0.M0())

	// T1
	t1 := T1{12}
	t1p := &T1{12}
	fmt.Println(t1.M1())
	fmt.Println(t1p.M1())
	fmt.Println(t1.M3())
	fmt.Println(t1p.M3())
	fmt.Println(T1.M1(t1))
	//fmt.Println(T1.M3(t1)) // Error: invalid method expression T1.M3 (needs pointer receiver: (*T1).M3)
	fmt.Println((*T1).M1(&t1))
	fmt.Println((*T1).M3(t1p))
	pt1 := PT1(t1p)
	fmt.Println(pt1.y)

	// T2
	fmt.Println(t.M0())
	fmt.Println(t.M1())
	fmt.Println(t.M3())
	fmt.Printf("(%v, %[1]T)\n", t.I)
	fmt.Println(t.M4())
	fmt.Println(t.y)
	t.y = 42
	fmt.Println(t.y)

	fmt.Println(t.z)  // t.z
	fmt.Println(t.x)  // (*t.T0).x
	fmt.Println(t.xp) // (*t.T0).xp
	fmt.Println(t.y)  // t.T1.y

	fmt.Println(p.z)  // (*p).z
	fmt.Println(p.y)  // (*p).T1.y
	fmt.Println(p.x)  // (*(*p).T0).x
	fmt.Println(p.xp) // (*(*p).T0).xp

	fmt.Println("????", q.x) // (*(*q).T0).x        (*q).x is a valid field selector
	fmt.Println("!!!", q.xp) // (*(*q).T0)).xp is a valid field selector
	//fmt.Println(*(*(*q).T0).xp) // (*(*q).T0).xp is a valid field selector, but xp is nil pointer
	// and can not be derefeenced

	fmt.Println(t.M2())    // (&t).M2()           M2 expects *T2 receiver
	fmt.Println(t.M1())    // (&t).M1()           M1 expects T1 receiver
	fmt.Println(t.M3())    // (&t).M3()           M3 expects *T1 receiver
	fmt.Println(t.M0())    // (&t).M0()           M0 expects *T0 receiver

	fmt.Println(p.M0())    // ((*p).T0).M0()      M0 expects *T0 receiver
	fmt.Println(p.M1())    // ((*p).T1).M1()      M1 expects T1 receiver
	fmt.Println(p.M2())    // p.M2()              M2 expects *T2 receiver
	fmt.Println(p.M3())    // p.M2()              M2 expects *T2 receiver
	fmt.Println(p.M4())

	fmt.Println((*q).M0()) // (*q).M0()           M0 expects *T0 receiver
	//fmt.Println(q.M0())    // (*q).M0 is valid but not a field selector
}
