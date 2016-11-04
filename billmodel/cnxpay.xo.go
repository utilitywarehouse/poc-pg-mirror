// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Cnxpay represents a row from 'equinox.cnxpay'.
type Cnxpay struct {
	Cnxpaydate       pq.NullTime     `json:"cnxpaydate"`       // cnxpaydate
	Cnxpaycctype     sql.NullInt64   `json:"cnxpaycctype"`     // cnxpaycctype
	Cnxpayccno       sql.NullString  `json:"cnxpayccno"`       // cnxpayccno
	Cnxpayccexp      sql.NullString  `json:"cnxpayccexp"`      // cnxpayccexp
	Cnxpayccstart    sql.NullString  `json:"cnxpayccstart"`    // cnxpayccstart
	Cnxpayccissue    sql.NullString  `json:"cnxpayccissue"`    // cnxpayccissue
	Cnxpayamount     sql.NullFloat64 `json:"cnxpayamount"`     // cnxpayamount
	Cnxpaychequeno   sql.NullString  `json:"cnxpaychequeno"`   // cnxpaychequeno
	Cnxpaytransaccod sql.NullString  `json:"cnxpaytransaccod"` // cnxpaytransaccod
	Cnxpayreffield   sql.NullString  `json:"cnxpayreffield"`   // cnxpayreffield
	Cnxpayexportflag sql.NullInt64   `json:"cnxpayexportflag"` // cnxpayexportflag
	Cnxpayauthorised sql.NullString  `json:"cnxpayauthorised"` // cnxpayauthorised
	Cnxpayauthamount sql.NullFloat64 `json:"cnxpayauthamount"` // cnxpayauthamount
	Cnxpayauthdate   pq.NullTime     `json:"cnxpayauthdate"`   // cnxpayauthdate
	Cnxpayauthtext   sql.NullString  `json:"cnxpayauthtext"`   // cnxpayauthtext
	Cnxpayenteredby  sql.NullString  `json:"cnxpayenteredby"`  // cnxpayenteredby
	Cnxpaycommentcod sql.NullString  `json:"cnxpaycommentcod"` // cnxpaycommentcod
	Cnxpayadvance    sql.NullFloat64 `json:"cnxpayadvance"`    // cnxpayadvance
	Cnxpayconncharge sql.NullFloat64 `json:"cnxpayconncharge"` // cnxpayconncharge
	Cnxpaypprice     sql.NullFloat64 `json:"cnxpaypprice"`     // cnxpaypprice
	Cnxpayprepay2bil pq.NullTime     `json:"cnxpayprepay2bil"` // cnxpayprepay2bil
	Cnxpayweborder   sql.NullInt64   `json:"cnxpayweborder"`   // cnxpayweborder
	Cnxmppuniquesys  sql.NullInt64   `json:"cnxmppuniquesys"`  // cnxmppuniquesys
	Cnxprosaccountno sql.NullInt64   `json:"cnxprosaccountno"` // cnxprosaccountno
	Cnxpaycv2        sql.NullString  `json:"cnxpaycv2"`        // cnxpaycv2
	Cnxpaypostcode   sql.NullString  `json:"cnxpaypostcode"`   // cnxpaypostcode
	Cnxpaycbcappfee  sql.NullInt64   `json:"cnxpaycbcappfee"`  // cnxpaycbcappfee
	Cnxpayspared1    pq.NullTime     `json:"cnxpayspared1"`    // cnxpayspared1
	Cnxpaybanked     pq.NullTime     `json:"cnxpaybanked"`     // cnxpaybanked
	Cnxpaypmm        pq.NullTime     `json:"cnxpaypmm"`        // cnxpaypmm
	Cnxpayrefund     pq.NullTime     `json:"cnxpayrefund"`     // cnxpayrefund
	Cnxpaytime       pq.NullTime     `json:"cnxpaytime"`       // cnxpaytime
	Cnxpaysparen1    sql.NullFloat64 `json:"cnxpaysparen1"`    // cnxpaysparen1
	Cnxpayvaliddate  pq.NullTime     `json:"cnxpayvaliddate"`  // cnxpayvaliddate
	Cnxpayvalidby    sql.NullString  `json:"cnxpayvalidby"`    // cnxpayvalidby
	Cnxpaymarked     sql.NullInt64   `json:"cnxpaymarked"`     // cnxpaymarked
	Cnxpaylastchange sql.NullString  `json:"cnxpaylastchange"` // cnxpaylastchange
	Cnxpaypsp        sql.NullString  `json:"cnxpaypsp"`        // cnxpaypsp
	Cnxpayref        sql.NullString  `json:"cnxpayref"`        // cnxpayref
	Cnxpayshopref    sql.NullString  `json:"cnxpayshopref"`    // cnxpayshopref
	Cnxpaycontract   sql.NullInt64   `json:"cnxpaycontract"`   // cnxpaycontract
	Cnxpayrecuniqsys sql.NullString  `json:"cnxpayrecuniqsys"` // cnxpayrecuniqsys
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Cnxpay exists in the database.
func (c *Cnxpay) Exists() bool {
	return c._exists
}

// Deleted provides information if the Cnxpay has been deleted from the database.
func (c *Cnxpay) Deleted() bool {
	return c._deleted
}

// Insert inserts the Cnxpay to the database.
func (c *Cnxpay) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cnxpay (` +
		`cnxpaydate, cnxpaycctype, cnxpayccno, cnxpayccexp, cnxpayccstart, cnxpayccissue, cnxpayamount, cnxpaychequeno, cnxpaytransaccod, cnxpayreffield, cnxpayexportflag, cnxpayauthorised, cnxpayauthamount, cnxpayauthdate, cnxpayauthtext, cnxpayenteredby, cnxpaycommentcod, cnxpayadvance, cnxpayconncharge, cnxpaypprice, cnxpayprepay2bil, cnxpayweborder, cnxmppuniquesys, cnxprosaccountno, cnxpaycv2, cnxpaypostcode, cnxpaycbcappfee, cnxpayspared1, cnxpaybanked, cnxpaypmm, cnxpayrefund, cnxpaytime, cnxpaysparen1, cnxpayvaliddate, cnxpayvalidby, cnxpaymarked, cnxpaylastchange, cnxpaypsp, cnxpayref, cnxpayshopref, cnxpaycontract, cnxpayrecuniqsys, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Cnxpaydate, c.Cnxpaycctype, c.Cnxpayccno, c.Cnxpayccexp, c.Cnxpayccstart, c.Cnxpayccissue, c.Cnxpayamount, c.Cnxpaychequeno, c.Cnxpaytransaccod, c.Cnxpayreffield, c.Cnxpayexportflag, c.Cnxpayauthorised, c.Cnxpayauthamount, c.Cnxpayauthdate, c.Cnxpayauthtext, c.Cnxpayenteredby, c.Cnxpaycommentcod, c.Cnxpayadvance, c.Cnxpayconncharge, c.Cnxpaypprice, c.Cnxpayprepay2bil, c.Cnxpayweborder, c.Cnxmppuniquesys, c.Cnxprosaccountno, c.Cnxpaycv2, c.Cnxpaypostcode, c.Cnxpaycbcappfee, c.Cnxpayspared1, c.Cnxpaybanked, c.Cnxpaypmm, c.Cnxpayrefund, c.Cnxpaytime, c.Cnxpaysparen1, c.Cnxpayvaliddate, c.Cnxpayvalidby, c.Cnxpaymarked, c.Cnxpaylastchange, c.Cnxpaypsp, c.Cnxpayref, c.Cnxpayshopref, c.Cnxpaycontract, c.Cnxpayrecuniqsys, c.EquinoxPrn, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Cnxpaydate, c.Cnxpaycctype, c.Cnxpayccno, c.Cnxpayccexp, c.Cnxpayccstart, c.Cnxpayccissue, c.Cnxpayamount, c.Cnxpaychequeno, c.Cnxpaytransaccod, c.Cnxpayreffield, c.Cnxpayexportflag, c.Cnxpayauthorised, c.Cnxpayauthamount, c.Cnxpayauthdate, c.Cnxpayauthtext, c.Cnxpayenteredby, c.Cnxpaycommentcod, c.Cnxpayadvance, c.Cnxpayconncharge, c.Cnxpaypprice, c.Cnxpayprepay2bil, c.Cnxpayweborder, c.Cnxmppuniquesys, c.Cnxprosaccountno, c.Cnxpaycv2, c.Cnxpaypostcode, c.Cnxpaycbcappfee, c.Cnxpayspared1, c.Cnxpaybanked, c.Cnxpaypmm, c.Cnxpayrefund, c.Cnxpaytime, c.Cnxpaysparen1, c.Cnxpayvaliddate, c.Cnxpayvalidby, c.Cnxpaymarked, c.Cnxpaylastchange, c.Cnxpaypsp, c.Cnxpayref, c.Cnxpayshopref, c.Cnxpaycontract, c.Cnxpayrecuniqsys, c.EquinoxPrn, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Cnxpay in the database.
func (c *Cnxpay) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.cnxpay SET (` +
		`cnxpaydate, cnxpaycctype, cnxpayccno, cnxpayccexp, cnxpayccstart, cnxpayccissue, cnxpayamount, cnxpaychequeno, cnxpaytransaccod, cnxpayreffield, cnxpayexportflag, cnxpayauthorised, cnxpayauthamount, cnxpayauthdate, cnxpayauthtext, cnxpayenteredby, cnxpaycommentcod, cnxpayadvance, cnxpayconncharge, cnxpaypprice, cnxpayprepay2bil, cnxpayweborder, cnxmppuniquesys, cnxprosaccountno, cnxpaycv2, cnxpaypostcode, cnxpaycbcappfee, cnxpayspared1, cnxpaybanked, cnxpaypmm, cnxpayrefund, cnxpaytime, cnxpaysparen1, cnxpayvaliddate, cnxpayvalidby, cnxpaymarked, cnxpaylastchange, cnxpaypsp, cnxpayref, cnxpayshopref, cnxpaycontract, cnxpayrecuniqsys, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44` +
		`) WHERE equinox_lrn = $45`

	// run query
	XOLog(sqlstr, c.Cnxpaydate, c.Cnxpaycctype, c.Cnxpayccno, c.Cnxpayccexp, c.Cnxpayccstart, c.Cnxpayccissue, c.Cnxpayamount, c.Cnxpaychequeno, c.Cnxpaytransaccod, c.Cnxpayreffield, c.Cnxpayexportflag, c.Cnxpayauthorised, c.Cnxpayauthamount, c.Cnxpayauthdate, c.Cnxpayauthtext, c.Cnxpayenteredby, c.Cnxpaycommentcod, c.Cnxpayadvance, c.Cnxpayconncharge, c.Cnxpaypprice, c.Cnxpayprepay2bil, c.Cnxpayweborder, c.Cnxmppuniquesys, c.Cnxprosaccountno, c.Cnxpaycv2, c.Cnxpaypostcode, c.Cnxpaycbcappfee, c.Cnxpayspared1, c.Cnxpaybanked, c.Cnxpaypmm, c.Cnxpayrefund, c.Cnxpaytime, c.Cnxpaysparen1, c.Cnxpayvaliddate, c.Cnxpayvalidby, c.Cnxpaymarked, c.Cnxpaylastchange, c.Cnxpaypsp, c.Cnxpayref, c.Cnxpayshopref, c.Cnxpaycontract, c.Cnxpayrecuniqsys, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Cnxpaydate, c.Cnxpaycctype, c.Cnxpayccno, c.Cnxpayccexp, c.Cnxpayccstart, c.Cnxpayccissue, c.Cnxpayamount, c.Cnxpaychequeno, c.Cnxpaytransaccod, c.Cnxpayreffield, c.Cnxpayexportflag, c.Cnxpayauthorised, c.Cnxpayauthamount, c.Cnxpayauthdate, c.Cnxpayauthtext, c.Cnxpayenteredby, c.Cnxpaycommentcod, c.Cnxpayadvance, c.Cnxpayconncharge, c.Cnxpaypprice, c.Cnxpayprepay2bil, c.Cnxpayweborder, c.Cnxmppuniquesys, c.Cnxprosaccountno, c.Cnxpaycv2, c.Cnxpaypostcode, c.Cnxpaycbcappfee, c.Cnxpayspared1, c.Cnxpaybanked, c.Cnxpaypmm, c.Cnxpayrefund, c.Cnxpaytime, c.Cnxpaysparen1, c.Cnxpayvaliddate, c.Cnxpayvalidby, c.Cnxpaymarked, c.Cnxpaylastchange, c.Cnxpaypsp, c.Cnxpayref, c.Cnxpayshopref, c.Cnxpaycontract, c.Cnxpayrecuniqsys, c.EquinoxPrn, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Cnxpay to the database.
func (c *Cnxpay) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Cnxpay.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Cnxpay) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.cnxpay (` +
		`cnxpaydate, cnxpaycctype, cnxpayccno, cnxpayccexp, cnxpayccstart, cnxpayccissue, cnxpayamount, cnxpaychequeno, cnxpaytransaccod, cnxpayreffield, cnxpayexportflag, cnxpayauthorised, cnxpayauthamount, cnxpayauthdate, cnxpayauthtext, cnxpayenteredby, cnxpaycommentcod, cnxpayadvance, cnxpayconncharge, cnxpaypprice, cnxpayprepay2bil, cnxpayweborder, cnxmppuniquesys, cnxprosaccountno, cnxpaycv2, cnxpaypostcode, cnxpaycbcappfee, cnxpayspared1, cnxpaybanked, cnxpaypmm, cnxpayrefund, cnxpaytime, cnxpaysparen1, cnxpayvaliddate, cnxpayvalidby, cnxpaymarked, cnxpaylastchange, cnxpaypsp, cnxpayref, cnxpayshopref, cnxpaycontract, cnxpayrecuniqsys, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`cnxpaydate, cnxpaycctype, cnxpayccno, cnxpayccexp, cnxpayccstart, cnxpayccissue, cnxpayamount, cnxpaychequeno, cnxpaytransaccod, cnxpayreffield, cnxpayexportflag, cnxpayauthorised, cnxpayauthamount, cnxpayauthdate, cnxpayauthtext, cnxpayenteredby, cnxpaycommentcod, cnxpayadvance, cnxpayconncharge, cnxpaypprice, cnxpayprepay2bil, cnxpayweborder, cnxmppuniquesys, cnxprosaccountno, cnxpaycv2, cnxpaypostcode, cnxpaycbcappfee, cnxpayspared1, cnxpaybanked, cnxpaypmm, cnxpayrefund, cnxpaytime, cnxpaysparen1, cnxpayvaliddate, cnxpayvalidby, cnxpaymarked, cnxpaylastchange, cnxpaypsp, cnxpayref, cnxpayshopref, cnxpaycontract, cnxpayrecuniqsys, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.cnxpaydate, EXCLUDED.cnxpaycctype, EXCLUDED.cnxpayccno, EXCLUDED.cnxpayccexp, EXCLUDED.cnxpayccstart, EXCLUDED.cnxpayccissue, EXCLUDED.cnxpayamount, EXCLUDED.cnxpaychequeno, EXCLUDED.cnxpaytransaccod, EXCLUDED.cnxpayreffield, EXCLUDED.cnxpayexportflag, EXCLUDED.cnxpayauthorised, EXCLUDED.cnxpayauthamount, EXCLUDED.cnxpayauthdate, EXCLUDED.cnxpayauthtext, EXCLUDED.cnxpayenteredby, EXCLUDED.cnxpaycommentcod, EXCLUDED.cnxpayadvance, EXCLUDED.cnxpayconncharge, EXCLUDED.cnxpaypprice, EXCLUDED.cnxpayprepay2bil, EXCLUDED.cnxpayweborder, EXCLUDED.cnxmppuniquesys, EXCLUDED.cnxprosaccountno, EXCLUDED.cnxpaycv2, EXCLUDED.cnxpaypostcode, EXCLUDED.cnxpaycbcappfee, EXCLUDED.cnxpayspared1, EXCLUDED.cnxpaybanked, EXCLUDED.cnxpaypmm, EXCLUDED.cnxpayrefund, EXCLUDED.cnxpaytime, EXCLUDED.cnxpaysparen1, EXCLUDED.cnxpayvaliddate, EXCLUDED.cnxpayvalidby, EXCLUDED.cnxpaymarked, EXCLUDED.cnxpaylastchange, EXCLUDED.cnxpaypsp, EXCLUDED.cnxpayref, EXCLUDED.cnxpayshopref, EXCLUDED.cnxpaycontract, EXCLUDED.cnxpayrecuniqsys, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Cnxpaydate, c.Cnxpaycctype, c.Cnxpayccno, c.Cnxpayccexp, c.Cnxpayccstart, c.Cnxpayccissue, c.Cnxpayamount, c.Cnxpaychequeno, c.Cnxpaytransaccod, c.Cnxpayreffield, c.Cnxpayexportflag, c.Cnxpayauthorised, c.Cnxpayauthamount, c.Cnxpayauthdate, c.Cnxpayauthtext, c.Cnxpayenteredby, c.Cnxpaycommentcod, c.Cnxpayadvance, c.Cnxpayconncharge, c.Cnxpaypprice, c.Cnxpayprepay2bil, c.Cnxpayweborder, c.Cnxmppuniquesys, c.Cnxprosaccountno, c.Cnxpaycv2, c.Cnxpaypostcode, c.Cnxpaycbcappfee, c.Cnxpayspared1, c.Cnxpaybanked, c.Cnxpaypmm, c.Cnxpayrefund, c.Cnxpaytime, c.Cnxpaysparen1, c.Cnxpayvaliddate, c.Cnxpayvalidby, c.Cnxpaymarked, c.Cnxpaylastchange, c.Cnxpaypsp, c.Cnxpayref, c.Cnxpayshopref, c.Cnxpaycontract, c.Cnxpayrecuniqsys, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Cnxpaydate, c.Cnxpaycctype, c.Cnxpayccno, c.Cnxpayccexp, c.Cnxpayccstart, c.Cnxpayccissue, c.Cnxpayamount, c.Cnxpaychequeno, c.Cnxpaytransaccod, c.Cnxpayreffield, c.Cnxpayexportflag, c.Cnxpayauthorised, c.Cnxpayauthamount, c.Cnxpayauthdate, c.Cnxpayauthtext, c.Cnxpayenteredby, c.Cnxpaycommentcod, c.Cnxpayadvance, c.Cnxpayconncharge, c.Cnxpaypprice, c.Cnxpayprepay2bil, c.Cnxpayweborder, c.Cnxmppuniquesys, c.Cnxprosaccountno, c.Cnxpaycv2, c.Cnxpaypostcode, c.Cnxpaycbcappfee, c.Cnxpayspared1, c.Cnxpaybanked, c.Cnxpaypmm, c.Cnxpayrefund, c.Cnxpaytime, c.Cnxpaysparen1, c.Cnxpayvaliddate, c.Cnxpayvalidby, c.Cnxpaymarked, c.Cnxpaylastchange, c.Cnxpaypsp, c.Cnxpayref, c.Cnxpayshopref, c.Cnxpaycontract, c.Cnxpayrecuniqsys, c.EquinoxPrn, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Cnxpay from the database.
func (c *Cnxpay) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.cnxpay WHERE equinox_lrn = $1`

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

// CnxpayByEquinoxLrn retrieves a row from 'equinox.cnxpay' as a Cnxpay.
//
// Generated from index 'cnxpay_pkey'.
func CnxpayByEquinoxLrn(db XODB, equinoxLrn int64) (*Cnxpay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxpaydate, cnxpaycctype, cnxpayccno, cnxpayccexp, cnxpayccstart, cnxpayccissue, cnxpayamount, cnxpaychequeno, cnxpaytransaccod, cnxpayreffield, cnxpayexportflag, cnxpayauthorised, cnxpayauthamount, cnxpayauthdate, cnxpayauthtext, cnxpayenteredby, cnxpaycommentcod, cnxpayadvance, cnxpayconncharge, cnxpaypprice, cnxpayprepay2bil, cnxpayweborder, cnxmppuniquesys, cnxprosaccountno, cnxpaycv2, cnxpaypostcode, cnxpaycbcappfee, cnxpayspared1, cnxpaybanked, cnxpaypmm, cnxpayrefund, cnxpaytime, cnxpaysparen1, cnxpayvaliddate, cnxpayvalidby, cnxpaymarked, cnxpaylastchange, cnxpaypsp, cnxpayref, cnxpayshopref, cnxpaycontract, cnxpayrecuniqsys, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.cnxpay ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Cnxpay{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxpaydate, &c.Cnxpaycctype, &c.Cnxpayccno, &c.Cnxpayccexp, &c.Cnxpayccstart, &c.Cnxpayccissue, &c.Cnxpayamount, &c.Cnxpaychequeno, &c.Cnxpaytransaccod, &c.Cnxpayreffield, &c.Cnxpayexportflag, &c.Cnxpayauthorised, &c.Cnxpayauthamount, &c.Cnxpayauthdate, &c.Cnxpayauthtext, &c.Cnxpayenteredby, &c.Cnxpaycommentcod, &c.Cnxpayadvance, &c.Cnxpayconncharge, &c.Cnxpaypprice, &c.Cnxpayprepay2bil, &c.Cnxpayweborder, &c.Cnxmppuniquesys, &c.Cnxprosaccountno, &c.Cnxpaycv2, &c.Cnxpaypostcode, &c.Cnxpaycbcappfee, &c.Cnxpayspared1, &c.Cnxpaybanked, &c.Cnxpaypmm, &c.Cnxpayrefund, &c.Cnxpaytime, &c.Cnxpaysparen1, &c.Cnxpayvaliddate, &c.Cnxpayvalidby, &c.Cnxpaymarked, &c.Cnxpaylastchange, &c.Cnxpaypsp, &c.Cnxpayref, &c.Cnxpayshopref, &c.Cnxpaycontract, &c.Cnxpayrecuniqsys, &c.EquinoxPrn, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
