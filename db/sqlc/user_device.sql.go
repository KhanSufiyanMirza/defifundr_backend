// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: user_device.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const countActiveDeviceTokensForUser = `-- name: CountActiveDeviceTokensForUser :one
SELECT COUNT(*) 
FROM user_device_tokens
WHERE 
    user_id = $1 
    AND is_active = true 
    AND (expires_at IS NULL OR expires_at > NOW())
`

func (q *Queries) CountActiveDeviceTokensForUser(ctx context.Context, userID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countActiveDeviceTokensForUser, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUserDeviceToken = `-- name: CreateUserDeviceToken :one
INSERT INTO user_device_tokens (
    id,
    user_id,
    device_token,
    platform,
    device_type,
    device_model,
    os_name,
    os_version,
    push_notification_token,
    is_active,
    is_verified,
    app_version,
    client_ip,
    expires_at
) VALUES (
    $1, 
    $2, 
    $3, 
    COALESCE($4, 'unknown'),
    COALESCE($5, 'unknown'),
    COALESCE($6, 'unknown'),
    COALESCE($7, 'unknown'),
    COALESCE($8, 'unknown'),
    $9,
    COALESCE($10, true),
    COALESCE($11, false),
    $12,
    $13,
    $14
) RETURNING id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked
`

type CreateUserDeviceTokenParams struct {
	ID                    uuid.UUID          `json:"id"`
	UserID                uuid.UUID          `json:"user_id"`
	DeviceToken           string             `json:"device_token"`
	Column4               interface{}        `json:"column_4"`
	Column5               interface{}        `json:"column_5"`
	Column6               interface{}        `json:"column_6"`
	Column7               interface{}        `json:"column_7"`
	Column8               interface{}        `json:"column_8"`
	PushNotificationToken pgtype.Text        `json:"push_notification_token"`
	Column10              interface{}        `json:"column_10"`
	Column11              interface{}        `json:"column_11"`
	AppVersion            pgtype.Text        `json:"app_version"`
	ClientIp              pgtype.Text        `json:"client_ip"`
	ExpiresAt             pgtype.Timestamptz `json:"expires_at"`
}

func (q *Queries) CreateUserDeviceToken(ctx context.Context, arg CreateUserDeviceTokenParams) (UserDeviceTokens, error) {
	row := q.db.QueryRow(ctx, createUserDeviceToken,
		arg.ID,
		arg.UserID,
		arg.DeviceToken,
		arg.Column4,
		arg.Column5,
		arg.Column6,
		arg.Column7,
		arg.Column8,
		arg.PushNotificationToken,
		arg.Column10,
		arg.Column11,
		arg.AppVersion,
		arg.ClientIp,
		arg.ExpiresAt,
	)
	var i UserDeviceTokens
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceToken,
		&i.Platform,
		&i.DeviceType,
		&i.DeviceModel,
		&i.OsName,
		&i.OsVersion,
		&i.PushNotificationToken,
		&i.IsActive,
		&i.IsVerified,
		&i.LastUsedAt,
		&i.FirstRegisteredAt,
		&i.AppVersion,
		&i.ClientIp,
		&i.ExpiresAt,
		&i.IsRevoked,
	)
	return i, err
}

const deleteExpiredDeviceTokens = `-- name: DeleteExpiredDeviceTokens :exec
DELETE FROM user_device_tokens
WHERE 
    expires_at < NOW() 
    AND is_active = false
`

func (q *Queries) DeleteExpiredDeviceTokens(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteExpiredDeviceTokens)
	return err
}

const getActiveDeviceTokensForUser = `-- name: GetActiveDeviceTokensForUser :many
SELECT id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked FROM user_device_tokens
WHERE user_id = $1 
  AND is_active = true 
  AND (expires_at IS NULL OR expires_at > NOW())
ORDER BY first_registered_at DESC
`

func (q *Queries) GetActiveDeviceTokensForUser(ctx context.Context, userID uuid.UUID) ([]UserDeviceTokens, error) {
	rows, err := q.db.Query(ctx, getActiveDeviceTokensForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserDeviceTokens{}
	for rows.Next() {
		var i UserDeviceTokens
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.DeviceToken,
			&i.Platform,
			&i.DeviceType,
			&i.DeviceModel,
			&i.OsName,
			&i.OsVersion,
			&i.PushNotificationToken,
			&i.IsActive,
			&i.IsVerified,
			&i.LastUsedAt,
			&i.FirstRegisteredAt,
			&i.AppVersion,
			&i.ClientIp,
			&i.ExpiresAt,
			&i.IsRevoked,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDeviceTokensByPlatform = `-- name: GetDeviceTokensByPlatform :many
