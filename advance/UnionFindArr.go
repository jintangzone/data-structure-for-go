package advance

import (
	"strconv"
	"strings"
)

type UnionFindArr struct {
	id []int
}

func (uf *UnionFindArr) GetSize() int {
	return len(uf.id)
}

func (uf *UnionFindArr) find(p int) int {
	if p < 0 || p >= len(uf.id) {
		panic("p is out of bound")
	}
	return uf.id[p]
}

func (uf *UnionFindArr) IsConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf * UnionFindArr) Union(p, q int) {
	pId := uf.find(p)
	qId := uf.find(q)
	if pId == qId {
		return
	}
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pId {
			uf.id[i] = qId
		}
	}
}

func (uf *UnionFindArr) ToString() string {
	str := make([]string, 0)
	ret := "union: ["
	for i := 0; i < len(uf.id); i++ {
		str = append(str, strconv.Itoa(i) + "=>" + strconv.Itoa(uf.id[i]))
	}
	ret = ret + strings.Join(str, ",")
	ret = ret + "]"
	return ret
}

func NewUnionFindArr(size int) *UnionFindArr {
	uf := new(UnionFindArr)
	uf.id = make([]int, size)
	for i := 0; i < size; i++ {
		uf.id[i] = i
	}
	return uf
}

