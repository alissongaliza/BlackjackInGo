package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateUser(user PostUser) {

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(user)
	req, err := http.NewRequest("POST", "http://localhost:8080/user", buf)
	if err != nil {
		log.Print(err)
	}

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Print(e)
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
}
