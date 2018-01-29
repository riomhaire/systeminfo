package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/riomhaire/systeminfo/usecases"
)

func Metrics(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("Request type %s\n", req.Header.Get("Accept"))

	stats := usecases.Metrics()

	b, _ := json.Marshal(stats)

	w.Write(b)

}
