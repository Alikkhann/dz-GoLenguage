package api

import (
	"bytes"
	"encoding/json"
	// "errors"
	// "fmt"
	"io"
	"myproject/config"
	"net/http"
)

	type getID struct {
		Metadata struct {
			ID string `json:"id"`  // тут была ошибка что я не полностью указала структуру json ответа которого я ждал я туда не добавил Metadata поэтому выскакивала пустая строка
		}`json:"metadata"`   // проверка тела ответа  fmt.Println("Ответ сервера:", string(body))
	}

func GetKey() string{
	key := config.NewConfig()
	if len(key.Key) == 0 {
		panic("Ключ не создан")
	}
	return key.Key
}

func Post(binss []byte) (*getID, error) {
	key := GetKey()

	request, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(binss))
	if err != nil { panic("Ошибка запроса!") }
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Master-Key", key)

	httpClient := &http.Client{}
	respClient, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	// if respClient.StatusCode != 201 {
	// 	fmt.Printf("Status Code: %d\n", respClient.StatusCode)
	// 	return nil, errors.New("ошибка ответа клиента")
	// }
	

	defer respClient.Body.Close()

	body, err := io.ReadAll(respClient.Body)
	if err != nil {
		return nil, err
	}
	var bin getID
	err = json.Unmarshal(body, &bin)
	if err != nil {
		return nil, err
	}
	return &bin, err
}

