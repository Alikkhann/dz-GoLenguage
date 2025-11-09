package api_test

import (
	"fmt"
	"encoding/json"
	"myproject/api"
	"myproject/bins"
	"testing"
	"time"
	"github.com/joho/godotenv"
)

func generateUniqueId() string {
	return fmt.Sprintf("test-%d", time.Now().UnixNano())
}

func generateUniqueName() string {
	return fmt.Sprintf("updated-test-%d", time.Now().UnixNano())
}

	var testData = bins.BinList{
		Bins: []bins.Bin{
			{
				Id: generateUniqueId(),
      	Private: true,
     	 	CreatedAt: time.Now(),
     	  Name: generateUniqueId(),
			},
		},
	}



func TestCreateBin(t *testing.T) {
	err := godotenv.Load()
			if err != nil {
				t.Error("Не удалось найти ENV файл")
			}
	var apimanager api.ApiManager = &api.ApiStruct{}
	key := apimanager.GetKey()
	data, err := json.Marshal(testData)
			if err != nil {
				t.Error(err)
			}
			if err != nil {
				t.Error(err)
			}
	id, err := apimanager.Post(data)
			if err != nil {
			t.Error(err)
			}
			if id == nil {
				t.Errorf("ожидалась длина ответа больше 0, получили %s", id)
			}
	t.Cleanup(func()  {
		apimanager.Delete(id.Metadata.ID, key)
	})
}


func TestGetBin(t *testing.T) {
		err := godotenv.Load()
			if err != nil {
				t.Error("Не удалось найти ENV файл")
			}
	var apimanager api.ApiManager = &api.ApiStruct{}
	key := apimanager.GetKey()
	data, err := json.Marshal(testData)
			if err != nil {
				t.Error(err)
			}
			if err != nil {
				t.Error(err)
			}
	id, err := apimanager.Post(data)
			if err != nil {
			t.Error(err)
			}
			if id == nil {
				t.Errorf("ожидалась длина ответа больше 0, получили %s", id)
			}	
	dataGet, err := apimanager.Get(id.Metadata.ID, key)
			if err != nil {
			t.Error(err)
			}
			if dataGet.Record.Bins[0].Name != testData.Bins[0].Name {
				t.Errorf("Ожидалось %s, получили %s", testData.Bins[0].Name, dataGet.Record.Bins[0].Name)
			}
	t.Cleanup(func() {
		apimanager.Delete(id.Metadata.ID, key)
	})
}



func TestPutBin(t *testing.T) {
		err := godotenv.Load()
			if err != nil {
				t.Error("Не удалось найти ENV файл")
			}
	var apimanager api.ApiManager = &api.ApiStruct{}
	key := apimanager.GetKey()
	data, err := json.Marshal(testData)
			if err != nil {
				t.Error(err)
			}
			if err != nil {
				t.Error(err)
			}
	id, err := apimanager.Post(data)
			if err != nil {
			t.Error(err)
			}
			if id == nil {
				t.Errorf("ожидалась длина ответа больше 0, получили %s", id)
			}

	testDataPut := bins.BinList{
		Bins: []bins.Bin{
			{
				Id: generateUniqueId(),
      	Private: true,
     	 	CreatedAt: time.Now(),
     	  Name: generateUniqueName(),
			},
		},
	}
	
	dataPut, err := json.Marshal(testDataPut)
			if err != nil {
				t.Error(err)
			}

	err = apimanager.Put(id.Metadata.ID, dataPut)
			if err != nil {
				t.Error(err)
			}

	dataGet, err := apimanager.Get(id.Metadata.ID, key)
			if err != nil {
			t.Error(err)
			}

	if testDataPut.Bins[0].Id != dataGet.Record.Bins[0].Id {
		t.Errorf("Ожидалось %v, получили %s", testDataPut.Bins[0].Name, dataGet.Record.Bins[0].Name)
	}

	t.Cleanup(func() {
		apimanager.Delete(id.Metadata.ID, key)
	})
}


func TestDeleteBin(t *testing.T) {
		err := godotenv.Load()
			if err != nil {
				t.Error("Не удалось найти ENV файл")
			}
	var apimanager api.ApiManager = &api.ApiStruct{}
	key := apimanager.GetKey()
	data, err := json.Marshal(testData)
			if err != nil {
				t.Error(err)
			}
			if err != nil {
				t.Error(err)
			}
	id, err := apimanager.Post(data)
			if err != nil {
			t.Error(err)
			}
			if id == nil {
				t.Errorf("ожидалась длина ответа больше 0, получили %s", id)
			}
	
	err = apimanager.Delete(id.Metadata.ID, key)
	if err != nil {
		t.Error(err)
	}

	_ , err = apimanager.Get(id.Metadata.ID, key)

		if err == nil {
			t.Errorf("Ожидалась ошибка, а получили данные")
		}
}