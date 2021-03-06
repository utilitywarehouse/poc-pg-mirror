// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// EdsmainBlb represents a row from 'equinox.edsmain_blb'.
type EdsmainBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllEdsmainBlb(db XODB, callback func(x EdsmainBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.edsmain_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		eb := EdsmainBlb{}

		// scan
		err = q.Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
		if err != nil {
			return err
		}
		if !callback(eb) {
			return nil
		}
	}

	return nil
}

// EdsmainBlbByBlbLrn retrieves a row from 'equinox.edsmain_blb' as a EdsmainBlb.
//
// Generated from index 'edsmain_blb_pkey'.
func EdsmainBlbByBlbLrn(db XODB, blbLrn int64) (*EdsmainBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.edsmain_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	eb := EdsmainBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
	if err != nil {
		return nil, err
	}

	return &eb, nil
}
