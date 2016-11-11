// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Valcom represents a row from 'equinox.valcom'.
type Valcom struct {
	Vcomfield       sql.NullString `json:"vcomfield"`       // vcomfield
	Vcomrange       sql.NullString `json:"vcomrange"`       // vcomrange
	Vcomvalue       sql.NullInt64  `json:"vcomvalue"`       // vcomvalue
	Vcomdatatype    sql.NullString `json:"vcomdatatype"`    // vcomdatatype
	Vcomoperation   sql.NullString `json:"vcomoperation"`   // vcomoperation
	Vcomdatefrom    pq.NullTime    `json:"vcomdatefrom"`    // vcomdatefrom
	Vcomdateto      pq.NullTime    `json:"vcomdateto"`      // vcomdateto
	Vcommandatory   sql.NullInt64  `json:"vcommandatory"`   // vcommandatory
	Vcomfillblank   sql.NullInt64  `json:"vcomfillblank"`   // vcomfillblank
	Vcomdateentered pq.NullTime    `json:"vcomdateentered"` // vcomdateentered
	Vcomenteredby   sql.NullString `json:"vcomenteredby"`   // vcomenteredby
	Vcommacrouse    sql.NullInt64  `json:"vcommacrouse"`    // vcommacrouse
	Vcommacrobf     sql.NullInt64  `json:"vcommacrobf"`     // vcommacrobf
	Vcommacroaf     sql.NullInt64  `json:"vcommacroaf"`     // vcommacroaf
	Vcomerror       sql.NullInt64  `json:"vcomerror"`       // vcomerror
	EquinoxPrn      sql.NullInt64  `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64          `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64  `json:"equinox_sec"`     // equinox_sec
}

func AllValcom(db XODB, callback func(x Valcom) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`vcomfield, vcomrange, vcomvalue, vcomdatatype, vcomoperation, vcomdatefrom, vcomdateto, vcommandatory, vcomfillblank, vcomdateentered, vcomenteredby, vcommacrouse, vcommacrobf, vcommacroaf, vcomerror, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.valcom `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		v := Valcom{}

		// scan
		err = q.Scan(&v.Vcomfield, &v.Vcomrange, &v.Vcomvalue, &v.Vcomdatatype, &v.Vcomoperation, &v.Vcomdatefrom, &v.Vcomdateto, &v.Vcommandatory, &v.Vcomfillblank, &v.Vcomdateentered, &v.Vcomenteredby, &v.Vcommacrouse, &v.Vcommacrobf, &v.Vcommacroaf, &v.Vcomerror, &v.EquinoxPrn, &v.EquinoxLrn, &v.EquinoxSec)
		if err != nil {
			return err
		}
		if !callback(v) {
			return nil
		}
	}

	return nil
}

// ValcomByEquinoxLrn retrieves a row from 'equinox.valcom' as a Valcom.
//
// Generated from index 'valcom_pkey'.
func ValcomByEquinoxLrn(db XODB, equinoxLrn int64) (*Valcom, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`vcomfield, vcomrange, vcomvalue, vcomdatatype, vcomoperation, vcomdatefrom, vcomdateto, vcommandatory, vcomfillblank, vcomdateentered, vcomenteredby, vcommacrouse, vcommacrobf, vcommacroaf, vcomerror, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.valcom ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	v := Valcom{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&v.Vcomfield, &v.Vcomrange, &v.Vcomvalue, &v.Vcomdatatype, &v.Vcomoperation, &v.Vcomdatefrom, &v.Vcomdateto, &v.Vcommandatory, &v.Vcomfillblank, &v.Vcomdateentered, &v.Vcomenteredby, &v.Vcommacrouse, &v.Vcommacrobf, &v.Vcommacroaf, &v.Vcomerror, &v.EquinoxPrn, &v.EquinoxLrn, &v.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &v, nil
}
