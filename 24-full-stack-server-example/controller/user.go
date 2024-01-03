package controller

import (
	"errors"
	"html/template"
	"math"
	"myapp/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// UserController struct holds the database connection and logger
type UserController struct {
    DB  *gorm.DB
    Log *logrus.Logger
}

// NewUserController creates a new instance of UserController
func NewUserController(db *gorm.DB, log *logrus.Logger) *UserController {
    return &UserController{DB: db, Log: log}
}


func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
    // Define a structure to hold template data, including error messages
    type templateData struct {
        User  model.User
        Error string
    }

    data := templateData{}

    // Extracting user ID from the request
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        data.Error = "User ID is required"
        uc.renderTemplate(w, "user.html", data)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        data.Error = "Invalid user ID"
        uc.renderTemplate(w, "user.html", data)
        return
    }

    // Retrieving user from the database
    user, err := model.GetUserById(uc.DB, id)
    if err != nil {
        if errors.Is(err, model.ErrUserNotFound) {
            data.Error = "User not found"
        } else {
            uc.Log.WithError(err).Error("Error retrieving user")
            data.Error = "An internal error occurred"
        }
        uc.renderTemplate(w, "user.html", data)
        return
    }

    data.User = user
    uc.renderTemplate(w, "user.html", data)
}

// renderTemplate is a helper function to parse and execute a template
func (uc *UserController) renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t, err := template.ParseFiles("view/" + tmpl)
    if err != nil {
        uc.Log.WithError(err).Error("Error parsing template")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if err := t.Execute(w, data); err != nil {
        uc.Log.WithError(err).Error("Error executing template")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}


// CreateUserForm displays the form for creating a new user
func (uc *UserController) CreateUserForm(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("view/create_user.html")
    if err != nil {
        uc.Log.WithError(err).Error("Error parsing create user template")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    t.Execute(w, nil)
}

// CreateUser handles the creation of a new user
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse form data
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Invalid form data", http.StatusBadRequest)
        return
    }

    name := r.FormValue("name")
    email := r.FormValue("email")

    // Server-side validation
    if len(name) < 3 {
        http.Error(w, "Name is too short, must be at least 3 characters", http.StatusBadRequest)
        return
    }

    if !isValidEmail(email) {
        http.Error(w, "Invalid email format", http.StatusBadRequest)
        return
    }

    user := model.User{Name: name, Email: email}

    // Create user in the database
    if err := model.CreateUser(uc.DB, user); err != nil {
        uc.Log.WithError(err).Error("Error creating user")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Redirect after successful creation or show a success message
    http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect to homepage or relevant page
}

func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
    // Handling pagination
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page <= 0 {
        page = 1
    }

    const itemsPerPage = 2
    users, totalCount, err := model.GetAllUsers(uc.DB, page, itemsPerPage)
    if err != nil {
        uc.Log.WithError(err).Error("Error retrieving users")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    totalPages := int(math.Ceil(float64(totalCount) / float64(itemsPerPage)))

    // Define custom template functions
    funcMap := template.FuncMap{
        "dec": minus1,
        "inc": plus1,
        "seq": sequence,
    }

    // Preparing data for the template
    data := struct {
        Users       []model.User
        CurrentPage int
        TotalPages  int
    }{
        Users:       users,
        CurrentPage: page,
        TotalPages:  totalPages,
    }

    // Parsing and executing the template
    t, err := template.New("list_users.html").Funcs(funcMap).ParseFiles("view/list_users.html")
    if err != nil {
        uc.Log.WithError(err).Error("Error parsing list users template")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if err := t.Execute(w, data); err != nil {
        uc.Log.WithError(err).Error("Error executing list users template")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}


func minus1(x int) int {
    return x - 1
}

func plus1(x int) int {
    return x + 1
}

func sequence(from, to int) []int {
    seq := make([]int, to - from + 1)
    for i := range seq {
        seq[i] = from + i
    }
    return seq
}

// isValidEmail checks if the email is valid
func isValidEmail(email string) bool {
    // Implement email validation logic
    return strings.Contains(email, "@") // Simplified check; consider a more robust validation
}
