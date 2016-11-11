// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import "database/sql"

// EquipBlb represents a row from 'equinox.equip_blb'.
type EquipBlb struct {
	BlbLrn  int64          `json:"blb_lrn"`  // blb_lrn
	BlbData []byte         `json:"blb_data"` // blb_data
	BlbText sql.NullString `json:"blb_text"` // blb_text
}

func AllEquipBlb(db XODB, callback func(x EquipBlb) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.equip_blb `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		eb := EquipBlb{}

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

// EquipBlbByBlbLrn retrieves a row from 'equinox.equip_blb' as a EquipBlb.
//
// Generated from index 'equip_blb_pkey'.
func EquipBlbByBlbLrn(db XODB, blbLrn int64) (*EquipBlb, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`blb_lrn, blb_data, blb_text ` +
		`FROM equinox.equip_blb ` +
		`WHERE blb_lrn = $1`

	// run query
	XOLog(sqlstr, blbLrn)
	eb := EquipBlb{}

	err = db.QueryRow(sqlstr, blbLrn).Scan(&eb.BlbLrn, &eb.BlbData, &eb.BlbText)
	if err != nil {
		return nil, err
	}

	return &eb, nil
}
