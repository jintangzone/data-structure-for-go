package advance

type QuickUnion struct {
	parent []int
}

func (qu *QuickUnion) GetSize() int {
	return len(qu.parent)
}

func (qu *QuickUnion) find(p int) int {
	if p < 0 || p >= len(qu.parent) {
		panic("p is out of range.")
	}
	for p != qu.parent[p] {
		p = qu.parent[p]
	}
	return p
}

func (qu *QuickUnion) IsConnected(p, q int) bool {
	return qu.find(p) == qu.find(q)
}

func (qu *QuickUnion) Union(p, q int) {
	pRoot := qu.find(p)
	qRoot := qu.find(q)
	if pRoot == qRoot {
		return
	}
	qu.parent[pRoot] = qRoot
}

func NewQuickUnion(size int) *QuickUnion {
	qu := new(QuickUnion)
	qu.parent = make([]int, size)
	for i := 0; i < size; i++ {
		qu.parent[i] = i
	}
	return qu
}