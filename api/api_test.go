package api_test

import (
	"encoding/json"
	"myproject/api"
	"myproject/bins"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

	var testDataPost = bins.BinList{
		Bins: []bins.Bin{
			{
				Id: "test-create-unique",
      	Private: true,
     	 	CreatedAt: time.Now(),
     	  Name: "test-create-bin",
			},
		},
	}

		var testDataPut = bins.BinList{
		Bins: []bins.Bin{
			{
				Id: "test-create-unique",
      	Private: true,
     	 	CreatedAt: time.Now(),
     	  Name: "test-create-bin",
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
	data, err := json.Marshal(testDataPost)
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
	data, err := json.Marshal(testDataPost)
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
			if dataGet.Record.Bins[0].Name != testDataPost.Bins[0].Name {
				t.Errorf("Ожидалось %s, получили %s", testDataPost.Bins[0].Name, dataGet.Record.Bins[0].Name)
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
	data, err := json.Marshal(testDataPost)
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
	
	dataPut, err := json.Marshal(testDataPost)
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
			
	if testDataPut.Bins[0].Name != dataGet.Record.Bins[0].Name {
		t.Errorf("Ожидалось %v, получили %s", testDataPut.Bins[0].Name, dataGet.Record.Bins[0].Name)
	}

	t.Cleanup(func() {
		apimanager.Delete(id.Metadata.ID, key)
	})
}