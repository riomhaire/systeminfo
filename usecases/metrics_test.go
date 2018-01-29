package usecases

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMetricsGathering(t *testing.T) {
	data := Metrics()
	if data.Host.Uptime == 0 { // Should def get some hostinfo
		t.Fatal("Expected at least some host data")
	}
	if data.Memory.Total <= 0 { // Should def get some memory info
		t.Fatal("Expected at least some memory")
	}

	b, _ := json.Marshal(data)
	j := string(b)
	t.Logf("%v\n", j)
	fmt.Printf("%v\n", j)
}
