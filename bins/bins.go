package bins


import (
	"errors"
	"fmt"
	"time"
)

type BinsManager interface {
	CreateBin() Bin
	PrintBinList()
}

type BinList struct{
	Bins []Bin					`json:"bins"`
}

type Bin struct{
	Id				string		`json:"id"`
	Private		bool			`json:"private"`
	CreatedAt time.Time	`json:"createdAt"`
	Name			string		`json:"name"`
}

func (binList BinList)PrintBinList() {
	for _, value := range binList.Bins{
		fmt.Printf("ID: %s | Имя: %s | Приватный: %t | Время создания: %s\n", value.Id, value.Name, value.Private,  value.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func (b BinList) CreateBin() Bin {
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
		Id:					id,			
		Private:		private,
		CreatedAt: 	time.Now(),
		Name:				name,
	}
	return bin
}

func idExam(value string) error {
	if value == "" {
		return errors.New("INVALID ID/NAME") 
	}
	return nil
}