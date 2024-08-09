package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
	"net/http/httptest"
	"reflect"
	"testing"
)

func insertTestUser(t *testing.T, userStore db.UserStore) *types.User {
	user, err := types.NewUserFromParams(&types.CreateUserParams{
		FirstName: "Nate",
		LastName:  "Alcedo",
		Email:     "natealcedo@gmail.com",
		Password:  "supersecurepassword",
	})

	if err != nil {
		t.Fatal(err)
	}

	_, err = userStore.InsertUser(context.Background(), user)

	if err != nil {
		t.Fatal(err)
	}

	return user
}

func TestHandleAuthenticate(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)
	insertedUser := insertTestUser(t, tdb.UserStore)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "natealcedo@gmail.com",
		Password: "supersecurepassword",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	expected := fiber.StatusOK
	actual := resp.StatusCode
	var authResponse AuthResponse

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		t.Fatal(err)
	}

	if authResponse.Token == "" {
		t.Fatal("expected token to be present")
	}

	// Comment out EncryptedPassword since the response doesn't have it
	insertedUser.EncryptedPassword = ""
	if !reflect.DeepEqual(authResponse.User, insertedUser) {
		t.Errorf("expected %v but got %v", insertedUser, authResponse.User)
	}

}
