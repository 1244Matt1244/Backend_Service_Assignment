Backend_Service_Assignment/
├── cmd/
│   └── main.go           # Entry point
├── internal/
│   ├── camera/
│   │   ├── camera_handler.go
│   │   ├── camera_service.go
│   │   ├── camera_service_test.go
│   │   └── cameras.csv
│   ├── mtg/
│   │   ├── fetch_cards.go
│   │   ├── mtg_service_test.go
│   │   └── mtg_cards.csv
│   ├── db/
│   │   ├── db.go
│   │   └── db_query.go   
│   └── utils/
│       └── utils.go
├── configs/
│   ├── docker-compose.yml
│   └── Dockerfile
├── routes/
│   └── routes.go
├── tests/
│   └── test.go
├── go.mod
├── go.sum
├── README.md
└── .gitignore
