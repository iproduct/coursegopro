package main

import (
	"fmt"
	"reflect"
)

type MyStack struct {
	t     reflect.Type `json:"type"`
	value reflect.Value `json:"value"`
}

func New(tp reflect.Type) *MyStack {
	return &MyStack{
		t:     tp,
		value: reflect.MakeSlice(reflect.SliceOf(tp), 0, 100),
	}
}

func (m *MyStack) Push(v interface{}) {
	if reflect.ValueOf(v).Type() != m.value.Type().Elem() {
		panic(fmt.Sprintf("Error putting %T '%[1]v' into a container of type %s", v, m.value.Type().Elem()))
	}

	m.value = reflect.Append(m.value, reflect.ValueOf(v))
}

func (m *MyStack) Pop() interface{} {
	if m.value.Len() == 0 {
		return nil
	}
	v := m.value.Index(m.value.Len() - 1) // []
	m.value = m.value.Slice(0, m.value.Len() - 1)
	return v.Interface()
}

func main() {
	val := 2.88
	fmt.Println(reflect.TypeOf(val))
	stack := New(reflect.TypeOf(val))
	fmt.Printf("%#v \n", reflect.TypeOf(*stack).String())
	for index, field := range reflect.VisibleFields(reflect.TypeOf(*stack)) {
		fmt.Printf("%#v -> %#v : %#v : %#v\n", index, field.Name, field.Type.Name(), field.Tag)
	}
	//stack.Push(val)
	//stack.Push(3.14)
	//stack.Push(2.895)
	//stack.Push(42.0)
	//stack.Push(135.5)
	//fmt.Println(stack.value.Index(0))
	//for result := stack.Pop(); result != nil; result = stack.Pop() {
	//	fmt.Printf("Result: %[1]f (%[1]T)\n", result)
	//}

}
