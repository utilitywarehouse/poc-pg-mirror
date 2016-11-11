{{- $short := (shortname .Name "err" "res" "sqlstr" "db" "XOLog") -}}
{{- $table := (schema .Schema .Table.TableName) -}}
{{- if .Comment -}}
// {{ .Comment }}
{{- else -}}
// {{ .Name }} represents a row from '{{ $table }}'.
{{- end }}
type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Name }} {{ retype .Type }} `json:"{{ .Col.ColumnName }}"` // {{ .Col.ColumnName }}
{{- end }}
}

func All{{ .Name }} (db XODB, callback func(x {{ .Name }}) bool) error {

	// sql query
	const sqlstr = `SELECT ` +
		`{{ colnames .Fields }} ` +
		`FROM {{ $table }} `

	q, err := db.Query(sqlstr)

	if err != nil {
		return err
	}
	defer q.Close()

	// load results
	for q.Next() {
		{{ $short }} := {{ .Name }}{}

		// scan
		err = q.Scan({{ fieldnames .Fields (print "&" $short) }})
		if err != nil {
			return err
		}
		if !callback({{ $short }}) {
			return nil
		}
	}

	return nil
}
