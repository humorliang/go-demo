package state

import (
	"fmt"
)

/*
状态模式（State Pattern）
   核心思想就是：当对象的状态改变时，同时改变其行为，很好理解！
   就拿扣扣来说，有几种状态，在线、隐身、忙碌等，每个状态对应不同的操作，
   而且你的好友也能看到你的状态， 所以，状态模式就两点：
		1、可以通过改变状态来获得不同的行为。
	   	2、你的好友能同时看到你的变化

1. Context类维护ConcreteState子类的实例，这个实例定义当前状态
2. 定义一个抽象接口定义以封装与Context特定状态的相关行为
3. ConcreteState子类的定义,每个子类实现与Context的一个状态的相关行为

// 相关模式：命令模式
*/

//eg:游戏人物在不同等级不同技能
//定义Context类
type CharacterLevel struct {
	//人物等级
	Level int
	//人物状态
	State State
}

//设置状态
func (cl *CharacterLevel) SetStatus(state State) {
	cl.State = state
}

//获取状态
func (cl *CharacterLevel) GetStatus() State {
	return cl.State
}

//设置等级
func (cl *CharacterLevel) SetLevel(level int) {
	cl.Level = level
}

//获取等级
func (cl *CharacterLevel) GetLevel() int {
	return cl.Level
}

//执行技能
func (cl *CharacterLevel) Skill() {
	//获取当前的状态  并执行相对应状态的技能
	cl.State.Skill(*cl)
}

//抽象状态 接口
type State interface {
	Skill(CharacterLevel)
}

//定义初级状态
type PrimaryLevel struct {
}

func (self *PrimaryLevel) Skill(cl CharacterLevel) {
	if cl.Level <= 10 {
		fmt.Println("当前等级：", cl.Level, " 此时为初级状态，只有一个技能")
	} else {
		//转换状态
		cl.SetStatus(new(MiddleLevel))
		cl.Skill()
	}
}

//定义中极
type MiddleLevel struct {
}

func (self *MiddleLevel) Skill(cl CharacterLevel) {
	if cl.Level <= 20 {
		fmt.Println("当前等级：", cl.Level, " 此时为中级状态，有五个技能")
	} else {
		//转换状态
		cl.SetStatus(new(AdvancedLevel))
		cl.Skill()
	}
}

//定义高级
type AdvancedLevel struct {
}

func (self *AdvancedLevel) Skill(cl CharacterLevel) {
	fmt.Println("当前等级：", cl.Level, " 此时为高级状态，有十个技能")
}
