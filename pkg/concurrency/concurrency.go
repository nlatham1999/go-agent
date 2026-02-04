package concurrency

type Concurrency struct {
}

type PoolTurtle struct {
	channelIn  chan interface{}
	channelOut chan interface{}
}
