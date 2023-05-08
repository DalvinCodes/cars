package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/DalvinCodes/cars/utils"
)

const traceIDHeaderName = "X-Trace-ID"

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// check for an existing trace ID in the request headers
		traceID := r.Header.Get(traceIDHeaderName)

		// generate a new id if not found
		if traceID == "" {
			traceID = utils.GenerateID()
			r.Header.Set(traceIDHeaderName, traceID)
		}

		// add the trace ID header to the response
		w.Header().Set(traceIDHeaderName, traceID)

		// create a context with the trace ID value
		ctx := context.WithValue(r.Context(), traceIDHeaderName, traceID)

		// call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))

		// extract the trace ID from the context
		traceID = ctx.Value(traceIDHeaderName).(string)

		// log the request details, including the trace ID
		log.Printf("Method: %s URI: %s  TraceID: %s Latency: %s IPAddress: %s", r.Method, r.RequestURI, traceID, time.Since(start), r.RemoteAddr)
	}
}

func AddHeaders(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next(w, r)
	}
}
