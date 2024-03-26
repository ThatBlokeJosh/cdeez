package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Apps() string {
	req, err := http.NewRequest("GET", "http://127.0.0.1:3000/apps", nil)
	Check(err)
	resp, err := http.DefaultClient.Do(req)
	Check(err)	
	body, err := io.ReadAll(resp.Body)
	Check(err)
	return string(body)
}

func Deploy(name, repo string) string {
	reqBody := []byte(fmt.Sprintf(`{"name": "%s", "repo": "%s"}`, name, repo))
	buf := bytes.NewReader(reqBody)	
	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/apps", buf)
	Check(err)
	resp, err := http.DefaultClient.Do(req)
	Check(err)
	body, err := io.ReadAll(resp.Body)
	Check(err)
	return string(body)
}


func Delete(name string) {
	reqBody := []byte(fmt.Sprintf(`{"name": "%s"}`, name))
	buf := bytes.NewReader(reqBody)	
	req, err := http.NewRequest("DELETE", "http://127.0.0.1:3000/apps", buf)
	Check(err)
	_, err = http.DefaultClient.Do(req)
	Check(err)
}


func Restart(name string) {
	reqBody := []byte(fmt.Sprintf(`{"name": "%s"}`, name))
	buf := bytes.NewReader(reqBody)	
	req, err := http.NewRequest("PUT", "http://127.0.0.1:3000/apps", buf)
	Check(err)
	_, err = http.DefaultClient.Do(req)
	Check(err)
}


func Stats(name string) string {
	reqBody := []byte(fmt.Sprintf(`{"name": "%s"}`, name))
	buf := bytes.NewReader(reqBody)	
	req, err := http.NewRequest("GET", "http://127.0.0.1:3000/apps/stats", buf)
	Check(err)
	resp, err := http.DefaultClient.Do(req)
	Check(err)
	body, err := io.ReadAll(resp.Body)
	Check(err)
	return string(body)
}


