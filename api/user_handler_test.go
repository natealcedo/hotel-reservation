package api

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/types"
	"net/http/httptest"
	"testing"
)

func TestHandlePostUser(t *testing.T) {
	testingDb := setup(t)
	defer testingDb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(testingDb.User)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		FirstName: "Nate",
		LastName:  "Alcedo",
		Email:     "natealcedo@gmail.com",
		Password:  "randompassword",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	expected := fiber.StatusOK
	actual := resp.StatusCode

	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}

	var user *types.User
	_ = json.NewDecoder(resp.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Errorf("expected user id to be set")
	}

	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expected encrypted password not to be returned from endpoint")
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected %d but got %d", fiber.StatusOK, resp.StatusCode)
	}

	if user.FirstName != params.FirstName {
		t.Errorf("expected %s but got %s", params.FirstName, user.FirstName)
	}

	if user.LastName != params.LastName {
		t.Errorf("expected %s but got %s", params.LastName, user.LastName)
	}

	if user.Email != params.Email {
		t.Errorf("expected %s but got %s", params.Email, user.Email)
	}
}

func TestHandleGetUsers(t *testing.T) {
	testingDb := setup(t)
	defer testingDb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(testingDb.User)
	app.Get("/", userHandler.HandleGetUsers)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	expected := fiber.StatusOK
	actual := resp.StatusCode

	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}

	var users *[]types.User
	_ = json.NewDecoder(resp.Body).Decode(&users)

	if len(*users) != 0 {
		t.Errorf("expected empty array but got %d", len(*users))
	}
}

func TestHandleGetUserById(t *testing.T) {
	testingDb := setup(t)
	defer testingDb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(testingDb.User)
	app.Post("/", userHandler.HandlePostUser)
	app.Get("/:id", userHandler.HandleGetUserById)

	params := types.CreateUserParams{
		FirstName: "Nate",
		LastName:  "Alcedo",
		Email:     "natealcedo@gmail.com",
		Password:  "randompassword",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	user := &types.User{}

	req = httptest.NewRequest("GET", "/"+user.ID.Hex(), nil)
	_ = json.NewDecoder(resp.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Errorf("expected user id to be set")
	}

	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expected encrypted password not to be returned from endpoint")
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("expected %d but got %d", fiber.StatusOK, resp.StatusCode)
	}

	if user.FirstName != params.FirstName {
		t.Errorf("expected %s but got %s", params.FirstName, user.FirstName)
	}

	if user.LastName != params.LastName {
		t.Errorf("expected %s but got %s", params.LastName, user.LastName)
	}

	if user.Email != params.Email {
		t.Errorf("expected %s but got %s", params.Email, user.Email)
	}
}

func TestHandleDeleteUserById(t *testing.T) {
	testingDb := setup(t)
	defer testingDb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(testingDb.User)
	app.Post("/", userHandler.HandlePostUser)
	app.Get("/:id", userHandler.HandleGetUserById)
	app.Delete("/:id", userHandler.HandleDeleteUserById)

	params := types.CreateUserParams{
		FirstName: "Nate",
		LastName:  "Alcedo",
		Email:     "natealcedo@gmail.com",
		Password:  "randompassword",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	user := &types.User{}
	_ = json.NewDecoder(resp.Body).Decode(&user)
	userID := user.ID.Hex()

	req = httptest.NewRequest("DELETE", "/"+userID, nil)
	req.Header.Add("Content-Type", "application/json")
	expected := fiber.StatusOK
	actual := resp.StatusCode
	resp, err = app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}

	var responseBody map[string]string
	_ = json.NewDecoder(resp.Body).Decode(&responseBody)

	expectedDeletedId := responseBody["deleted"]

	if expectedDeletedId != userID {
		t.Errorf("expected %s but got %s", userID, expectedDeletedId)
	}
}

func TestHandlePutUserById(t *testing.T) {
	testingDb := setup(t)
	defer testingDb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(testingDb.User)
	app.Post("/", userHandler.HandlePostUser)
	app.Get("/:id", userHandler.HandleGetUserById)
	app.Put("/:id", userHandler.HandlePutUserById)

	// Create User
	createUserParams := types.CreateUserParams{
		FirstName: "Nate",
		LastName:  "Alcedo",
		Email:     "natealcedo@gmail.com",
		Password:  "randompassword",
	}

	b, _ := json.Marshal(createUserParams)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	// Created User
	user := &types.User{}
	_ = json.NewDecoder(resp.Body).Decode(&user)
	userID := user.ID.Hex()

	// Update User
	updateUserParams := types.UpdateUserParams{
		FirstName: "Bob",
		LastName:  "Marley",
	}

	updateBytes, _ := json.Marshal(updateUserParams)
	updateReq := httptest.NewRequest("PUT", "/"+userID, bytes.NewReader(updateBytes))
	updateReq.Header.Add("Content-Type", "application/json")
	resp, err = app.Test(updateReq)

	if err != nil {
		t.Fatal(err)
	}

	expected, actual := fiber.StatusOK, resp.StatusCode

	if expected != actual {
		t.Errorf("expected %d but got %d", expected, actual)
	}

	// Get Updated User
	updatedUser := &types.User{}
	req = httptest.NewRequest("GET", "/"+userID, nil)
	req.Header.Add("Content-Type", "application/json")

	// The error here: I didn't send the GET REQUEST to the server after updating the user
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	_ = json.NewDecoder(resp.Body).Decode(&updatedUser)

	if len(updatedUser.ID) == 0 {
		t.Errorf("expected user id to be set")
	}

	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expected encrypted password not to be returned from endpoint")
	}

	if updatedUser.FirstName != updateUserParams.FirstName {
		t.Errorf("expected %s but got %s", updateUserParams.FirstName, updatedUser.FirstName)
	}

	if updatedUser.LastName != updateUserParams.LastName {
		t.Errorf("expected %s but got %s", updateUserParams.LastName, updatedUser.LastName)
	}

	if user.Email != createUserParams.Email {
		t.Errorf("expected %s but got %s", createUserParams.Email, user.Email)
	}
}
