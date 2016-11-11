// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// College represents a row from 'equinox.college'.
type College struct {
	CollID       sql.NullInt64  `json:"coll_id"`       // coll_id
	CollName     sql.NullString `json:"coll_name"`     // coll_name
	CollAddr1    sql.NullString `json:"coll_addr1"`    // coll_addr1
	CollAddr2    sql.NullString `json:"coll_addr2"`    // coll_addr2
	CollAddr3    sql.NullString `json:"coll_addr3"`    // coll_addr3
	CollAddr4    sql.NullString `json:"coll_addr4"`    // coll_addr4
	CollAddr5    sql.NullString `json:"coll_addr5"`    // coll_addr5
	CollPcode    sql.NullString `json:"coll_pcode"`    // coll_pcode
	CollTel1     sql.NullString `json:"coll_tel1"`     // coll_tel1
	CollTel2     sql.NullString `json:"coll_tel2"`     // coll_tel2
	CollFax      sql.NullString `json:"coll_fax"`      // coll_fax
	CollCapacity sql.NullInt64  `json:"coll_capacity"` // coll_capacity
	CollTrainer  sql.NullString `json:"coll_trainer"`  // coll_trainer
	CollTime1    pq.NullTime    `json:"coll_time1"`    // coll_time1
	CollTime2    pq.NullTime    `json:"coll_time2"`    // coll_time2
	EquinoxLrn   int64          `json:"equinox_lrn"`   // equinox_lrn
	EquinoxSec   sql.NullInt64  `json:"equinox_sec"`   // equinox_sec
}

func AllCollege(db XODB, callback func(x College) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`coll_id, coll_name, coll_addr1, coll_addr2, coll_addr3, coll_addr4, coll_addr5, coll_pcode, coll_tel1, coll_tel2, coll_fax, coll_capacity, coll_trainer, coll_time1, coll_time2, equinox_lrn, equinox_sec ` +
		`FROM equinox.college `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := College{}

		// scan
		err = q.Scan(&c.CollID, &c.CollName, &c.CollAddr1, &c.CollAddr2, &c.CollAddr3, &c.CollAddr4, &c.CollAddr5, &c.CollPcode, &c.CollTel1, &c.CollTel2, &c.CollFax, &c.CollCapacity, &c.CollTrainer, &c.CollTime1, &c.CollTime2, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CollegeByEquinoxLrn retrieves a row from 'equinox.college' as a College.
//
// Generated from index 'college_pkey'.
func CollegeByEquinoxLrn(db XODB, equinoxLrn int64) (*College, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`coll_id, coll_name, coll_addr1, coll_addr2, coll_addr3, coll_addr4, coll_addr5, coll_pcode, coll_tel1, coll_tel2, coll_fax, coll_capacity, coll_trainer, coll_time1, coll_time2, equinox_lrn, equinox_sec ` +
		`FROM equinox.college ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := College{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.CollID, &c.CollName, &c.CollAddr1, &c.CollAddr2, &c.CollAddr3, &c.CollAddr4, &c.CollAddr5, &c.CollPcode, &c.CollTel1, &c.CollTel2, &c.CollFax, &c.CollCapacity, &c.CollTrainer, &c.CollTime1, &c.CollTime2, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
