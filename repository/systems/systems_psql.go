package systems

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sysapi/models"
	"net/http"
	"time"
)

func GetAll(db *sql.DB) models.Tables {
	var t []models.Systems
	stmt := "select * from systems"
	rows, _ := db.Query(stmt)
	defer rows.Close()
	for rows.Next() {
		var sys models.Systems
		_ = rows.Scan(&sys.ID, &sys.Hostname, &sys.Domain, &sys.Device_Type, &sys.Machine_ID, &sys.Device_Status, &sys.Scan_Timestamp)
		t = append(t, sys)
	}
	var tbls models.Tables
	tbls.Systems = t
	return tbls

}

func Update(db *sql.DB, r *http.Request) error {
	// json.NewDecoder(r.Body).Decode(&t)
	//var t models.Tables
	var result models.Tables
	scan_timestamp := time.Now()

	json.NewDecoder(r.Body).Decode(&result)
	fmt.Printf("POST: %v\n", &result)
	// txn, err := db.Begin()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("tx error: %v\n", err)
		return err
	}
	stmt, _ := tx.Prepare("insert into systems (hostname, domain, device_type, machine_id, device_status,scan_timestamp) VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT (machine_id) DO UPDATE SET hostname = $1 RETURNING id  ;")
	defer stmt.Close()

	for _, v := range result.Systems {

		// fmt.Printf("K = %v, V = %v\n", k, v.Hostname)

		_, err := stmt.Exec(v.Hostname, v.Domain, v.Device_Type, v.Machine_ID, v.Device_Status, scan_timestamp)
		if err != nil {
			fmt.Printf("TX Exec error: %v\n", err)
			return err
		}
	}
	tx.Commit()
	return nil
}
