package util

import (
	"context"
	"fmt"

	"github.com/kiki-ki/lesson-ent/ent"
)

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		err = rollback(tx, err)
		return err
	}
	if cerr := tx.Commit(); cerr != nil {
		return fmt.Errorf("%w: committing transaction: %v", err, cerr)
	}
	return nil
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
