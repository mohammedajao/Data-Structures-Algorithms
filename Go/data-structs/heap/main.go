// Implementing min heap

package main

type Heap struct {
	data   []int
	head   int
	length int
}

func NewHeap() *Heap {
	return &Heap{
		data: make([]int, 0),
	}
}

func (h *Heap) GetMinimum() int {
	return h.data[0]
}

func (h *Heap) Insert(item int) {
	for i := 0; i <= h.length; i++ {
		if h.data[(2*i)+1] == 0 {
			h.data[(2*i)+1] = item
		} else if h.data[(2*i)+2] == 0 {
			h.data[(2*i)+2] = item
		}
	}
}

func main() {

}
