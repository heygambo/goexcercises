package main

import (
	"fmt"
)

type Language struct {
	Name string
	Speed int
}

type Coder struct {
	Name string
	Language Language
}

func (c Coder) Code (n int) int {
  return n * c.Language.Speed
}

func (c Coder) GetName () string {
	return c.Name
}

func (c Coder) String() string {
	return fmt.Sprintf("<%s lang:%s speed:%d>", c.Name, c.Language.Name, c.Language.Speed)
}

func main() {
	javascript := Language{
		Name: "JavaScript",
		Speed: 10,
	}
	golang := Language{
		Name: "Go",
		Speed: 1,
	}
	ruby := Language{
		Name: "Ruby",
		Speed: 5,
	}
	swift := Language{
		Name: "Swift",
		Speed: 15,
	}

	artur := Coder{
		Name: "Artur",
		Language: javascript,
	}
	nick := Coder{
		Name: "Nick",
		Language: javascript,
	}
	marcin := Coder{
		Name: "Marcin",
		Language: swift,
	}
	gambo := Coder{
		Name: "Gambo",
		Language: golang,
	}
	evgenii := Coder{
		Name: "Evgenii",
		Language: ruby,
	}

	coders := make([]Coder, 5)

	_ = append(coders, artur)
	_ = append(coders, nick)
	_ = append(coders, marcin)
	_ = append(coders, gambo)
	_ = append(coders, evgenii)
	
	// fmt.Println(coders)
	for _, c := range coders {
		fmt.Println("Coder:", c.Code(5))
	}
}
