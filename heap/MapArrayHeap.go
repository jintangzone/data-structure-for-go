package heap

type MaxArrayHeap struct {
	data []int
}

func (mah *MaxArrayHeap) GetSize() int {
	return len(mah.data)
}

func (mah *MaxArrayHeap) IsEmpty() bool {
	return mah.GetSize() != 0
}

func (mah *MaxArrayHeap) Add(e int)  {
	mah.data = append(mah.data, e)
	mah.siftUp(len(mah.data)-1)
}

func (mah *MaxArrayHeap) ExtractMax() int {
	ret := mah.FindMax()
	mah.swap(0, len(mah.data)-1)
	mah.data = mah.data[:len(mah.data)-1]
	mah.siftDown(0)
	return ret
}

func (mah *MaxArrayHeap) FindMax() int {
	if mah.GetSize() == 0 {
		panic("Can't findMax when heap")
	}
	return mah.data[0]
}

func (mah *MaxArrayHeap) parent(index int) int {
	if index == 0 {
		panic("index-0 doesn't have parent.")
	}
	return (index-1)/2
}

func (mah *MaxArrayHeap) leftChild(index int) int {
	return index*2+1
}

func (mah *MaxArrayHeap) rightChild(index int) int {
	return index*2+2
}

func (mah *MaxArrayHeap) siftUp(index int)  {
	// 首先当前节点必须大于零，因为一直上浮到堆顶，index是0，则退出
	// 每一次都比较当前节点是否大于父亲节点
	for index > 0 && mah.data[mah.parent(index)] < mah.data[index] {

		// 与父亲节点交换
		mah.swap(index, mah.parent(index))

		// 然后把父亲节点当成当前节点，继续上浮
		index = mah.parent(index)
	}
}

func (mah *MaxArrayHeap) siftDown(index int) {
	for mah.leftChild(index) < len(mah.data) {

		max := mah.leftChild(index) // 先让左节点左最大值
		right := mah.rightChild(index)

		if right < len(mah.data) {
			// 比较左右节点的大小, 取最大值得索引
			if mah.data[right] > mah.data[max] {
				max = right
			}
		}

		// 当当前节点的值，比左右节点中最大值得节点都大，则没必要调整了
		if mah.data[index] >= mah.data[max] {
			break
		}

		// 否则交换当前节点与左右节点中的最大值节点
		mah.swap(index, max)

		// 然后接续下沉
		index = max
	}
}

func (mah *MaxArrayHeap) swap(i, j int)  {
	if i < 0 || i >= len(mah.data) || j < 0 || j >= len(mah.data) {
		panic("index is illegal.")
	}

	t := mah.data[i]
	mah.data[i] = mah.data[j]
	mah.data[j] = t
}

func NewMaxArrayHeap(capacity int) *MaxArrayHeap {
	return &MaxArrayHeap{
		data: make([]int, capacity),
	}
}
