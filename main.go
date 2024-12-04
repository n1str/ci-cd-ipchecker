package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		log.Fatalf("Ошибка при получении IP-адреса: %v\n", err)
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v\n", err)
	}

	fmt.Printf("Ваш IP-адрес: %s\n", ip)
}
