package main

type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte

	select {
	case v = <-p:
	default: //返回
		v = make([]byte, 10) //返回失败，新建
	}
	return v
}

func (p pool) put(b []byte) {
	select {
	case p <- b: //放回
	default: //返回失败，放弃
	}

}
