package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/riomhaire/systeminfo/usecases"
)

func Metrics(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestType := req.Header.Get("Accept")

	requestType = strings.ToLower(requestType)
	if strings.Contains(requestType, "text/plain") {
		w.Header().Set("Content-Type", "text/plain")
		stats := usecases.PrometheusMetrics()
		w.Write([]byte(stats))
	} else {
		w.Header().Set("Content-Type", "application/json")
		stats := usecases.Metrics()
		b, _ := json.Marshal(stats)
		w.Write(b)
	}

}
