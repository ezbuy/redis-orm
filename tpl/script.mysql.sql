{{- define "script.mysql"}}{{- $obj := . -}}
{{- if ne $obj.DbTable ""}}
DROP TABLE IF EXISTS `{{$obj.DbTable}}`;
CREATE TABLE `{{$obj.DbTable}}` (
	{{- range $i, $field := $obj.Fields}}
		{{- if eq (add $i 1) (len $obj.Fields)}}
	{{$field.SQLColumn "mysql"}}
		{{- else}}
	{{$field.SQLColumn "mysql"}}, 
		{{- end}}
	{{- end}}
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

{{range $i, $unique := $obj.Uniques}}
CREATE UNIQUE INDEX `{{$unique.Name | camel2name}}` ON `{{$obj.DbTable}}`(
	{{- range $i, $f := $unique.Fields -}}
		{{- if eq (add $i 1) (len $unique.Fields) -}}
			`{{- $f.Name | camel2name -}}`
		{{- else -}}
			`{{- $f.Name | camel2name -}}`,
		{{- end -}}
	{{- end -}}
);
{{- end}}

{{range $i, $index := $obj.Indexes}}
CREATE INDEX `{{$index.Name | camel2name}}` ON `{{$obj.DbTable}}`(
	{{- range $i, $f := $index.Fields -}}
		{{- if eq (add $i 1) (len $index.Fields) -}}
			`{{- $f.Name | camel2name -}}`
		{{- else -}}
			`{{- $f.Name | camel2name -}}`,
		{{- end -}}
	{{- end -}}
);
{{- end}}

{{range $i, $index := $obj.Ranges}}
CREATE INDEX `{{$index.Name | camel2name}}` ON `{{$obj.DbTable}}`(
	{{- range $i, $f := $index.Fields -}}
		{{- if eq (add $i 1) (len $index.Fields) -}}
			`{{- $f.Name | camel2name -}}`
		{{- else -}}
			`{{- $f.Name | camel2name -}}`,
		{{- end -}}
	{{- end -}}
);
{{- end}}
{{- end}}

{{- if ne $obj.DbView ""}}
DROP VIEW IF EXISTS `{{$obj.DbView}}`;
CREATE VIEW `{{$obj.DbView}}` AS {{$obj.ImportSQL}};
{{- end}}

{{end}}
