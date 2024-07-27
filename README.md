# Hotel Reservation JSON api backend

This is the project in [here](https://fulltimegodev.teachable.com/courses/full-time-go-dev/lectures/46484616)

## Project outline
- users -> book rom from a hotel
- admins -> going to check reservations/bookings
- authentication and authorization -> JWT tokens
- hotels -> CRUD API / JSON
- rooms -> CRUD API / JSON
- scripts -> database management 
  - seeding
  - migrations

## Resources
### Mongodb Driver
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongo
```bash
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber
```bash
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a docker container
```bash
docker run --name mongodb -d mongo:latest -p 27017:27017 
```