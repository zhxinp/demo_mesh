package main

import "net/http"

func main() {
	http.HandleFunc("/validate",serveValidate)
	http.HandleFunc("/admitpilot",serveAdmitPilot)
	http.HandleFunc("/admitmixer",serveAdmitMixer)
	http.ListenAndServe("localhost:8000",nil)

}

func serveValidate(w http.ResponseWriter, r *http.Request) {

}

func serveAdmitPilot(w http.ResponseWriter, r *http.Request) {

}

func serveAdmitMixer(w http.ResponseWriter, r *http.Request) {

}
