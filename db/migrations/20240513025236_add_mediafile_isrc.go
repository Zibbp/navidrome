package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddMediafileIsrc, downAddMediafileIsrc)
}

func upAddMediafileIsrc(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='media_file' AND column_name='isrc') THEN
			ALTER TABLE media_file ADD COLUMN isrc varchar DEFAULT '';
	END IF;

	IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='media_file' AND column_name='upc') THEN
			ALTER TABLE media_file ADD COLUMN upc varchar DEFAULT '';
	END IF;

	IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='album' AND column_name='upc') THEN
			ALTER TABLE album ADD COLUMN upc varchar DEFAULT '';
	END IF;
END $$;
	`)
	return err
}

func downAddMediafileIsrc(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
alter table media_file
	drop isrc;
alter table media_file
	drop upc;
alter table album
	drop upc;
	`)
	return err
}
