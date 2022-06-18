package helpers

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfErr(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfErr(errorCommit)
	}
}
