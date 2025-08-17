package bins


import (
	"errors"
	"fmt"
	"time"
)


type BinList struct{
	Bins []Bin
}

type Bin struct{
	id				string
	private		bool
	createdAt time.Time
	name			string
}

func (binList BinList)PrintBinList() {
	for _, value := range binList.Bins{
		fmt.Printf("ID: %s | Имя: %s | Приватный: %t | Время создания: %s\n", value.id, value.name, value.private,  value.createdAt.Format("2006-01-02 15:04:05"))
	}
}

func CreatBin() Bin {
	var id, name string
	var private bool
	for {
	fmt.Print("Введите свой ID: ")
	fmt.Scanln(&id)
	err := idExam(id)
	if err == nil {
		break
	}
	fmt.Println(err)
	}
	fmt.Print("Укажите, будет ли файл приватным или нет: ")
	fmt.Scanln(&private)
	for {
	fmt.Print("Укажите название файла: ")
	fmt.Scanln(&name)
  err := idExam(name)
	if err == nil {
		break
	}
	fmt.Println(err)
	}

	bin := Bin{
		id:					id,			
		private:		private,
		createdAt: 	time.Now(),
		name:				name,
	}
	return bin
}

func idExam(value string) error {
	if value == "" {
		return errors.New("INVALID ID/NAME") 
	}
	return nil
}