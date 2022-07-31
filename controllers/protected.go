package controllers

import (
	"database/sql"
	"fmt"
	"sysapi/models"
	"sysapi/repository/systems"
	"sysapi/utils"
	"net/http"
)

func (c Controller) GetSystems(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		all := Systems.GetAll(db)
		utils.ResponseJSON(w, all)
		fmt.Printf("success. \n")

	}
}
func (c Controller) UpdateSystems(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// json.NewDecoder(r.Body).Decode(&user)
		err := Systems.Update(db, r)
		if err != nil {
			fmt.Printf("System.Update error: %v\n", err)
			var e models.Error
			e.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, e)
		}

		fmt.Printf("success. \n")

	}
}
