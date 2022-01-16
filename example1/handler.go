package example1

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

/*
	Source from ["Top level logging · Developer 2.0"](https://developer20.com/top-level-logging/)
*/

type Handler struct {
	logger log.Logger
	srv    Service
}

func (h *Handler) operation(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Errorf("cannot read the body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := http.Request{}
	if err = json.Unmarshal(body, &req); err != nil {
		h.logger.Errorf("cannot read the body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.srv.Operation(r.Context(), req.Param1, req.Param2)
	if err != nil {
		if errors.Is(err, ErrEmptyResult) {
			// this shouldn't happen but when it does, we're ignoring such cases
			s.logger.Infof("the result is zero")
			return
		}

		h.logger.Errorf("cannot execute the operation: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return the success response
}
