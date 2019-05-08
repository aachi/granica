/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package tests

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"gitlab.com/mikrowezel/backend/granica/pkg/authentication"
	authServ "gitlab.com/mikrowezel/backend/granica/pkg/authentication"
	"gitlab.com/mikrowezel/backend/granica/pkg/test/bootstrap"
)

var (
	th            bootstrap.TestHandler
	authURL       string
	signupURL     string
	updateURL     string
	cancelURL     string
	signinURL     string
	signoutURL    string
	createURL     string
	removeURL     string
	user1         = "25369b95-dc28-44ad-9de5-c8f9868547b8"
	user2         = "41ad0ab8-9be4-4dc2-baff-d5888b2ef65f"
	user1Username = "admin"
	user1Role     = "admin"
	user2Username = "admin"
	user2Role     = "admin"
)

func init() {
	authURL = fmt.Sprintf("%s", th.API.ServerURL)
	signupURL = fmt.Sprintf("%s/sign-up", th.API.ServerURL)
	updateURL = fmt.Sprintf("%s/update", th.API.ServerURL)
	cancelURL = fmt.Sprintf("%s/cancel", th.API.ServerURL)
	signinURL = fmt.Sprintf("%s/sign-in", th.API.ServerURL)
	signoutURL = fmt.Sprintf("%s/sign-out", th.API.ServerURL)
	createURL = fmt.Sprintf("%s/create", th.API.ServerURL)
	removeURL = fmt.Sprintf("%s/remove", th.API.ServerURL)
	authServ.Run()
}

func TestMain(m *testing.M) {
	th.Start(m)
}

