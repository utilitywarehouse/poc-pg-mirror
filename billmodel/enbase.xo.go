// Package billmodel contains the types for schema 'equinox'.
package billmodel

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"

	"github.com/lib/pq"
)

// Enbase represents a row from 'equinox.enbase'.
type Enbase struct {
	Endate           pq.NullTime     `json:"endate"`           // endate
	Enregion         sql.NullInt64   `json:"enregion"`         // enregion
	Ensedd           sql.NullInt64   `json:"ensedd"`           // ensedd
	Ensecash         sql.NullInt64   `json:"ensecash"`         // ensecash
	Ensepp           sql.NullInt64   `json:"ensepp"`           // ensepp
	Enhodd           sql.NullInt64   `json:"enhodd"`           // enhodd
	Enhocash         sql.NullInt64   `json:"enhocash"`         // enhocash
	Enhopp           sql.NullInt64   `json:"enhopp"`           // enhopp
	Engodd           sql.NullInt64   `json:"engodd"`           // engodd
	Engocash         sql.NullInt64   `json:"engocash"`         // engocash
	Engopp           sql.NullInt64   `json:"engopp"`           // engopp
	Engohdd          sql.NullInt64   `json:"engohdd"`          // engohdd
	Engohcash        sql.NullInt64   `json:"engohcash"`        // engohcash
	Engohpp          sql.NullInt64   `json:"engohpp"`          // engohpp
	Enseddaq         sql.NullFloat64 `json:"enseddaq"`         // enseddaq
	Ensecashaq       sql.NullFloat64 `json:"ensecashaq"`       // ensecashaq
	Enseppaq         sql.NullFloat64 `json:"enseppaq"`         // enseppaq
	Enhoddaq         sql.NullFloat64 `json:"enhoddaq"`         // enhoddaq
	Enhocashaq       sql.NullFloat64 `json:"enhocashaq"`       // enhocashaq
	Enhoppaq         sql.NullFloat64 `json:"enhoppaq"`         // enhoppaq
	Engoddaq         sql.NullFloat64 `json:"engoddaq"`         // engoddaq
	Engocashaq       sql.NullFloat64 `json:"engocashaq"`       // engocashaq
	Engoppaq         sql.NullFloat64 `json:"engoppaq"`         // engoppaq
	Engohddaq        sql.NullFloat64 `json:"engohddaq"`        // engohddaq
	Engohcashaq      sql.NullFloat64 `json:"engohcashaq"`      // engohcashaq
	Engohppaq        sql.NullFloat64 `json:"engohppaq"`        // engohppaq
	Enseddelec       sql.NullInt64   `json:"enseddelec"`       // enseddelec
	Enseddeleceac    sql.NullFloat64 `json:"enseddeleceac"`    // enseddeleceac
	Ensedde7         sql.NullInt64   `json:"ensedde7"`         // ensedde7
	Ensedde7deac     sql.NullFloat64 `json:"ensedde7deac"`     // ensedde7deac
	Ensedde7neac     sql.NullFloat64 `json:"ensedde7neac"`     // ensedde7neac
	Ensecashelec     sql.NullInt64   `json:"ensecashelec"`     // ensecashelec
	Ensecasheleceac  sql.NullFloat64 `json:"ensecasheleceac"`  // ensecasheleceac
	Ensecashe7       sql.NullInt64   `json:"ensecashe7"`       // ensecashe7
	Ensecashe7deac   sql.NullFloat64 `json:"ensecashe7deac"`   // ensecashe7deac
	Ensecashe7neac   sql.NullFloat64 `json:"ensecashe7neac"`   // ensecashe7neac
	Enseppelec       sql.NullInt64   `json:"enseppelec"`       // enseppelec
	Enseppeleceac    sql.NullFloat64 `json:"enseppeleceac"`    // enseppeleceac
	Enseppe7         sql.NullInt64   `json:"enseppe7"`         // enseppe7
	Enseppe7deac     sql.NullFloat64 `json:"enseppe7deac"`     // enseppe7deac
	Enseppe7neac     sql.NullFloat64 `json:"enseppe7neac"`     // enseppe7neac
	Enhoddelec       sql.NullInt64   `json:"enhoddelec"`       // enhoddelec
	Enhoddeleceac    sql.NullFloat64 `json:"enhoddeleceac"`    // enhoddeleceac
	Enhodde7         sql.NullInt64   `json:"enhodde7"`         // enhodde7
	Enhodde7deac     sql.NullFloat64 `json:"enhodde7deac"`     // enhodde7deac
	Enhodde7neac     sql.NullFloat64 `json:"enhodde7neac"`     // enhodde7neac
	Enhocashelec     sql.NullInt64   `json:"enhocashelec"`     // enhocashelec
	Enhocasheleceac  sql.NullFloat64 `json:"enhocasheleceac"`  // enhocasheleceac
	Enhocashe7       sql.NullInt64   `json:"enhocashe7"`       // enhocashe7
	Enhocashe7deac   sql.NullFloat64 `json:"enhocashe7deac"`   // enhocashe7deac
	Enhocashe7neac   sql.NullFloat64 `json:"enhocashe7neac"`   // enhocashe7neac
	Enhoppelec       sql.NullInt64   `json:"enhoppelec"`       // enhoppelec
	Enhoppeleceac    sql.NullFloat64 `json:"enhoppeleceac"`    // enhoppeleceac
	Enhoppe7         sql.NullInt64   `json:"enhoppe7"`         // enhoppe7
	Enhoppe7deac     sql.NullFloat64 `json:"enhoppe7deac"`     // enhoppe7deac
	Enhoppe7neac     sql.NullFloat64 `json:"enhoppe7neac"`     // enhoppe7neac
	Engoddelec       sql.NullInt64   `json:"engoddelec"`       // engoddelec
	Engohddelec      sql.NullInt64   `json:"engohddelec"`      // engohddelec
	Engoddeleceac    sql.NullFloat64 `json:"engoddeleceac"`    // engoddeleceac
	Engohddeleceac   sql.NullFloat64 `json:"engohddeleceac"`   // engohddeleceac
	Engodde7         sql.NullInt64   `json:"engodde7"`         // engodde7
	Engohdde7        sql.NullInt64   `json:"engohdde7"`        // engohdde7
	Engodde7deac     sql.NullFloat64 `json:"engodde7deac"`     // engodde7deac
	Engohdde7deac    sql.NullFloat64 `json:"engohdde7deac"`    // engohdde7deac
	Engodde7neac     sql.NullFloat64 `json:"engodde7neac"`     // engodde7neac
	Engohdde7neac    sql.NullFloat64 `json:"engohdde7neac"`    // engohdde7neac
	Engocashelec     sql.NullInt64   `json:"engocashelec"`     // engocashelec
	Engohcashelec    sql.NullInt64   `json:"engohcashelec"`    // engohcashelec
	Engocasheleceac  sql.NullFloat64 `json:"engocasheleceac"`  // engocasheleceac
	Engohcasheleceac sql.NullFloat64 `json:"engohcasheleceac"` // engohcasheleceac
	Engocashe7       sql.NullInt64   `json:"engocashe7"`       // engocashe7
	Engohcashe7      sql.NullInt64   `json:"engohcashe7"`      // engohcashe7
	Engocashe7deac   sql.NullFloat64 `json:"engocashe7deac"`   // engocashe7deac
	Engohcashe7deac  sql.NullFloat64 `json:"engohcashe7deac"`  // engohcashe7deac
	Engocashe7neac   sql.NullFloat64 `json:"engocashe7neac"`   // engocashe7neac
	Engohcashe7neac  sql.NullFloat64 `json:"engohcashe7neac"`  // engohcashe7neac
	Engoppelec       sql.NullInt64   `json:"engoppelec"`       // engoppelec
	Engohppelec      sql.NullInt64   `json:"engohppelec"`      // engohppelec
	Engoppeleceac    sql.NullFloat64 `json:"engoppeleceac"`    // engoppeleceac
	Engohppeleceac   sql.NullFloat64 `json:"engohppeleceac"`   // engohppeleceac
	Engoppe7         sql.NullInt64   `json:"engoppe7"`         // engoppe7
	Engohppe7        sql.NullInt64   `json:"engohppe7"`        // engohppe7
	Engoppe7deac     sql.NullFloat64 `json:"engoppe7deac"`     // engoppe7deac
	Engohppe7deac    sql.NullFloat64 `json:"engohppe7deac"`    // engohppe7deac
	Engoppe7neac     sql.NullFloat64 `json:"engoppe7neac"`     // engoppe7neac
	Engohppe7neac    sql.NullFloat64 `json:"engohppe7neac"`    // engohppe7neac
	Entotaldualfdisc sql.NullInt64   `json:"entotaldualfdisc"` // entotaldualfdisc
	Engoldcrossover  sql.NullInt64   `json:"engoldcrossover"`  // engoldcrossover
	Eneleccrossover  sql.NullInt64   `json:"eneleccrossover"`  // eneleccrossover
	Ene7crossover    sql.NullInt64   `json:"ene7crossover"`    // ene7crossover
	Gogasddsplit     sql.NullString  `json:"gogasddsplit"`     // gogasddsplit
	Goelecddsplit    sql.NullString  `json:"goelecddsplit"`    // goelecddsplit
	Goe7ddsplitd     sql.NullString  `json:"goe7ddsplitd"`     // goe7ddsplitd
	Goe7ddsplitn     sql.NullString  `json:"goe7ddsplitn"`     // goe7ddsplitn
	Gogasddsplitaq   sql.NullString  `json:"gogasddsplitaq"`   // gogasddsplitaq
	Goelecddspliteac sql.NullString  `json:"goelecddspliteac"` // goelecddspliteac
	Goe7ddspliteac   sql.NullString  `json:"goe7ddspliteac"`   // goe7ddspliteac
	Goe7ddspliteacn  sql.NullString  `json:"goe7ddspliteacn"`  // goe7ddspliteacn
	Enavbill         sql.NullFloat64 `json:"enavbill"`         // enavbill
	Enavebill        sql.NullFloat64 `json:"enavebill"`        // enavebill
	Enave7bill       sql.NullFloat64 `json:"enave7bill"`       // enave7bill
	Endfdiscount     sql.NullFloat64 `json:"endfdiscount"`     // endfdiscount
	Endfnos          sql.NullInt64   `json:"endfnos"`          // endfnos
	Enn1             sql.NullInt64   `json:"enn1"`             // enn1
	Enn2             sql.NullInt64   `json:"enn2"`             // enn2
	EquinoxLrn       int64           `json:"equinox_lrn"`      // equinox_lrn
	EquinoxSec       sql.NullInt64   `json:"equinox_sec"`      // equinox_sec
}

