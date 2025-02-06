package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func WriteJSONResponse(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	logger.Info("RESPONSE", "code:", status)
	return json.NewEncoder(w).Encode(v)
}
