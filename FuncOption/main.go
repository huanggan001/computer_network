package main

import "fmt"

type config struct {
	a string
	b string
	c int
}

// NewConfig 最简单的方式，直接写死，但有多个字段时怎么处理
func NewConfig(a, b string, c int) *config {
	return &config{a, b, c}
}

type ConfigOption func(*config)

func WithSetConfigC(C int) func(*config) {
	return func(c *config) {
		c.c = C
	}
}

var defaultC = 18

func NewConfigOption(a, b string, opts ...ConfigOption) *config {
	cc := &config{
		a: a,
		b: b,
		c: defaultC,
	}
	for _, opt := range opts {
		opt(cc)
	}
	return cc
}

// 进阶用法，通过接口的形式
type serviceConfigOption interface {
	apply(*config)
}

// 定义funcOption
type funcOption struct {
	f func(*config)
}

func (f *funcOption) apply(c *config) {
	f.f(c)
}

func NewfuncOption(f func(f *config)) serviceConfigOption {
	return &funcOption{f: f}
}

func WithSetConfigCByInterface(C int) serviceConfigOption {
	return NewfuncOption(func(f *config) {
		f.c = C
	})
}

func NewConfigOptionByInterface(a, b string, opts ...serviceConfigOption) *config {
	cc := &config{
		a: a,
		b: b,
		c: defaultC,
	}

	for _, opt := range opts {
		opt.apply(cc)
	}
	return cc
}

func main() {
	////不设置c的值，就使用默认c值
	//cc := NewConfigOption("hh", "gg")
	//fmt.Println(cc)
	////设置c的值
	//cc2 := NewConfigOption("hh", "gg", WithSetConfigC(20))
	//fmt.Println(cc2)

	//不设置c的值，就使用默认c值
	cc := NewConfigOptionByInterface("hh", "gg")
	fmt.Println(cc)
	//设置c的值
	cc2 := NewConfigOptionByInterface("hh", "gg", WithSetConfigCByInterface(22))
	fmt.Println(cc2)
}
