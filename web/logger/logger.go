package logger

import (
	"github.com/golang/glog"
	"net/http"
	"time"
)

func Logger(inner func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner(w, r)

		glog.V(2).Infof(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}
