package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lindolfomoizinho/api-go/api_err"
)

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := GetUserRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		api_err.NotFound(w, r)
		return
	}

	h.store.RLock()

	u, ok := h.store.M[matches[1]]
	h.store.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	jsonBytes, err := json.Marshal(u)

	if err != nil {
		api_err.InternalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
