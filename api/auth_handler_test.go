package api

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db/fixtures"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandleAuthenticate(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)
	insertedUser := fixtures.AddUser(tdb.Store, "nate", "alcedo", false)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.User)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "nate@alcedo.com",
		Password: "nate_alcedo",
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

func TestHandleAuthenticateFailure(t *testing.T) {
	tdb := setup(t)
	defer tdb.tearDown(t)
	fixtures.AddUser(tdb.Store, "Nate", "Alcedo", false)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.User)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "natealcedo@gmail.com",
		Password: "wrong password",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	expected := fiber.StatusUnauthorized
	actual := res.StatusCode

	if expected != actual {
		t.Errorf("expected %v but got %v", expected, actual)
	}

	var genResp genericResponse
	if err := json.NewDecoder(res.Body).Decode(&genResp); err != nil {
		t.Fatal(err)
	}

	if genResp.Type != "error" {
		t.Fatalf("expected %v but got %v", "error", genResp.Type)
	}

	if genResp.Msg != invalidCredentialsMsg {
		t.Fatalf("expected %v but got %v", invalidCredentialsMsg, genResp.Msg)
	}

}
