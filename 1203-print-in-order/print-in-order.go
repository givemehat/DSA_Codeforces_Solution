type Foo struct {
    ch1 chan struct{}
    ch2 chan struct{}
}

func NewFoo() *Foo {
    return &Foo{
        ch1: make(chan struct{}),
        ch2: make(chan struct{}),
    }
}

func (f *Foo) First(printFirst func()) {
    printFirst()
    close(f.ch1)
}

func (f *Foo) Second(printSecond func()) {
    <-f.ch1

    printSecond()
    close(f.ch2)
}

func (f *Foo) Third(printThird func()) {
    <-f.ch2

    printThird()
}