package core

type Key interface {
	Then(k Key) int
	HashCode() int
}

