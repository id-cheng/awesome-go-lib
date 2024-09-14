package main

import (
	"bytes"
	"encoding/json"
	"ergo.services/ergo/act"
	"ergo.services/ergo/gen"
	"net/http"
)

func factory_MyWebWebWorker() gen.ProcessBehavior {
	return &MyWebWebWorker{}
}

type MyWebWebWorker struct {
	act.WebWorker
}

// Init invoked on a start this process.
func (w *MyWebWebWorker) Init(args ...any) error {
	w.Log().Info("started web worker process with args %v", args)
	return nil
}

// Handle GET requests. For the other HTTP methods (POST, PATCH, etc)
// you need to add the accoring callback-method implementation. See act.WebWorkerBehavior.

func (w *MyWebWebWorker) HandleGet(from gen.PID, writer http.ResponseWriter, request *http.Request) error {
	var buf bytes.Buffer

	w.Log().Info("got HTTP request %q", request.URL.Path)
	writer.Header().Set("Content-Type", "application/json")
	// response JSON message with information about this process
	info, _ := w.Info()
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.Encode(info)
	writer.Write(buf.Bytes())
	return nil
}
