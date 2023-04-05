package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
	Паттерн позволяет
	Плюcы:
	S и O принципы SOLID.
	Минусы:
	Необходимость создания новых структур.
*/

type IDevice interface {
	setArchitecture(name string)
	setRam(ram int)
	getArchitecture() string
	getRam() int
}

func (g *Device) setArchitecture(name string) {
	g.architecture = name
}

func (g *Device) getArchitecture() string {
	return g.architecture
}

func (g *Device) setRam(ram int) {
	g.ram = ram
}

func (g *Device) getRam() int {
	return g.ram
}

type Server struct {
	Device
}

func newServer() IDevice {
	return &Server{
		Device: Device{
			architecture: "x86_64",
			ram:          128,
		},
	}
}

type Mobile struct {
	Device
}

func newMobile() IDevice {
	return &Mobile{
		Device: Device{
			architecture: "armv8",
			ram:          8,
		},
	}
}
func getGun(deviceType string) (IDevice, error) {
	if deviceType == "server" {
		return newServer(), nil
	}
	if deviceType == "mobile" {
		return newMobile(), nil
	}
	return nil, fmt.Errorf("Wrong device type passed")
}
