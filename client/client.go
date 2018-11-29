package client

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Client interface {
	// URL у нас путь на сервере, например /bruteforce1414/study-go, без самого домена
	Get(url string) ([]byte, error)
}

type httpClient struct {
	// это домен сервера, например http://github.com, к нему добавляя url из метода Get что б получить полный адрес страницы
	baseUrl string
	//
	client *http.Client
}

func NewHttpClient(base string) *httpClient {

	return &httpClient {
		baseUrl: base,
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}



}

func (c *httpClient) Get(url string) ([]byte, error) {
	resp, err := c.client.Get(c.baseUrl + url) // тут пакет path должен быть =)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // обязательно читаем всё
	// handling error and doing stuff with body that needs to be unit tested
	return body, err
}
