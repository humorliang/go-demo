package mem

import "container/list"

/*
备忘录模式（Memento pattern）
   在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。
   这样以后就可将该对象恢复到原先保存的状态。

相关：原形模式
*/

// 文本编辑
type Text struct {
	Value string
}

// 写
func (t *Text) Write(value string) {
	t.Value = value
}

// 读
func (t *Text) Read() string {
	return t.Value
}

//备忘结构
type Memento struct {
	Value string
}

// 备忘
func (t *Text) SaveToMemento() *Memento {
	return &Memento{
		Value: t.Value,
	}
}

//重备忘中进行恢复
func (t *Text) RestoreFromMemento(m *Memento) bool {
	if m != nil {
		t.Value = m.Value
		return true
	}
	return false
}

// 管理备忘录
type Storage struct {
	//创建一个双向链表
	*list.List
}

//返回链表的最后一个元素并删除
func (s *Storage) RightPop() *list.Element {
	//返回最后一个元素
	ele := s.Back()
	if ele != nil {
		s.Remove(ele)
	}
	return ele
}
