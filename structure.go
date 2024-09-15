Backend_Service_Assignment/
├── .gitignore                  # Ignore unnecessary files for Git
├── Dockerfile                  # Docker configuration file
├── docker-compose.yml          # Docker Compose configuration
├── go.mod                      # Go module dependencies
├── go.sum                      # Go module checksums
├── main.go                     # Main application entry point
├── README.md                   # Project documentation
├── utils.go                    # Utility functions like error handling
├── internal/
│   ├── models/
│   │   └── models.go           # Struct definitions for MTGCard, Camera, etc.
│   ├── mtg/
│   │   ├── mtg_service.go      # Service for fetching and storing MTG cards
│   │   ├── mtg_service_test.go # Tests for MTG service
│   │   └── test.go             # Additional tests for the MTG service
│   └── camera/
│       ├── camera_service.go   # Service for handling camera data (e.g., parsing CSV)
│       ├── camera_handler.go   # Handlers for camera-related API routes
│       └── db.query.go         # Queries for camera data within a certain radius
└── db.go                       # Database connection logic
