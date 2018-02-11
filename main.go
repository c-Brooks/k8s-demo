package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"

	"github.com/gogap/logrus"
)

func main() {

	logrus.Infof("serving on port 8080")

	http.HandleFunc("/", handleRoot)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)

	nodeName := os.Getenv("NODE_NAME")
	podName := os.Getenv("POD_NAME")

	resp["nodeName"] = nodeName
	resp["hostname"], _ = os.Hostname()
	resp["podName"] = podName

	// // this won't block, which lets us really pile on the requests
	// x := int64(0)
	// go cpuIntensive(&x)

	json, err := json.Marshal(resp)
	if err != nil {
		logrus.Errorf("could not retrieve hostname: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	logrus.Info(string(json))

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// pointers all escape to heap + many calls to scheduler
func cpuIntensive(p *int64) {
	for i := int64(1); i <= 10000000; i++ {
		*p = i

		// calling the scheduler is especially cpu intensive
		runtime.Gosched()
	}
	logrus.Info("done cpu task")
}
