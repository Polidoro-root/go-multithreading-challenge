package main

import (
	"io"
	"net/http"
	"time"
)

func GetAddressByCEPFromURL(ch chan<- string, url string) {
	req, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	ch <- string(res)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go GetAddressByCEPFromURL(ch1, "https://cdn.apicep.com/file/apicep/18400-180.json")
	go GetAddressByCEPFromURL(ch2, "http://viacep.com.br/ws/18400180/json/")

	select {
	case apicepRes := <-ch1:
		println("apicep.com resposta: \n", apicepRes)
	case viacepRes := <-ch2:
		println("viacep.com.br resposta: \n", viacepRes)
	case <-time.After(time.Second):
		panic("Não foi possível realizar a consulta a tempo, tente novamente")
	}
}
