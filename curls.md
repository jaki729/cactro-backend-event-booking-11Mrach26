# Example curl commands for Booking Backend

## Register Organizer
```sh
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"alicepass","role":"organizer"}'
```

## Register Customer
```sh
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"bob","password":"bobpass","role":"customer"}'
```

## Login Organizer
```sh
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"alicepass"}'
```

## Login Customer
```sh
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"bob","password":"bobpass"}'
```

## List Events (Organizer or Customer)
```sh
curl -X GET http://localhost:8080/events \
  -H "Authorization: Basic YWxpY2U6YWxpY2VwYXNz"
```

## Create Event (Organizer only)
```sh
curl -X POST http://localhost:8080/events/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic YWxpY2U6YWxpY2VwYXNz" \
  -d '{"name":"Go Conference","description":"A conference about Go"}'
```

## Update Event (Organizer only)
```sh
curl -X POST http://localhost:8080/events/update \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic YWxpY2U6YWxpY2VwYXNz" \
  -d '{"id":1,"name":"Go Conference 2026","description":"Updated description"}'
```

## Delete Event (Organizer only)
```sh
curl -X POST http://localhost:8080/events/delete \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic YWxpY2U6YWxpY2VwYXNz" \
  -d '{"id":1}'
```

## Create Booking (Customer only)
```sh
curl -X POST http://localhost:8080/bookings/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic Ym9iOmJvYnBhc3M=" \
  -d '{"event_id":1}'
```

## List My Bookings (Customer only)
```sh
curl -X GET http://localhost:8080/bookings/my \
  -H "Authorization: Basic Ym9iOmJvYnBhc3M="
```

---

**Notes:**
- Replace `alice`/`bob` and passwords as needed.
- The `id` and `event_id` values must match actual event IDs.
- Register users and create events before booking.