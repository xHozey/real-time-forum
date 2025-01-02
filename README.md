real-time-forum/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── api/                    # API handlers
│   │   ├── auth.go            # Authentication handlers
│   │   ├── posts.go           # Post-related handlers
│   │   ├── comments.go        # Comment-related handlers
│   │   └── messages.go        # Private message handlers
│   ├── middleware/
│   │   ├── auth.go            # Authentication middleware
│   │   └── logger.go          # Logging middleware
│   ├── models/
│   │   ├── user.go            # User model
│   │   ├── post.go            # Post model
│   │   ├── comment.go         # Comment model
│   │   └── message.go         # Private message model
│   ├── database/
│   │   ├── sqlite.go          # Database connection and initialization
│   │   └── migrations/        # SQL migration files
│   └── websocket/
│       ├── client.go          # WebSocket client handling
│       ├── hub.go             # WebSocket hub for managing connections
│       └── message.go         # WebSocket message handling
├── web/
│   ├── static/
│   │   ├── css/
│   │   │   ├── main.css       # Main stylesheet
│   │   │   ├── auth.css       # Auth-related styles
│   │   │   ├── posts.css      # Post-related styles
│   │   │   └── chat.css       # Chat-related styles
│   │   ├── js/
│   │   │   ├── main.js        # Main JavaScript file
│   │   │   ├── auth.js        # Authentication handling
│   │   │   ├── router.js      # SPA routing
│   │   │   ├── posts.js       # Post-related functionality
│   │   │   ├── comments.js    # Comment handling
│   │   │   ├── chat.js        # Private messaging
│   │   │   └── websocket.js   # WebSocket client implementation
│   │   └── img/              # Images and assets
│   └── index.html            # Single HTML file for SPA
├── configs/
│   └── config.go             # Configuration handling
├── pkg/
│   └── utils/               # Shared utilities
├── .gitignore
├── go.mod
├── go.sum
└── README.md