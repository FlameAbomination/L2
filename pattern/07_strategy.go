package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
	Паттерн применяется для переключения между алгоритмами.
	Плюcы:
	Замена наследования на композицию
	O принцип SOLID.
	Минусы:
	Код переусложняется, если существует небольшое количество алгоритмов.
*/

type SortingAlgorithm interface {
	sort(c []int)
}
type BubbleSort struct {
}

func (s *BubbleSort) sort(c []int) {

}

type QuickSort struct {
}

func (s *QuickSort) sort(c []int) {

}

type Array struct {
	sortingAlgorithm SortingAlgorithm
	data             []int
}

func (a *Array) getData() {

}

func initArray(e SortingAlgorithm) *Array {
	return &Array{
		sortingAlgorithm: e,
	}
}

func (a *Array) setSortingAlgorithm(s SortingAlgorithm) {
	a.sortingAlgorithm = s
}

func sorting() {
	var bubbleSort *BubbleSort
	var quickSort *QuickSort
	array := initArray(bubbleSort)
	array.getData()
	array.sortingAlgorithm.sort(array.data)
	array.getData()
	array.setSortingAlgorithm(quickSort)
	array.sortingAlgorithm.sort(array.data)
}
