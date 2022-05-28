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
			fmt.Fprintf(w, "Kullanıcı adı: %s, email: %s, pwd: %s, pwdConfirm: %s", uName, eMail, pwd, pwdConfirm)
		}
	})

	http.ListenAndServe(":8080", mux)
}
