package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Command_pattern
	Паттерн Command преобразовывает запрос на выполнение действия в отдельный объект-команду.
	Плюcы:
	S и O принципы SOLID.
	Возможность создания последовательности комманд.
	Минусы:
	Создаётся слой между отправителем и приёмником.
*/
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}
type AddCommand struct {
	game GameRule
}
type MoveCommand struct {
	game GameRule
}
type StopCommand struct {
	game GameRule
}

func (c *MoveCommand) execute() {
	c.game.move()
}

func (c *AddCommand) execute() {
	c.game.add()
}
func (c *StopCommand) execute() {
	c.game.stop()
}

type GameRule interface {
	move()
	add()
	stop()
}

type Game struct {
	isMoving  bool
	armyCount int
}

func (t *Game) move() {
	t.isMoving = true
	fmt.Println("Army moved")
}

func (t *Game) add() {
	t.armyCount++
}

func (t *Game) stop() {
	t.isMoving = false
	fmt.Println("Army stopped")
}

func main() {
	game := &Game{}

	addCommand := &AddCommand{
		game: game,
	}

	moveCommand := &MoveCommand{
		game: game,
	}
	stopCommand := &StopCommand{
		game: game,
	}
	addButton := &Button{
		command: addCommand,
	}
	addButton.press()

	moveButton := &Button{
		command: moveCommand,
	}
	moveButton.press()

	stopButton := &Button{
		command: stopCommand,
	}
	stopButton.press()
}
