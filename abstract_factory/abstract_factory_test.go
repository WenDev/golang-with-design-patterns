package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	c1 := Creator1{}
	c2 := Creator2{}
	a1 := c1.CreateProductA()
	a2 := c2.CreateProductA()
	b1 := c1.CreateProductB()
	b2 := c2.CreateProductB()
	a1.DoSomethingA()
	a2.DoSomethingA()
	b1.DoSomethingB()
	b2.DoSomethingB()
}
