package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestRouter(t *testing.T) {
  r := newRouter()

  mockServer := httptest.NewServer(r)

  resp, err := http.Get(mockServer.URL + "/")

  if err != nil {
    t.Fatal(err)
  }

  if resp.StatusCode != http.StatusOK {
    t.Errorf("Status should be ok, got %d", resp.StatusCode)
  }

  // Further notes on testing a static site: 
  // https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/
}

func TestNonExistentRoute(t *testing.T) {
  r := newRouter()

  mockServer := httptest.NewServer(r)

  resp, err := http.Post(mockServer.URL + "/", "", nil)

  if err != nil {
    t.Fatal(err)
  }

  if resp.StatusCode != http.StatusMethodNotAllowed {
    t.Errorf("Status code should be 405, got %d", resp.StatusCode)
  }

  defer resp.Body.Close()
  b, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    t.Fatal(err)
  }

  respString := string(b)
  expected := ""

  if respString != expected {
    t.Errorf("Response should be %s, but we got %s", expected, respString)
  }

}
