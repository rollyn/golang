package lib


import (
		"net/http"
)


func Index() http.HandlerFunc  {
	return http.HandlerFunc(func(w http.ResponseWriter, 
								 r *http.Request) {
    	http.ServeFile(w, r, "./static/"+r.URL.Path[1:])
  	})
	
}