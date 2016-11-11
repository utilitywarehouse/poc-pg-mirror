// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// StaffenqBlb represents a row from 'equinox.staffenq_blb'.
type StaffenqBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllStaffenqBlb(db XODB, callback func(x StaffenqBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.staffenq_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		sb := StaffenqBlb{}

		// scan
		err = q.Scan(&sb.BlbLrn, &sb.BlbData, &sb.BlbText)
		if err != nil {
			return err
		}
		if !callback(sb) {
			return nil
		}
	}

	return nil
}

// StaffenqBlbByBlbLrn retrieves a row from 'equinox.staffenq_blb' as a StaffenqBlb.
//
// Generated from index 'staffenq_blb_pkey'.
func StaffenqBlbByBlbLrn(db XODB, blbLrn int64) (*StaffenqBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.staffenq_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	sb := StaffenqBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&sb.BlbLrn, &sb.BlbData, &sb.BlbText)
	if err != nil {
		return nil, err
	}

	return &sb, nil
}
