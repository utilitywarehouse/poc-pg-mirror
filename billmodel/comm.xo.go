// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Comm represents a row from 'equinox.comms'.
type Comm struct {
	Comuniquesys    sql.NullFloat64 `json:"comuniquesys"`    // comuniquesys
	Comdate         pq.NullTime     `json:"comdate"`         // comdate
	Comtime         pq.NullTime     `json:"comtime"`         // comtime
	Comprobcode     sql.NullString  `json:"comprobcode"`     // comprobcode
	Comdetails      sql.NullInt64   `json:"comdetails"`      // comdetails
	Comenteredby    sql.NullString  `json:"comenteredby"`    // comenteredby
	Comcliservice   sql.NullString  `json:"comcliservice"`   // comcliservice
	Comadditional1  sql.NullInt64   `json:"comadditional1"`  // comadditional1
	Comadditional2  sql.NullInt64   `json:"comadditional2"`  // comadditional2
	Comadditional3  sql.NullInt64   `json:"comadditional3"`  // comadditional3
	Comadditional4  sql.NullInt64   `json:"comadditional4"`  // comadditional4
	Comadditional5  sql.NullInt64   `json:"comadditional5"`  // comadditional5
	Comadditional6  sql.NullInt64   `json:"comadditional6"`  // comadditional6
	Comadditional7  sql.NullInt64   `json:"comadditional7"`  // comadditional7
	Comadditional8  sql.NullInt64   `json:"comadditional8"`  // comadditional8
	Comadditional9  sql.NullInt64   `json:"comadditional9"`  // comadditional9
	Comadditional10 sql.NullInt64   `json:"comadditional10"` // comadditional10
	Compriority     sql.NullString  `json:"compriority"`     // compriority
	Comstatus       sql.NullInt64   `json:"comstatus"`       // comstatus
	Comtaskto       sql.NullString  `json:"comtaskto"`       // comtaskto
	Comcompletedate pq.NullTime     `json:"comcompletedate"` // comcompletedate
	Comcompleteby   sql.NullString  `json:"comcompleteby"`   // comcompleteby
	Comhistoric     sql.NullInt64   `json:"comhistoric"`     // comhistoric
	Comactiondate   pq.NullTime     `json:"comactiondate"`   // comactiondate
	Comactiontime   pq.NullTime     `json:"comactiontime"`   // comactiontime
	Comdip          sql.NullString  `json:"comdip"`          // comdip
	Comdeptid       sql.NullString  `json:"comdeptid"`       // comdeptid
	Comlastchange   sql.NullString  `json:"comlastchange"`   // comlastchange
	Comletterpath   sql.NullInt64   `json:"comletterpath"`   // comletterpath
	Comadddata      sql.NullInt64   `json:"comadddata"`      // comadddata
	Comcliuniquesys sql.NullFloat64 `json:"comcliuniquesys"` // comcliuniquesys
	Comcompletetime pq.NullTime     `json:"comcompletetime"` // comcompletetime
	Comadduniqueid  sql.NullFloat64 `json:"comadduniqueid"`  // comadduniqueid
	Comarchcounter  sql.NullInt64   `json:"comarchcounter"`  // comarchcounter
	EquinoxPrn      sql.NullInt64   `json:"equinox_prn"`     // equinox_prn
	EquinoxLrn      int64           `json:"equinox_lrn"`     // equinox_lrn
	EquinoxSec      sql.NullInt64   `json:"equinox_sec"`     // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Comm exists in the database.
func (c *Comm) Exists() bool {
	return c._exists
}

// Deleted provides information if the Comm has been deleted from the database.
func (c *Comm) Deleted() bool {
	return c._deleted
}

// Insert inserts the Comm to the database.
func (c *Comm) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.comms (` +
		`comuniquesys, comdate, comtime, comprobcode, comdetails, comenteredby, comcliservice, comadditional1, comadditional2, comadditional3, comadditional4, comadditional5, comadditional6, comadditional7, comadditional8, comadditional9, comadditional10, compriority, comstatus, comtaskto, comcompletedate, comcompleteby, comhistoric, comactiondate, comactiontime, comdip, comdeptid, comlastchange, comletterpath, comadddata, comcliuniquesys, comcompletetime, comadduniqueid, comarchcounter, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Comuniquesys, c.Comdate, c.Comtime, c.Comprobcode, c.Comdetails, c.Comenteredby, c.Comcliservice, c.Comadditional1, c.Comadditional2, c.Comadditional3, c.Comadditional4, c.Comadditional5, c.Comadditional6, c.Comadditional7, c.Comadditional8, c.Comadditional9, c.Comadditional10, c.Compriority, c.Comstatus, c.Comtaskto, c.Comcompletedate, c.Comcompleteby, c.Comhistoric, c.Comactiondate, c.Comactiontime, c.Comdip, c.Comdeptid, c.Comlastchange, c.Comletterpath, c.Comadddata, c.Comcliuniquesys, c.Comcompletetime, c.Comadduniqueid, c.Comarchcounter, c.EquinoxPrn, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Comuniquesys, c.Comdate, c.Comtime, c.Comprobcode, c.Comdetails, c.Comenteredby, c.Comcliservice, c.Comadditional1, c.Comadditional2, c.Comadditional3, c.Comadditional4, c.Comadditional5, c.Comadditional6, c.Comadditional7, c.Comadditional8, c.Comadditional9, c.Comadditional10, c.Compriority, c.Comstatus, c.Comtaskto, c.Comcompletedate, c.Comcompleteby, c.Comhistoric, c.Comactiondate, c.Comactiontime, c.Comdip, c.Comdeptid, c.Comlastchange, c.Comletterpath, c.Comadddata, c.Comcliuniquesys, c.Comcompletetime, c.Comadduniqueid, c.Comarchcounter, c.EquinoxPrn, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Comm in the database.
func (c *Comm) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.comms SET (` +
		`comuniquesys, comdate, comtime, comprobcode, comdetails, comenteredby, comcliservice, comadditional1, comadditional2, comadditional3, comadditional4, comadditional5, comadditional6, comadditional7, comadditional8, comadditional9, comadditional10, compriority, comstatus, comtaskto, comcompletedate, comcompleteby, comhistoric, comactiondate, comactiontime, comdip, comdeptid, comlastchange, comletterpath, comadddata, comcliuniquesys, comcompletetime, comadduniqueid, comarchcounter, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36` +
		`) WHERE equinox_lrn = $37`

	// run query
	XOLog(sqlstr, c.Comuniquesys, c.Comdate, c.Comtime, c.Comprobcode, c.Comdetails, c.Comenteredby, c.Comcliservice, c.Comadditional1, c.Comadditional2, c.Comadditional3, c.Comadditional4, c.Comadditional5, c.Comadditional6, c.Comadditional7, c.Comadditional8, c.Comadditional9, c.Comadditional10, c.Compriority, c.Comstatus, c.Comtaskto, c.Comcompletedate, c.Comcompleteby, c.Comhistoric, c.Comactiondate, c.Comactiontime, c.Comdip, c.Comdeptid, c.Comlastchange, c.Comletterpath, c.Comadddata, c.Comcliuniquesys, c.Comcompletetime, c.Comadduniqueid, c.Comarchcounter, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Comuniquesys, c.Comdate, c.Comtime, c.Comprobcode, c.Comdetails, c.Comenteredby, c.Comcliservice, c.Comadditional1, c.Comadditional2, c.Comadditional3, c.Comadditional4, c.Comadditional5, c.Comadditional6, c.Comadditional7, c.Comadditional8, c.Comadditional9, c.Comadditional10, c.Compriority, c.Comstatus, c.Comtaskto, c.Comcompletedate, c.Comcompleteby, c.Comhistoric, c.Comactiondate, c.Comactiontime, c.Comdip, c.Comdeptid, c.Comlastchange, c.Comletterpath, c.Comadddata, c.Comcliuniquesys, c.Comcompletetime, c.Comadduniqueid, c.Comarchcounter, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Comm to the database.
func (c *Comm) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Comm.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Comm) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.comms (` +
		`comuniquesys, comdate, comtime, comprobcode, comdetails, comenteredby, comcliservice, comadditional1, comadditional2, comadditional3, comadditional4, comadditional5, comadditional6, comadditional7, comadditional8, comadditional9, comadditional10, compriority, comstatus, comtaskto, comcompletedate, comcompleteby, comhistoric, comactiondate, comactiontime, comdip, comdeptid, comlastchange, comletterpath, comadddata, comcliuniquesys, comcompletetime, comadduniqueid, comarchcounter, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`comuniquesys, comdate, comtime, comprobcode, comdetails, comenteredby, comcliservice, comadditional1, comadditional2, comadditional3, comadditional4, comadditional5, comadditional6, comadditional7, comadditional8, comadditional9, comadditional10, compriority, comstatus, comtaskto, comcompletedate, comcompleteby, comhistoric, comactiondate, comactiontime, comdip, comdeptid, comlastchange, comletterpath, comadddata, comcliuniquesys, comcompletetime, comadduniqueid, comarchcounter, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.comuniquesys, EXCLUDED.comdate, EXCLUDED.comtime, EXCLUDED.comprobcode, EXCLUDED.comdetails, EXCLUDED.comenteredby, EXCLUDED.comcliservice, EXCLUDED.comadditional1, EXCLUDED.comadditional2, EXCLUDED.comadditional3, EXCLUDED.comadditional4, EXCLUDED.comadditional5, EXCLUDED.comadditional6, EXCLUDED.comadditional7, EXCLUDED.comadditional8, EXCLUDED.comadditional9, EXCLUDED.comadditional10, EXCLUDED.compriority, EXCLUDED.comstatus, EXCLUDED.comtaskto, EXCLUDED.comcompletedate, EXCLUDED.comcompleteby, EXCLUDED.comhistoric, EXCLUDED.comactiondate, EXCLUDED.comactiontime, EXCLUDED.comdip, EXCLUDED.comdeptid, EXCLUDED.comlastchange, EXCLUDED.comletterpath, EXCLUDED.comadddata, EXCLUDED.comcliuniquesys, EXCLUDED.comcompletetime, EXCLUDED.comadduniqueid, EXCLUDED.comarchcounter, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Comuniquesys, c.Comdate, c.Comtime, c.Comprobcode, c.Comdetails, c.Comenteredby, c.Comcliservice, c.Comadditional1, c.Comadditional2, c.Comadditional3, c.Comadditional4, c.Comadditional5, c.Comadditional6, c.Comadditional7, c.Comadditional8, c.Comadditional9, c.Comadditional10, c.Compriority, c.Comstatus, c.Comtaskto, c.Comcompletedate, c.Comcompleteby, c.Comhistoric, c.Comactiondate, c.Comactiontime, c.Comdip, c.Comdeptid, c.Comlastchange, c.Comletterpath, c.Comadddata, c.Comcliuniquesys, c.Comcompletetime, c.Comadduniqueid, c.Comarchcounter, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Comuniquesys, c.Comdate, c.Comtime, c.Comprobcode, c.Comdetails, c.Comenteredby, c.Comcliservice, c.Comadditional1, c.Comadditional2, c.Comadditional3, c.Comadditional4, c.Comadditional5, c.Comadditional6, c.Comadditional7, c.Comadditional8, c.Comadditional9, c.Comadditional10, c.Compriority, c.Comstatus, c.Comtaskto, c.Comcompletedate, c.Comcompleteby, c.Comhistoric, c.Comactiondate, c.Comactiontime, c.Comdip, c.Comdeptid, c.Comlastchange, c.Comletterpath, c.Comadddata, c.Comcliuniquesys, c.Comcompletetime, c.Comadduniqueid, c.Comarchcounter, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Comm from the database.
func (c *Comm) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.comms WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// CommByEquinoxLrn retrieves a row from 'equinox.comms' as a Comm.
//
// Generated from index 'comms_pkey'.
func CommByEquinoxLrn(db XODB, equinoxLrn int64) (*Comm, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`comuniquesys, comdate, comtime, comprobcode, comdetails, comenteredby, comcliservice, comadditional1, comadditional2, comadditional3, comadditional4, comadditional5, comadditional6, comadditional7, comadditional8, comadditional9, comadditional10, compriority, comstatus, comtaskto, comcompletedate, comcompleteby, comhistoric, comactiondate, comactiontime, comdip, comdeptid, comlastchange, comletterpath, comadddata, comcliuniquesys, comcompletetime, comadduniqueid, comarchcounter, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.comms ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Comm{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Comuniquesys, &c.Comdate, &c.Comtime, &c.Comprobcode, &c.Comdetails, &c.Comenteredby, &c.Comcliservice, &c.Comadditional1, &c.Comadditional2, &c.Comadditional3, &c.Comadditional4, &c.Comadditional5, &c.Comadditional6, &c.Comadditional7, &c.Comadditional8, &c.Comadditional9, &c.Comadditional10, &c.Compriority, &c.Comstatus, &c.Comtaskto, &c.Comcompletedate, &c.Comcompleteby, &c.Comhistoric, &c.Comactiondate, &c.Comactiontime, &c.Comdip, &c.Comdeptid, &c.Comlastchange, &c.Comletterpath, &c.Comadddata, &c.Comcliuniquesys, &c.Comcompletetime, &c.Comadduniqueid, &c.Comarchcounter, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
