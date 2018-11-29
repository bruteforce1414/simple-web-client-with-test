package client_test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/bruteforce1414/simple-web-client-with-test/client"
)

func TestHttpClient_Get(t *testing.T) {
	a:=assert.New(t)

	mux:= http.NewServeMux()
	server:= httptest.NewServer(mux)
	defer server.Close()

	mux.HandleFunc("/a/11526/vyezzhaya-vecherom-na-velosipede-na-dorogu-imejte-s-soboj-komplekt-ispravno-rabotayuschego-sveta-na-vsyu-poezdku", func(w http.ResponseWriter, r *http.Request) {
		// а это функция-handler как в сервере и ты уже знаешь как тут писать ответ клиенту
	})

	// нам нужно передать в функцию создания клиента URL сервера, к которому нужно делать запросы
	clientTest := client.NewHttpClient("http://kstatida.ru") // незабудь добавить импортированный пакет перед функция создания клиента

	repos, err := clientTest.Get("/a/11526/vyezzhaya-vecherom-na-velosipede-na-dorogu-imejte-s-soboj-komplekt-ispravno-rabotayuschego-sveta-na-vsyu-poezdku")
	if err!= nil {
		t.Fatal(err)
	}
	a.Equal(len(repos),43565)


}