package render

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/20326/vega/pkg/errors"
)

type (
	// result data
	DataModel map[string]interface{}

	// Result represents a common-used result struct.
	Result struct {
		Code int    `json:"code"` // return code
		Msg  string `json:"msg"`  // message
		// ErrorMsg string      `json:"errorMsg,omitempty"` // show message to ui
		Result interface{} `json:"result"` // data object
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
	r.Code = errors.CodeErr
	r.Msg = err.Error()
}

// ErrorCode writes the json-encoded error message to the response.
func ErrorCode(w http.ResponseWriter, err error, status int) {
	JSON(w, &errors.Error{Message: err.Error()}, status)
}

// InternalError writes the json-encoded error message to the response
// with a 500 internal server error.
func InternalError(w http.ResponseWriter) {
	JSON(w, &errors.ErrInternalError, 500)
}

// NotImplemented writes the json-encoded error message to the
// response with a 501 not found status code.
func NotImplemented(w http.ResponseWriter) {
	JSON(w, &errors.ErrNotImplemented, 501)
}

// NotFound writes the json-encoded error message to the response
// with a 404 not found status code.
func NotFound(w http.ResponseWriter) {
	JSON(w, &errors.ErrNotFound, 404)
}

// Unauthorized writes the json-encoded error message to the response
// with a 401 unauthorized status code.
func Unauthorized(w http.ResponseWriter) {
	JSON(w, &errors.ErrUnauthorized, 401)
}

// Forbidden writes the json-encoded error message to the response
// with a 403 forbidden status code.
func Forbidden(w http.ResponseWriter) {
	JSON(w, &errors.ErrForbidden, 403)
}

// BadRequest writes the json-encoded error message to the response
// with a 400 bad request status code.
func BadRequest(w http.ResponseWriter) {
	JSON(w, &errors.ErrBadRequest, 400)
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
