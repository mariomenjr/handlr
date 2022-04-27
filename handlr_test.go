package handlr

import (
	"net/http/httptest"
	"testing"
)

func TestStartHandlr(t *testing.T) {
	h := New()

	s := httptest.NewServer(&h.router)
	defer s.Close()

	t.Logf("Server started at %s", s.URL)
}
