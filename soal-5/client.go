package main

import (
	"net/http"
    "fmt"
    "time"
    "bytes"
    "encoding/json"
    "math/rand"
)

func main() {
    done := make(chan bool)
    go forever()
    <-done
}

func forever() {
	var i = 0
    for {
    	i = i + 1
        fmt.Printf("%v+\n", time.Now())

        sendReq(i)

        time.Sleep(time.Minute)
    }
}

func sendReq(i int) {
	client := &http.Client{}
	body, _ := json.Marshal(map[string]int{
		"counter": i,
	})
	reqBody := bytes.NewBuffer(body)
	req, _ := http.NewRequest("POST","http://localhost:2021/receiver", reqBody)

	xRandom := randomString(8)

	req.Header.Add("X-RANDOM", xRandom)
	resp, _ := client.Do(req)

	defer resp.Body.Close()
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

  	b := make([]byte, length)
  	for i := range b {
    	b[i] = charset[rand.Intn(len(charset))]
  	}
  	return string(b)
}