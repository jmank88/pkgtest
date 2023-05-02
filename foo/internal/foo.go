package internal

type Foo struct{}

func (*Foo) method() {}

type FooTest struct {
	Foo
}

func (f *FooTest) Method() { f.method() }
