package handler

import (
	"bytes"
	"fmt"
	"net/http"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		return fmt.Errorf("could not read request body: %s", err)
	}
	fmt.Println(buf.String())
	return nil
}