SELECT id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked FROM user_device_tokens
WHERE 
    user_id = $1 
    AND platform = $2 
    AND is_active = true 
    AND (expires_at IS NULL OR expires_at > NOW())
ORDER BY first_registered_at DESC
`

type GetDeviceTokensByPlatformParams struct {
	UserID   uuid.UUID `json:"user_id"`
	Platform string    `json:"platform"`
}

func (q *Queries) GetDeviceTokensByPlatform(ctx context.Context, arg GetDeviceTokensByPlatformParams) ([]UserDeviceTokens, error) {
	rows, err := q.db.Query(ctx, getDeviceTokensByPlatform, arg.UserID, arg.Platform)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserDeviceTokens{}
	for rows.Next() {
		var i UserDeviceTokens
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.DeviceToken,
			&i.Platform,
			&i.DeviceType,
			&i.DeviceModel,
			&i.OsName,
			&i.OsVersion,
			&i.PushNotificationToken,
			&i.IsActive,
			&i.IsVerified,
			&i.LastUsedAt,
			&i.FirstRegisteredAt,
			&i.AppVersion,
			&i.ClientIp,
			&i.ExpiresAt,
			&i.IsRevoked,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserDeviceTokenByDeviceToken = `-- name: GetUserDeviceTokenByDeviceToken :one
SELECT id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked FROM user_device_tokens
WHERE device_token = $1
`

func (q *Queries) GetUserDeviceTokenByDeviceToken(ctx context.Context, deviceToken string) (UserDeviceTokens, error) {
	row := q.db.QueryRow(ctx, getUserDeviceTokenByDeviceToken, deviceToken)
	var i UserDeviceTokens
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceToken,
		&i.Platform,
		&i.DeviceType,
		&i.DeviceModel,
		&i.OsName,
		&i.OsVersion,
		&i.PushNotificationToken,
		&i.IsActive,
		&i.IsVerified,
		&i.LastUsedAt,
		&i.FirstRegisteredAt,
		&i.AppVersion,
		&i.ClientIp,
		&i.ExpiresAt,
		&i.IsRevoked,
	)
	return i, err
}

const getUserDeviceTokenByID = `-- name: GetUserDeviceTokenByID :one
SELECT id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked FROM user_device_tokens
WHERE id = $1
`

func (q *Queries) GetUserDeviceTokenByID(ctx context.Context, id uuid.UUID) (UserDeviceTokens, error) {
	row := q.db.QueryRow(ctx, getUserDeviceTokenByID, id)
	var i UserDeviceTokens
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceToken,
		&i.Platform,
		&i.DeviceType,
		&i.DeviceModel,
		&i.OsName,
		&i.OsVersion,
		&i.PushNotificationToken,
		&i.IsActive,
		&i.IsVerified,
		&i.LastUsedAt,
		&i.FirstRegisteredAt,
		&i.AppVersion,
		&i.ClientIp,
		&i.ExpiresAt,
		&i.IsRevoked,
	)
	return i, err
}

const revokeDeviceToken = `-- name: RevokeDeviceToken :one
UPDATE user_device_tokens
SET 
    is_active = false,
    is_revoked = true
WHERE id = $1
RETURNING id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked
`

func (q *Queries) RevokeDeviceToken(ctx context.Context, id uuid.UUID) (UserDeviceTokens, error) {
	row := q.db.QueryRow(ctx, revokeDeviceToken, id)
	var i UserDeviceTokens
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceToken,
		&i.Platform,
		&i.DeviceType,
		&i.DeviceModel,
		&i.OsName,
		&i.OsVersion,
		&i.PushNotificationToken,
		&i.IsActive,
		&i.IsVerified,
		&i.LastUsedAt,
		&i.FirstRegisteredAt,
		&i.AppVersion,
		&i.ClientIp,
		&i.ExpiresAt,
		&i.IsRevoked,
	)
	return i, err
}

const searchDeviceTokens = `-- name: SearchDeviceTokens :many
SELECT id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked FROM user_device_tokens
WHERE 
    user_id = $1 
    AND (
        COALESCE(device_type, '') ILIKE '%' || $2 || '%' 
        OR COALESCE(device_model, '') ILIKE '%' || $2 || '%' 
        OR platform ILIKE '%' || $2 || '%'
    )
ORDER BY first_registered_at DESC
LIMIT $3 OFFSET $4
`

type SearchDeviceTokensParams struct {
	UserID  uuid.UUID   `json:"user_id"`
	Column2 pgtype.Text `json:"column_2"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

