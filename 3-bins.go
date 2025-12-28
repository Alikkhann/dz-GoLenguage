 package main

// import (
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"myproject/api"
// 	"myproject/bins"
// 	"myproject/files"
// 	"myproject/storage"
// 	"strings"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	action := flag.String("action", "", "create(POST)/get/updata(PUT)/delete")
// 	flag.Parse()
// 	err := godotenv.Load()
// 			if err != nil {
// 				fmt.Println("Не удалось найти ENV файл")
// 			}
// 			var filesmanager files.FilesManager = &files.Files{}
// 			binList := bins.BinList{}
// 			var binsmanager bins.BinsManager = binList
// 			var apimanager api.ApiManager = &api.ApiStruct{}
			
// 		if *action == "create" || *action == "update" {
// 		data, err := filesmanager.ReadAnyFile("file.json")
// 				if err == nil {
// 					err = json.Unmarshal(data, &binList)
// 				}
// 				if err != nil {
// 					fmt.Print("Ошибка разбора JSON")
// 				}
// 		var n int
// 		fmt.Println("Функция создания файлов типа `ключ` - `значение`.")
// 		fmt.Println("Сколько списков вы хотите создать?")
// 		fmt.Scan(&n)
// 		for i := 0; i < n; i++ {
// 			creatBin := binsmanager.CreateBin()
// 			binList.Bins = append(binList.Bins, creatBin)
// 		}
// 		data, err = storage.ToBytes(binList)
// 				if err != nil {
// 					fmt.Println("Ошибка!")
// 					return
// 				}
// 		filesmanager.WriteFile(data, "file.json")
	

// 	if *action == "create" {
// 		id, err := apimanager.Post(data)
// 				if err != nil {
// 					fmt.Println("Не удалось вернуть ID")
// 					fmt.Println(err)
// 					return
// 				}
// 		idToBytes, err := storage.ToBytes(id.Metadata.ID)
// 				if err != nil {
// 					fmt.Println("Не удалось преобразовать в байты ID")
// 					return
// 				}
// 		idToTrim := strings.Trim(string(idToBytes), "\"")
// 		filesmanager.WriteFile([]byte(idToTrim), "file.id")
// 	}
// 	if *action == "update" {
// 		readId, err := filesmanager.ReadAnyFile("file.id")
// 				if err != nil {
// 					fmt.Println("Не удалось прочитать ID")
// 					fmt.Println("Сначала создайте(отправьте) create запрос, а потом читайте ID")
// 					return
// 				}
// 		err = apimanager.Put(string(readId), data)
// 				if err != nil {
// 					fmt.Println("Не удалось обновить данные при PUT запросе")
// 					fmt.Println(err)
// 					return
// 				}
// 		}
// 	}else {
// 	if *action == "get" {
// 	  key := apimanager.GetKey()
// 		readId, err := filesmanager.ReadAnyFile("file.id")
// 				if err != nil {
// 					fmt.Println("Не удалось прочитать ID")
// 					fmt.Println("Сначала создайте(отправьте) create запрос, а потом читайте ID")
// 					return
// 				}
// 		dataGet, err := apimanager.Get(string(readId), key)
// 				if err != nil {
// 					fmt.Println("Не удалось выполнить GET запрос")
// 					return
// 				}
// 		dataToBytes, err := storage.ToBytes(dataGet)
// 				if err != nil {
// 					fmt.Println("Не удалось преобразовать в байты данные GET зароса")
// 					return
// 				}
// 		fmt.Println(string(dataToBytes))
// 			}
// 		}
// 	if *action == "delete" {
// 		key := apimanager.GetKey()

// 	  readId, err := filesmanager.ReadAnyFile("file.id")
// 				if err != nil {
// 					fmt.Println("Не удалось прочитать ID")
// 					fmt.Println("Сначала создайте(отправьте) create запрос, а потом читайте ID")
// 					return
// 				}

// 		err = apimanager.Delete(string(readId), key)
// 				if err != nil {
// 					fmt.Println("Не удалось удалить Bin")
// 					fmt.Println(err)
// 					return
// 				}
			
// 		}
// 		if *action == "" {
// 			fmt.Println("Выберите действие (go run 3-bins --action='create/get/update')")
// 		}

	




//---

//filesmanager.WriteFile(data, "file.json")
// id, err := api.Post(data)
// if err != nil {
// 	fmt.Println(err)
// }

// key := api.GetKey()
// myId := id.Metadata.ID

// dataGet, err := api.Get(myId, key)
// if err != nil {
// 	fmt.Println(err)
// }

// butesData, err := storage.ToBytes(dataGet)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(butesData))
	
	// filesmanager.WriteFile(butesData, "dataGet")
	// binsmanager.PrintBinList()
// }
