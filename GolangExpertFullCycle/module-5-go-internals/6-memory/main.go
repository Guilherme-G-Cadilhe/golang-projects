package main

import (
	"fmt"
	"sync"
)

type Span struct {
	size      int
	allocated bool
}

type mheap struct {
	spans []*Span
	lock  sync.Mutex
}

type mcentral struct {
	sizeSpans []*Span
	lock      sync.Mutex
}

type mcache struct {
	localSpans []*Span
}

func NewHeap(size int) *mheap {
	h := &mheap{}
	for i := 0; i < size; i++ {
		h.spans = append(h.spans, &Span{size: i + 1})
	}
	return h
}

func (h *mheap) GetSpan(size int) *Span {
	// s = span
	h.lock.Lock()
	defer h.lock.Unlock()
	for _, s := range h.spans {
		if !s.allocated && s.size == size {
			s.allocated = true
			return s
		}
	}
	return nil
}

func (mc *mcentral) getSpanFromCentral(size int) *Span {
	// s = span
	mc.lock.Lock()
	defer mc.lock.Unlock()
	for _, s := range mc.sizeSpans {
		if !s.allocated && s.size == size {
			s.allocated = true
			return s
		}
	}
	return nil
}

func (mc *mcache) getSpanFromCache(size int) *Span {
	// s = span
	for _, s := range mc.localSpans {
		if !s.allocated && s.size == size {
			s.allocated = true
			return s
		}
	}
	return nil
}

func main() {

	// mcache -> mcentral -> mheap
	// Pede 5, se inicia com 4 dá sem memoria.
	heap := NewHeap(4)
	mcentral := &mcentral{
		sizeSpans: heap.spans,
	}
	mcache := &mcache{}

	requestSpan := mcache.getSpanFromCache(5)
	if requestSpan == nil {
		requestSpan = mcentral.getSpanFromCentral(5)
	}
	if requestSpan == nil {
		requestSpan = heap.GetSpan(5)
	}
	if requestSpan == nil {
		panic("Sem memória disponível")
	}

	fmt.Println("Span alocado:", requestSpan)

}
