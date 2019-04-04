package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	// Pulling in the pq
	_ "github.com/lib/pq"

	"github.com/dwbelliston/strategy-one-api/app/handler"
	"github.com/dwbelliston/strategy-one-api/config"
	"github.com/gorilla/mux"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	// Connect to db
	db, err := MakeDbConnection(config)

	if err != nil {
		fmt.Println("MakeDbConnection error")
		fmt.Println(err)
		os.Exit(1)
	}

	a.Router = mux.NewRouter()
	a.DB = db
	a.setRouters()
}

// MakeDbConnection used to connect to redshift
func MakeDbConnection(config *config.Config) (*sql.DB, error) {
	url := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v",
		config.DB.Username,
		config.DB.Password,
		config.DB.Endpoint,
		config.DB.Port,
		config.DB.Name)

	var err error
	var db *sql.DB

	if db, err = sql.Open("postgres", url); err != nil {
		return nil, fmt.Errorf("DB connect error : (%v)", err)
	}

	return db, nil
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the shape
	a.Get("/api/shapes", a.GetAllShapes)
	a.Post("/api/shapes", a.CreateShape)
	// a.Get("/shapes/{title}", a.GetShape)
	// a.Put("/shapes/{title}", a.UpdateShape)
	// a.Delete("/shapes/{title}", a.DeleteShape)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

/*
** Shapes Handlers
 */
func (a *App) GetAllShapes(w http.ResponseWriter, r *http.Request) {
	handler.GetAllShapes(a.DB, w, r)
}

func (a *App) CreateShape(w http.ResponseWriter, r *http.Request) {
	handler.CreateShape(a.DB, w, r)
}

// func (a *App) GetShape(w http.ResponseWriter, r *http.Request) {
// 	handler.GetShape(w, r)
// }

// func (a *App) UpdateShape(w http.ResponseWriter, r *http.Request) {
// 	handler.UpdateShape(w, r)
// }

// func (a *App) DeleteShape(w http.ResponseWriter, r *http.Request) {
// 	handler.DeleteShape(w, r)
// }

// Run the app on it's router
func (a *App) Run(host string) {
	fmt.Println("Runing on ", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
