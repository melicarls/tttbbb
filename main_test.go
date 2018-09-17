package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHandler(t *testing.T) {
  r := newRouter()

  mockServer := httptest.NewServer(r)

  resp, err := http.Get(mockServer.URL + "/")

  if err != nil {
    t.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
    t.Errorf("Status should ")
  }
}
