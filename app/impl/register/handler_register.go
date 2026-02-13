package register

import "net/http"

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	http.Handle("/getAll", http.HandlerFunc(func(GetAll) {}))
}
