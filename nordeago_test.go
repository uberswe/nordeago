// MIT License
//
// Copyright (c) 2018 Markus Tenghamn
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package nordeago

import "testing"

func TestReplaceVariable(t *testing.T) {
	valid := "http://someurl/goes/here/this_is_a_token/"
	text := ReplaceVariable("http://someurl/goes/here/{{token}}/", "token", "this_is_a_token")
	if text != valid {
		t.Errorf("replaceVariable was incorrect, got: %s, want: %s.", text, valid)
	}
}

func TestBearerAuthHeader(t *testing.T) {
	valid := "Bearer somelonghashstringthatshouldgohere"
	text := BearerAuthHeader("somelonghashstringthatshouldgohere")
	if text != valid {
		t.Errorf("bearerAuthHeader was incorrect, got: %s, want: %s.", text, valid)
	}
}
