// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build go1.7

package traceutil_test

import (
	"log"
	"net/http"

	"cloud.google.com/go/trace"
	"cloud.google.com/go/trace/traceutil"
)

var traceClient *trace.Client

func ExampleHTTPClient_Do() {
	client := traceutil.NewHTTPClient(traceClient, nil) // traceClient is a *trace.Client

	req, _ := http.NewRequest("GET", "https://metadata/users", nil)
	if _, err := client.Do(req); err != nil {
		log.Fatal(err)
	}
}

func ExampleHTTPClient_Do_propagation() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		span := traceClient.SpanFromRequest(r) // traceClient is a *trace.Client
		defer span.Finish()

		client := traceutil.NewHTTPClient(traceClient, nil)
		ctx := trace.NewContext(r.Context(), span)

		req, _ := http.NewRequest("GET", "https://metadata/users", nil)
		req = req.WithContext(ctx)

		// The outgoing request will be traced with r's trace ID.
		if _, err := client.Do(req); err != nil {
			log.Fatal(err)
		}
	})
}