func (q *Queries) SearchDeviceTokens(ctx context.Context, arg SearchDeviceTokensParams) ([]UserDeviceTokens, error) {
	rows, err := q.db.Query(ctx, searchDeviceTokens,
		arg.UserID,
		arg.Column2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserDeviceTokens{}
	for rows.Next() {
		var i UserDeviceTokens
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.DeviceToken,
			&i.Platform,
			&i.DeviceType,
			&i.DeviceModel,
			&i.OsName,
			&i.OsVersion,
			&i.PushNotificationToken,
			&i.IsActive,
			&i.IsVerified,
			&i.LastUsedAt,
			&i.FirstRegisteredAt,
			&i.AppVersion,
			&i.ClientIp,
			&i.ExpiresAt,
			&i.IsRevoked,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDeviceTokenDetails = `-- name: UpdateDeviceTokenDetails :one
UPDATE user_device_tokens
SET 
    platform = COALESCE($2, platform),
    device_type = COALESCE($3, device_type),
    device_model = COALESCE($4, device_model),
    os_name = COALESCE($5, os_name),
    os_version = COALESCE($6, os_version),
    app_version = COALESCE($7, app_version),
    client_ip = COALESCE($8, client_ip),
    last_used_at = NOW(),
    is_verified = COALESCE($9, is_verified)
WHERE id = $1
RETURNING id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked
`

type UpdateDeviceTokenDetailsParams struct {
	ID          uuid.UUID   `json:"id"`
	Platform    string      `json:"platform"`
	DeviceType  pgtype.Text `json:"device_type"`
	DeviceModel pgtype.Text `json:"device_model"`
	OsName      pgtype.Text `json:"os_name"`
	OsVersion   pgtype.Text `json:"os_version"`
	AppVersion  pgtype.Text `json:"app_version"`
	ClientIp    pgtype.Text `json:"client_ip"`
	IsVerified  bool        `json:"is_verified"`
}

func (q *Queries) UpdateDeviceTokenDetails(ctx context.Context, arg UpdateDeviceTokenDetailsParams) (UserDeviceTokens, error) {
	row := q.db.QueryRow(ctx, updateDeviceTokenDetails,
		arg.ID,
		arg.Platform,
		arg.DeviceType,
		arg.DeviceModel,
		arg.OsName,
		arg.OsVersion,
		arg.AppVersion,
		arg.ClientIp,
		arg.IsVerified,
	)
	var i UserDeviceTokens
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceToken,
		&i.Platform,
		&i.DeviceType,
		&i.DeviceModel,
		&i.OsName,
		&i.OsVersion,
		&i.PushNotificationToken,
		&i.IsActive,
		&i.IsVerified,
		&i.LastUsedAt,
		&i.FirstRegisteredAt,
		&i.AppVersion,
		&i.ClientIp,
		&i.ExpiresAt,
		&i.IsRevoked,
	)
	return i, err
}

const updateDeviceTokenLastUsed = `-- name: UpdateDeviceTokenLastUsed :one
UPDATE user_device_tokens
SET 
    last_used_at = NOW(),
    client_ip = COALESCE($2, client_ip)
WHERE id = $1
RETURNING id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked
`

type UpdateDeviceTokenLastUsedParams struct {
	ID       uuid.UUID   `json:"id"`
	ClientIp pgtype.Text `json:"client_ip"`
}

func (q *Queries) UpdateDeviceTokenLastUsed(ctx context.Context, arg UpdateDeviceTokenLastUsedParams) (UserDeviceTokens, error) {
	row := q.db.QueryRow(ctx, updateDeviceTokenLastUsed, arg.ID, arg.ClientIp)
	var i UserDeviceTokens
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceToken,
		&i.Platform,
		&i.DeviceType,
		&i.DeviceModel,
		&i.OsName,
		&i.OsVersion,
		&i.PushNotificationToken,
		&i.IsActive,
		&i.IsVerified,
		&i.LastUsedAt,
		&i.FirstRegisteredAt,
		&i.AppVersion,
		&i.ClientIp,
		&i.ExpiresAt,
		&i.IsRevoked,
	)
	return i, err
}

const updateDeviceTokenPushNotificationToken = `-- name: UpdateDeviceTokenPushNotificationToken :one
UPDATE user_device_tokens
SET 
    push_notification_token = $2,
    is_verified = true
WHERE id = $1
RETURNING id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked
`

type UpdateDeviceTokenPushNotificationTokenParams struct {
	ID                    uuid.UUID   `json:"id"`
	PushNotificationToken pgtype.Text `json:"push_notification_token"`
}

func (q *Queries) UpdateDeviceTokenPushNotificationToken(ctx context.Context, arg UpdateDeviceTokenPushNotificationTokenParams) (UserDeviceTokens, error) {
	row := q.db.QueryRow(ctx, updateDeviceTokenPushNotificationToken, arg.ID, arg.PushNotificationToken)
	var i UserDeviceTokens
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceToken,
		&i.Platform,
		&i.DeviceType,
		&i.DeviceModel,
		&i.OsName,
		&i.OsVersion,
		&i.PushNotificationToken,
		&i.IsActive,
		&i.IsVerified,
		&i.LastUsedAt,
		&i.FirstRegisteredAt,
		&i.AppVersion,
		&i.ClientIp,
		&i.ExpiresAt,
		&i.IsRevoked,
	)
	return i, err
}

