package handlr

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type unitTests struct {
	r  *http.Request
	w  *httptest.ResponseRecorder
	m  string
	rt func()
	fh func() *ActionHandler
}

func TestRegisterRoute(t *testing.T) {
	h := New()

	h.RouteFunc("/", func(r *Router) {
		t.Logf("Route registered correctly.")
	})
}

func TestRegisterHandler(t *testing.T) {
	h := New()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Handler for %s registered correctly.", r.URL.Path)
	})

	h.mux.Handle("/", &h.router)
	h.mux.ServeHTTP(w, r)
}

func TestBuildPath(t *testing.T) {
	h := New()

	f := func(bp string, rp string) bool {
		pass := bp == rp
		if pass {
			t.Logf(bp)
		} else {
			t.Fatalf("Wrong path built: %s should've been %s.", bp, rp)
		}
		return pass
	}

	h.RouteFunc("a", func(r1 *Router) {
		r1.RouteFunc("b", func(r2 *Router) {
			r2.RouteFunc("c", func(r3 *Router) {
				r3.RouteFunc("/", func(r4 *Router) {
					if f(r4.buildPath(), "/a/b/c") && f(r3.buildPath(), "/a/b/c") && f(r2.buildPath(), "/a/b") && f(r1.buildPath(), "/a") {
						t.Logf("Paths generated correctly.")
					}
				})
			})
		})
	})
}

func TestFindHandler(t *testing.T) {
	h := New()

	tests := []unitTests{
		{
			r: httptest.NewRequest(http.MethodGet, "/a/b", nil),
			w: httptest.NewRecorder(),
			m: "Handler matched correctly.",
			rt: func() {
				h.RouteFunc("/", func(r1 *Router) {
					r1.RouteFunc("/a", func(r2 *Router) {
						r2.HandlerFunc("/b", func(w http.ResponseWriter, r *http.Request) {
							t.Logf("Handler for %s matched correctly.", r.URL.Path)
						})
					})
				})
			},
			fh: func() *ActionHandler {
				e := len(h.router.children) - 1                         // For /
				a := len(h.router.children[e].children) - 1             // For /a
				b := len(h.router.children[e].children[a].children) - 1 // For /b

				return h.router.children[e].children[a].children[b].handler
			},
		},
		{
			r: httptest.NewRequest(http.MethodGet, "/x/y/z", nil),
			w: httptest.NewRecorder(),
			m: "Handler matched correctly.",
			rt: func() {
				h.HandlerFunc("/x/y/z", func(w http.ResponseWriter, r *http.Request) {
					t.Logf("Handler for %s matched correctly.", r.URL.Path)
				})
			},
			fh: func() *ActionHandler {
				return h.router.children[1].handler
			},
		},
	}

	h.mux.Handle("/", &h.router)

	for i := 0; i < len(tests); i++ {
		u := tests[i]

		u.rt()

		fh := u.fh()
		rh := h.router.findHandler(u.r)

		if rh == fh {
			h.mux.ServeHTTP(u.w, u.r)
		} else {
			t.Fatalf("Handler couldn't be found: %d != %d", fh, rh)
		}
	}

	t.Logf("Handlers found without problems.")
}
