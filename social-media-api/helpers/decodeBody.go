package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MalformedRequest struct {
	status int
	msg    string
}

func (mr *MalformedRequest) Error() string {
	return mr.msg
}

func (mr *MalformedRequest) Status() int {
	return mr.status
}

func DecodeBodyToJson(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()

	err := decode.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			log.Println(msg)
			return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			log.Println(msg)
			return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			log.Println(msg)
			return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			log.Println(msg)
			return &MalformedRequest{status: http.StatusBadRequest, msg: msg}

		default:
			return err
		}
	}
	return nil
}

func UnmarshalJson(jsondata string, dst interface{}) {
	err := json.Unmarshal([]byte(jsondata), &dst)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func MarshalJson(i interface{}, w http.ResponseWriter) {
	jsonData, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error :", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
