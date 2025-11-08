package api_test

import (
	"encoding/json"
	"myproject/api"
	"myproject/bins"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

	var testData = bins.BinList{
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