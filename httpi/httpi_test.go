package httpi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	req, err := http.NewRequest("GET", "/list/?lat=37.486714&long=-122.226306&radius=.5", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(List)
	handler.ServeHTTP(w, req)

	type args struct {
		w   http.ResponseWriter
		req *http.Request
	}
	a := args{
		w:   w,
		req: req,
	}
	tests := []struct {
		name string
		args args
	}{
		{"HTTP-List", a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			List(tt.args.w, tt.args.req)
		})
	}
}
