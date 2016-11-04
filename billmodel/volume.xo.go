// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// Volume represents a row from 'equinox.volume'.
type Volume struct {
	Voluniquesys     sql.NullFloat64 `json:"voluniquesys"`     // voluniquesys
	Volchrgoption    sql.NullString  `json:"volchrgoption"`    // volchrgoption
	Volfromto        sql.NullString  `json:"volfromto"`        // volfromto
	Volchargetype    sql.NullString  `json:"volchargetype"`    // volchargetype
	Volminsorpounds  sql.NullInt64   `json:"volminsorpounds"`  // volminsorpounds
	Voldiscountorvol sql.NullInt64   `json:"voldiscountorvol"` // voldiscountorvol
	Volvatrate       sql.NullFloat64 `json:"volvatrate"`       // volvatrate
	Volthresholdrat1 sql.NullInt64   `json:"volthresholdrat1"` // volthresholdrat1
	Volthresholdrat2 sql.NullInt64   `json:"volthresholdrat2"` // volthresholdrat2
	Volthresholdrat3 sql.NullInt64   `json:"volthresholdrat3"` // volthresholdrat3
	Volthresholdrat4 sql.NullInt64   `json:"volthresholdrat4"` // volthresholdrat4
	Volthresholdrat5 sql.NullInt64   `json:"volthresholdrat5"` // volthresholdrat5
	Volthresholdrat6 sql.NullInt64   `json:"volthresholdrat6"` // volthresholdrat6
	Volthresholdrat7 sql.NullInt64   `json:"volthresholdrat7"` // volthresholdrat7
	Volthresholdrat8 sql.NullInt64   `json:"volthresholdrat8"` // volthresholdrat8
	Volthresholdrat9 sql.NullInt64   `json:"volthresholdrat9"` // volthresholdrat9
	Volthreshper1    sql.NullFloat64 `json:"volthreshper1"`    // volthreshper1
	Volthreshper2    sql.NullFloat64 `json:"volthreshper2"`    // volthreshper2
	Volthreshper3    sql.NullFloat64 `json:"volthreshper3"`    // volthreshper3
	Volthreshper4    sql.NullFloat64 `json:"volthreshper4"`    // volthreshper4
	Volthreshper5    sql.NullFloat64 `json:"volthreshper5"`    // volthreshper5
	Volthreshper6    sql.NullFloat64 `json:"volthreshper6"`    // volthreshper6
	Volthreshper7    sql.NullFloat64 `json:"volthreshper7"`    // volthreshper7
	Volthreshper8    sql.NullFloat64 `json:"volthreshper8"`    // volthreshper8
	Volthreshper9    sql.NullFloat64 `json:"volthreshper9"`    // volthreshper9
	Volthreshper10   sql.NullFloat64 `json:"volthreshper10"`   // volthreshper10
	Voldayopt1       sql.NullInt64   `json:"voldayopt1"`       // voldayopt1
	Voldayopt2       sql.NullInt64   `json:"voldayopt2"`       // voldayopt2
	Voldayopt3       sql.NullInt64   `json:"voldayopt3"`       // voldayopt3
	Voldayopt4       sql.NullInt64   `json:"voldayopt4"`       // voldayopt4
	Voldayopt5       sql.NullInt64   `json:"voldayopt5"`       // voldayopt5
	Voldayopt6       sql.NullInt64   `json:"voldayopt6"`       // voldayopt6
	Voldayopt7       sql.NullInt64   `json:"voldayopt7"`       // voldayopt7
	Voldayopt8       sql.NullInt64   `json:"voldayopt8"`       // voldayopt8
	Voldayopt9       sql.NullInt64   `json:"voldayopt9"`       // voldayopt9
	Voldayopt10      sql.NullInt64   `json:"voldayopt10"`      // voldayopt10
	Voldayopt11      sql.NullInt64   `json:"voldayopt11"`      // voldayopt11
	Voldayopt12      sql.NullInt64   `json:"voldayopt12"`      // voldayopt12
	Voldayopt13      sql.NullInt64   `json:"voldayopt13"`      // voldayopt13
	Voldayopt14      sql.NullInt64   `json:"voldayopt14"`      // voldayopt14
	Voldayopt15      sql.NullInt64   `json:"voldayopt15"`      // voldayopt15
	Voldayopt16      sql.NullInt64   `json:"voldayopt16"`      // voldayopt16
	Voldayopt17      sql.NullInt64   `json:"voldayopt17"`      // voldayopt17
	Voldayopt18      sql.NullInt64   `json:"voldayopt18"`      // voldayopt18
	Voldayopt19      sql.NullInt64   `json:"voldayopt19"`      // voldayopt19
	Voldayopt20      sql.NullInt64   `json:"voldayopt20"`      // voldayopt20
	Voltimeopt1      pq.NullTime     `json:"voltimeopt1"`      // voltimeopt1
	Voltimeopt2      pq.NullTime     `json:"voltimeopt2"`      // voltimeopt2
	Voltimeopt3      pq.NullTime     `json:"voltimeopt3"`      // voltimeopt3
	Voltimeopt4      pq.NullTime     `json:"voltimeopt4"`      // voltimeopt4
	Voltimeopt5      pq.NullTime     `json:"voltimeopt5"`      // voltimeopt5
	Voltimeopt6      pq.NullTime     `json:"voltimeopt6"`      // voltimeopt6
	Voltimeopt7      pq.NullTime     `json:"voltimeopt7"`      // voltimeopt7
	Voltimeopt8      pq.NullTime     `json:"voltimeopt8"`      // voltimeopt8
	Voltimeopt9      pq.NullTime     `json:"voltimeopt9"`      // voltimeopt9
	Voltimeopt10     pq.NullTime     `json:"voltimeopt10"`     // voltimeopt10
	Voltimeopt11     pq.NullTime     `json:"voltimeopt11"`     // voltimeopt11
	Voltimeopt12     pq.NullTime     `json:"voltimeopt12"`     // voltimeopt12
	Voltimeopt13     pq.NullTime     `json:"voltimeopt13"`     // voltimeopt13
	Voltimeopt14     pq.NullTime     `json:"voltimeopt14"`     // voltimeopt14
	Voltimeopt15     pq.NullTime     `json:"voltimeopt15"`     // voltimeopt15
	Voltimeopt16     pq.NullTime     `json:"voltimeopt16"`     // voltimeopt16
	Voltimeopt17     pq.NullTime     `json:"voltimeopt17"`     // voltimeopt17
	Voltimeopt18     pq.NullTime     `json:"voltimeopt18"`     // voltimeopt18
	Voltimeopt19     pq.NullTime     `json:"voltimeopt19"`     // voltimeopt19
	Voltimeopt20     pq.NullTime     `json:"voltimeopt20"`     // voltimeopt20
	Voltimestype1    sql.NullInt64   `json:"voltimestype1"`    // voltimestype1
	Voltimestype2    sql.NullInt64   `json:"voltimestype2"`    // voltimestype2
	Voltimestype3    sql.NullInt64   `json:"voltimestype3"`    // voltimestype3
	Voltimestype4    sql.NullInt64   `json:"voltimestype4"`    // voltimestype4
	Voltimestype5    sql.NullInt64   `json:"voltimestype5"`    // voltimestype5
	Voltimestype6    sql.NullInt64   `json:"voltimestype6"`    // voltimestype6
	Voltimestype7    sql.NullInt64   `json:"voltimestype7"`    // voltimestype7
	Voltimestype8    sql.NullInt64   `json:"voltimestype8"`    // voltimestype8
	Voltimestype9    sql.NullInt64   `json:"voltimestype9"`    // voltimestype9
	Voltimestype10   sql.NullInt64   `json:"voltimestype10"`   // voltimestype10
	Voltimestype11   sql.NullInt64   `json:"voltimestype11"`   // voltimestype11
	Voltimestype12   sql.NullInt64   `json:"voltimestype12"`   // voltimestype12
	Voltimestype13   sql.NullInt64   `json:"voltimestype13"`   // voltimestype13
	Voltimestype14   sql.NullInt64   `json:"voltimestype14"`   // voltimestype14
	Voltimestype15   sql.NullInt64   `json:"voltimestype15"`   // voltimestype15
	Voltimestype16   sql.NullInt64   `json:"voltimestype16"`   // voltimestype16
	Voltimestype17   sql.NullInt64   `json:"voltimestype17"`   // voltimestype17
	Voltimestype18   sql.NullInt64   `json:"voltimestype18"`   // voltimestype18
	Voltimestype19   sql.NullInt64   `json:"voltimestype19"`   // voltimestype19
	Voltimestype20   sql.NullInt64   `json:"voltimestype20"`   // voltimestype20
	Volalt           sql.NullString  `json:"volalt"`           // volalt
	Volsparec1       sql.NullString  `json:"volsparec1"`       // volsparec1
	Volsparen1       sql.NullFloat64 `json:"volsparen1"`       // volsparen1
	Volspared1       pq.NullTime     `json:"volspared1"`       // volspared1
	EquinoxPrn       sql.NullInt64   `json:"equinox_prn"`      // equinox_prn
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Volume exists in the database.
func (v *Volume) Exists() bool {
	return v._exists
}

// Deleted provides information if the Volume has been deleted from the database.
func (v *Volume) Deleted() bool {
	return v._deleted
}

// Insert inserts the Volume to the database.
func (v *Volume) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if v._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.volume (` +
		`voluniquesys, volchrgoption, volfromto, volchargetype, volminsorpounds, voldiscountorvol, volvatrate, volthresholdrat1, volthresholdrat2, volthresholdrat3, volthresholdrat4, volthresholdrat5, volthresholdrat6, volthresholdrat7, volthresholdrat8, volthresholdrat9, volthreshper1, volthreshper2, volthreshper3, volthreshper4, volthreshper5, volthreshper6, volthreshper7, volthreshper8, volthreshper9, volthreshper10, voldayopt1, voldayopt2, voldayopt3, voldayopt4, voldayopt5, voldayopt6, voldayopt7, voldayopt8, voldayopt9, voldayopt10, voldayopt11, voldayopt12, voldayopt13, voldayopt14, voldayopt15, voldayopt16, voldayopt17, voldayopt18, voldayopt19, voldayopt20, voltimeopt1, voltimeopt2, voltimeopt3, voltimeopt4, voltimeopt5, voltimeopt6, voltimeopt7, voltimeopt8, voltimeopt9, voltimeopt10, voltimeopt11, voltimeopt12, voltimeopt13, voltimeopt14, voltimeopt15, voltimeopt16, voltimeopt17, voltimeopt18, voltimeopt19, voltimeopt20, voltimestype1, voltimestype2, voltimestype3, voltimestype4, voltimestype5, voltimestype6, voltimestype7, voltimestype8, voltimestype9, voltimestype10, voltimestype11, voltimestype12, voltimestype13, voltimestype14, voltimestype15, voltimestype16, voltimestype17, voltimestype18, voltimestype19, voltimestype20, volalt, volsparec1, volsparen1, volspared1, equinox_prn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81, $82, $83, $84, $85, $86, $87, $88, $89, $90, $91, $92` +
		`) RETURNING equinox_lrn`

	// run query
	XOLog(sqlstr, v.Voluniquesys, v.Volchrgoption, v.Volfromto, v.Volchargetype, v.Volminsorpounds, v.Voldiscountorvol, v.Volvatrate, v.Volthresholdrat1, v.Volthresholdrat2, v.Volthresholdrat3, v.Volthresholdrat4, v.Volthresholdrat5, v.Volthresholdrat6, v.Volthresholdrat7, v.Volthresholdrat8, v.Volthresholdrat9, v.Volthreshper1, v.Volthreshper2, v.Volthreshper3, v.Volthreshper4, v.Volthreshper5, v.Volthreshper6, v.Volthreshper7, v.Volthreshper8, v.Volthreshper9, v.Volthreshper10, v.Voldayopt1, v.Voldayopt2, v.Voldayopt3, v.Voldayopt4, v.Voldayopt5, v.Voldayopt6, v.Voldayopt7, v.Voldayopt8, v.Voldayopt9, v.Voldayopt10, v.Voldayopt11, v.Voldayopt12, v.Voldayopt13, v.Voldayopt14, v.Voldayopt15, v.Voldayopt16, v.Voldayopt17, v.Voldayopt18, v.Voldayopt19, v.Voldayopt20, v.Voltimeopt1, v.Voltimeopt2, v.Voltimeopt3, v.Voltimeopt4, v.Voltimeopt5, v.Voltimeopt6, v.Voltimeopt7, v.Voltimeopt8, v.Voltimeopt9, v.Voltimeopt10, v.Voltimeopt11, v.Voltimeopt12, v.Voltimeopt13, v.Voltimeopt14, v.Voltimeopt15, v.Voltimeopt16, v.Voltimeopt17, v.Voltimeopt18, v.Voltimeopt19, v.Voltimeopt20, v.Voltimestype1, v.Voltimestype2, v.Voltimestype3, v.Voltimestype4, v.Voltimestype5, v.Voltimestype6, v.Voltimestype7, v.Voltimestype8, v.Voltimestype9, v.Voltimestype10, v.Voltimestype11, v.Voltimestype12, v.Voltimestype13, v.Voltimestype14, v.Voltimestype15, v.Voltimestype16, v.Voltimestype17, v.Voltimestype18, v.Voltimestype19, v.Voltimestype20, v.Volalt, v.Volsparec1, v.Volsparen1, v.Volspared1, v.EquinoxPrn, v.EquinoxSec)
	err = db.QueryRow(sqlstr, v.Voluniquesys, v.Volchrgoption, v.Volfromto, v.Volchargetype, v.Volminsorpounds, v.Voldiscountorvol, v.Volvatrate, v.Volthresholdrat1, v.Volthresholdrat2, v.Volthresholdrat3, v.Volthresholdrat4, v.Volthresholdrat5, v.Volthresholdrat6, v.Volthresholdrat7, v.Volthresholdrat8, v.Volthresholdrat9, v.Volthreshper1, v.Volthreshper2, v.Volthreshper3, v.Volthreshper4, v.Volthreshper5, v.Volthreshper6, v.Volthreshper7, v.Volthreshper8, v.Volthreshper9, v.Volthreshper10, v.Voldayopt1, v.Voldayopt2, v.Voldayopt3, v.Voldayopt4, v.Voldayopt5, v.Voldayopt6, v.Voldayopt7, v.Voldayopt8, v.Voldayopt9, v.Voldayopt10, v.Voldayopt11, v.Voldayopt12, v.Voldayopt13, v.Voldayopt14, v.Voldayopt15, v.Voldayopt16, v.Voldayopt17, v.Voldayopt18, v.Voldayopt19, v.Voldayopt20, v.Voltimeopt1, v.Voltimeopt2, v.Voltimeopt3, v.Voltimeopt4, v.Voltimeopt5, v.Voltimeopt6, v.Voltimeopt7, v.Voltimeopt8, v.Voltimeopt9, v.Voltimeopt10, v.Voltimeopt11, v.Voltimeopt12, v.Voltimeopt13, v.Voltimeopt14, v.Voltimeopt15, v.Voltimeopt16, v.Voltimeopt17, v.Voltimeopt18, v.Voltimeopt19, v.Voltimeopt20, v.Voltimestype1, v.Voltimestype2, v.Voltimestype3, v.Voltimestype4, v.Voltimestype5, v.Voltimestype6, v.Voltimestype7, v.Voltimestype8, v.Voltimestype9, v.Voltimestype10, v.Voltimestype11, v.Voltimestype12, v.Voltimestype13, v.Voltimestype14, v.Voltimestype15, v.Voltimestype16, v.Voltimestype17, v.Voltimestype18, v.Voltimestype19, v.Voltimestype20, v.Volalt, v.Volsparec1, v.Volsparen1, v.Volspared1, v.EquinoxPrn, v.EquinoxSec).Scan(&v.EquinoxLrn)
	if err != nil {
		return err
	}

	// set existence
	v._exists = true

	return nil
}

// Update updates the Volume in the database.
func (v *Volume) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !v._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if v._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE equinox.volume SET (` +
		`voluniquesys, volchrgoption, volfromto, volchargetype, volminsorpounds, voldiscountorvol, volvatrate, volthresholdrat1, volthresholdrat2, volthresholdrat3, volthresholdrat4, volthresholdrat5, volthresholdrat6, volthresholdrat7, volthresholdrat8, volthresholdrat9, volthreshper1, volthreshper2, volthreshper3, volthreshper4, volthreshper5, volthreshper6, volthreshper7, volthreshper8, volthreshper9, volthreshper10, voldayopt1, voldayopt2, voldayopt3, voldayopt4, voldayopt5, voldayopt6, voldayopt7, voldayopt8, voldayopt9, voldayopt10, voldayopt11, voldayopt12, voldayopt13, voldayopt14, voldayopt15, voldayopt16, voldayopt17, voldayopt18, voldayopt19, voldayopt20, voltimeopt1, voltimeopt2, voltimeopt3, voltimeopt4, voltimeopt5, voltimeopt6, voltimeopt7, voltimeopt8, voltimeopt9, voltimeopt10, voltimeopt11, voltimeopt12, voltimeopt13, voltimeopt14, voltimeopt15, voltimeopt16, voltimeopt17, voltimeopt18, voltimeopt19, voltimeopt20, voltimestype1, voltimestype2, voltimestype3, voltimestype4, voltimestype5, voltimestype6, voltimestype7, voltimestype8, voltimestype9, voltimestype10, voltimestype11, voltimestype12, voltimestype13, voltimestype14, voltimestype15, voltimestype16, voltimestype17, voltimestype18, voltimestype19, voltimestype20, volalt, volsparec1, volsparen1, volspared1, equinox_prn, equinox_sec` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81, $82, $83, $84, $85, $86, $87, $88, $89, $90, $91, $92` +
		`) WHERE equinox_lrn = $93`

	// run query
	XOLog(sqlstr, v.Voluniquesys, v.Volchrgoption, v.Volfromto, v.Volchargetype, v.Volminsorpounds, v.Voldiscountorvol, v.Volvatrate, v.Volthresholdrat1, v.Volthresholdrat2, v.Volthresholdrat3, v.Volthresholdrat4, v.Volthresholdrat5, v.Volthresholdrat6, v.Volthresholdrat7, v.Volthresholdrat8, v.Volthresholdrat9, v.Volthreshper1, v.Volthreshper2, v.Volthreshper3, v.Volthreshper4, v.Volthreshper5, v.Volthreshper6, v.Volthreshper7, v.Volthreshper8, v.Volthreshper9, v.Volthreshper10, v.Voldayopt1, v.Voldayopt2, v.Voldayopt3, v.Voldayopt4, v.Voldayopt5, v.Voldayopt6, v.Voldayopt7, v.Voldayopt8, v.Voldayopt9, v.Voldayopt10, v.Voldayopt11, v.Voldayopt12, v.Voldayopt13, v.Voldayopt14, v.Voldayopt15, v.Voldayopt16, v.Voldayopt17, v.Voldayopt18, v.Voldayopt19, v.Voldayopt20, v.Voltimeopt1, v.Voltimeopt2, v.Voltimeopt3, v.Voltimeopt4, v.Voltimeopt5, v.Voltimeopt6, v.Voltimeopt7, v.Voltimeopt8, v.Voltimeopt9, v.Voltimeopt10, v.Voltimeopt11, v.Voltimeopt12, v.Voltimeopt13, v.Voltimeopt14, v.Voltimeopt15, v.Voltimeopt16, v.Voltimeopt17, v.Voltimeopt18, v.Voltimeopt19, v.Voltimeopt20, v.Voltimestype1, v.Voltimestype2, v.Voltimestype3, v.Voltimestype4, v.Voltimestype5, v.Voltimestype6, v.Voltimestype7, v.Voltimestype8, v.Voltimestype9, v.Voltimestype10, v.Voltimestype11, v.Voltimestype12, v.Voltimestype13, v.Voltimestype14, v.Voltimestype15, v.Voltimestype16, v.Voltimestype17, v.Voltimestype18, v.Voltimestype19, v.Voltimestype20, v.Volalt, v.Volsparec1, v.Volsparen1, v.Volspared1, v.EquinoxPrn, v.EquinoxSec, v.EquinoxLrn)
	_, err = db.Exec(sqlstr, v.Voluniquesys, v.Volchrgoption, v.Volfromto, v.Volchargetype, v.Volminsorpounds, v.Voldiscountorvol, v.Volvatrate, v.Volthresholdrat1, v.Volthresholdrat2, v.Volthresholdrat3, v.Volthresholdrat4, v.Volthresholdrat5, v.Volthresholdrat6, v.Volthresholdrat7, v.Volthresholdrat8, v.Volthresholdrat9, v.Volthreshper1, v.Volthreshper2, v.Volthreshper3, v.Volthreshper4, v.Volthreshper5, v.Volthreshper6, v.Volthreshper7, v.Volthreshper8, v.Volthreshper9, v.Volthreshper10, v.Voldayopt1, v.Voldayopt2, v.Voldayopt3, v.Voldayopt4, v.Voldayopt5, v.Voldayopt6, v.Voldayopt7, v.Voldayopt8, v.Voldayopt9, v.Voldayopt10, v.Voldayopt11, v.Voldayopt12, v.Voldayopt13, v.Voldayopt14, v.Voldayopt15, v.Voldayopt16, v.Voldayopt17, v.Voldayopt18, v.Voldayopt19, v.Voldayopt20, v.Voltimeopt1, v.Voltimeopt2, v.Voltimeopt3, v.Voltimeopt4, v.Voltimeopt5, v.Voltimeopt6, v.Voltimeopt7, v.Voltimeopt8, v.Voltimeopt9, v.Voltimeopt10, v.Voltimeopt11, v.Voltimeopt12, v.Voltimeopt13, v.Voltimeopt14, v.Voltimeopt15, v.Voltimeopt16, v.Voltimeopt17, v.Voltimeopt18, v.Voltimeopt19, v.Voltimeopt20, v.Voltimestype1, v.Voltimestype2, v.Voltimestype3, v.Voltimestype4, v.Voltimestype5, v.Voltimestype6, v.Voltimestype7, v.Voltimestype8, v.Voltimestype9, v.Voltimestype10, v.Voltimestype11, v.Voltimestype12, v.Voltimestype13, v.Voltimestype14, v.Voltimestype15, v.Voltimestype16, v.Voltimestype17, v.Voltimestype18, v.Voltimestype19, v.Voltimestype20, v.Volalt, v.Volsparec1, v.Volsparen1, v.Volspared1, v.EquinoxPrn, v.EquinoxSec, v.EquinoxLrn)
	return err
}

// Save saves the Volume to the database.
func (v *Volume) Save(db XODB) error {
	if v.Exists() {
		return v.Update(db)
	}

	return v.Insert(db)
}

// Upsert performs an upsert for Volume.
//
// NOTE: PostgreSQL 9.5+ only
func (v *Volume) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if v._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO equinox.volume (` +
		`voluniquesys, volchrgoption, volfromto, volchargetype, volminsorpounds, voldiscountorvol, volvatrate, volthresholdrat1, volthresholdrat2, volthresholdrat3, volthresholdrat4, volthresholdrat5, volthresholdrat6, volthresholdrat7, volthresholdrat8, volthresholdrat9, volthreshper1, volthreshper2, volthreshper3, volthreshper4, volthreshper5, volthreshper6, volthreshper7, volthreshper8, volthreshper9, volthreshper10, voldayopt1, voldayopt2, voldayopt3, voldayopt4, voldayopt5, voldayopt6, voldayopt7, voldayopt8, voldayopt9, voldayopt10, voldayopt11, voldayopt12, voldayopt13, voldayopt14, voldayopt15, voldayopt16, voldayopt17, voldayopt18, voldayopt19, voldayopt20, voltimeopt1, voltimeopt2, voltimeopt3, voltimeopt4, voltimeopt5, voltimeopt6, voltimeopt7, voltimeopt8, voltimeopt9, voltimeopt10, voltimeopt11, voltimeopt12, voltimeopt13, voltimeopt14, voltimeopt15, voltimeopt16, voltimeopt17, voltimeopt18, voltimeopt19, voltimeopt20, voltimestype1, voltimestype2, voltimestype3, voltimestype4, voltimestype5, voltimestype6, voltimestype7, voltimestype8, voltimestype9, voltimestype10, voltimestype11, voltimestype12, voltimestype13, voltimestype14, voltimestype15, voltimestype16, voltimestype17, voltimestype18, voltimestype19, voltimestype20, volalt, volsparec1, volsparen1, volspared1, equinox_prn, equinox_lrn, equinox_sec` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81, $82, $83, $84, $85, $86, $87, $88, $89, $90, $91, $92, $93` +
		`) ON CONFLICT (equinox_lrn) DO UPDATE SET (` +
		`voluniquesys, volchrgoption, volfromto, volchargetype, volminsorpounds, voldiscountorvol, volvatrate, volthresholdrat1, volthresholdrat2, volthresholdrat3, volthresholdrat4, volthresholdrat5, volthresholdrat6, volthresholdrat7, volthresholdrat8, volthresholdrat9, volthreshper1, volthreshper2, volthreshper3, volthreshper4, volthreshper5, volthreshper6, volthreshper7, volthreshper8, volthreshper9, volthreshper10, voldayopt1, voldayopt2, voldayopt3, voldayopt4, voldayopt5, voldayopt6, voldayopt7, voldayopt8, voldayopt9, voldayopt10, voldayopt11, voldayopt12, voldayopt13, voldayopt14, voldayopt15, voldayopt16, voldayopt17, voldayopt18, voldayopt19, voldayopt20, voltimeopt1, voltimeopt2, voltimeopt3, voltimeopt4, voltimeopt5, voltimeopt6, voltimeopt7, voltimeopt8, voltimeopt9, voltimeopt10, voltimeopt11, voltimeopt12, voltimeopt13, voltimeopt14, voltimeopt15, voltimeopt16, voltimeopt17, voltimeopt18, voltimeopt19, voltimeopt20, voltimestype1, voltimestype2, voltimestype3, voltimestype4, voltimestype5, voltimestype6, voltimestype7, voltimestype8, voltimestype9, voltimestype10, voltimestype11, voltimestype12, voltimestype13, voltimestype14, voltimestype15, voltimestype16, voltimestype17, voltimestype18, voltimestype19, voltimestype20, volalt, volsparec1, volsparen1, volspared1, equinox_prn, equinox_lrn, equinox_sec` +
		`) = (` +
		`EXCLUDED.voluniquesys, EXCLUDED.volchrgoption, EXCLUDED.volfromto, EXCLUDED.volchargetype, EXCLUDED.volminsorpounds, EXCLUDED.voldiscountorvol, EXCLUDED.volvatrate, EXCLUDED.volthresholdrat1, EXCLUDED.volthresholdrat2, EXCLUDED.volthresholdrat3, EXCLUDED.volthresholdrat4, EXCLUDED.volthresholdrat5, EXCLUDED.volthresholdrat6, EXCLUDED.volthresholdrat7, EXCLUDED.volthresholdrat8, EXCLUDED.volthresholdrat9, EXCLUDED.volthreshper1, EXCLUDED.volthreshper2, EXCLUDED.volthreshper3, EXCLUDED.volthreshper4, EXCLUDED.volthreshper5, EXCLUDED.volthreshper6, EXCLUDED.volthreshper7, EXCLUDED.volthreshper8, EXCLUDED.volthreshper9, EXCLUDED.volthreshper10, EXCLUDED.voldayopt1, EXCLUDED.voldayopt2, EXCLUDED.voldayopt3, EXCLUDED.voldayopt4, EXCLUDED.voldayopt5, EXCLUDED.voldayopt6, EXCLUDED.voldayopt7, EXCLUDED.voldayopt8, EXCLUDED.voldayopt9, EXCLUDED.voldayopt10, EXCLUDED.voldayopt11, EXCLUDED.voldayopt12, EXCLUDED.voldayopt13, EXCLUDED.voldayopt14, EXCLUDED.voldayopt15, EXCLUDED.voldayopt16, EXCLUDED.voldayopt17, EXCLUDED.voldayopt18, EXCLUDED.voldayopt19, EXCLUDED.voldayopt20, EXCLUDED.voltimeopt1, EXCLUDED.voltimeopt2, EXCLUDED.voltimeopt3, EXCLUDED.voltimeopt4, EXCLUDED.voltimeopt5, EXCLUDED.voltimeopt6, EXCLUDED.voltimeopt7, EXCLUDED.voltimeopt8, EXCLUDED.voltimeopt9, EXCLUDED.voltimeopt10, EXCLUDED.voltimeopt11, EXCLUDED.voltimeopt12, EXCLUDED.voltimeopt13, EXCLUDED.voltimeopt14, EXCLUDED.voltimeopt15, EXCLUDED.voltimeopt16, EXCLUDED.voltimeopt17, EXCLUDED.voltimeopt18, EXCLUDED.voltimeopt19, EXCLUDED.voltimeopt20, EXCLUDED.voltimestype1, EXCLUDED.voltimestype2, EXCLUDED.voltimestype3, EXCLUDED.voltimestype4, EXCLUDED.voltimestype5, EXCLUDED.voltimestype6, EXCLUDED.voltimestype7, EXCLUDED.voltimestype8, EXCLUDED.voltimestype9, EXCLUDED.voltimestype10, EXCLUDED.voltimestype11, EXCLUDED.voltimestype12, EXCLUDED.voltimestype13, EXCLUDED.voltimestype14, EXCLUDED.voltimestype15, EXCLUDED.voltimestype16, EXCLUDED.voltimestype17, EXCLUDED.voltimestype18, EXCLUDED.voltimestype19, EXCLUDED.voltimestype20, EXCLUDED.volalt, EXCLUDED.volsparec1, EXCLUDED.volsparen1, EXCLUDED.volspared1, EXCLUDED.equinox_prn, EXCLUDED.equinox_lrn, EXCLUDED.equinox_sec` +
		`)`

	// run query
	XOLog(sqlstr, v.Voluniquesys, v.Volchrgoption, v.Volfromto, v.Volchargetype, v.Volminsorpounds, v.Voldiscountorvol, v.Volvatrate, v.Volthresholdrat1, v.Volthresholdrat2, v.Volthresholdrat3, v.Volthresholdrat4, v.Volthresholdrat5, v.Volthresholdrat6, v.Volthresholdrat7, v.Volthresholdrat8, v.Volthresholdrat9, v.Volthreshper1, v.Volthreshper2, v.Volthreshper3, v.Volthreshper4, v.Volthreshper5, v.Volthreshper6, v.Volthreshper7, v.Volthreshper8, v.Volthreshper9, v.Volthreshper10, v.Voldayopt1, v.Voldayopt2, v.Voldayopt3, v.Voldayopt4, v.Voldayopt5, v.Voldayopt6, v.Voldayopt7, v.Voldayopt8, v.Voldayopt9, v.Voldayopt10, v.Voldayopt11, v.Voldayopt12, v.Voldayopt13, v.Voldayopt14, v.Voldayopt15, v.Voldayopt16, v.Voldayopt17, v.Voldayopt18, v.Voldayopt19, v.Voldayopt20, v.Voltimeopt1, v.Voltimeopt2, v.Voltimeopt3, v.Voltimeopt4, v.Voltimeopt5, v.Voltimeopt6, v.Voltimeopt7, v.Voltimeopt8, v.Voltimeopt9, v.Voltimeopt10, v.Voltimeopt11, v.Voltimeopt12, v.Voltimeopt13, v.Voltimeopt14, v.Voltimeopt15, v.Voltimeopt16, v.Voltimeopt17, v.Voltimeopt18, v.Voltimeopt19, v.Voltimeopt20, v.Voltimestype1, v.Voltimestype2, v.Voltimestype3, v.Voltimestype4, v.Voltimestype5, v.Voltimestype6, v.Voltimestype7, v.Voltimestype8, v.Voltimestype9, v.Voltimestype10, v.Voltimestype11, v.Voltimestype12, v.Voltimestype13, v.Voltimestype14, v.Voltimestype15, v.Voltimestype16, v.Voltimestype17, v.Voltimestype18, v.Voltimestype19, v.Voltimestype20, v.Volalt, v.Volsparec1, v.Volsparen1, v.Volspared1, v.EquinoxPrn, v.EquinoxLrn, v.EquinoxSec)
	_, err = db.Exec(sqlstr, v.Voluniquesys, v.Volchrgoption, v.Volfromto, v.Volchargetype, v.Volminsorpounds, v.Voldiscountorvol, v.Volvatrate, v.Volthresholdrat1, v.Volthresholdrat2, v.Volthresholdrat3, v.Volthresholdrat4, v.Volthresholdrat5, v.Volthresholdrat6, v.Volthresholdrat7, v.Volthresholdrat8, v.Volthresholdrat9, v.Volthreshper1, v.Volthreshper2, v.Volthreshper3, v.Volthreshper4, v.Volthreshper5, v.Volthreshper6, v.Volthreshper7, v.Volthreshper8, v.Volthreshper9, v.Volthreshper10, v.Voldayopt1, v.Voldayopt2, v.Voldayopt3, v.Voldayopt4, v.Voldayopt5, v.Voldayopt6, v.Voldayopt7, v.Voldayopt8, v.Voldayopt9, v.Voldayopt10, v.Voldayopt11, v.Voldayopt12, v.Voldayopt13, v.Voldayopt14, v.Voldayopt15, v.Voldayopt16, v.Voldayopt17, v.Voldayopt18, v.Voldayopt19, v.Voldayopt20, v.Voltimeopt1, v.Voltimeopt2, v.Voltimeopt3, v.Voltimeopt4, v.Voltimeopt5, v.Voltimeopt6, v.Voltimeopt7, v.Voltimeopt8, v.Voltimeopt9, v.Voltimeopt10, v.Voltimeopt11, v.Voltimeopt12, v.Voltimeopt13, v.Voltimeopt14, v.Voltimeopt15, v.Voltimeopt16, v.Voltimeopt17, v.Voltimeopt18, v.Voltimeopt19, v.Voltimeopt20, v.Voltimestype1, v.Voltimestype2, v.Voltimestype3, v.Voltimestype4, v.Voltimestype5, v.Voltimestype6, v.Voltimestype7, v.Voltimestype8, v.Voltimestype9, v.Voltimestype10, v.Voltimestype11, v.Voltimestype12, v.Voltimestype13, v.Voltimestype14, v.Voltimestype15, v.Voltimestype16, v.Voltimestype17, v.Voltimestype18, v.Voltimestype19, v.Voltimestype20, v.Volalt, v.Volsparec1, v.Volsparen1, v.Volspared1, v.EquinoxPrn, v.EquinoxLrn, v.EquinoxSec)
	if err != nil {
		return err
	}

	// set existence
	v._exists = true

	return nil
}

// Delete deletes the Volume from the database.
func (v *Volume) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !v._exists {
		return nil
	}

	// if deleted, bail
	if v._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM equinox.volume WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, v.EquinoxLrn)
	_, err = db.Exec(sqlstr, v.EquinoxLrn)
	if err != nil {
		return err
	}

	// set deleted
	v._deleted = true

	return nil
}

// VolumeByEquinoxLrn retrieves a row from 'equinox.volume' as a Volume.
//
// Generated from index 'volume_pkey'.
func VolumeByEquinoxLrn(db XODB, equinoxLrn int64) (*Volume, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`voluniquesys, volchrgoption, volfromto, volchargetype, volminsorpounds, voldiscountorvol, volvatrate, volthresholdrat1, volthresholdrat2, volthresholdrat3, volthresholdrat4, volthresholdrat5, volthresholdrat6, volthresholdrat7, volthresholdrat8, volthresholdrat9, volthreshper1, volthreshper2, volthreshper3, volthreshper4, volthreshper5, volthreshper6, volthreshper7, volthreshper8, volthreshper9, volthreshper10, voldayopt1, voldayopt2, voldayopt3, voldayopt4, voldayopt5, voldayopt6, voldayopt7, voldayopt8, voldayopt9, voldayopt10, voldayopt11, voldayopt12, voldayopt13, voldayopt14, voldayopt15, voldayopt16, voldayopt17, voldayopt18, voldayopt19, voldayopt20, voltimeopt1, voltimeopt2, voltimeopt3, voltimeopt4, voltimeopt5, voltimeopt6, voltimeopt7, voltimeopt8, voltimeopt9, voltimeopt10, voltimeopt11, voltimeopt12, voltimeopt13, voltimeopt14, voltimeopt15, voltimeopt16, voltimeopt17, voltimeopt18, voltimeopt19, voltimeopt20, voltimestype1, voltimestype2, voltimestype3, voltimestype4, voltimestype5, voltimestype6, voltimestype7, voltimestype8, voltimestype9, voltimestype10, voltimestype11, voltimestype12, voltimestype13, voltimestype14, voltimestype15, voltimestype16, voltimestype17, voltimestype18, voltimestype19, voltimestype20, volalt, volsparec1, volsparen1, volspared1, equinox_prn, equinox_lrn, equinox_sec ` +
		`FROM equinox.volume ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	v := Volume{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&v.Voluniquesys, &v.Volchrgoption, &v.Volfromto, &v.Volchargetype, &v.Volminsorpounds, &v.Voldiscountorvol, &v.Volvatrate, &v.Volthresholdrat1, &v.Volthresholdrat2, &v.Volthresholdrat3, &v.Volthresholdrat4, &v.Volthresholdrat5, &v.Volthresholdrat6, &v.Volthresholdrat7, &v.Volthresholdrat8, &v.Volthresholdrat9, &v.Volthreshper1, &v.Volthreshper2, &v.Volthreshper3, &v.Volthreshper4, &v.Volthreshper5, &v.Volthreshper6, &v.Volthreshper7, &v.Volthreshper8, &v.Volthreshper9, &v.Volthreshper10, &v.Voldayopt1, &v.Voldayopt2, &v.Voldayopt3, &v.Voldayopt4, &v.Voldayopt5, &v.Voldayopt6, &v.Voldayopt7, &v.Voldayopt8, &v.Voldayopt9, &v.Voldayopt10, &v.Voldayopt11, &v.Voldayopt12, &v.Voldayopt13, &v.Voldayopt14, &v.Voldayopt15, &v.Voldayopt16, &v.Voldayopt17, &v.Voldayopt18, &v.Voldayopt19, &v.Voldayopt20, &v.Voltimeopt1, &v.Voltimeopt2, &v.Voltimeopt3, &v.Voltimeopt4, &v.Voltimeopt5, &v.Voltimeopt6, &v.Voltimeopt7, &v.Voltimeopt8, &v.Voltimeopt9, &v.Voltimeopt10, &v.Voltimeopt11, &v.Voltimeopt12, &v.Voltimeopt13, &v.Voltimeopt14, &v.Voltimeopt15, &v.Voltimeopt16, &v.Voltimeopt17, &v.Voltimeopt18, &v.Voltimeopt19, &v.Voltimeopt20, &v.Voltimestype1, &v.Voltimestype2, &v.Voltimestype3, &v.Voltimestype4, &v.Voltimestype5, &v.Voltimestype6, &v.Voltimestype7, &v.Voltimestype8, &v.Voltimestype9, &v.Voltimestype10, &v.Voltimestype11, &v.Voltimestype12, &v.Voltimestype13, &v.Voltimestype14, &v.Voltimestype15, &v.Voltimestype16, &v.Voltimestype17, &v.Voltimestype18, &v.Voltimestype19, &v.Voltimestype20, &v.Volalt, &v.Volsparec1, &v.Volsparen1, &v.Volspared1, &v.EquinoxPrn, &v.EquinoxLrn, &v.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &v, nil
}