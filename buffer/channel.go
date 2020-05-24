package buffer

import "errors"

type ChanBuffer struct {
	Channel  chan int // 信道
	Capacity int      // 容量
}

func NewChanBuffer() *ChanBuffer {
	cap := int(1e6)
	chanBuffer := ChanBuffer{
		Capacity: cap,
		Channel:  make(chan int, cap),
	}
	return &chanBuffer
}

func (c *ChanBuffer) Put(val int) {
	c.Channel <- val
}

func (c *ChanBuffer) Get() (int, error) {
	if len(c.Channel) > 0 {
		val := <-c.Channel
		return val, nil
	}
	return 0, errors.New("")

}
