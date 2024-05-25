package httpserver

import "net/http"

// modified https://gist.github.com/husobee/fd23681261a39699ee37

type middleware func(HandlerFunc) HandlerFunc

type middlewares []middleware

func (h middlewares) Wrap(hf HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = buildChain(hf, h...)(w, r)
	}
}

func buildChain(f HandlerFunc, m ...middleware) HandlerFunc {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise nest the handlerfuncs
	return m[0](buildChain(f, m[1:cap(m)]...))
}
