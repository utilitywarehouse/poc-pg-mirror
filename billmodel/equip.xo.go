// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Equip represents a row from 'equinox.equip'.
type Equip struct {
	Equipsysnumber   sql.NullFloat64 `json:"equipsysnumber"`   // equipsysnumber
	Equipmanufacture sql.NullString  `json:"equipmanufacture"` // equipmanufacture
	Equipmodel       sql.NullString  `json:"equipmodel"`       // equipmodel
	Equipserialno    sql.NullString  `json:"equipserialno"`    // equipserialno
	Equiptype        sql.NullString  `json:"equiptype"`        // equiptype
	Equipname        sql.NullString  `json:"equipname"`        // equipname
	Equipmanagement  sql.NullString  `json:"equipmanagement"`  // equipmanagement
	Equipinstalldate pq.NullTime     `json:"equipinstalldate"` // equipinstalldate
	Equipengineer    sql.NullString  `json:"equipengineer"`    // equipengineer
	Equipservicecall sql.NullString  `json:"equipservicecall"` // equipservicecall
	Equiporiunsubpri sql.NullFloat64 `json:"equiporiunsubpri"` // equiporiunsubpri
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllEquip(db XODB, callback func(x Equip) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`equipsysnumber, equipmanufacture, equipmodel, equipserialno, equiptype, equipname, equipmanagement, equipinstalldate, equipengineer, equipservicecall, equiporiunsubpri, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.equip `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		e := Equip{}

		// scan
		err = q.Scan(&e.Equipsysnumber, &e.Equipmanufacture, &e.Equipmodel, &e.Equipserialno, &e.Equiptype, &e.Equipname, &e.Equipmanagement, &e.Equipinstalldate, &e.Equipengineer, &e.Equipservicecall, &e.Equiporiunsubpri, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(e) {
			return nil
		}
	}

	return nil
}

// EquipByEquinoxLrn retrieves a row from 'equinox.equip' as a Equip.
//
// Generated from index 'equip_pkey'.
func EquipByEquinoxLrn(db XODB, equinoxLrn int64) (*Equip, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`equipsysnumber, equipmanufacture, equipmodel, equipserialno, equiptype, equipname, equipmanagement, equipinstalldate, equipengineer, equipservicecall, equiporiunsubpri, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.equip ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Equip{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Equipsysnumber, &e.Equipmanufacture, &e.Equipmodel, &e.Equipserialno, &e.Equiptype, &e.Equipname, &e.Equipmanagement, &e.Equipinstalldate, &e.Equipengineer, &e.Equipservicecall, &e.Equiporiunsubpri, &e.EquinoxPrn, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
