package main

import (
	"encoding/json"
    "fmt"
    "log"
    "os"
    "net/http"
)

type Hit struct {
    Counter  int
}

func receiver(w http.ResponseWriter, r *http.Request) {
	var h Hit

	err := json.NewDecoder(r.Body).Decode(&h)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)
    line := fmt.Sprintf("Success : POST http://localhost:2021/receiver {\"counter\": %d, \"X-RANDOM\":\"%s\"}", h.Counter, r.Header.Get("X-RANDOM"))

    log.Println(line)
    fmt.Println(line)

    fmt.Fprintf(w, "%+v", h)
}

func main() {
    http.HandleFunc("/receiver", receiver)
    log.Fatal(http.ListenAndServe(":2021", nil))
}