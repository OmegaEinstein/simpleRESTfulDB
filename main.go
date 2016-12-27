package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
	"html"
	"sync"
)

type Hello struct {

}

type db struct {
	l sync.Mutex
	db_map map[string]interface{}

}
var DB_g = db{db_map:make(map[string]interface{})}


func (he Hello) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	url_list := strings.Split(html.EscapeString(r.URL.Path), "/")
	//w.Write([]byte(url_list[2]))
	if url_list[1] == "get" && len(url_list) == 3 {
		DB_g.l.Lock()
		defer DB_g.l.Unlock()
		v, exist := DB_g.db_map[url_list[2]]
		if exist {
			v_str := v.(string)
			w.Write([]byte(v_str))


		}else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("key %s not found", url_list[2])))
		}
	}else if url_list[1] == "set" && len(url_list) == 4 {
		DB_g.l.Lock()
		defer DB_g.l.Unlock()
		DB_g.db_map[url_list[2]] = url_list[3]
	}else {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(fmt.Sprintf("error request length")))
	}
	return


}

func main()  {
	var h Hello
	err := http.ListenAndServe("0.0.0.0:4000", h)
	if err != nil {
		log.Fatal(err)
	}

}
