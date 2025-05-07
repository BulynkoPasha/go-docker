package main

import (
  "fmt"
  "log"
  "net/http"
  "sync"
  "time"
)

var (
  mu       sync.Mutex
  visits   int
)

func handler(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  visits++
  mu.Unlock()

  clientIP := r.RemoteAddr
  currentTime := time.Now().Format("02-Jan-2006 15:04:05")

  fmt.Fprintf(w, "Привет, мир!\n")
  fmt.Fprintf(w, "Текущее время: %s\n", currentTime)
  fmt.Fprintf(w, "Ваш IP: %s\n", clientIP)
  fmt.Fprintf(w, "Количество посещений: %d\n", visits)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintln(w, "Сервис работает корректно.")
}

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/status", statusHandler)
  fmt.Println("апуск demo-приложения на порту 8888...")
  log.Fatal(http.ListenAndServe(":8888", nil))
}
