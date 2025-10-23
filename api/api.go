package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"


	// "errors"
	// "fmt"
	"io"
	"myproject/bins"
	"myproject/config"
	"net/http"
)

type ApiManager interface {
	GetKey() string
	Post([]byte) (*getID, error)
	Get(string, string) (*BinResponse, error)
	Put(string, []byte) error
}

type ApiStruct struct {}

type getID struct {
	Metadata struct {
		ID string `json:"id"` // тут была ошибка что я не полностью указала структуру json ответа которого я ждал я туда не добавил Metadata поэтому выскакивала пустая строка
	} `json:"metadata"` // проверка тела ответа  fmt.Println("Ответ сервера:", string(body))
}

type BinResponse struct {
	Record   bins.BinList `json:"record"`
	Metadata struct {
		ID        string `json:"id"`
		Private   bool   `json:"private"`
		CreatedAt string `json:"createdAt"`
		Name      string `json:"name"`
	} `json:"metadata"`
}

func (keys *ApiStruct) GetKey() string {
	key := config.NewConfig()
			if len(key.Key) == 0 {
				panic("Ключ не создан")
			}
	return key.Key
}

func (post *ApiStruct) Post(binss []byte) (*getID, error) {
	key := post.GetKey()
	request, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(binss))
			if err != nil {
				panic("Ошибка запроса!")
			}
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
	var bin getID      // вытаскиваем то json тегу id в структуру
	err = json.Unmarshal(body, &bin)
			if err != nil {
				return nil, err
			}

	return &bin, err
}

func (get *ApiStruct) Get(id string, key string) (*BinResponse, error) {
			if id == "" {
				return nil, fmt.Errorf("ID не передан") //универсальный способ создавать ошибки с нужными тебе сообщениями и параметрами.
																								//Дает подробную ошибку для разработки, дебага и логирования. errors.New("some error") — только статичная строка без параметров.
			}
			if key == "" {
				return nil, fmt.Errorf("ключ не передан")
			}
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%v", id) // формирую правильный url добавляя свой ID
	request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return nil, err
			}
	request.Header.Set("X-Master-Key", key)
	client := &http.Client{}
	response, err := client.Do(request)
			if err != nil {
				return nil, err
			}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}
			if response.StatusCode != 200 {
				bodyText := string(body)
				return nil, fmt.Errorf("HTTP error: %d, ответ: %s", response.StatusCode, bodyText)
			}
	// fmt.Println(string(body))
	var resp BinResponse
	err = json.Unmarshal(body, &resp)
			if err != nil {
				return nil, err
			}
	return &resp, nil
}

func (put *ApiStruct) Put(id string, binss []byte) error {
	key := put.GetKey()
			if id == "" {
				return errors.New("ID не передан")
			}
			if key == "" {
				return errors.New("Ключ не передан")
			}
	url := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", id)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(binss))
			if err != nil {
				return err
			}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", key)

	client := http.Client{}
	resp, err := client.Do(req)
			if err != nil {
				return err
			}
	defer resp.Body.Close()
			if resp.StatusCode != 200 {
				// return err
				    b, _ := io.ReadAll(resp.Body)
    				return fmt.Errorf("ошибка PUT: код=%v, ответ=%s", resp.StatusCode, b)
			}

	return nil
}
