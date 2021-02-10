package list

import (
	"errors"
	"fmt"
)

//节点的数据结构
type Node struct {
	pre  *Node
	next *Node
	data interface{}
}

//链表的数据结构
type DoubleList struct {
	first *Node
	last  *Node
	size  int
}

//迭代器结构
type LinkedIterator struct {
	list   *DoubleList
	cursor *Node
	end    *Node
}

// 给链表末尾新增一个节点
func (list *DoubleList) Add(obj ...interface{}) error {
	for _, elem := range obj {
		p := new(Node)
		p.data = elem
		list.last.next = p
		p.pre = list.last
		list.last = p
		list.size++
	}
	return nil
}

//向指定位置加入元素
func (list *DoubleList) Insert(location int, obj interface{}) error {
	if location <= 0 || location > list.size {
		return errors.New("位置超出")
	}
	//找相应位置元素
	p := list.first
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	//在相应位置元素进行赋值
	q := new(Node)
	q.next = p.next
	q.pre = p
	p.next = q
	q.data = obj
	list.size++
	return nil
}

//向指定位置修改元素
func (list *DoubleList) Set(location int, obj interface{}) error {
	if location <= 0 || location > list.size {
		return errors.New("位置超出")
	}
	//找相应位置元素
	p := list.first
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	//赋值
	p.data = obj
	return errors.New("no error")
}

//是否存在某元素
func (list *DoubleList) Contain(obj interface{}) bool {
	//找相应位置元素
	p := list.first
	for p.next != nil && p.data != obj {
		p = p.next
	}
	//若已经走到尽头该值不等于要找的值则返回false否则返回false
	if p.next == nil || p.data != obj {
		return false
	}
	return true
}

//是否为空
func (list *DoubleList) IsEmpty() bool {
	if list.last == list.first {
		return true
	}
	return false
}

//查看某一位置上的元素
func (list *DoubleList) Get(location int) (interface{}, error) {
	if location <= 0 || location > list.size {
		return nil, errors.New("位置超出")
	}
	//找相应位置元素
	p := list.first
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	return p.data, nil
}

//判断是否相等
//这里使用迭代器对里面元素进行一一比较
func (list *DoubleList) Equals(list1 List) bool {
	if list.Size() != list1.Size() {
		return false
	}
	it := list.Iterator()
	it1 := list1.Iterator()
	for it.HasNext() && it1.HasNext() {
		elem, _ := it.Next()
		elem1, _ := it1.Next()
		if elem != elem1 {
			break
		}
	}
	if it.HasNext() == false && it1.HasNext() == false {
		return true
	}
	return false
}

//转换为Slice类型
func (list *DoubleList) ToSlice() []interface{} {
	//定义一个切片进入循环将值一一赋给切片你即可
	p := list.first.next
	sli := make([]interface{}, list.size)
	i := 0
	//如果为空链表则直接返回一个空切片
	if list.size == 0 {
		return sli
	}
	for p != nil {
		sli[i] = p.data
		p = p.next
		i++
	}
	return sli
}

//输出当前list的长度
func (list *DoubleList) Size() int {
	return list.size
}

// 打印(遍历)链表
func (list *DoubleList) Print() {
	fmt.Printf("\n链表输出为:")
	p := list.first.next
	for p != nil {
		fmt.Print(p.data)
		fmt.Print(" ")
		p = p.next
	}

}

// 创建一个空的双链表
func NewDoubleList() (list *DoubleList) {
	p := &DoubleList{nil, nil, 0}
	p.last = new(Node)
	p.first = p.last
	return p
}

//迭代器
func (list *DoubleList) Iterator() Iterator {
	it := new(LinkedIterator)
	it.list = list
	it.cursor = list.first.next
	it.end = list.first.next
	return it
}

//判断是否存在下一个元素
func (it *LinkedIterator) HasNext() bool {
	//如果当前下标等于它的大小说明没有下一个元素了
	return it.cursor != nil
}

//返回下一个元素
func (it *LinkedIterator) Next() (interface{}, error) {
	//首先获取当前下标的位置
	i := it.cursor
	if i == nil {
		return nil, errors.New("没有这样的索引")
	}
	//下标位置往后移
	it.cursor = it.cursor.next
	it.end = i
	return it.end.data, nil
}

func Doubletest() {
	list1 := NewDoubleList()
	list2 := NewDoubleList()
	list3 := NewDoubleList()
	list3.Add(1, 2, 3, 4)
	list1.Add(1, 2, 3, 4)
	list2.Add(1, 2, 3, 4)
	list1.Print()
	fmt.Println(list1.Equals(list2))
	list1.Insert(2, 5)
	list1.Print()
	fmt.Println(list1.Contain(5))
	list1.Set(3, 1)
	list1.Print()
	fmt.Println(list1.IsEmpty())
	fmt.Println(list1.Get(4))
	fmt.Println(list1.Equals(list2))
	fmt.Println(list1.ToSlice())
	fmt.Println(list1.Size())
}
