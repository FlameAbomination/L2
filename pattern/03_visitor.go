package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/
/*
	Применяется для добавления функциональности без изменения кода объекта.
	Плюcы:
	S и O принципы SOLID.
	Минусы:
	При изменении полей структуры появляется необходимость изменения функции.
	Если структура находится в другом модуле, то к её приватным полям не будт доступа.
	В Go нет перегрузки функций, поэтому для каждого типа необходимо создавать отдельную функцию.

*/

type IfNode struct {
}
type AllocateNode struct {
}
type OperationNode struct {
}

type Visitor interface {
	visitIfNode(node IfNode) string
	visitAllocateNode(node AllocateNode) string
	visitOperationNode(node OperationNode) string
}
