package api_test

import (
	"testing"
	"myproject/files"
	"myproject/api"
	"github.com/joho/godotenv"
)

func TestCreateBin(t *testing.T) {
	err := godotenv.Load()
			if err != nil {
				t.Error("Не удалось найти ENV файл")
			}
			
			var filesmanager files.FilesManager = &files.Files{}
			var apimanager api.ApiManager = &api.ApiStruct{}
			
	key := apimanager.GetKey()
	data, err := filesmanager.ReadAnyFile("../file.json")
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
	apimanager.Delete(id.Metadata.ID, key)
}