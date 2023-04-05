package pattern

/*
	Реализовать паттерн «строитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Builder_pattern
*/
/*
	Используется для создания структур с одинаковым набором полей.
	К плюсам паттерна относится возможность контроля создания структур.
	К минусам относится необходимость создания новых структур.
*/
type Device struct {
	architecture string
	drives       int
	ram          int
}

type ServerBuilder struct {
	architecture string
	drives       int
	ram          int
}

type MobileBuilder struct {
	architecture string
	drives       int
	ram          int
}

func newServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (b *ServerBuilder) setArchitecture() {
	b.architecture = "x86_64"
}

func (b *ServerBuilder) setDrives() {
	b.drives = 8
}

func (b *ServerBuilder) setRam() {
	b.ram = 128
}

func (b *ServerBuilder) getDevice() Device {
	return Device{
		architecture: b.architecture,
		drives:       b.drives,
		ram:          b.ram,
	}
}

func newMobileBuilder() *MobileBuilder {
	return &MobileBuilder{}
}

func (b *MobileBuilder) setArchitecture() {
	b.architecture = "armv8"
}

func (b *MobileBuilder) setDrives() {
	b.drives = 1
}

func (b *MobileBuilder) setRam() {
	b.ram = 8
}

func (b *MobileBuilder) getDevice() Device {
	return Device{
		architecture: b.architecture,
		drives:       b.drives,
		ram:          b.ram,
	}
}

type IBuilder interface {
	setArchitecture()
	setDrives()
	setRam()
	getDevice() Device
}

func getBuilder(builderType string) IBuilder {
	if builderType == "server" {
		return newServerBuilder()
	}

	if builderType == "mobile" {
		return newMobileBuilder()
	}
	return nil
}
