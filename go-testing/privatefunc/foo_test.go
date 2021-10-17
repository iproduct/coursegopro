package privatefunc

import (
	"testing"
)

func TestDoSomething(t *testing.T) {
	DoSomething(true, "Hello")
}

func TestDoPrivate(t *testing.T) {
	doPrivateThings("hello")
}
