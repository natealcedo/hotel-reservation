package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/natealcedo/hotel-reservation/api"
	"github.com/natealcedo/hotel-reservation/db"
	"github.com/natealcedo/hotel-reservation/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The address of the server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Hotel:   db.NewMongoHotelStore(client),
		Room:    db.NewMongoRoomStore(client, db.NewMongoHotelStore(client)),
		Booking: db.NewMongoBookingStore(client),
	}

	app := fiber.New(config)
	auth := app.Group("/api")
	apiV1 := app.Group("/api/v1", middleware.JWTAuthentication(store.User))
	admin := apiV1.Group("/admin", middleware.AdminAuth)

	// handlers
	userHandler := api.NewUserHandler(store.User)
	hotelHandler := api.NewHotelHandler(store)
	authHandler := api.NewAuthHandler(store.User)
	roomHandler := api.NewRoomHandler(store)
	bookingHandler := api.NewBookingHandler(store)

	// Auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// Users
	apiV1.Get("/users/:id", userHandler.HandleGetUserById)
	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Post("/users", userHandler.HandlePostUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUserById)
	apiV1.Put("/users/:id", userHandler.HandlePutUserById)

	// Hotels
	apiV1.Get("/hotels", hotelHandler.HandleGetHotels)
	apiV1.Get("/hotels/:id", hotelHandler.HandleGetHotelById)
	apiV1.Get("/hotels/:id/rooms", hotelHandler.HandleGetRooms)

	// Rooms
	apiV1.Get("/rooms", roomHandler.HandleGetRooms)
	apiV1.Post("/rooms/:id/book", roomHandler.HandleBookRoom)

	// Bookings
	apiV1.Get("/bookings/:id", bookingHandler.HandleGetBookingByID)
	apiV1.Get("/bookings/:id/cancel", bookingHandler.UpdateBookingByID)

	// Admins
	admin.Get("/bookings", bookingHandler.HandleGetBookings)

	err = app.Listen(*listenAddr)

	if err != nil {
		panic(err)
	}
}
