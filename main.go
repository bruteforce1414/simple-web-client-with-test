package main

import (
	"fmt"
	"github.com/bruteforce1414/simple-web-client-with-test/client"
)



func main() {

	client:=client.NewHttpClient("http://kstatida.ru")
	body, _:=(*client).Get("/a/11526/vyezzhaya-vecherom-na-velosipede-na-dorogu-imejte-s-soboj-komplekt-ispravno-rabotayuschego-sveta-na-vsyu-poezdku")
	fmt.Println("Количество байт на странице:", len(body))

}

