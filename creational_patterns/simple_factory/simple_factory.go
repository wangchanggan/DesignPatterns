/*
在简单工厂模式中，可以根据参数的不同返回不同类的实例。

简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。
*/

package simple_factory

type Factory struct {

}

type A struct {

}

type B struct {

}

type Interface interface {
	Say() string
}

func (a *A) Say() string {
	return "A"
}
func (b *B) Say() string {
	return "B"
}

func (factory *Factory) Create(name string) Interface {
	switch name {
	case `a`:
		return new(A)
	case `b`:
		return new(B)
	default:
		panic(`name not exists`)
	}
	return nil
}