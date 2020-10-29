package command

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/minamijoyo/tfmigrate/config"
	"github.com/minamijoyo/tfmigrate/tfmigrate"
)

// FileRunner is a runner for a single migration file.
type FileRunner struct {
	// A path to migration file.
	filename string
	// A definition of migration.
	mc *tfmigrate.MigrationConfig
	// A migrator instance to be run.
	m tfmigrate.Migrator
}

// NewFileRunner returns a new FileRunner instance.
func NewFileRunner(filename string, option *tfmigrate.MigratorOption) (*FileRunner, error) {
	log.Printf("[INFO] [runner] load migration file: %s\n", filename)
	mc, err := loadMigrationFile(filename)
	if err != nil {
		return nil, err
	}

	m, err := mc.Migrator.NewMigrator(option)
	if err != nil {
		return nil, err
	}

	r := &FileRunner{
		filename: filename,
		mc:       mc,
		m:        m,
	}

	return r, nil
}

// loadMigrationFile is a helper function which reads and parses a migration file.
func loadMigrationFile(filename string) (*tfmigrate.MigrationConfig, error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config, err := config.ParseMigrationFile(filename, source)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// Plan plans a single migration.
func (r *FileRunner) Plan(ctx context.Context) error {
	return r.m.Plan(ctx)
}

// Apply applies a single migration..
func (r *FileRunner) Apply(ctx context.Context) error {
	return r.m.Apply(ctx)
}

// MigrationConfig returns an instance of migration.
// This is required for metadata stored in history
func (r *FileRunner) MigrationConfig() *tfmigrate.MigrationConfig {
	return r.mc
}
