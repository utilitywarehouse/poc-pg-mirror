// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Billrun represents a row from 'equinox.billruns'.
type Billrun struct {
	Brnumber         sql.NullFloat64 `json:"brnumber"`         // brnumber
	Brperiod         sql.NullString  `json:"brperiod"`         // brperiod
	Brtype           sql.NullInt64   `json:"brtype"`           // brtype
	Brtypetxt        sql.NullString  `json:"brtypetxt"`        // brtypetxt
	Brrundate        pq.NullTime     `json:"brrundate"`        // brrundate
	Brrunstarttime   pq.NullTime     `json:"brrunstarttime"`   // brrunstarttime
	Brrunenddate     pq.NullTime     `json:"brrunenddate"`     // brrunenddate
	Brrunendtime     pq.NullTime     `json:"brrunendtime"`     // brrunendtime
	Brrunby          sql.NullString  `json:"brrunby"`          // brrunby
	Brcomplete       sql.NullInt64   `json:"brcomplete"`       // brcomplete
	Brcloseddate     pq.NullTime     `json:"brcloseddate"`     // brcloseddate
	Brclosedby       sql.NullString  `json:"brclosedby"`       // brclosedby
	Brlastbilldate   pq.NullTime     `json:"brlastbilldate"`   // brlastbilldate
	Brmonthstobill   sql.NullInt64   `json:"brmonthstobill"`   // brmonthstobill
	Brlowestbillno   sql.NullFloat64 `json:"brlowestbillno"`   // brlowestbillno
	Brhighestbillno  sql.NullFloat64 `json:"brhighestbillno"`  // brhighestbillno
	Brnoofbills      sql.NullFloat64 `json:"brnoofbills"`      // brnoofbills
	Brprevos         sql.NullFloat64 `json:"brprevos"`         // brprevos
	Brwholesale      sql.NullFloat64 `json:"brwholesale"`      // brwholesale
	Brretail         sql.NullFloat64 `json:"brretail"`         // brretail
	Brretail2        sql.NullFloat64 `json:"brretail2"`        // brretail2
	Broneoff         sql.NullFloat64 `json:"broneoff"`         // broneoff
	Brgreendeal      sql.NullFloat64 `json:"brgreendeal"`      // brgreendeal
	Brmonthly        sql.NullFloat64 `json:"brmonthly"`        // brmonthly
	Brmonthlyw       sql.NullFloat64 `json:"brmonthlyw"`       // brmonthlyw
	Brcreditall      sql.NullFloat64 `json:"brcreditall"`      // brcreditall
	Brcreditunits    sql.NullFloat64 `json:"brcreditunits"`    // brcreditunits
	Brdebitall       sql.NullFloat64 `json:"brdebitall"`       // brdebitall
	Brvattp          sql.NullFloat64 `json:"brvattp"`          // brvattp
	Brvatgp          sql.NullFloat64 `json:"brvatgp"`          // brvatgp
	Brvatep          sql.NullFloat64 `json:"brvatep"`          // brvatep
	Brnettp          sql.NullFloat64 `json:"brnettp"`          // brnettp
	Brnetgp          sql.NullFloat64 `json:"brnetgp"`          // brnetgp
	Brnetep          sql.NullFloat64 `json:"brnetep"`          // brnetep
	Brdiscount       sql.NullFloat64 `json:"brdiscount"`       // brdiscount
	Brsurcharge      sql.NullFloat64 `json:"brsurcharge"`      // brsurcharge
	Brclubfee        sql.NullFloat64 `json:"brclubfee"`        // brclubfee
	Brclubdisc       sql.NullFloat64 `json:"brclubdisc"`       // brclubdisc
	Brtotal          sql.NullFloat64 `json:"brtotal"`          // brtotal
	Brcgadiscount    sql.NullFloat64 `json:"brcgadiscount"`    // brcgadiscount
	Brchgecallcount  sql.NullFloat64 `json:"brchgecallcount"`  // brchgecallcount
	Brchgecallsecs   sql.NullFloat64 `json:"brchgecallsecs"`   // brchgecallsecs
	Brnoofcreditnote sql.NullFloat64 `json:"brnoofcreditnote"` // brnoofcreditnote
	Brcardcashbak    sql.NullFloat64 `json:"brcardcashbak"`    // brcardcashbak
	Braffcashbak     sql.NullFloat64 `json:"braffcashbak"`     // braffcashbak
	Brbydd           sql.NullFloat64 `json:"brbydd"`           // brbydd
	Brbydd1          sql.NullFloat64 `json:"brbydd1"`          // brbydd1
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Billrun exists in the database.
func (b *Billrun) Exists() bool {
	return b._exists
}

// Deleted provides information if the Billrun has been deleted from the database.
func (b *Billrun) Deleted() bool {
	return b._deleted
}

// Insert inserts the Billrun to the database.
func (b *Billrun) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.billruns (` +
		`brnumber, brperiod, brtype, brtypetxt, brrundate, brrunstarttime, brrunenddate, brrunendtime, brrunby, brcomplete, brcloseddate, brclosedby, brlastbilldate, brmonthstobill, brlowestbillno, brhighestbillno, brnoofbills, brprevos, brwholesale, brretail, brretail2, broneoff, brgreendeal, brmonthly, brmonthlyw, brcreditall, brcreditunits, brdebitall, brvattp, brvatgp, brvatep, brnettp, brnetgp, brnetep, brdiscount, brsurcharge, brclubfee, brclubdisc, brtotal, brcgadiscount, brchgecallcount, brchgecallsecs, brnoofcreditnote, brcardcashbak, braffcashbak, brbydd, brbydd1, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, b.Brnumber, b.Brperiod, b.Brtype, b.Brtypetxt, b.Brrundate, b.Brrunstarttime, b.Brrunenddate, b.Brrunendtime, b.Brrunby, b.Brcomplete, b.Brcloseddate, b.Brclosedby, b.Brlastbilldate, b.Brmonthstobill, b.Brlowestbillno, b.Brhighestbillno, b.Brnoofbills, b.Brprevos, b.Brwholesale, b.Brretail, b.Brretail2, b.Broneoff, b.Brgreendeal, b.Brmonthly, b.Brmonthlyw, b.Brcreditall, b.Brcreditunits, b.Brdebitall, b.Brvattp, b.Brvatgp, b.Brvatep, b.Brnettp, b.Brnetgp, b.Brnetep, b.Brdiscount, b.Brsurcharge, b.Brclubfee, b.Brclubdisc, b.Brtotal, b.Brcgadiscount, b.Brchgecallcount, b.Brchgecallsecs, b.Brnoofcreditnote, b.Brcardcashbak, b.Braffcashbak, b.Brbydd, b.Brbydd1, b.EquinoxSec)
	err = db.QueryRow(sqlstr, b.Brnumber, b.Brperiod, b.Brtype, b.Brtypetxt, b.Brrundate, b.Brrunstarttime, b.Brrunenddate, b.Brrunendtime, b.Brrunby, b.Brcomplete, b.Brcloseddate, b.Brclosedby, b.Brlastbilldate, b.Brmonthstobill, b.Brlowestbillno, b.Brhighestbillno, b.Brnoofbills, b.Brprevos, b.Brwholesale, b.Brretail, b.Brretail2, b.Broneoff, b.Brgreendeal, b.Brmonthly, b.Brmonthlyw, b.Brcreditall, b.Brcreditunits, b.Brdebitall, b.Brvattp, b.Brvatgp, b.Brvatep, b.Brnettp, b.Brnetgp, b.Brnetep, b.Brdiscount, b.Brsurcharge, b.Brclubfee, b.Brclubdisc, b.Brtotal, b.Brcgadiscount, b.Brchgecallcount, b.Brchgecallsecs, b.Brnoofcreditnote, b.Brcardcashbak, b.Braffcashbak, b.Brbydd, b.Brbydd1, b.EquinoxSec).Scan(&b.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Update updates the Billrun in the database.
func (b *Billrun) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if b._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.billruns SET (` +
		`brnumber, brperiod, brtype, brtypetxt, brrundate, brrunstarttime, brrunenddate, brrunendtime, brrunby, brcomplete, brcloseddate, brclosedby, brlastbilldate, brmonthstobill, brlowestbillno, brhighestbillno, brnoofbills, brprevos, brwholesale, brretail, brretail2, broneoff, brgreendeal, brmonthly, brmonthlyw, brcreditall, brcreditunits, brdebitall, brvattp, brvatgp, brvatep, brnettp, brnetgp, brnetep, brdiscount, brsurcharge, brclubfee, brclubdisc, brtotal, brcgadiscount, brchgecallcount, brchgecallsecs, brnoofcreditnote, brcardcashbak, braffcashbak, brbydd, brbydd1, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48` +
		`) WHERE equinox_lrn = $49`

	// run query
	XOLog(sqlstr, b.Brnumber, b.Brperiod, b.Brtype, b.Brtypetxt, b.Brrundate, b.Brrunstarttime, b.Brrunenddate, b.Brrunendtime, b.Brrunby, b.Brcomplete, b.Brcloseddate, b.Brclosedby, b.Brlastbilldate, b.Brmonthstobill, b.Brlowestbillno, b.Brhighestbillno, b.Brnoofbills, b.Brprevos, b.Brwholesale, b.Brretail, b.Brretail2, b.Broneoff, b.Brgreendeal, b.Brmonthly, b.Brmonthlyw, b.Brcreditall, b.Brcreditunits, b.Brdebitall, b.Brvattp, b.Brvatgp, b.Brvatep, b.Brnettp, b.Brnetgp, b.Brnetep, b.Brdiscount, b.Brsurcharge, b.Brclubfee, b.Brclubdisc, b.Brtotal, b.Brcgadiscount, b.Brchgecallcount, b.Brchgecallsecs, b.Brnoofcreditnote, b.Brcardcashbak, b.Braffcashbak, b.Brbydd, b.Brbydd1, b.EquinoxSec, b.EquinoxLrn)
	_, err = db.Exec(sqlstr, b.Brnumber, b.Brperiod, b.Brtype, b.Brtypetxt, b.Brrundate, b.Brrunstarttime, b.Brrunenddate, b.Brrunendtime, b.Brrunby, b.Brcomplete, b.Brcloseddate, b.Brclosedby, b.Brlastbilldate, b.Brmonthstobill, b.Brlowestbillno, b.Brhighestbillno, b.Brnoofbills, b.Brprevos, b.Brwholesale, b.Brretail, b.Brretail2, b.Broneoff, b.Brgreendeal, b.Brmonthly, b.Brmonthlyw, b.Brcreditall, b.Brcreditunits, b.Brdebitall, b.Brvattp, b.Brvatgp, b.Brvatep, b.Brnettp, b.Brnetgp, b.Brnetep, b.Brdiscount, b.Brsurcharge, b.Brclubfee, b.Brclubdisc, b.Brtotal, b.Brcgadiscount, b.Brchgecallcount, b.Brchgecallsecs, b.Brnoofcreditnote, b.Brcardcashbak, b.Braffcashbak, b.Brbydd, b.Brbydd1, b.EquinoxSec, b.EquinoxLrn)
	return err
}

// Save saves the Billrun to the database.
func (b *Billrun) Save(db XODB) error {
	if b.Exists() {
		return b.Update(db)
	}

	return b.Insert(db)
}

// Upsert performs an upsert for Billrun.
//
// NOTE: PostgreSQL 9.5+ only
func (b *Billrun) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.billruns (` +
		`brnumber, brperiod, brtype, brtypetxt, brrundate, brrunstarttime, brrunenddate, brrunendtime, brrunby, brcomplete, brcloseddate, brclosedby, brlastbilldate, brmonthstobill, brlowestbillno, brhighestbillno, brnoofbills, brprevos, brwholesale, brretail, brretail2, broneoff, brgreendeal, brmonthly, brmonthlyw, brcreditall, brcreditunits, brdebitall, brvattp, brvatgp, brvatep, brnettp, brnetgp, brnetep, brdiscount, brsurcharge, brclubfee, brclubdisc, brtotal, brcgadiscount, brchgecallcount, brchgecallsecs, brnoofcreditnote, brcardcashbak, braffcashbak, brbydd, brbydd1, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`brnumber, brperiod, brtype, brtypetxt, brrundate, brrunstarttime, brrunenddate, brrunendtime, brrunby, brcomplete, brcloseddate, brclosedby, brlastbilldate, brmonthstobill, brlowestbillno, brhighestbillno, brnoofbills, brprevos, brwholesale, brretail, brretail2, broneoff, brgreendeal, brmonthly, brmonthlyw, brcreditall, brcreditunits, brdebitall, brvattp, brvatgp, brvatep, brnettp, brnetgp, brnetep, brdiscount, brsurcharge, brclubfee, brclubdisc, brtotal, brcgadiscount, brchgecallcount, brchgecallsecs, brnoofcreditnote, brcardcashbak, braffcashbak, brbydd, brbydd1, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.brnumber, EXCLUDED.brperiod, EXCLUDED.brtype, EXCLUDED.brtypetxt, EXCLUDED.brrundate, EXCLUDED.brrunstarttime, EXCLUDED.brrunenddate, EXCLUDED.brrunendtime, EXCLUDED.brrunby, EXCLUDED.brcomplete, EXCLUDED.brcloseddate, EXCLUDED.brclosedby, EXCLUDED.brlastbilldate, EXCLUDED.brmonthstobill, EXCLUDED.brlowestbillno, EXCLUDED.brhighestbillno, EXCLUDED.brnoofbills, EXCLUDED.brprevos, EXCLUDED.brwholesale, EXCLUDED.brretail, EXCLUDED.brretail2, EXCLUDED.broneoff, EXCLUDED.brgreendeal, EXCLUDED.brmonthly, EXCLUDED.brmonthlyw, EXCLUDED.brcreditall, EXCLUDED.brcreditunits, EXCLUDED.brdebitall, EXCLUDED.brvattp, EXCLUDED.brvatgp, EXCLUDED.brvatep, EXCLUDED.brnettp, EXCLUDED.brnetgp, EXCLUDED.brnetep, EXCLUDED.brdiscount, EXCLUDED.brsurcharge, EXCLUDED.brclubfee, EXCLUDED.brclubdisc, EXCLUDED.brtotal, EXCLUDED.brcgadiscount, EXCLUDED.brchgecallcount, EXCLUDED.brchgecallsecs, EXCLUDED.brnoofcreditnote, EXCLUDED.brcardcashbak, EXCLUDED.braffcashbak, EXCLUDED.brbydd, EXCLUDED.brbydd1, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, b.Brnumber, b.Brperiod, b.Brtype, b.Brtypetxt, b.Brrundate, b.Brrunstarttime, b.Brrunenddate, b.Brrunendtime, b.Brrunby, b.Brcomplete, b.Brcloseddate, b.Brclosedby, b.Brlastbilldate, b.Brmonthstobill, b.Brlowestbillno, b.Brhighestbillno, b.Brnoofbills, b.Brprevos, b.Brwholesale, b.Brretail, b.Brretail2, b.Broneoff, b.Brgreendeal, b.Brmonthly, b.Brmonthlyw, b.Brcreditall, b.Brcreditunits, b.Brdebitall, b.Brvattp, b.Brvatgp, b.Brvatep, b.Brnettp, b.Brnetgp, b.Brnetep, b.Brdiscount, b.Brsurcharge, b.Brclubfee, b.Brclubdisc, b.Brtotal, b.Brcgadiscount, b.Brchgecallcount, b.Brchgecallsecs, b.Brnoofcreditnote, b.Brcardcashbak, b.Braffcashbak, b.Brbydd, b.Brbydd1, b.EquinoxLrn, b.EquinoxSec)
	_, err = db.Exec(sqlstr, b.Brnumber, b.Brperiod, b.Brtype, b.Brtypetxt, b.Brrundate, b.Brrunstarttime, b.Brrunenddate, b.Brrunendtime, b.Brrunby, b.Brcomplete, b.Brcloseddate, b.Brclosedby, b.Brlastbilldate, b.Brmonthstobill, b.Brlowestbillno, b.Brhighestbillno, b.Brnoofbills, b.Brprevos, b.Brwholesale, b.Brretail, b.Brretail2, b.Broneoff, b.Brgreendeal, b.Brmonthly, b.Brmonthlyw, b.Brcreditall, b.Brcreditunits, b.Brdebitall, b.Brvattp, b.Brvatgp, b.Brvatep, b.Brnettp, b.Brnetgp, b.Brnetep, b.Brdiscount, b.Brsurcharge, b.Brclubfee, b.Brclubdisc, b.Brtotal, b.Brcgadiscount, b.Brchgecallcount, b.Brchgecallsecs, b.Brnoofcreditnote, b.Brcardcashbak, b.Braffcashbak, b.Brbydd, b.Brbydd1, b.EquinoxLrn, b.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
}

// Delete deletes the Billrun from the database.
func (b *Billrun) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return nil
	}

	// if deleted, bail
	if b._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.billruns WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, b.EquinoxLrn)
	_, err = db.Exec(sqlstr, b.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	b._deleted = true

	return nil
}

