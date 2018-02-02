package fs

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/ezbuy/redis-orm/parser"
	"github.com/ezbuy/redis-orm/tpl"
)

var RedisOrmTemplate *template.Template

func generate_templates(obj *parser.MetaObject) []string {
	tpls := []string{}
	if obj.Relation != nil {
		tpls = append(tpls, "relation")
	} else {
		if obj.DbView != "" {
			tpls = append(tpls, "view")
		}

		if obj.DbTable != "" {
			tpls = append(tpls, "object")
		}

		if obj.DbSource() == "" {
			tpls = append(tpls, "query")
		}

	}
	return tpls
}

func ExecuteMetaObjectCodeTemplate(output string, obj *parser.MetaObject) error {
	for _, tpl := range generate_templates(obj) {
		filename := filepath.Join(output, strings.Join([]string{"gen", tpl, camel2sep(obj.Name, "."), "go"}, "."))
		fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer fd.Close()
		if err := RedisOrmTemplate.ExecuteTemplate(fd, tpl, obj); err != nil {
			return err
		}

		oscmd := exec.Command("gofmt", "-w", filename)
		oscmd.Run()
	}
	return nil
}

func ExecuteMetaObjectScriptTemplate(output string, driver string, obj *parser.MetaObject) error {
	filename := filepath.Join(output, strings.Join([]string{"gen", "script", driver, camel2sep(obj.Name, "."), "sql"}, "."))
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()
	if err := RedisOrmTemplate.ExecuteTemplate(fd, strings.Join([]string{"script", driver}, "."), obj); err != nil {
		return err
	}
	return nil
}

func ExecuteConfigTemplate(output, db string, packageName string) error {
	filename := filepath.Join(output, strings.Join([]string{"gen", "conf", db, "go"}, "."))
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()
	if err := RedisOrmTemplate.ExecuteTemplate(fd, strings.Join([]string{"conf", db}, "."), map[string]interface{}{
		"GoPackage": packageName,
	}); err != nil {
		return err
	}

	oscmd := exec.Command("gofmt", "-w", filename)
	oscmd.Run()
	return nil
}

func init() {
	funcMap := template.FuncMap{
		"add":        Add,
		"sub":        Sub,
		"divide":     Divide,
		"multiply":   Multiply,
		"camel2name": parser.Camel2Name,
		"camel2sep":  camel2sep,
	}

	RedisOrmTemplate = template.New("redis-orm").Funcs(funcMap)
	files := []string{
		"tpl/conf.elastic.gogo",
		"tpl/conf.mongo.gogo",
		"tpl/conf.mssql.gogo",
		"tpl/conf.mysql.gogo",
		"tpl/conf.orm.gogo",
		"tpl/conf.redis.gogo",
		"tpl/object.db.gogo",
		"tpl/object.db.read.gogo",
		"tpl/object.db.write.gogo",
		"tpl/object.db.query.gogo",
		"tpl/object.elastic.gogo",
		"tpl/object.functions.gogo",
		"tpl/object.gogo",
		"tpl/object.index.gogo",
		"tpl/object.mongo.gogo",
		"tpl/object.primary.key.gogo",
		"tpl/object.range.gogo",
		"tpl/object.redis.gogo",
		"tpl/object.redis.manager.gogo",
		"tpl/object.redis.pipeline.gogo",
		"tpl/object.redis.read.gogo",
		"tpl/object.redis.sync.gogo",
		"tpl/object.redis.write.gogo",
		"tpl/object.relation.gogo",
		"tpl/object.unqiue.gogo",
		"tpl/query.gogo",
		"tpl/relation.db.read.gogo",
		"tpl/relation.functions.gogo",
		"tpl/relation.geo.gogo",
		"tpl/relation.geo.sync.gogo",
		"tpl/relation.gogo",
		"tpl/relation.list.gogo",
		"tpl/relation.list.sync.gogo",
		"tpl/relation.manager.gogo",
		"tpl/relation.pair.gogo",
		"tpl/relation.pair.sync.gogo",
		"tpl/relation.pipeline.gogo",
		"tpl/relation.set.gogo",
		"tpl/relation.set.sync.gogo",
		"tpl/relation.zset.gogo",
		"tpl/relation.zset.sync.gogo",
		"tpl/script.mysql.sql",
		"tpl/view.gogo",
	}
	for _, fname := range files {
		data, err := tpl.Asset(fname)
		if err != nil {
			panic(err)
		}
		_, err = RedisOrmTemplate.Parse(string(data))
		if err != nil {
			fmt.Println(fname)
			panic(err)
		}
	}
}

func camel2sep(s string, sep string) string {
	nameBuf := bytes.NewBuffer(nil)
	for i := range s {
		n := rune(s[i]) // always ASCII?
		if unicode.IsUpper(n) {
			if i > 0 {
				nameBuf.WriteString(sep)
			}
			n = unicode.ToLower(n)
		}
		nameBuf.WriteRune(n)
	}
	return nameBuf.String()
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Divide(a, b int) int {
	return a / b
}

func Multiply(a, b int) int {
	return a * b
}
