package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	firstItem *ListItem
	lastItem  *ListItem
	len       int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.firstItem
}

func (l *list) Back() *ListItem {
	return l.lastItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := ListItem{Value: v}
	if l.firstItem != nil {
		item.Next = l.firstItem
		l.firstItem.Prev = &item
	} else {
		l.lastItem = &item
	}
	l.firstItem = &item
	l.len++
	return l.firstItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := ListItem{Value: v}
	if l.lastItem != nil {
		item.Prev = l.lastItem
		l.lastItem.Next = &item
	} else {
		l.firstItem = &item
	}
	l.lastItem = &item
	l.len++
	return l.lastItem
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.firstItem = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.lastItem = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.firstItem == i {
		return
	}

	i.Prev.Next = i.Next
	if i.Next == nil {
		l.lastItem = i.Prev
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if l.firstItem != nil {
		l.firstItem.Prev = i
		i.Next = l.firstItem
	}

	i.Prev = nil
	l.firstItem = i
}

func NewList() List {
	return new(list)
}
