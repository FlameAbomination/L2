package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
	Паттерн используется для передачи структуры через последовательность обрабатывающих функций.
	Плюcы:
	S и O принципы SOLID.
	Возможность изменения последовательности обработки.
	Минусы:
	Необходимо отслеживать правильность цепочки.
*/
import "fmt"

type Package struct {
	name             string
	registrationDone bool
	paymentDone      bool
	recieved         bool
}

type Department interface {
	execute(*Package)
	setNext(Department)
}

type Registration struct {
	next Department
}
type Payment struct {
	next Department
}
type Mail struct {
	next Department
}

func (r *Registration) execute(p *Package) {
	if p.registrationDone {
		fmt.Println("Registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Registration) setNext(next Department) {
	r.next = next
}

func (r *Payment) execute(p *Package) {
	if p.registrationDone {
		fmt.Println("Payment already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Paying new delivery")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Payment) setNext(next Department) {
	r.next = next
}

func (r *Mail) execute(p *Package) {
	if p.registrationDone {
		fmt.Println("Delivery already done")
		r.next.execute(p)
		return
	}
	fmt.Println("REceiving new package")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *Mail) setNext(next Department) {
	r.next = next
}

func deliverPackage() {

	mail := &Mail{}

	payment := &Payment{}
	payment.setNext(mail)

	registration := &Registration{}
	registration.setNext(payment)

	delivery := &Package{name: "abc"}
	registration.execute(delivery)
}
