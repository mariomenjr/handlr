package handlr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterRoute(t *testing.T) {
	h := New()

	h.Route("/", func(r *Router) {
		t.Log("Route registered correctly.")
	})
}

func TestRegisterHandler(t *testing.T) {
	h := New()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.Handler("/", func(w http.ResponseWriter, r *http.Request) {
		t.Log("Handler registered correctly.")
	})

	(*h.router.children[0].handler)(w, r)
}

func TestBuildPath(t *testing.T) {
	h := New()

	f := func(bp string, rp string) bool {
		pass := bp == rp
		if pass {
			t.Log(bp)
		} else {
			t.Errorf("Wrong path built: %s should've been %s.", bp, rp)
		}
		return pass
	}

	h.Route("a", func(r1 *Router) {
		r1.Route("b", func(r2 *Router) {
			r2.Route("c", func(r3 *Router) {
				r3.Route("/", func(r4 *Router) {
					if f(r4.buildPath(), "/a/b/c") && f(r3.buildPath(), "/a/b/c") && f(r2.buildPath(), "/a/b") && f(r1.buildPath(), "/a") {
						t.Log("Paths generated correctly.")
					}
				})
			})
		})
	})
}

// TODO: Simplify
func TestFindHandler(t *testing.T) {
	h := New()

	r := httptest.NewRequest(http.MethodGet, "/a/b", nil)
	w := httptest.NewRecorder()

	// Complicated registration

	h.Route("/", func(r1 *Router) {
		r1.Route("/a", func(r2 *Router) {
			r2.Handler("/b", func(w http.ResponseWriter, r *http.Request) {
				t.Log("Handler matched correctly.")
			})
		})
	})

	e := len(h.router.children) - 1                         // For /
	a := len(h.router.children[e].children) - 1             // For /a
	b := len(h.router.children[e].children[a].children) - 1 // For /b

	rh1 := h.router.children[e].children[a].children[b].handler
	fh1 := h.router.findHandler(r)

	if fh1 == rh1 {
		h.mux.ServeHTTP(w, r)
	} else {
		t.Errorf("Handler couldn't be found: %d != %d", fh1, rh1)
	}

	// Simple registration

	r = httptest.NewRequest(http.MethodGet, "/x/y/z", nil)

	h.Handler("/x/y/z", func(w http.ResponseWriter, r *http.Request) {
		t.Log("Handler matched correctly.")
	})

	rh2 := h.router.children[1].handler
	fh2 := h.router.findHandler(r)

	if fh2 == rh2 {
		h.mux.ServeHTTP(w, r)
	} else {
		t.Errorf("Handler couldn't be found: %d != %d", fh2, rh2)
	}
}
