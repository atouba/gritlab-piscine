package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPr(l List) {
	node := l.Head
	for node != nil {
		node = node.Next
	}
}

// ---- why this doesn't work
// func ListPushBack(l *List, data interface{}) {
//   node := l.Head
//   for node != nil  {
//     node = node.Next
//   }
//   node = &NodeL{Data: data, Next: nil}
//
//   fm . rintln("print list:")
//   ListPr(*l)
// }

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		return
	}
	it := l.Head
	for it.Next != nil {
		it = it.Next
	}
	it.Next = newNode
	l.Tail = newNode
}
