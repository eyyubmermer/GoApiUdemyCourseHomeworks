package main

import (
	"fmt"
	"net/http"

	. "eyyub/utils"
)

func main() {

	mux := http.NewServeMux()

	uName, eMail, pwd, pwdConfirm := "", "", "", ""

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		uName = r.FormValue("username")
		eMail = r.FormValue("email")
		pwd = r.FormValue("pwd")
		pwdConfirm = r.FormValue("pwdConfirm")

		if IsEmpty(uName) || IsEmpty(eMail) || IsEmpty(pwd) || IsEmpty(pwdConfirm) {
			fmt.Fprint(w, "İstenilen bilgiler boş olamaz.")
		} else if pwd != pwdConfirm {
			fmt.Fprint(w, "Şifreler uyuşmuyor.")
		} else {
			fmt.Fprintf(w, "Kayıt başarılı. \nKullanıcı adı: %s, \nemail: %s, \npwd: %s, \npwdConfirm: %s", uName, eMail, pwd, pwdConfirm)
		}
	})
	logName, logpwd := "", ""

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		logName = r.FormValue("username")
		logpwd = r.FormValue("pwd")

		if logName == uName && logpwd == pwd {
			fmt.Fprintf(w, "Giriş başarılı.")
		} else {
			fmt.Fprintf(w, "Giriş başarısız.")
		}
	})

	http.ListenAndServe(":8080", mux)
}
