package chrome

import "testing"

func TestEvalPromise(t *testing.T) {
	conn := testingConn(t, false)
	defer conn.Close()

	var res string
	err := conn.EvalPromise("(function() {return Promise.resolve('hi')})();", &res)
	if err != nil {
		t.Error(err)
	} else if res != "hi" {
		t.Errorf("bad result: %v", res)
	}
}
