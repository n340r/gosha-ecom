package httpapi

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
  writeJSON(w, http.StatusOK, map[string]string{"ok": "true"})
}

func Products(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    return
  }

  resp := MockGetProductsResponse()

  ids := make([]int, 0, 4)

  if idsStr := r.URL.Query().Get("ids"); idsStr != "" {
    for _, part := range strings.Split(idsStr, ",") {
      part = strings.TrimSpace(part)
      if part == "" {
        continue
      }

      id, err := strconv.Atoi(part)
      if err != nil {
        http.Error(w, "invalid ids", http.StatusBadRequest)
        return
      }

      if !slices.Contains(ids, id) {
        ids = append(ids, id)
      }
    }
  }

  if idStr := r.URL.Query().Get("id"); idStr != "" {
    id, err := strconv.Atoi(strings.TrimSpace(idStr))
    if err != nil {
      http.Error(w, "invalid id", http.StatusBadRequest)
      return
    }

    if !slices.Contains(ids, id) {
      ids = append(ids, id)
    }
  }

  if len(ids) > 0 {
    filtered := make([]Product, 0, len(resp.Products))
    for _, p := range resp.Products {
      if slices.Contains(ids, p.ID) {
        filtered = append(filtered, p)
      }
    }

    resp.Products = filtered
    resp.Pagination.TotalCount = len(filtered)
    resp.Pagination.CurrentPage = 1// mock
    resp.Pagination.TotalPageCount = 1// mock
  }

  writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)

  enc := json.NewEncoder(w)
  enc.SetEscapeHTML(false)
  _ = enc.Encode(v)
}
