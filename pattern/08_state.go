package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
	Паттерн применяется для упрощения работы с состояниями.
	Плюcы:
	S и O принципы SOLID.
	Минусы:
	Лишняя сложность при малом количестве состояний.
*/

type State interface {
	GetName() string
	Freeze()
	Heat()
}

type StateContext struct {
	State
}

type SolidState struct {
	state StateContext
	name  string
}
type GasState struct {
	state StateContext
	name  string
}
type LiquidState struct {
	state StateContext
	name  string
}

func (c *StateContext) SetState(state State) {
	c = &StateContext{state}
}

func (c SolidState) Freeze() {
	fmt.Println("Nothing happens")
}

func (c SolidState) Heat() {
	c.state.SetState(LiquidState{c.state, "Liquid"})
}

func (c SolidState) GetName() string {
	return c.name
}

func (c LiquidState) Freeze() {
	c.state.SetState(SolidState{c.state, "Solid"})
}

func (c LiquidState) Heat() {
	c.state.SetState(GasState{c.state, "Gas"})
}

func (c LiquidState) GetName() string {
	return c.name
}

func (c GasState) Freeze() {
	c.state.SetState(LiquidState{c.state, "Liquid"})
}

func (c GasState) Heat() {
	fmt.Println("Nothing happens")
}

func (c GasState) GetName() string {
	return c.name
}

func Test() {
	var sc StateContext
	sc = StateContext{SolidState{sc, "Solid"}}
	sc.Heat()
	sc.Heat()
	sc.Heat()
	sc.Freeze()
	sc.Freeze()
	sc.Freeze()
}
