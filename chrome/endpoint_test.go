package chrome

import "testing"

func TestEndpoints(t *testing.T) {
	host := testingHost(t)
	list, err := Endpoints(host)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range list {
		if item.ID == "" {
			t.Error("endpoint missing ID field")
		}
	}
}
