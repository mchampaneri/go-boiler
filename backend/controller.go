package main

import (
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	View(w, r, nil, "index.html")
}

// Authentication Related Functions //

func registerUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		View(w, r, nil, "auth/register.html")
	} else if r.Method == "POST" {
		r.ParseForm()
		user := &User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		user.RegisterUser()
		http.Redirect(w, r, "/users", 302)
	}

}

func allUsers(w http.ResponseWriter, r *http.Request) {
	users := AllUsers()
	data := make(map[string]interface{})
	data["users"] = users
	View(w, r, data, "users.html")
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		View(w, r, nil, "auth/login.html")
	} else if r.Method == "POST" {
		r.ParseForm()
		user := &User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		pass, su := user.LoginUser()
		if pass == true {
			issueSession(w,r,su)
			http.Redirect(w, r, "/users", 302)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	}
}

func issueSession(w http.ResponseWriter, r *http.Request, su *User) bool {
	session, _ := UserSession.Get(r, "mvc-user-session")

	session.Values["id"] = su.Id
	session.Values["name"] = su.Name
	session.Values["email"] = su.Email
	session.Values["auth"] = true
	session.Values["role"] = su.Role
	session.Values["profile_pic"] = su.ProfilePic
	session.Values["cover_pic"] = su.CoverPic

	if su.Status == AccountActive {
		session.Values["active"] = true
		session.Values["message"] = "Welcome"
	}else{
		session.Values["active"] = false
		session.Values["message"] = "Please Active Your Account by Verifying Your Email Address"
	}

	session.Save(r, w)
	return  true
}

func logout(req *http.Request, w http.ResponseWriter) bool {
	session, err := UserSession.Get(req, "mvc-user-session")
	if err == nil {
		for k := range session.Values {
			delete(session.Values, k)
		}
		session.Options.MaxAge = -1
		session.Save(req, w)
		http.Redirect(w, req, "/", http.StatusMovedPermanently)
		return true
	}
	return false
}



