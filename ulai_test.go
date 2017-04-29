package ulai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChat(t *testing.T) {
	expected := &ChatResponse{
		Status: "success",
		Result: "poyopi!",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(expected)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		return
	}))
	defer ts.Close()

	client := NewClient()
	client.SetUri(ts.URL)
	client.SetKey("secret")

	res, err := client.Chat(context.Background(), "poyoyo?")
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}
	if res != expected.Result {
		t.Errorf("want %q but %q", expected.Result, res)
	}
}
