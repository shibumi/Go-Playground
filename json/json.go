package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r, err := http.Get("https://security.archlinux.org/all.json")
	if err != nil {
		log.Fatal("Ops..we shouldn't be here")
	}
	defer r.Body.Close()
	if r.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		var m []map[string]interface{}
		err3 := json.Unmarshal(bodyBytes, &m)
		if err3 != nil {
			log.Fatal("Ops..")
		} else {
            log.Println(m[0])
		}
	}
}
