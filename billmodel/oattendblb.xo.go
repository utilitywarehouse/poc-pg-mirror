// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// OAttendBlb represents a row from 'equinox.o_attend_blb'.
type OAttendBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllOAttendBlb(db XODB, callback func(x OAttendBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.o_attend_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		oab := OAttendBlb{}

		// scan
		err = q.Scan(&oab.BlbLrn, &oab.BlbData, &oab.BlbText)
		if err != nil {
			return err
		}
		if !callback(oab) {
			return nil
		}
	}

	return nil
}

// OAttendBlbByBlbLrn retrieves a row from 'equinox.o_attend_blb' as a OAttendBlb.
//
// Generated from index 'o_attend_blb_pkey'.
func OAttendBlbByBlbLrn(db XODB, blbLrn int64) (*OAttendBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.o_attend_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	oab := OAttendBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&oab.BlbLrn, &oab.BlbData, &oab.BlbText)
	if err != nil {
		return nil, err
	}

	return &oab, nil
}
