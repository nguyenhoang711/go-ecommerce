-- name: GetUserByEmailSQLC :one
SELECT usr_email, usr_id FROM `ecom_user` WHERE usr_email = ? LIMIT 1;

-- name: UpdateUserStatusByUserId :exec
UPDATE `ecom_user`
SET usr_status = $2,
    usr_updated_at = $3
WHERE usr_id = $1;