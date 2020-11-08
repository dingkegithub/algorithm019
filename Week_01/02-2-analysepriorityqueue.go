package Week_01

import (
	"os"
	"sync"

	"github.com/dingkegithub/algorithm019/utils"
)

/*
java 文档中说明优先队列实现的其中一个接口就是 Queue<E>,

它使用优先堆组织数据，实际底层存储时一个数组


主要的方法

- boolean add(E e): 插入特定元素到队列中,会抛出异常

- clear(): 清除队列中的所有元素

- comparator(): 比较器，用来优先级比较，若为null 表示用自然排序

- boolean offer(E e): 插入元素，异常时返回false

- boolean remove(Object o): 删除元素，会抛出异常

翻看源码发现底层存储数据他定义了

```
// 底层是一个数组
E[] storage;
```

经过源码分析，本质来说 java 优先队列是用堆来实现的，需要额外提供一个 compare 用来比较优先级,

下面根据原理来实现
*/
type PriorityQueue struct {
	// 底层存储
	E []interface{}

	// 比较函数
	compare utils.Compare

	// 实际数组使用量
	used int

	mutex sync.RWMutex
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		compare: utils.CompareNum,
		E:       make([]interface{}, 10, 10),
	}
}

// clear all instance in queue
func (p *PriorityQueue) Clear() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for i := range p.E {
		p.E[i] = nil
	}
	p.used = 0
}

// add element into queue, failed with error
func (p *PriorityQueue) Add(e interface{}) (bool, error) {
	if e == nil {
		return false, os.ErrInvalid
	}

	return p.Offer(e), nil
}

// 添加元素，成功返回 true， 异常时返回false
func (p *PriorityQueue) Offer(e interface{}) bool {
	if e == nil {
		return false
	}

	// 放到空位置
	slot := p.findSlot(-1)

	p.E[slot] = e
	p.used += 1

	// 根据各元素优先级调整堆
	p.bubleUp(slot)
	return true
}

// 删除元素
// 获取并删除队列头元素，成功返回队列头元素， 异常时返回error
func (p *PriorityQueue) Remove() (interface{}, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.used == 0 {
		return nil, ErrBlank
	}

	d := p.E[0]
	// 删除队列头元素
	p.remove(0)
	return d, nil
}

// 删除元素
// 获取并删除队列头元素，成功返回队列头元素， 异常时返回nil
func (p *PriorityQueue) Poll() interface{} {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.used == 0 {
		return nil
	}

	d := p.E[0]
	// 删除队列头元素
	p.remove(0)
	return d
}

// 获取但不删除元素，成功返回队列头元素， 异常时返回error
func (p *PriorityQueue) Element() (interface{}, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.used == 0 {
		return nil, ErrBlank
	}

	return p.E[0], nil

}

// 获取但不删除元素，成功返回队列头元素， 异常时返回nil
func (p *PriorityQueue) Peek() interface{} {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.used == 0 {
		return nil
	}

	return p.E[0]
}

func (p *PriorityQueue) findSlot(start int) int {
	slot := 0

	if p.used == len(p.E) {
		p.resize()
		slot = p.used
	} else {

		for slot = start + 1; slot < len(p.E); slot++ {
			if p.E[slot] == nil {
				break
			}
		}
	}

	return slot
}

// 确保堆的容量
func (p *PriorityQueue) resize() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	data := make([]interface{}, 2*len(p.E))
	for i, v := range p.E {
		data[i] = v
	}

	p.E = data
}

// 删除元素
// 堆中删除元素要自上而下调整, 优先级大的向上调整
func (p *PriorityQueue) remove(idx int) {
	i := idx

	for p.E[i] != nil {
		// 找到左子节点
		child := (i << 1) + 1

		// 	没有子节点，置空当前节点，成功删除返回
		if child >= len(p.E) {
			p.E[i] = nil
			break
		}

		// 只有左子节点，没有右子节点时，下面将提升左子节点到父节点位置
		if child+1 >= len(p.E) || p.E[child+1] == nil {

		} else if p.E[child] == nil || p.compare(p.E[child], p.E[child+1]) < 0 {
			// 左子节点为空，或者左子节点优先级大于右子节点时，child指向右子节点，下面将提升右子节点为父节点
			child += 1
		}

		// 子节点提升为父节点，相当于父节点删除，向下继续搜寻，填补子节点空缺(因为子节点补了父节点)
		p.E[i] = p.E[child]
		i = child
	}

	p.used -= 1
}

// 堆调整
// 优先堆的调整需要字底向上调整
func (p *PriorityQueue) bubleUp(slot int) {
	idx := slot

	for idx > 0 {
		// 求出父节点的位置
		parent := (idx - 1) >> 1

		// 父节点的优先级大于子节点，则调整结束
		if p.compare(p.E[parent], p.E[idx]) >= 0 {
			break
		}

		// 父节点优先级小于子节点，子节点向上调整，父节点放到子节点位置
		p.E[parent], p.E[idx] = p.E[idx], p.E[parent]
		idx = parent
	}
}
