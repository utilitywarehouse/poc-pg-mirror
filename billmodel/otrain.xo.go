// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// OTrain represents a row from 'equinox.o_train'.
type OTrain struct {
	OtEventid  sql.NullString `json:"ot_eventid"`  // ot_eventid
	OtExecid   sql.NullString `json:"ot_execid"`   // ot_execid
	OtName     sql.NullString `json:"ot_name"`     // ot_name
	EquinoxPrn sql.NullInt64  `json:"equinox_prn"` // equinox_prn
	EquinoxLrn int64          `json:"equinox_lrn"` // equinox_lrn
	EquinoxSec sql.NullInt64  `json:"equinox_sec"` // equinox_sec
}

func AllOTrain(db XODB, callback func(x OTrain) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`ot_eventid, ot_execid, ot_name, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.o_train `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		ot := OTrain{}

		// scan
		err = q.Scan(&ot.OtEventid, &ot.OtExecid, &ot.OtName, &ot.EquinoxPrn, &ot.EquinoxLrn, &ot.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(ot) {
			return nil
		}
	}

	return nil
}

// OTrainByEquinoxLrn retrieves a row from 'equinox.o_train' as a OTrain.
//
// Generated from index 'o_train_pkey'.
func OTrainByEquinoxLrn(db XODB, equinoxLrn int64) (*OTrain, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`ot_eventid, ot_execid, ot_name, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.o_train ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	ot := OTrain{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&ot.OtEventid, &ot.OtExecid, &ot.OtName, &ot.EquinoxPrn, &ot.EquinoxLrn, &ot.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &ot, nil
}
