// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 00003_repo_ecom_user.sql

package repos

import (
	"context"
	"database/sql"
)

const getUserByEmailSQLC = `-- name: GetUserByEmailSQLC :one
SELECT user_email, user_id FROM ` + "`" + `pre_go_acc_user_9999` + "`" + ` WHERE user_email = ? LIMIT 1
`

type GetUserByEmailSQLCRow struct {
	UserEmail sql.NullString
	UserID    uint64
}

func (q *Queries) GetUserByEmailSQLC(ctx context.Context, userEmail sql.NullString) (GetUserByEmailSQLCRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmailSQLC, userEmail)
	var i GetUserByEmailSQLCRow
	err := row.Scan(&i.UserEmail, &i.UserID)
	return i, err
}

const updateUserStatusByUserId = `-- name: UpdateUserStatusByUserId :exec
UPDATE ` + "`" + `pre_go_acc_user_9999` + "`" + `
SET user_state = $2,
    updated_at = $3
WHERE user_id = $1
`

func (q *Queries) UpdateUserStatusByUserId(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateUserStatusByUserId)
	return err
}
