package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JessicaEspejo10/go-fundamentals-web-users/internal/user"
	"github.com/JessicaEspejo10/go-fundamentals-web-users/pkg/transport"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, endpoints user.Endpoints) {
	router.HandleFunc("/users", UserServer(ctx, endpoints))

}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(w http.ResponseWriter, e *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		trans := transport.New(w, r, ctx)

		switch r.Method {
		case http.MethodGet:
			trans.Server(
				transport.Endpoint(endpoints.GetAll),
				decodeGetAllUser,
				encodeResponse,
				encodeError)
			return

		case http.MethodPost:
			trans.Server(
				transport.Endpoint(endpoints.Create),
				decodeCreateUser,
				encodeResponse,
				encodeError)
			return

		}
		InvalidMethod(w)
	}
}

func decodeGetAllUser(ctx context.Context, r *http.Request) (interface{}, error) {

	return nil, nil
}

func decodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request format: '%v'", err.Error())
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	status := http.StatusOK
	w.WriteHeader(status)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	fmt.Fprintf(w, `{"status": %d, "data": %s}`, status, data)
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	status := http.StatusInternalServerError
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	fmt.Fprintf(w, `{"status":%d," message":%s}`, status, err.Error())
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, "not found")

}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, status, message)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {

}
