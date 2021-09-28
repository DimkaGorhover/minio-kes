// Copyright 2020 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package http

import (
	"net/http"
	"time"
)

// Timeout returns an HTTP handler that aborts f
// after the given time limit.
//
// The request times out when it takes longer then
// the given time limit to read the request body
// and write a response back to the client.
//
// Once the timeout exceeds, any further Write call
// by f to the http.ResponseWriter will return
// http.ErrHandlerTimeout. Further, if the timeout
// exceeds before f writes an HTTP status code then
// Timeout will return 503 ServiceUnavailable to the
// client.
//
// Timeout cancels the request context before aborting f.
func Timeout(after time.Duration, f http.HandlerFunc) http.HandlerFunc {
	const Message = `{"message":"request timeout exceeded"}`
	var handler = http.TimeoutHandler(f, after, Message)
	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}