// EnbaseByEquinoxLrn retrieves a row from 'equinox.enbase' as a Enbase.
//
// Generated from index 'enbase_pkey'.
func EnbaseByEquinoxLrn(db XODB, equinoxLrn int64) (*Enbase, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`endate, enregion, ensedd, ensecash, ensepp, enhodd, enhocash, enhopp, engodd, engocash, engopp, engohdd, engohcash, engohpp, enseddaq, ensecashaq, enseppaq, enhoddaq, enhocashaq, enhoppaq, engoddaq, engocashaq, engoppaq, engohddaq, engohcashaq, engohppaq, enseddelec, enseddeleceac, ensedde7, ensedde7deac, ensedde7neac, ensecashelec, ensecasheleceac, ensecashe7, ensecashe7deac, ensecashe7neac, enseppelec, enseppeleceac, enseppe7, enseppe7deac, enseppe7neac, enhoddelec, enhoddeleceac, enhodde7, enhodde7deac, enhodde7neac, enhocashelec, enhocasheleceac, enhocashe7, enhocashe7deac, enhocashe7neac, enhoppelec, enhoppeleceac, enhoppe7, enhoppe7deac, enhoppe7neac, engoddelec, engohddelec, engoddeleceac, engohddeleceac, engodde7, engohdde7, engodde7deac, engohdde7deac, engodde7neac, engohdde7neac, engocashelec, engohcashelec, engocasheleceac, engohcasheleceac, engocashe7, engohcashe7, engocashe7deac, engohcashe7deac, engocashe7neac, engohcashe7neac, engoppelec, engohppelec, engoppeleceac, engohppeleceac, engoppe7, engohppe7, engoppe7deac, engohppe7deac, engoppe7neac, engohppe7neac, entotaldualfdisc, engoldcrossover, eneleccrossover, ene7crossover, gogasddsplit, goelecddsplit, goe7ddsplitd, goe7ddsplitn, gogasddsplitaq, goelecddspliteac, goe7ddspliteac, goe7ddspliteacn, enavbill, enavebill, enave7bill, endfdiscount, endfnos, enn1, enn2, equinox_lrn, equinox_sec ` +
		`FROM equinox.enbase ` +
		`WHERE equinox_lrn = $1`

	// run query
	XOLog(sqlstr, equinoxLrn)
	e := Enbase{}

	err = db.QueryRow(sqlstr, equinoxLrn).Scan(&e.Endate, &e.Enregion, &e.Ensedd, &e.Ensecash, &e.Ensepp, &e.Enhodd, &e.Enhocash, &e.Enhopp, &e.Engodd, &e.Engocash, &e.Engopp, &e.Engohdd, &e.Engohcash, &e.Engohpp, &e.Enseddaq, &e.Ensecashaq, &e.Enseppaq, &e.Enhoddaq, &e.Enhocashaq, &e.Enhoppaq, &e.Engoddaq, &e.Engocashaq, &e.Engoppaq, &e.Engohddaq, &e.Engohcashaq, &e.Engohppaq, &e.Enseddelec, &e.Enseddeleceac, &e.Ensedde7, &e.Ensedde7deac, &e.Ensedde7neac, &e.Ensecashelec, &e.Ensecasheleceac, &e.Ensecashe7, &e.Ensecashe7deac, &e.Ensecashe7neac, &e.Enseppelec, &e.Enseppeleceac, &e.Enseppe7, &e.Enseppe7deac, &e.Enseppe7neac, &e.Enhoddelec, &e.Enhoddeleceac, &e.Enhodde7, &e.Enhodde7deac, &e.Enhodde7neac, &e.Enhocashelec, &e.Enhocasheleceac, &e.Enhocashe7, &e.Enhocashe7deac, &e.Enhocashe7neac, &e.Enhoppelec, &e.Enhoppeleceac, &e.Enhoppe7, &e.Enhoppe7deac, &e.Enhoppe7neac, &e.Engoddelec, &e.Engohddelec, &e.Engoddeleceac, &e.Engohddeleceac, &e.Engodde7, &e.Engohdde7, &e.Engodde7deac, &e.Engohdde7deac, &e.Engodde7neac, &e.Engohdde7neac, &e.Engocashelec, &e.Engohcashelec, &e.Engocasheleceac, &e.Engohcasheleceac, &e.Engocashe7, &e.Engohcashe7, &e.Engocashe7deac, &e.Engohcashe7deac, &e.Engocashe7neac, &e.Engohcashe7neac, &e.Engoppelec, &e.Engohppelec, &e.Engoppeleceac, &e.Engohppeleceac, &e.Engoppe7, &e.Engohppe7, &e.Engoppe7deac, &e.Engohppe7deac, &e.Engoppe7neac, &e.Engohppe7neac, &e.Entotaldualfdisc, &e.Engoldcrossover, &e.Eneleccrossover, &e.Ene7crossover, &e.Gogasddsplit, &e.Goelecddsplit, &e.Goe7ddsplitd, &e.Goe7ddsplitn, &e.Gogasddsplitaq, &e.Goelecddspliteac, &e.Goe7ddspliteac, &e.Goe7ddspliteacn, &e.Enavbill, &e.Enavebill, &e.Enave7bill, &e.Endfdiscount, &e.Endfnos, &e.Enn1, &e.Enn2, &e.EquinoxLrn, &e.EquinoxSec)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
