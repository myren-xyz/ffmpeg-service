package main

import (
	"net/http"
	"strings"
)

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if origin ends in .myren.xyz , allow it
		origin := r.Header.Get("origin")
		if origin == "" {
			origin = r.Header.Get("Origin")
		}
		if strings.HasSuffix(origin, ".myren.xyz") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// allow credentials
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		next.ServeHTTP(w, r)
	}
}