func TestSignup(t *testing.T) {
	log.Println("[INFO] TestSignup...")
	userJSON := `
	{
		"data": {
			"name": "Arthur",
			"username": "aquaman",
			"password": "sevenseas",
			"email": "arthurcurry@gmail.com"
		}
	}
	`
	th.Reader = strings.NewReader(userJSON)
	request, _ := http.NewRequest("POST", signupURL, th.Reader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Error(err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
}

func TestSignupAlreadySignuped(t *testing.T) {
	log.Println("[INFO] TestSignupAlreadySignuped...")
	// userJSON := `{"data": {"name": "admin", "username": "Admin", "password": "password", "email": "admin@gmail.com"}}`
	userJSON := `
	{
		"data": {
			"name": "admin",
			"username": "Admin",
			"password": "password",
			"email": "admin@gmail.com"
		}
	}
	`
	th.Reader = strings.NewReader(userJSON)
	request, _ := http.NewRequest("POST", signupURL, th.Reader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode == http.StatusCreated {
		t.Errorf("Status: %d | Expected: 201-StatusCreated", res.StatusCode)
	}
}

func TestUpdate(t *testing.T) {
	log.Println("[INFO] TestUpdate...")
	if true {
		t.Errorf("Not implemented test: %s", "TestUpdate")
	}
}

func TestCancel(t *testing.T) {
	log.Println("[INFO] TestCancel...")
	if true {
		t.Errorf("Not implemented test: %s", "TestCancel")
	}
}

func TestSigninByUsername(t *testing.T) {
	log.Println("[INFO] TestSigninByUsername...")
	userJSON := `
	{
		"data": {
			"username": "admin",
			"password": "darkknight"
		}
	}
	`
	th.Reader = strings.NewReader(userJSON)
	request, _ := http.NewRequest("POST", signinURL, th.Reader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}

func TestSigninByEmail(t *testing.T) {
	log.Println("[INFO] TestSigninByEmail...")
	userJSON := `
	{
		"data": {
			"email": "admin@gmail.com",
			"password": "darkknight"
		}
	}
	`
	th.Reader = strings.NewReader(userJSON)
	request, _ := http.NewRequest("POST", signinURL, th.Reader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}

func TestSigninBadCredentials(t *testing.T) {
	log.Println("[INFO] TestSigninByEmail...")
	userJSON := `
	{
		"data": {
			"username": "Flash",
			"email": "barryallen@gmail.com",
			"password": "speedy"
		}
	}
	`
	th.Reader = strings.NewReader(userJSON)
	request, _ := http.NewRequest("POST", signinURL, th.Reader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode == http.StatusOK {
		t.Errorf("Status: %d | Expected: 401-StatusUnauthorized", res.StatusCode)
	}
}

func TestSignout(t *testing.T) {
	log.Println("[INFO] TestSignout...")
	if true {
		t.Errorf("Not implemented test: %s", "TestSignout")
	}
}

func TestCreateUser(t *testing.T) {
	// Not implemented.
	if 1 == 200 {
		t.Errorf("Status: %d | Expected: 200-StatusOk", -1)
	}
}

func TestRemove(t *testing.T) {
	log.Println("[INFO] TestRemove...")
	th.Reader = strings.NewReader("")
	userURL := fmt.Sprintf("%s", user1)
	request, _ := http.NewRequest("DELETE", userURL, th.Reader)
	th.AuthorizeRequest(request, user1, user1Username, user1Role)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusNoContent", res.StatusCode)
	}
}

func TestGetUsers(t *testing.T) {
	log.Println("[INFO] TestGetUsers...")
	th.Reader = strings.NewReader("")
	request, _ := http.NewRequest("GET", authURL, th.Reader)
	th.AuthorizeRequest(request, user1, user1Username, user1Role)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}

func TestGet(t *testing.T) {
	log.Println("[INFO] TestGet...")
	th.Reader = strings.NewReader("")
	userURL := fmt.Sprintf("%s", user1)
	request, _ := http.NewRequest("GET", userURL, th.Reader)
	th.AuthorizeRequest(request, user1, user1Username, user1Role)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}

func TestGetByUsername(t *testing.T) {
	log.Println("[INFO] TestGetByUsername...")
	th.Reader = strings.NewReader("")
	userURL := fmt.Sprintf("%s", "admin")
	request, _ := http.NewRequest("GET", userURL, th.Reader)
	th.AuthorizeRequest(request, user1, user1Username, user1Role)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status: %d | Expected: 200-StatusOk", res.StatusCode)
	}
}

func TestUpdateUser(t *testing.T) {
	log.Println("[INFO] TestUpdateUser...")
	userJSON := `
	{
		"data": {
			"username": "administrator",
			"password": "password",
			"email": "administrator@gmail.com",
			"firstName": "Tim",
			"middleNames": "",
			"lastName": "Drake",
			"startedAt": 1735693261
		}
	}
	`
	th.Reader = strings.NewReader(userJSON)
	userURL := fmt.Sprintf("%s", user1)
	request, _ := http.NewRequest("PUT", userURL, th.Reader)
	th.AuthorizeRequest(request, user1, user1Username, user1Role)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 204-StatusNoContent", res.StatusCode)
	}
}

func TestUpdateUserWithDatabaseVerify(t *testing.T) {
	log.Println("[INFO] TestUpdateUserWithDatabaseVerify...")
	newUsername := "administrator"
	newEmail := "administrator@gmail.com"
	userJSON := fmt.Sprintf(`
	{
		"data": {
			"username": "%s",
			"password": "%s",
			"email": "administrator@gmail.com",
			"firstName": "Tim",
			"middleNames": "",
			"lastName": "Drake",
			"startedAt": 1735693261
		}
	}
	`, newUsername, newEmail)
	th.Reader = strings.NewReader(userJSON)
	userURL := fmt.Sprintf("%s", user1)
	request, _ := http.NewRequest("PUT", userURL, th.Reader)
	th.AuthorizeRequest(request, user1, user1Username, user1Role)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	if res.StatusCode == http.StatusNoContent {
		userRepo, err := authentication.InitRepo(th.MainConfig)
		if err != nil {
			log.Fatal(err)
			return
		}
		user, err := userRepo.Get(user1)
		if err == nil {
			if user.Username == newUsername && user.Email == newEmail {
				log.Println("[INFO] User update: ok.")
			} else {
				error := fmt.Sprintf("Username: '%s' | Expected: '%s' - ", user.Username, newUsername)
				error += fmt.Sprintf("Email: '%s' | Expected: '%s'", user.Email, newEmail)
				t.Error(error)
			}
		} else {
			t.Error(err.Error())
		}
	} else {
		t.Errorf("Status: %d | Expected: 204-StatusNoContent", res.StatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	log.Println("[INFO] TestDeleteUserWithDatabaseVerify...")
	th.Reader = strings.NewReader("")
	userURL := fmt.Sprintf("%s", user1)
	request, _ := http.NewRequest("DELETE", userURL, th.Reader)
	th.AuthorizeRequest(request, user1, user1Username, user1Role)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("Status: %d | Expected: 200-StatusNoContent", res.StatusCode)
	}
}

func TestDeleteUserWithDatabaseVerify(t *testing.T) {
	log.Println("[INFO] TestDeleteUser...")
	th.Reader = strings.NewReader("")
	userURL := fmt.Sprintf("%s", user1)
	request, _ := http.NewRequest("DELETE", userURL, th.Reader)
	th.AuthorizeRequest(request, user1, user1Username, user1Role)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error executing request: %s", err.Error())
		return
	}
	if res.StatusCode == http.StatusNoContent {
		userRepo, err := authentication.InitRepo(th.MainConfig)
		if err != nil {
			log.Fatal(err)
			return
		}
		user, err := userRepo.Get(user1)
		if err != nil {
			log.Println("[INFO] TestDeleteUser: ok")
		} else {
			t.Errorf("User: %s | Expected: 'nil'", user.Username)
		}
	} else {
		t.Errorf("Status: %d | Expected: 204-StatusNoContent", res.StatusCode)
	}
}
