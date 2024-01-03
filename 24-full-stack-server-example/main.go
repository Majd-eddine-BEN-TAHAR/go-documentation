package main

import (
	"html/template"
	"myapp/controller"
	"myapp/model"
	"net/http"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    // Initialize Logrus logger
    log := logrus.New()
    log.SetLevel(logrus.InfoLevel)
    log.SetFormatter(&logrus.JSONFormatter{})

    // Initialize GORM with a SQLite database
    db, err := gorm.Open(sqlite.Open("myapp.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to the database: ", err)
    }

    // Auto migrate our User model
    if err := db.AutoMigrate(&model.User{}); err != nil {
        log.Fatal("failed to migrate database: ", err)
    }

    // Initialize UserController with DB and logger
    userController := controller.NewUserController(db, log)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))


	// Setup routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user", userController.GetUser)
	http.HandleFunc("/create_user", userController.CreateUser)
	http.HandleFunc("/new_user", userController.CreateUserForm)
	http.HandleFunc("/list_users", userController.ListUsers)

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}


// homeHandler serves the homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("view/home.html")
    if err != nil {
        // Handle error
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    t.Execute(w, nil)
}

