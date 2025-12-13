package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"catalog-service/internal/httpapi"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/healthz", httpapi.Healthz)
  mux.HandleFunc("/products", httpapi.Products)

  addr := envOr("ADDR", ":8080")

  server := &http.Server{
    Addr: addr,
    Handler: withCORS(mux),
    ReadHeaderTimeout: 5 * time.Second,
  }

  log.Printf("📒 [catalog-service] listening on %s\n", addr)

  if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    log.Fatal(err)
  }
}

func envOr(key, dev string) string {
  if v := os.Getenv(key); v != "" {
    return v
  }

  return dev
}

func withCORS(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusNoContent)
      return
    }

    next.ServeHTTP(w, r)
  })
}

