package render

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

const (
	CodeOk      = 0  // OK
	CodeErr     = -1 // general error
	CodeAuthErr = -2 // unauthenticated request
)

type (
	// result data
	DataModel map[string]interface{}

	// Result represents a common-used result struct.
	Result struct {
		Code   int         `json:"code"`             // return code
		Msg    string      `json:"msg"`              // message
		Data   interface{} `json:"data,omitempty"`   // data object
		Result interface{} `json:"result,omitempty"` // result data object
	}
)

var (
	// indent the json-encoded API responses
	indent bool
)

func init() {
	indent, _ = strconv.ParseBool(
		os.Getenv("HTTP_JSON_INDENT"),
	)
}

// NewResult creates a result with Code=0, Msg="", Data=nil.
func NewResult() *Result {
	return &Result{
		Code:   0,
		Msg:    "",
		Result: &DataModel{},
	}
}

func (r *Result) Error(err error) {
	r.Code = CodeErr
	r.Msg = err.Error()
}

func (r *Result) AuthError(err error) {
	r.Code = CodeAuthErr
	r.Msg = err.Error()
}

// ErrorCode writes the json-encoded error message to the response.
func ErrorJSON(w http.ResponseWriter, code int, err string, status int) {
	JSON(w, &Result{
		Code:   code,
		Msg:    err,
		Result: &DataModel{},
	}, status)
}

// JSON writes the json-encoded error message to the response
// with a 400 bad request status code.
func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	if indent {
		enc.SetIndent("", "  ")
	}
	_ = enc.Encode(v)
}

// GzJSON writes the json-encoded error message to the response
// with a 400 bad request status code.
func GzJSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	w.WriteHeader(status)

	gz := gzip.NewWriter(w)
	err := json.NewEncoder(gz).Encode(v)
	if nil != err {
		return
	}
	err = gz.Close()
	if nil != err {
		return
	}
}
