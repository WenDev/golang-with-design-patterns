package factory

import "testing"

func TestConcreteCreator_CreateProduct(t *testing.T) {
	c := &ConcreteCreator{}
	c1 := c.CreateProduct("1")
	c1.Method()
	c2 := c.CreateProduct("2")
	c2.Method()
}
