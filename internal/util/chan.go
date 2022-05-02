package util

type Chan[T any] chan T

func (ch *Chan[T]) Make()                { *ch = make(Chan[T]) }
func (ch Chan[T]) RecieveOnly() <-chan T { return ch }
func (ch Chan[T]) SendOnly() chan<- T    { return ch }
