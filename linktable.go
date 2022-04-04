package main

import "fmt"

type LinkNode struct {
	val  int
	next *LinkNode
}

type LinkTable struct {
	pHead *LinkNode
	pTail *LinkNode
	count int
}

func CreateLinkTable() *LinkTable {
	link := &LinkTable{
		pHead: nil,
		pTail: nil,
		count: 0,
	}
	return link
}

func (l *LinkTable) DeleteLinkTable() {
	l.pHead = nil
	l.pTail = nil
	l.count = 0
}

func (l *LinkTable) Add(v int) {
	node := LinkNode{v, nil}

	if l.count == 0 {
		l.pHead = &node
		l.pTail = &node
	} else {
		l.pTail.next = &node
		l.pTail = &node
	}
	l.count++
	fmt.Printf("Add value %d\n", v)
	return
}

func (l *LinkTable) Delete(v int) {
	cur := l.pHead

	if l.count == 0 {
		return
	}

	if cur.val == v {
		if l.pHead.next != nil {
			l.pHead = l.pHead.next
		} else {
			l.pHead = nil
		}
		l.count--
		return
	}

	for cur.next != nil {
		if cur.next.val == v {
			fmt.Printf("Delete value %d\n", v)
			if cur.next.next != nil {
				cur.next = cur.next.next
			} else {
				cur.next = nil
			}
			l.count--
		}
		cur = cur.next
	}
}

func (l *LinkTable) Find(v int) *LinkNode {
	if l.count == 0 {
		return nil
	} else {
		cur := l.pHead

		for cur.next != nil {
			if cur.val == v {
				return cur
			}
			cur = cur.next
		}
		return nil
	}
}

func (l *LinkTable) GetNextNode(v int) *LinkNode {
	if l.count == 0 || l.count == 1 {
		return nil
	} else {
		cur := l.pHead

		for cur.next != nil {
			if cur.val == v {
				return cur.next
			}
			cur = cur.next
		}
		return nil
	}
}

func (l *LinkTable) Print() {
	prt := l.pHead

	fmt.Print("LinkTable List: ")
	for prt.next != nil {
		fmt.Printf("%d ", prt.val)
		prt = prt.next
	}
	fmt.Println(prt.val)
}

func main() {
	var link *LinkTable
	link = CreateLinkTable()
	link.Add(1)
	link.Add(8)
	link.Add(3)
	link.Add(4)
	link.Add(6)
	link.Add(2)
	link.Add(5)
	link.Add(7)
	link.Print()
	link.Delete(2)
	link.Print()
	node := link.Find(3)
	fmt.Println(*node)
	next := link.GetNextNode(6)
	fmt.Println(*next)
}
