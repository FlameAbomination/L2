package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/
/*
	Паттерн предназначен для сокрытия сложной реализации и создания более простого интерфейса для пользователя.
	Для примера используется набор функций для абстрактного компилятора.
	Для пользователя важен только полученный код без доступа к внутренним функциям.
	К минусам можно отнести то, что фасад мжет включать в себя слишком много функций.
*/
type Program struct {
}

type MemoryAllocator interface {
	AllocateMemory(trade Program) error
}

type ProgramValidator interface {
	Validate(trade Program) error
}

type ProgramPreprocessor interface {
	Preprocess(trade Program) error
}

type ProgramTokenizer interface {
	Tokenizer(trade Program) error
}

type CompilerFacade struct {
	MemoryAllocator
	ProgramValidator
	ProgramPreprocessor
	ProgramTokenizer
}

func (compiler *CompilerFacade) ProcessCode(program Program) error {
	err := compiler.Validate(program)
	if err != nil {
		return err
	}
	err = compiler.AllocateMemory(program)
	if err != nil {
		return err
	}
	err = compiler.Preprocess(program)
	if err != nil {
		return err
	}
	err = compiler.Tokenizer(program)
	if err != nil {
		return err
	}
	return nil
}
