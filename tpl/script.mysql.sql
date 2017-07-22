{{- define "script.mysql"}}{{- $obj := . -}}
{{- if ne $obj.DbTable ""}}
CREATE TABLE `{{$obj.DbTable}}` (
	{{- range $i, $field := $obj.Fields}}
	{{$field.SQLColumn "mysql"}},
	{{- end}}
	{{$obj.PrimaryKey.SQLColumn "mysql"}}
	{{- range $i, $unique := $obj.Uniques}}
	{{- if not $unique.HasPrimaryKey}},
	UNIQUE KEY `uniq_{{$unique.Name | camel2name}}` (
		{{- range $i, $f := $unique.Fields -}}
			{{- if eq (add $i 1) (len $unique.Fields) -}}
				`{{- $f.Name | camel2name -}}`
			{{- else -}}
				`{{- $f.Name | camel2name -}}`,
			{{- end -}}
		{{- end -}}
	)
	{{- end}}
	{{- end}}
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '{{$obj.Comment}}';

{{- range $i, $index := $obj.Indexes}}
{{- if not $index.HasPrimaryKey}}
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

{{- range $i, $index := $obj.Ranges}}
{{- if not $index.HasPrimaryKey}}
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
{{- end}}

{{- if ne $obj.DbView ""}}
DROP VIEW IF EXISTS `{{$obj.DbView}}`;
CREATE VIEW `{{$obj.DbView}}` AS {{$obj.ImportSQL}};
{{- end}}

{{end}}
