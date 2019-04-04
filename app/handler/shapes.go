package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dwbelliston/strategy-one-api/app/model"
	"github.com/dwbelliston/strategy-one-api/config"
)

var configModule = config.GetConfig()

// GetAllShapes handler
func GetAllShapes(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("api > GetAllShapes")

	shapes := []model.Shape{}

	// Find all sql
	rows, rerr := db.Query(configModule.SQL["GET_ALL_SHAPES"])
	if rerr != nil {
		fmt.Println("GetAllShapes query error")
		fmt.Println(rerr)
		respondError(w, http.StatusInternalServerError, rerr.Error())
		return
	}

	defer rows.Close()

	var shape model.Shape

	for rows.Next() {
		err := rows.Scan(&shape.ID, &shape.Title, &shape.Sides, &shape.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(shape)
		log.Println(shape.ID, shape.Title, shape.Sides, shape.CreatedAt)
		shapes = append(shapes, shape)
	}

	var err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	respondJSON(w, http.StatusOK, shapes)
}

// CreateShape post to create
func CreateShape(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	shape := model.Shape{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&shape); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	fmt.Println(shape)

	// insert into sql
	sqlStatement := `INSERT INTO shapes (title, sides) VALUES ($1, $2)`

	var _, execErr = db.Exec(sqlStatement, shape.Title, shape.Sides)
	if execErr != nil {
		fmt.Println("CreateShape insert error")
		fmt.Println(execErr)
		respondError(w, http.StatusInternalServerError, execErr.Error())
		return
	}

	respondJSON(w, http.StatusCreated, shape)
}

// func GetShape(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	title := vars["title"]
// 	shape := getShapeOr404(db, title, w, r)
// 	if shape == nil {
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, shape)
// }

// func UpdateShape(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	title := vars["title"]
// 	shape := getShapeOr404(db, title, w, r)
// 	if shape == nil {
// 		return
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&shape); err != nil {
// 		respondError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := db.Save(&shape).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusOK, shape)
// }

// func DeleteShape(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	title := vars["title"]
// 	shape := getShapeOr404(db, title, w, r)
// 	if shape == nil {
// 		return
// 	}
// 	if err := db.Delete(&shape).Error; err != nil {
// 		respondError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	respondJSON(w, http.StatusNoContent, nil)
// }

// getShapeOr404 gets a shape instance if exists, or respond the 404 error otherwise
func getShapeOr404(title string, w http.ResponseWriter, r *http.Request) *model.Shape {
	shape := model.Shape{}
	// Get sql

	// if err := SQL for get; err != nil {
	// 	respondError(w, http.StatusNotFound, err.Error())
	// 	return nil
	// }
	return &shape
}