// BillrunByEquinoxLrn retrieves a row from 'equinox.billruns' as a Billrun.
//
// Generated from index 'billruns_pkey'.
func BillrunByEquinoxLrn(db XODB, equinoxLrn int64) (*Billrun, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`brnumber, brperiod, brtype, brtypetxt, brrundate, brrunstarttime, brrunenddate, brrunendtime, brrunby, brcomplete, brcloseddate, brclosedby, brlastbilldate, brmonthstobill, brlowestbillno, brhighestbillno, brnoofbills, brprevos, brwholesale, brretail, brretail2, broneoff, brgreendeal, brmonthly, brmonthlyw, brcreditall, brcreditunits, brdebitall, brvattp, brvatgp, brvatep, brnettp, brnetgp, brnetep, brdiscount, brsurcharge, brclubfee, brclubdisc, brtotal, brcgadiscount, brchgecallcount, brchgecallsecs, brnoofcreditnote, brcardcashbak, braffcashbak, brbydd, brbydd1, equinox_lrn, equinox_sec ` +
		`FROM equinox.billruns ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	b := Billrun{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&b.Brnumber, &b.Brperiod, &b.Brtype, &b.Brtypetxt, &b.Brrundate, &b.Brrunstarttime, &b.Brrunenddate, &b.Brrunendtime, &b.Brrunby, &b.Brcomplete, &b.Brcloseddate, &b.Brclosedby, &b.Brlastbilldate, &b.Brmonthstobill, &b.Brlowestbillno, &b.Brhighestbillno, &b.Brnoofbills, &b.Brprevos, &b.Brwholesale, &b.Brretail, &b.Brretail2, &b.Broneoff, &b.Brgreendeal, &b.Brmonthly, &b.Brmonthlyw, &b.Brcreditall, &b.Brcreditunits, &b.Brdebitall, &b.Brvattp, &b.Brvatgp, &b.Brvatep, &b.Brnettp, &b.Brnetgp, &b.Brnetep, &b.Brdiscount, &b.Brsurcharge, &b.Brclubfee, &b.Brclubdisc, &b.Brtotal, &b.Brcgadiscount, &b.Brchgecallcount, &b.Brchgecallsecs, &b.Brnoofcreditnote, &b.Brcardcashbak, &b.Braffcashbak, &b.Brbydd, &b.Brbydd1, &b.EquinoxLrn, &b.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &b, nil
}