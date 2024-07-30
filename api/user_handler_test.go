package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http/httptest"
	"testing"
)

const testDBName = "test-hotel-reservation"

type testDb struct {
	db.UserStore
}

func setup(t *testing.T) *testDb {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))

	if err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}

	return &testDb{
		UserStore: db.NewMongoUserStore(client, testDBName),
	}

}

func (testingDb *testDb) tearDown(t *testing.T) {
	if err := testingDb.UserStore.Drop(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestHandlePostUser(t *testing.T) {
	testingDb := setup(t)
	defer testingDb.tearDown(t)

	app := fiber.New()
	userHandler := NewUserHandler(testingDb.UserStore)
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
	userHandler := NewUserHandler(testingDb.UserStore)
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
