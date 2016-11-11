// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Cnxcomm represents a row from 'equinox.cnxcomm'.
type Cnxcomm struct {
	Cnxcomdate       pq.NullTime     `json:"cnxcomdate"`       // cnxcomdate
	Cnxcomtime       pq.NullTime     `json:"cnxcomtime"`       // cnxcomtime
	Cnxcomuser       sql.NullString  `json:"cnxcomuser"`       // cnxcomuser
	Cnxcomcomment    sql.NullInt64   `json:"cnxcomcomment"`    // cnxcomcomment
	Cnxcomreceived   pq.NullTime     `json:"cnxcomreceived"`   // cnxcomreceived
	Cnxcomreasoncode sql.NullString  `json:"cnxcomreasoncode"` // cnxcomreasoncode
	Cnxcomactioncode sql.NullString  `json:"cnxcomactioncode"` // cnxcomactioncode
	Cnxcomsupervisor sql.NullInt64   `json:"cnxcomsupervisor"` // cnxcomsupervisor
	Cnxcomlastaction pq.NullTime     `json:"cnxcomlastaction"` // cnxcomlastaction
	Cnxcomresolved   pq.NullTime     `json:"cnxcomresolved"`   // cnxcomresolved
	Cnxcomdead       sql.NullInt64   `json:"cnxcomdead"`       // cnxcomdead
	Cnxcomdropped    sql.NullInt64   `json:"cnxcomdropped"`    // cnxcomdropped
	Cnxcomserial     sql.NullInt64   `json:"cnxcomserial"`     // cnxcomserial
	Cnxcomservice    sql.NullString  `json:"cnxcomservice"`    // cnxcomservice
	Cnxcomapptype    sql.NullString  `json:"cnxcomapptype"`    // cnxcomapptype
	Cnxcomsparec1    sql.NullString  `json:"cnxcomsparec1"`    // cnxcomsparec1
	Cnxcomdueaction  sql.NullFloat64 `json:"cnxcomdueaction"`  // cnxcomdueaction
	Cnxcomspared1    pq.NullTime     `json:"cnxcomspared1"`    // cnxcomspared1
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

func AllCnxcomm(db XODB, callback func(x Cnxcomm) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`cnxcomdate, cnxcomtime, cnxcomuser, cnxcomcomment, cnxcomreceived, cnxcomreasoncode, cnxcomactioncode, cnxcomsupervisor, cnxcomlastaction, cnxcomresolved, cnxcomdead, cnxcomdropped, cnxcomserial, cnxcomservice, cnxcomapptype, cnxcomsparec1, cnxcomdueaction, cnxcomspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxcomm `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		c := Cnxcomm{}

		// scan
		err = q.Scan(&c.Cnxcomdate, &c.Cnxcomtime, &c.Cnxcomuser, &c.Cnxcomcomment, &c.Cnxcomreceived, &c.Cnxcomreasoncode, &c.Cnxcomactioncode, &c.Cnxcomsupervisor, &c.Cnxcomlastaction, &c.Cnxcomresolved, &c.Cnxcomdead, &c.Cnxcomdropped, &c.Cnxcomserial, &c.Cnxcomservice, &c.Cnxcomapptype, &c.Cnxcomsparec1, &c.Cnxcomdueaction, &c.Cnxcomspared1, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(c) {
			return nil
		}
	}

	return nil
}

// CnxcommByEquinoxLrn retrieves a row from 'equinox.cnxcomm' as a Cnxcomm.
//
// Generated from index 'cnxcomm_pkey'.
func CnxcommByEquinoxLrn(db XODB, equinoxLrn int64) (*Cnxcomm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxcomdate, cnxcomtime, cnxcomuser, cnxcomcomment, cnxcomreceived, cnxcomreasoncode, cnxcomactioncode, cnxcomsupervisor, cnxcomlastaction, cnxcomresolved, cnxcomdead, cnxcomdropped, cnxcomserial, cnxcomservice, cnxcomapptype, cnxcomsparec1, cnxcomdueaction, cnxcomspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxcomm ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cnxcomm{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxcomdate, &c.Cnxcomtime, &c.Cnxcomuser, &c.Cnxcomcomment, &c.Cnxcomreceived, &c.Cnxcomreasoncode, &c.Cnxcomactioncode, &c.Cnxcomsupervisor, &c.Cnxcomlastaction, &c.Cnxcomresolved, &c.Cnxcomdead, &c.Cnxcomdropped, &c.Cnxcomserial, &c.Cnxcomservice, &c.Cnxcomapptype, &c.Cnxcomsparec1, &c.Cnxcomdueaction, &c.Cnxcomspared1, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
