package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andkolbe/chirper-gin/models"
	"github.com/stretchr/testify/assert"
)

type mockUserModel struct{}

func (m *mockUserModel) GetAllUsers() ([]models.User, error) {
	var users []models.User

	users = append(users, models.User{ID: 100, Name: "Odin", Email: "odin@aol.com"})
	users = append(users, models.User{ID: 101, Name: "Winnie", Email: "winnie@aol.com"})

	return users, nil
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}


// func TestUsersIndex(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/users", nil)


// 	env := Env{users: &mockUserModel{}}


// 	expected := "100, Odin, odin@aol.com\n101, Winnie, winnie@aol.com\n"
// 	if expected != w.Body.String() {
// 		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
// 	}

// }