package migration

import (
	"micro_blog/migration/contract"
	"micro_blog/migration/version"
)

var lastMigration contract.Migration = version.M202101252200Init{}

func GetMigrationRange() []contract.Migration {
	var migrationRange []contract.Migration
	var m = lastMigration
	migrationRange = []contract.Migration{m}
	for m.PreviousMigration() != nil {
		migrationRange = append(migrationRange, *m.PreviousMigration())
		m = *m.PreviousMigration()
	}
	reverse(migrationRange)
	return migrationRange
}

func reverse(migrations []contract.Migration) {
	for i, j := 0, len(migrations)-1; i < j; i, j = i+1, j-1 {
		migrations[i], migrations[j] = migrations[j], migrations[i]
	}
}
