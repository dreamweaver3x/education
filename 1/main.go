package main

type Human struct {
	name string
	age  int
	Action
}

type Action struct {
	run  bool
	fly  bool
	play bool
}
