# go-poc
A simple user management web server built in Go as proof of concept of my Go journey


user-mgmt-poc/
├── go.mod                 # module definition (go mod init user-mgmt-poc)
├── go.sum                 # auto-generated dependency checksums
├── main.go                # entrypoint, server setup, routes
├── handlers.go            # HTTP handler functions (create, get, update, delete)
├── models.go              # User struct & maybe validation
├── storage.go             # in-memory DB (map + mutex), CRUD logic
├── main_test.go           # tests for handlers
└── README.md              # notes on running/testing the service