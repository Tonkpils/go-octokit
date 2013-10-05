package octokat

import (
	"github.com/bmizerany/assert"
	"github.com/octokit/octokat/hyper"
	"net/http"
	"testing"
)

func TestClient_User(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("user.json"))
	})

	user, _ := client.User("", nil)

	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, 25, user.PublicGists)

	mux.HandleFunc("/users/jingweno", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("user.json"))
	})

	user, _ = client.User("jingweno", nil)

	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, hyper.Link("https://api.github.com/users/jingweno/repos"), user.ReposURL)
}

func TestUser_UpdateUser(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testBody(t, r, `{"name":"name","email":"email"}`)
		respondWith(w, loadFixture("user.json"))
	})

	var userToUpdate = User{
		Name:  "name",
		Email: "email",
	}

	user, _ := client.UpdateUser(userToUpdate, nil)
	assert.Equal(t, 169064, user.ID)
}

func TestUser_AllUsers(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("users.json"))
	})

	users, _ := client.AllUsers(nil)
	assert.Equal(t, 1, len(users))
}
