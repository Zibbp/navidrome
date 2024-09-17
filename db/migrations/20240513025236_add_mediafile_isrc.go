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
-- Add isrc to media_file if it doesn't exist
SELECT CASE 
	WHEN NOT EXISTS (SELECT 1 FROM pragma_table_info('media_file') WHERE name = 'isrc') THEN
			ALTER TABLE media_file ADD COLUMN isrc TEXT DEFAULT '';
END;

-- Add upc to media_file if it doesn't exist
SELECT CASE 
	WHEN NOT EXISTS (SELECT 1 FROM pragma_table_info('media_file') WHERE name = 'upc') THEN
			ALTER TABLE media_file ADD COLUMN upc TEXT DEFAULT '';
END;

-- Add upc to album if it doesn't exist
SELECT CASE 
	WHEN NOT EXISTS (SELECT 1 FROM pragma_table_info('album') WHERE name = 'upc') THEN
			ALTER TABLE album ADD COLUMN upc TEXT DEFAULT '';
END;
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
