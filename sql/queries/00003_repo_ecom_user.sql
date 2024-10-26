-- name: GetUserByEmailSQLC :one
SELECT user_email, user_id FROM `pre_go_acc_user_9999` WHERE user_email = ? LIMIT 1;

-- name: UpdateUserStatusByUserId :exec
UPDATE `pre_go_acc_user_9999`
SET user_state = $2,
    updated_at = $3
WHERE user_id = $1;