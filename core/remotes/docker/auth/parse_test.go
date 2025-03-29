/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package auth

import (
	"net/http"
	"testing"
)

func FuzzParseAuthHeader(f *testing.F) {
	f.Add(`Digest realm="https://example.com/token",service="example.com",scope="repository:foo/bar:pull,push"`)
	f.Fuzz(func(t *testing.T, v string) {
		h := http.Header{http.CanonicalHeaderKey("WWW-Authenticate"): []string{v}}
		_ = ParseAuthHeader(h)
	})
}
