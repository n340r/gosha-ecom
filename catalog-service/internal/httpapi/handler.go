package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
  writeJSON(w, http.StatusOK, map[string]string{"ok": "true"})
}

func Products(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
  }

  resp := MockGetProductsResponse()

  if idStr := r.URL.Query().Get("id"); idStr != "" {
    id, err := strconv.Atoi(idStr)
    if err != nil {
      http.Error(w, "invalid id", http.StatusBadRequest)
      return
    }

    filtered := make([]Product, 0, len(resp.Products))
    for _, p := range resp.Products {
      if p.ID == id {
        filtered = append(filtered, p)
      }
    }

    resp.Products = filtered
    resp.Pagination.TotalCount = len(filtered)
    resp.Pagination.TotalPageCount = 1
  }

  writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
  w.Header().Set("Content-Type", "application-json")
  w.WriteHeader(status)

  enc := json.NewEncoder(w)
  enc.SetEscapeHTML(false)
  _ = enc.Encode(v)
}
