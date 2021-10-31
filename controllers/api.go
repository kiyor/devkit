package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Resp struct {
	Code int
	Data interface{}
}

func NewResp(w http.ResponseWriter, data interface{}, code ...int) []byte {
	c := 0
	if len(code) > 0 {
		c = code[0]
	}
	r := &Resp{
		Code: c,
		Data: data,
	}
	b, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	w.Header().Add("content-type", "application/json")
	w.Write(b)
	return b
}

func NewErrResp(w http.ResponseWriter, code int, err error) []byte {
	return NewResp(w, err.Error(), code)
}

func callback(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()
	log.Println(string(body))
}

func (c *Controller) ApiHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	action := q.Get("action")
	switch action {
	case "ping":
		c.apiPing(w, r)
	// add your case here
	default:
		w.Write([]byte("api ok"))
	}
}

func (c *Controller) apiPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func toJSON(i interface{}) string {
	b, _ := json.MarshalIndent(i, "", "  ")
	return string(b)
}
