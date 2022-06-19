package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jesse0michael/go-request"
)

type HandlerFunc[I any, O any] func(context.Context, I) (O, error)

func IO[I any, O any](h HandlerFunc[I, O]) http.Handler {
	return io[I, O]{fun: h}
}

type io[I any, O any] struct {
	fun HandlerFunc[I, O]
}

func (h io[I, O]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var in I

	if err := request.Decode(r, &in); err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := h.fun(r.Context(), in)

	if httpErr, ok := err.(interface {
		Error() string
		HTTPStatusCode() int
	}); ok {
		out := struct {
			Error string `json:"error"`
		}{Error: httpErr.Error()}

		b, err := json.Marshal(out)
		if err != nil {
			// TODO: metric = failed to parse error
			println(err.Error())
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(httpErr.HTTPStatusCode())
		w.Write(b)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
