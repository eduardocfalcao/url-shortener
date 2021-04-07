package migrations

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
	"time"
)

func Create(name string) error {
	version := time.Now().Format("20060102150405")

	in := struct {
		Version string
		Name    string
	}{
		Version: version,
		Name:    name,
	}

	var out bytes.Buffer

	t := template.Must(template.ParseFiles("./migrations/migration_template.tpl"))
	err := t.Execute(&out, in)
	if err != nil {
		return fmt.Errorf("Unable to execute template: %w", err)
	}

	f, err := os.Create(fmt.Sprintf("./migrations/%s_%s.go", version, name))
	if err != nil {
		return fmt.Errorf("Unable to create migration file: %w", err)
	}
	defer f.Close()

	if _, err := f.WriteString(out.String()); err != nil {
		return fmt.Errorf("Unable to write to migration file: %w", err)
	}

	fmt.Println("Generated new migration files...", f.Name())
	return nil
}
