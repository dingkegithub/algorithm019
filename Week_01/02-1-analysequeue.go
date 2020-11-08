package Week_01

/*
java 中的queue是一个接口，主要提供了两组方法，

```
+--------+------------+------------------------+
|        |抛出异常    |   特殊值代表失败       |
+--------+------------+------------------------+
| Insert | add(e)     |   offer(e)             |
+--------+------------+------------------------+
| Remove | remove()   |   poll()               |
+--------+------------+------------------------+
| Examine| element()  |  peek()                |
+--------+------------+------------------------+
```

- boolean add(E e):  插入元素到队列，空间不足时抛出异常

- boolean offer(E e):  插入元素到队列，空间不足返回false

- E element(): 获取队列头元素，但不删除元素, 队列空抛出异常

- E peek():  获取队列头元素，但不删除元素, 队列空返回 null

- E remove()：获取和删除队列头的元素，队列空抛出异常

- E poll()：获取和删除队列头的元素，队列空返回null

这里转换成 go 进行实现
*/
type Queue interface {
	// 添加元素，成功返回 true， 异常时返回error
	Add(e interface{}) (bool, error)

	// 添加元素，成功返回 true， 异常时返回false
	Offer(e interface{}) bool

	// 获取并删除队列头元素，成功返回队列头元素， 异常时返回error
	Remove() (interface{}, error)

	// 获取并删除队列头元素，成功返回队列头元素， 异常时返回nil
	Poll() interface{}

	// 获取但不删除元素，成功返回队列头元素， 异常时返回error
	Element() (interface{}, error)

	// 获取但不删除元素，成功返回队列头元素， 异常时返回nil
	Peek() interface{}
}