const upsertUserDeviceToken = `-- name: UpsertUserDeviceToken :one
INSERT INTO user_device_tokens (
    id,
    user_id,
    device_token,
    platform,
    device_type,
    device_model,
    os_name,
    os_version,
    push_notification_token,
    is_active,
    is_verified,
    app_version,
    client_ip,
    expires_at
) VALUES (
    $1, 
    $2, 
    $3, 
    COALESCE($4, 'unknown'),
    COALESCE($5, 'unknown'),
    COALESCE($6, 'unknown'),
    COALESCE($7, 'unknown'),
    COALESCE($8, 'unknown'),
    $9,
    COALESCE($10, true),
    COALESCE($11, false),
    $12,
    $13,
    $14
) ON CONFLICT (device_token) DO UPDATE SET
    platform = COALESCE(EXCLUDED.platform, user_device_tokens.platform),
    device_type = COALESCE(EXCLUDED.device_type, user_device_tokens.device_type),
    device_model = COALESCE(EXCLUDED.device_model, user_device_tokens.device_model),
    os_name = COALESCE(EXCLUDED.os_name, user_device_tokens.os_name),
    os_version = COALESCE(EXCLUDED.os_version, user_device_tokens.os_version),
    push_notification_token = COALESCE(EXCLUDED.push_notification_token, user_device_tokens.push_notification_token),
    is_active = COALESCE(EXCLUDED.is_active, user_device_tokens.is_active),
    is_verified = COALESCE(EXCLUDED.is_verified, user_device_tokens.is_verified),
    app_version = COALESCE(EXCLUDED.app_version, user_device_tokens.app_version),
    client_ip = COALESCE(EXCLUDED.client_ip, user_device_tokens.client_ip),
    expires_at = COALESCE(EXCLUDED.expires_at, user_device_tokens.expires_at),
    last_used_at = NOW()
RETURNING id, user_id, device_token, platform, device_type, device_model, os_name, os_version, push_notification_token, is_active, is_verified, last_used_at, first_registered_at, app_version, client_ip, expires_at, is_revoked
`

type UpsertUserDeviceTokenParams struct {
	ID                    uuid.UUID          `json:"id"`
	UserID                uuid.UUID          `json:"user_id"`
	DeviceToken           string             `json:"device_token"`
	Column4               interface{}        `json:"column_4"`
	Column5               interface{}        `json:"column_5"`
	Column6               interface{}        `json:"column_6"`
	Column7               interface{}        `json:"column_7"`
	Column8               interface{}        `json:"column_8"`
	PushNotificationToken pgtype.Text        `json:"push_notification_token"`
	Column10              interface{}        `json:"column_10"`
	Column11              interface{}        `json:"column_11"`
	AppVersion            pgtype.Text        `json:"app_version"`
	ClientIp              pgtype.Text        `json:"client_ip"`
	ExpiresAt             pgtype.Timestamptz `json:"expires_at"`
}

func (q *Queries) UpsertUserDeviceToken(ctx context.Context, arg UpsertUserDeviceTokenParams) (UserDeviceTokens, error) {
	row := q.db.QueryRow(ctx, upsertUserDeviceToken,
		arg.ID,
		arg.UserID,
		arg.DeviceToken,
		arg.Column4,
		arg.Column5,
		arg.Column6,
		arg.Column7,
		arg.Column8,
		arg.PushNotificationToken,
		arg.Column10,
		arg.Column11,
		arg.AppVersion,
		arg.ClientIp,
		arg.ExpiresAt,
	)
	var i UserDeviceTokens
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DeviceToken,
		&i.Platform,
		&i.DeviceType,
		&i.DeviceModel,
		&i.OsName,
		&i.OsVersion,
		&i.PushNotificationToken,
		&i.IsActive,
		&i.IsVerified,
		&i.LastUsedAt,
		&i.FirstRegisteredAt,
		&i.AppVersion,
		&i.ClientIp,
		&i.ExpiresAt,
		&i.IsRevoked,
	)
	return i, err
}
