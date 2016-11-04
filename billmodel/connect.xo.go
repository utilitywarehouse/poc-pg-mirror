// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Connect represents a row from 'equinox.connect'.
type Connect struct {
	Cnxaccountno     sql.NullString  `json:"cnxaccountno"`     // cnxaccountno
	Cnxbtbillsupp    sql.NullInt64   `json:"cnxbtbillsupp"`    // cnxbtbillsupp
	Cnxletterflag    sql.NullInt64   `json:"cnxletterflag"`    // cnxletterflag
	Cnxactpending    sql.NullString  `json:"cnxactpending"`    // cnxactpending
	Cnxspmp          sql.NullInt64   `json:"cnxspmp"`          // cnxspmp
	Cnxcomment       sql.NullInt64   `json:"cnxcomment"`       // cnxcomment
	Cnxcomment2      sql.NullString  `json:"cnxcomment2"`      // cnxcomment2
	Cnxcomment3      sql.NullString  `json:"cnxcomment3"`      // cnxcomment3
	Cnxconnqueries   sql.NullString  `json:"cnxconnqueries"`   // cnxconnqueries
	Cnxaddrcheck     sql.NullInt64   `json:"cnxaddrcheck"`     // cnxaddrcheck
	Cnxappform       sql.NullInt64   `json:"cnxappform"`       // cnxappform
	Cnxcbccreditcard pq.NullTime     `json:"cnxcbccreditcard"` // cnxcbccreditcard
	Cnxcontacttel    sql.NullString  `json:"cnxcontacttel"`    // cnxcontacttel
	Cnxcoudatetobill pq.NullTime     `json:"cnxcoudatetobill"` // cnxcoudatetobill
	Cnxreferral      sql.NullString  `json:"cnxreferral"`      // cnxreferral
	Cnxownerortenant sql.NullString  `json:"cnxownerortenant"` // cnxownerortenant
	Cnxinitservices  sql.NullString  `json:"cnxinitservices"`  // cnxinitservices
	Cnxexpress       sql.NullInt64   `json:"cnxexpress"`       // cnxexpress
	Cnxeligablecust  sql.NullInt64   `json:"cnxeligablecust"`  // cnxeligablecust
	Cnxappformno     sql.NullString  `json:"cnxappformno"`     // cnxappformno
	Cnxrequestservs  sql.NullString  `json:"cnxrequestservs"`  // cnxrequestservs
	Cnxdatetolr      pq.NullTime     `json:"cnxdatetolr"`      // cnxdatetolr
	Cnxtenancyend    pq.NullTime     `json:"cnxtenancyend"`    // cnxtenancyend
	Cnxcbcbarred     pq.NullTime     `json:"cnxcbcbarred"`     // cnxcbcbarred
	Cnxversion       sql.NullInt64   `json:"cnxversion"`       // cnxversion
	Cnxprotectorlvl  sql.NullFloat64 `json:"cnxprotectorlvl"`  // cnxprotectorlvl
	Cnxcreditscore   sql.NullInt64   `json:"cnxcreditscore"`   // cnxcreditscore
	Cnxexorder       sql.NullInt64   `json:"cnxexorder"`       // cnxexorder
	Cnxsurname2      sql.NullString  `json:"cnxsurname2"`      // cnxsurname2
	Cnxtitle2        sql.NullString  `json:"cnxtitle2"`        // cnxtitle2
	Cnxfirstname2    sql.NullString  `json:"cnxfirstname2"`    // cnxfirstname2
	Cnxrelationship  sql.NullString  `json:"cnxrelationship"`  // cnxrelationship
	Cnxcreditcheck   sql.NullString  `json:"cnxcreditcheck"`   // cnxcreditcheck
	Cnxchecking      sql.NullString  `json:"cnxchecking"`      // cnxchecking
	Cnxvaliddate     pq.NullTime     `json:"cnxvaliddate"`     // cnxvaliddate
	Cnxvalidby       sql.NullString  `json:"cnxvalidby"`       // cnxvalidby
	Cnxappsource     sql.NullString  `json:"cnxappsource"`     // cnxappsource
	Cnxcampaign      sql.NullString  `json:"cnxcampaign"`      // cnxcampaign
	Cnxeligablesrvs  sql.NullInt64   `json:"cnxeligablesrvs"`  // cnxeligablesrvs
	Cnxgiftsettled   pq.NullTime     `json:"cnxgiftsettled"`   // cnxgiftsettled
	Cnxservicelvl    sql.NullString  `json:"cnxservicelvl"`    // cnxservicelvl
	Cnxapptype       sql.NullInt64   `json:"cnxapptype"`       // cnxapptype
	Cnxbusnature     sql.NullString  `json:"cnxbusnature"`     // cnxbusnature
	Cnxcomregno      sql.NullString  `json:"cnxcomregno"`      // cnxcomregno
	Cnxyrstrading    sql.NullInt64   `json:"cnxyrstrading"`    // cnxyrstrading
	Cnxcomcredscore  sql.NullInt64   `json:"cnxcomcredscore"`  // cnxcomcredscore
	Cnxsiccode       sql.NullString  `json:"cnxsiccode"`       // cnxsiccode
	Cnxcompanypos    sql.NullString  `json:"cnxcompanypos"`    // cnxcompanypos
	Cnxownerfirst    sql.NullInt64   `json:"cnxownerfirst"`    // cnxownerfirst
	Cnxownerlast     sql.NullInt64   `json:"cnxownerlast"`     // cnxownerlast
	Cnxownertitle    sql.NullInt64   `json:"cnxownertitle"`    // cnxownertitle
	Cnxownerdob      sql.NullInt64   `json:"cnxownerdob"`      // cnxownerdob
	Cnxowneradd1     sql.NullInt64   `json:"cnxowneradd1"`     // cnxowneradd1
	Cnxowneradd2     sql.NullInt64   `json:"cnxowneradd2"`     // cnxowneradd2
	Cnxowneradd3     sql.NullInt64   `json:"cnxowneradd3"`     // cnxowneradd3
	Cnxownercity     sql.NullInt64   `json:"cnxownercity"`     // cnxownercity
	Cnxownercnty     sql.NullInt64   `json:"cnxownercnty"`     // cnxownercnty
	Cnxownerpcod     sql.NullInt64   `json:"cnxownerpcod"`     // cnxownerpcod
	Cnxowneraddver   sql.NullInt64   `json:"cnxowneraddver"`   // cnxowneraddver
	Cnxidcheck       sql.NullString  `json:"cnxidcheck"`       // cnxidcheck
	Cnxelectoral     sql.NullString  `json:"cnxelectoral"`     // cnxelectoral
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Connect exists in the database.
func (c *Connect) Exists() bool {
	return c._exists
}

// Deleted provides information if the Connect has been deleted from the database.
func (c *Connect) Deleted() bool {
	return c._deleted
}

// Insert inserts the Connect to the database.
func (c *Connect) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.connect (` +
		`cnxaccountno, cnxbtbillsupp, cnxletterflag, cnxactpending, cnxspmp, cnxcomment, cnxcomment2, cnxcomment3, cnxconnqueries, cnxaddrcheck, cnxappform, cnxcbccreditcard, cnxcontacttel, cnxcoudatetobill, cnxreferral, cnxownerortenant, cnxinitservices, cnxexpress, cnxeligablecust, cnxappformno, cnxrequestservs, cnxdatetolr, cnxtenancyend, cnxcbcbarred, cnxversion, cnxprotectorlvl, cnxcreditscore, cnxexorder, cnxsurname2, cnxtitle2, cnxfirstname2, cnxrelationship, cnxcreditcheck, cnxchecking, cnxvaliddate, cnxvalidby, cnxappsource, cnxcampaign, cnxeligablesrvs, cnxgiftsettled, cnxservicelvl, cnxapptype, cnxbusnature, cnxcomregno, cnxyrstrading, cnxcomcredscore, cnxsiccode, cnxcompanypos, cnxownerfirst, cnxownerlast, cnxownertitle, cnxownerdob, cnxowneradd1, cnxowneradd2, cnxowneradd3, cnxownercity, cnxownercnty, cnxownerpcod, cnxowneraddver, cnxidcheck, cnxelectoral, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, c.Cnxaccountno, c.Cnxbtbillsupp, c.Cnxletterflag, c.Cnxactpending, c.Cnxspmp, c.Cnxcomment, c.Cnxcomment2, c.Cnxcomment3, c.Cnxconnqueries, c.Cnxaddrcheck, c.Cnxappform, c.Cnxcbccreditcard, c.Cnxcontacttel, c.Cnxcoudatetobill, c.Cnxreferral, c.Cnxownerortenant, c.Cnxinitservices, c.Cnxexpress, c.Cnxeligablecust, c.Cnxappformno, c.Cnxrequestservs, c.Cnxdatetolr, c.Cnxtenancyend, c.Cnxcbcbarred, c.Cnxversion, c.Cnxprotectorlvl, c.Cnxcreditscore, c.Cnxexorder, c.Cnxsurname2, c.Cnxtitle2, c.Cnxfirstname2, c.Cnxrelationship, c.Cnxcreditcheck, c.Cnxchecking, c.Cnxvaliddate, c.Cnxvalidby, c.Cnxappsource, c.Cnxcampaign, c.Cnxeligablesrvs, c.Cnxgiftsettled, c.Cnxservicelvl, c.Cnxapptype, c.Cnxbusnature, c.Cnxcomregno, c.Cnxyrstrading, c.Cnxcomcredscore, c.Cnxsiccode, c.Cnxcompanypos, c.Cnxownerfirst, c.Cnxownerlast, c.Cnxownertitle, c.Cnxownerdob, c.Cnxowneradd1, c.Cnxowneradd2, c.Cnxowneradd3, c.Cnxownercity, c.Cnxownercnty, c.Cnxownerpcod, c.Cnxowneraddver, c.Cnxidcheck, c.Cnxelectoral, c.EquinoxSec)
	err = db.QueryRow(sqlstr, c.Cnxaccountno, c.Cnxbtbillsupp, c.Cnxletterflag, c.Cnxactpending, c.Cnxspmp, c.Cnxcomment, c.Cnxcomment2, c.Cnxcomment3, c.Cnxconnqueries, c.Cnxaddrcheck, c.Cnxappform, c.Cnxcbccreditcard, c.Cnxcontacttel, c.Cnxcoudatetobill, c.Cnxreferral, c.Cnxownerortenant, c.Cnxinitservices, c.Cnxexpress, c.Cnxeligablecust, c.Cnxappformno, c.Cnxrequestservs, c.Cnxdatetolr, c.Cnxtenancyend, c.Cnxcbcbarred, c.Cnxversion, c.Cnxprotectorlvl, c.Cnxcreditscore, c.Cnxexorder, c.Cnxsurname2, c.Cnxtitle2, c.Cnxfirstname2, c.Cnxrelationship, c.Cnxcreditcheck, c.Cnxchecking, c.Cnxvaliddate, c.Cnxvalidby, c.Cnxappsource, c.Cnxcampaign, c.Cnxeligablesrvs, c.Cnxgiftsettled, c.Cnxservicelvl, c.Cnxapptype, c.Cnxbusnature, c.Cnxcomregno, c.Cnxyrstrading, c.Cnxcomcredscore, c.Cnxsiccode, c.Cnxcompanypos, c.Cnxownerfirst, c.Cnxownerlast, c.Cnxownertitle, c.Cnxownerdob, c.Cnxowneradd1, c.Cnxowneradd2, c.Cnxowneradd3, c.Cnxownercity, c.Cnxownercnty, c.Cnxownerpcod, c.Cnxowneraddver, c.Cnxidcheck, c.Cnxelectoral, c.EquinoxSec).Scan(&c.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Update updates the Connect in the database.
func (c *Connect) Update(db XODB) error {
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
	const sqlstr = `UPDATE equinox.connect SET (` +
		`cnxaccountno, cnxbtbillsupp, cnxletterflag, cnxactpending, cnxspmp, cnxcomment, cnxcomment2, cnxcomment3, cnxconnqueries, cnxaddrcheck, cnxappform, cnxcbccreditcard, cnxcontacttel, cnxcoudatetobill, cnxreferral, cnxownerortenant, cnxinitservices, cnxexpress, cnxeligablecust, cnxappformno, cnxrequestservs, cnxdatetolr, cnxtenancyend, cnxcbcbarred, cnxversion, cnxprotectorlvl, cnxcreditscore, cnxexorder, cnxsurname2, cnxtitle2, cnxfirstname2, cnxrelationship, cnxcreditcheck, cnxchecking, cnxvaliddate, cnxvalidby, cnxappsource, cnxcampaign, cnxeligablesrvs, cnxgiftsettled, cnxservicelvl, cnxapptype, cnxbusnature, cnxcomregno, cnxyrstrading, cnxcomcredscore, cnxsiccode, cnxcompanypos, cnxownerfirst, cnxownerlast, cnxownertitle, cnxownerdob, cnxowneradd1, cnxowneradd2, cnxowneradd3, cnxownercity, cnxownercnty, cnxownerpcod, cnxowneraddver, cnxidcheck, cnxelectoral, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62` +
		`) WHERE equinox_lrn = $63`

	// run query
	XOLog(sqlstr, c.Cnxaccountno, c.Cnxbtbillsupp, c.Cnxletterflag, c.Cnxactpending, c.Cnxspmp, c.Cnxcomment, c.Cnxcomment2, c.Cnxcomment3, c.Cnxconnqueries, c.Cnxaddrcheck, c.Cnxappform, c.Cnxcbccreditcard, c.Cnxcontacttel, c.Cnxcoudatetobill, c.Cnxreferral, c.Cnxownerortenant, c.Cnxinitservices, c.Cnxexpress, c.Cnxeligablecust, c.Cnxappformno, c.Cnxrequestservs, c.Cnxdatetolr, c.Cnxtenancyend, c.Cnxcbcbarred, c.Cnxversion, c.Cnxprotectorlvl, c.Cnxcreditscore, c.Cnxexorder, c.Cnxsurname2, c.Cnxtitle2, c.Cnxfirstname2, c.Cnxrelationship, c.Cnxcreditcheck, c.Cnxchecking, c.Cnxvaliddate, c.Cnxvalidby, c.Cnxappsource, c.Cnxcampaign, c.Cnxeligablesrvs, c.Cnxgiftsettled, c.Cnxservicelvl, c.Cnxapptype, c.Cnxbusnature, c.Cnxcomregno, c.Cnxyrstrading, c.Cnxcomcredscore, c.Cnxsiccode, c.Cnxcompanypos, c.Cnxownerfirst, c.Cnxownerlast, c.Cnxownertitle, c.Cnxownerdob, c.Cnxowneradd1, c.Cnxowneradd2, c.Cnxowneradd3, c.Cnxownercity, c.Cnxownercnty, c.Cnxownerpcod, c.Cnxowneraddver, c.Cnxidcheck, c.Cnxelectoral, c.EquinoxSec, c.EquinoxLrn)
	_, err = db.Exec(sqlstr, c.Cnxaccountno, c.Cnxbtbillsupp, c.Cnxletterflag, c.Cnxactpending, c.Cnxspmp, c.Cnxcomment, c.Cnxcomment2, c.Cnxcomment3, c.Cnxconnqueries, c.Cnxaddrcheck, c.Cnxappform, c.Cnxcbccreditcard, c.Cnxcontacttel, c.Cnxcoudatetobill, c.Cnxreferral, c.Cnxownerortenant, c.Cnxinitservices, c.Cnxexpress, c.Cnxeligablecust, c.Cnxappformno, c.Cnxrequestservs, c.Cnxdatetolr, c.Cnxtenancyend, c.Cnxcbcbarred, c.Cnxversion, c.Cnxprotectorlvl, c.Cnxcreditscore, c.Cnxexorder, c.Cnxsurname2, c.Cnxtitle2, c.Cnxfirstname2, c.Cnxrelationship, c.Cnxcreditcheck, c.Cnxchecking, c.Cnxvaliddate, c.Cnxvalidby, c.Cnxappsource, c.Cnxcampaign, c.Cnxeligablesrvs, c.Cnxgiftsettled, c.Cnxservicelvl, c.Cnxapptype, c.Cnxbusnature, c.Cnxcomregno, c.Cnxyrstrading, c.Cnxcomcredscore, c.Cnxsiccode, c.Cnxcompanypos, c.Cnxownerfirst, c.Cnxownerlast, c.Cnxownertitle, c.Cnxownerdob, c.Cnxowneradd1, c.Cnxowneradd2, c.Cnxowneradd3, c.Cnxownercity, c.Cnxownercnty, c.Cnxownerpcod, c.Cnxowneraddver, c.Cnxidcheck, c.Cnxelectoral, c.EquinoxSec, c.EquinoxLrn)
	return err
}

// Save saves the Connect to the database.
func (c *Connect) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Upsert performs an upsert for Connect.
//
// NOTE: PostgreSQL 9.5+ only
func (c *Connect) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.connect (` +
		`cnxaccountno, cnxbtbillsupp, cnxletterflag, cnxactpending, cnxspmp, cnxcomment, cnxcomment2, cnxcomment3, cnxconnqueries, cnxaddrcheck, cnxappform, cnxcbccreditcard, cnxcontacttel, cnxcoudatetobill, cnxreferral, cnxownerortenant, cnxinitservices, cnxexpress, cnxeligablecust, cnxappformno, cnxrequestservs, cnxdatetolr, cnxtenancyend, cnxcbcbarred, cnxversion, cnxprotectorlvl, cnxcreditscore, cnxexorder, cnxsurname2, cnxtitle2, cnxfirstname2, cnxrelationship, cnxcreditcheck, cnxchecking, cnxvaliddate, cnxvalidby, cnxappsource, cnxcampaign, cnxeligablesrvs, cnxgiftsettled, cnxservicelvl, cnxapptype, cnxbusnature, cnxcomregno, cnxyrstrading, cnxcomcredscore, cnxsiccode, cnxcompanypos, cnxownerfirst, cnxownerlast, cnxownertitle, cnxownerdob, cnxowneradd1, cnxowneradd2, cnxowneradd3, cnxownercity, cnxownercnty, cnxownerpcod, cnxowneraddver, cnxidcheck, cnxelectoral, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`cnxaccountno, cnxbtbillsupp, cnxletterflag, cnxactpending, cnxspmp, cnxcomment, cnxcomment2, cnxcomment3, cnxconnqueries, cnxaddrcheck, cnxappform, cnxcbccreditcard, cnxcontacttel, cnxcoudatetobill, cnxreferral, cnxownerortenant, cnxinitservices, cnxexpress, cnxeligablecust, cnxappformno, cnxrequestservs, cnxdatetolr, cnxtenancyend, cnxcbcbarred, cnxversion, cnxprotectorlvl, cnxcreditscore, cnxexorder, cnxsurname2, cnxtitle2, cnxfirstname2, cnxrelationship, cnxcreditcheck, cnxchecking, cnxvaliddate, cnxvalidby, cnxappsource, cnxcampaign, cnxeligablesrvs, cnxgiftsettled, cnxservicelvl, cnxapptype, cnxbusnature, cnxcomregno, cnxyrstrading, cnxcomcredscore, cnxsiccode, cnxcompanypos, cnxownerfirst, cnxownerlast, cnxownertitle, cnxownerdob, cnxowneradd1, cnxowneradd2, cnxowneradd3, cnxownercity, cnxownercnty, cnxownerpcod, cnxowneraddver, cnxidcheck, cnxelectoral, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.cnxaccountno, EXCLUDED.cnxbtbillsupp, EXCLUDED.cnxletterflag, EXCLUDED.cnxactpending, EXCLUDED.cnxspmp, EXCLUDED.cnxcomment, EXCLUDED.cnxcomment2, EXCLUDED.cnxcomment3, EXCLUDED.cnxconnqueries, EXCLUDED.cnxaddrcheck, EXCLUDED.cnxappform, EXCLUDED.cnxcbccreditcard, EXCLUDED.cnxcontacttel, EXCLUDED.cnxcoudatetobill, EXCLUDED.cnxreferral, EXCLUDED.cnxownerortenant, EXCLUDED.cnxinitservices, EXCLUDED.cnxexpress, EXCLUDED.cnxeligablecust, EXCLUDED.cnxappformno, EXCLUDED.cnxrequestservs, EXCLUDED.cnxdatetolr, EXCLUDED.cnxtenancyend, EXCLUDED.cnxcbcbarred, EXCLUDED.cnxversion, EXCLUDED.cnxprotectorlvl, EXCLUDED.cnxcreditscore, EXCLUDED.cnxexorder, EXCLUDED.cnxsurname2, EXCLUDED.cnxtitle2, EXCLUDED.cnxfirstname2, EXCLUDED.cnxrelationship, EXCLUDED.cnxcreditcheck, EXCLUDED.cnxchecking, EXCLUDED.cnxvaliddate, EXCLUDED.cnxvalidby, EXCLUDED.cnxappsource, EXCLUDED.cnxcampaign, EXCLUDED.cnxeligablesrvs, EXCLUDED.cnxgiftsettled, EXCLUDED.cnxservicelvl, EXCLUDED.cnxapptype, EXCLUDED.cnxbusnature, EXCLUDED.cnxcomregno, EXCLUDED.cnxyrstrading, EXCLUDED.cnxcomcredscore, EXCLUDED.cnxsiccode, EXCLUDED.cnxcompanypos, EXCLUDED.cnxownerfirst, EXCLUDED.cnxownerlast, EXCLUDED.cnxownertitle, EXCLUDED.cnxownerdob, EXCLUDED.cnxowneradd1, EXCLUDED.cnxowneradd2, EXCLUDED.cnxowneradd3, EXCLUDED.cnxownercity, EXCLUDED.cnxownercnty, EXCLUDED.cnxownerpcod, EXCLUDED.cnxowneraddver, EXCLUDED.cnxidcheck, EXCLUDED.cnxelectoral, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, c.Cnxaccountno, c.Cnxbtbillsupp, c.Cnxletterflag, c.Cnxactpending, c.Cnxspmp, c.Cnxcomment, c.Cnxcomment2, c.Cnxcomment3, c.Cnxconnqueries, c.Cnxaddrcheck, c.Cnxappform, c.Cnxcbccreditcard, c.Cnxcontacttel, c.Cnxcoudatetobill, c.Cnxreferral, c.Cnxownerortenant, c.Cnxinitservices, c.Cnxexpress, c.Cnxeligablecust, c.Cnxappformno, c.Cnxrequestservs, c.Cnxdatetolr, c.Cnxtenancyend, c.Cnxcbcbarred, c.Cnxversion, c.Cnxprotectorlvl, c.Cnxcreditscore, c.Cnxexorder, c.Cnxsurname2, c.Cnxtitle2, c.Cnxfirstname2, c.Cnxrelationship, c.Cnxcreditcheck, c.Cnxchecking, c.Cnxvaliddate, c.Cnxvalidby, c.Cnxappsource, c.Cnxcampaign, c.Cnxeligablesrvs, c.Cnxgiftsettled, c.Cnxservicelvl, c.Cnxapptype, c.Cnxbusnature, c.Cnxcomregno, c.Cnxyrstrading, c.Cnxcomcredscore, c.Cnxsiccode, c.Cnxcompanypos, c.Cnxownerfirst, c.Cnxownerlast, c.Cnxownertitle, c.Cnxownerdob, c.Cnxowneradd1, c.Cnxowneradd2, c.Cnxowneradd3, c.Cnxownercity, c.Cnxownercnty, c.Cnxownerpcod, c.Cnxowneraddver, c.Cnxidcheck, c.Cnxelectoral, c.EquinoxLrn, c.EquinoxSec)
	_, err = db.Exec(sqlstr, c.Cnxaccountno, c.Cnxbtbillsupp, c.Cnxletterflag, c.Cnxactpending, c.Cnxspmp, c.Cnxcomment, c.Cnxcomment2, c.Cnxcomment3, c.Cnxconnqueries, c.Cnxaddrcheck, c.Cnxappform, c.Cnxcbccreditcard, c.Cnxcontacttel, c.Cnxcoudatetobill, c.Cnxreferral, c.Cnxownerortenant, c.Cnxinitservices, c.Cnxexpress, c.Cnxeligablecust, c.Cnxappformno, c.Cnxrequestservs, c.Cnxdatetolr, c.Cnxtenancyend, c.Cnxcbcbarred, c.Cnxversion, c.Cnxprotectorlvl, c.Cnxcreditscore, c.Cnxexorder, c.Cnxsurname2, c.Cnxtitle2, c.Cnxfirstname2, c.Cnxrelationship, c.Cnxcreditcheck, c.Cnxchecking, c.Cnxvaliddate, c.Cnxvalidby, c.Cnxappsource, c.Cnxcampaign, c.Cnxeligablesrvs, c.Cnxgiftsettled, c.Cnxservicelvl, c.Cnxapptype, c.Cnxbusnature, c.Cnxcomregno, c.Cnxyrstrading, c.Cnxcomcredscore, c.Cnxsiccode, c.Cnxcompanypos, c.Cnxownerfirst, c.Cnxownerlast, c.Cnxownertitle, c.Cnxownerdob, c.Cnxowneradd1, c.Cnxowneradd2, c.Cnxowneradd3, c.Cnxownercity, c.Cnxownercnty, c.Cnxownerpcod, c.Cnxowneraddver, c.Cnxidcheck, c.Cnxelectoral, c.EquinoxLrn, c.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	c._exists = true

	return nil
}

// Delete deletes the Connect from the database.
func (c *Connect) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM equinox.connect WHERE equinox_lrn = $1`

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

// ConnectByEquinoxLrn retrieves a row from 'equinox.connect' as a Connect.
//
// Generated from index 'connect_pkey'.
func ConnectByEquinoxLrn(db XODB, equinoxLrn int64) (*Connect, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`cnxaccountno, cnxbtbillsupp, cnxletterflag, cnxactpending, cnxspmp, cnxcomment, cnxcomment2, cnxcomment3, cnxconnqueries, cnxaddrcheck, cnxappform, cnxcbccreditcard, cnxcontacttel, cnxcoudatetobill, cnxreferral, cnxownerortenant, cnxinitservices, cnxexpress, cnxeligablecust, cnxappformno, cnxrequestservs, cnxdatetolr, cnxtenancyend, cnxcbcbarred, cnxversion, cnxprotectorlvl, cnxcreditscore, cnxexorder, cnxsurname2, cnxtitle2, cnxfirstname2, cnxrelationship, cnxcreditcheck, cnxchecking, cnxvaliddate, cnxvalidby, cnxappsource, cnxcampaign, cnxeligablesrvs, cnxgiftsettled, cnxservicelvl, cnxapptype, cnxbusnature, cnxcomregno, cnxyrstrading, cnxcomcredscore, cnxsiccode, cnxcompanypos, cnxownerfirst, cnxownerlast, cnxownertitle, cnxownerdob, cnxowneradd1, cnxowneradd2, cnxowneradd3, cnxownercity, cnxownercnty, cnxownerpcod, cnxowneraddver, cnxidcheck, cnxelectoral, equinox_lrn, equinox_sec ` +
		`FROM equinox.connect ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	c := Connect{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&c.Cnxaccountno, &c.Cnxbtbillsupp, &c.Cnxletterflag, &c.Cnxactpending, &c.Cnxspmp, &c.Cnxcomment, &c.Cnxcomment2, &c.Cnxcomment3, &c.Cnxconnqueries, &c.Cnxaddrcheck, &c.Cnxappform, &c.Cnxcbccreditcard, &c.Cnxcontacttel, &c.Cnxcoudatetobill, &c.Cnxreferral, &c.Cnxownerortenant, &c.Cnxinitservices, &c.Cnxexpress, &c.Cnxeligablecust, &c.Cnxappformno, &c.Cnxrequestservs, &c.Cnxdatetolr, &c.Cnxtenancyend, &c.Cnxcbcbarred, &c.Cnxversion, &c.Cnxprotectorlvl, &c.Cnxcreditscore, &c.Cnxexorder, &c.Cnxsurname2, &c.Cnxtitle2, &c.Cnxfirstname2, &c.Cnxrelationship, &c.Cnxcreditcheck, &c.Cnxchecking, &c.Cnxvaliddate, &c.Cnxvalidby, &c.Cnxappsource, &c.Cnxcampaign, &c.Cnxeligablesrvs, &c.Cnxgiftsettled, &c.Cnxservicelvl, &c.Cnxapptype, &c.Cnxbusnature, &c.Cnxcomregno, &c.Cnxyrstrading, &c.Cnxcomcredscore, &c.Cnxsiccode, &c.Cnxcompanypos, &c.Cnxownerfirst, &c.Cnxownerlast, &c.Cnxownertitle, &c.Cnxownerdob, &c.Cnxowneradd1, &c.Cnxowneradd2, &c.Cnxowneradd3, &c.Cnxownercity, &c.Cnxownercnty, &c.Cnxownerpcod, &c.Cnxowneraddver, &c.Cnxidcheck, &c.Cnxelectoral, &c.EquinoxLrn, &c.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &c, nil
}