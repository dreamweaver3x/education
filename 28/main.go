package main

import "fmt"

type greeter interface {
	hello()
}

type engGreet struct {
	msg string
}

func (e engGreet) hello() {
	if e.msg == "hello" {
		fmt.Println("hello to you as well!")
	}
}

type rusGreet struct {
	msg string
}

func (r rusGreet) helloInRus() {
	if r.msg == "привет" {
		fmt.Println("и тебе привет!")
	}
}

type rusAdapter struct {
	rusLang *rusGreet
}

func (r rusAdapter) hello() {
	if r.rusLang.msg == "hello" {
		r.rusLang.msg = "привет"
		r.rusLang.helloInRus()
	}
}

type client struct {
}

func (c client) greet(g greeter) {
	g.hello()
}

func main() {
	engMsg := engGreet{msg: "hello"}
	rusMsg := rusGreet{msg: "hello"}
	rusAd := rusAdapter{rusLang: &rusMsg}

	c := client{}

	c.greet(engMsg)
	c.greet(rusAd)

}
