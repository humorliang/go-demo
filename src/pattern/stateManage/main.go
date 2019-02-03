package main

import "pattern/stateManage/state"

func main() {
	//初始化
	charaLevel := &state.CharacterLevel{Level: 0, State: new(state.PrimaryLevel)}
	charaLevel.SetLevel(5)
	charaLevel.Skill()
	charaLevel.SetLevel(15)
	charaLevel.Skill()
	charaLevel.SetLevel(25)
	charaLevel.Skill()
	charaLevel.SetLevel(35)
	charaLevel.Skill()
	/*
	当前等级： 5  此时为初级状态，只有一个技能
	当前等级： 15  此时为中级状态，有五个技能
	当前等级： 25  此时为高级状态，有十个技能
	当前等级： 35  此时为高级状态，有十个技能
	*/
}
