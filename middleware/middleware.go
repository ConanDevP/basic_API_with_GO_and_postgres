package middleware

import (
	"log"
	"net/http"
	"time"
)

type FuncHandler func(http.ResponseWriter, *http.Request)
 /**Crea los log sobre peticiones **/
func Log(f FuncHandler)FuncHandler{
	return func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		log.Printf("Metodo %q URL %q",r.Method, r.URL.Path)
		f(w,r)
			log.Printf("Tardo %b Milisegundos en ejecutarce", time.Since(inicio).Milliseconds())
	}
}


func Authenticatos(f FuncHandler)FuncHandler{
	return  func(w http.ResponseWriter, r *http.Request) {
		tockend := r.Header.Get("Authorization")
		if tockend != "Algun-token"{
			//responder
			forbidden(w,r)
			return
		}
		f(w,r)
	}

}

func forbidden(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autorización"))
}