# mariomenjr/handlr

[![GoDoc](https://godoc.org/github.com/mariomenjr/handlr?status.svg)](https://godoc.org/github.com/mariomenjr/handlr)
[![CircleCI](https://circleci.com/gh/mariomenjr/handlr/tree/main.svg?style=svg)](https://circleci.com/gh/mariomenjr/handlr/tree/main)

Easily manage routes and handlers on top of *http.ServeMux.

---

## Install

As you would do with almost any other Go package.

```bash
go get -u github.com/mariomenjr/handlr
```

---

## Examples

### Handlers

You can register paths and handlers directly:

```go
// main.go

func main() {
	h := handlr.New()

	h.HandleFunc("/feed", feedHandler)

	r.Start(1993)
}

func feedHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

### Routes

You can also register `Route`s which allow you to organize your handlers (or even sub-Routes) into multiple files in an elegant way.

```
├── main.go
├── feed.route.go
├── acccount.route.go
```

```go
// main.go

func main() {
	h := handlr.New()

	h.RouteFunc("/feed", feedRoute)
	h.RouteFunc("/account", accountRoute)
	
	r.Start(1993)
}
```

```go
// feed.route.go

func feedRoute(r *handlr.Router) {
	r.HandleFunc("/latest", latestHandler)

	r.RouteFunc("/custom", feedCustomRoute)
}

func latestHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}

func feedCustomRoute(r *handlr.Router) {
	r.HandleFunc("/monthly", feedCustomMonthlyHandler)
}

func feedCustomMonthlyHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

```go
// account.route.go

func accountRoute(r *handlr.Router) {
	r.HandleFunc("/profile", accountProfileHandlr)
	r.HandleFunc("/settings", accountSettingsHandlr)
}

func accountProfileHandlr(w http.ResponseWriter, r *http.Request) {
	// ...
}

func accountSettingsHandlr(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

---

## License 
The source code of this project is under [MIT License](https://opensource.org/licenses/MIT).
