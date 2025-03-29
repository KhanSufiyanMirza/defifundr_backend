// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	// Blocks all sessions for a specific user
	BlockAllUserSessions(ctx context.Context, userID uuid.UUID) error
	// Blocks all expired sessions
	BlockExpiredSessions(ctx context.Context) error
	// Blocks a session (marks it as invalid)
	BlockSession(ctx context.Context, id uuid.UUID) error
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CountActiveDeviceTokensForUser(ctx context.Context, userID uuid.UUID) (int64, error)
	CountActiveOTPsForUser(ctx context.Context, arg CountActiveOTPsForUserParams) (int64, error)
	// Counts the number of active sessions
	CountActiveSessions(ctx context.Context) (int64, error)
	// Counts the number of active sessions for a specific user
	CountActiveSessionsByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
	// Counts the number of users matching a search query
	CountSearchUsers(ctx context.Context, dollar_1 pgtype.Text) (int64, error)
	// Counts the number of waitlist entries matching a search query
	CountSearchWaitlist(ctx context.Context, dollar_1 pgtype.Text) (int64, error)
	// Counts the total number of users (useful for pagination)
	CountUsers(ctx context.Context) (int64, error)
	// Counts users filtered by account type
	CountUsersByAccountType(ctx context.Context, accountType string) (int64, error)
	// Counts the total number of waitlist entries matching filters
	CountWaitlistEntries(ctx context.Context, arg CountWaitlistEntriesParams) (int64, error)
	CreateOTPVerification(ctx context.Context, arg CreateOTPVerificationParams) (OtpVerifications, error)
	// Creates a new session and returns the created session record
	CreateSession(ctx context.Context, arg CreateSessionParams) (Sessions, error)
	// Creates a new user record and returns the created user
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	CreateUserDeviceToken(ctx context.Context, arg CreateUserDeviceTokenParams) (UserDeviceTokens, error)
	// Creates a new waitlist entry and returns the created entry
	CreateWaitlistEntry(ctx context.Context, arg CreateWaitlistEntryParams) (Waitlist, error)
	DeleteExpiredDeviceTokens(ctx context.Context) error
	DeleteExpiredOTPs(ctx context.Context) error
	// Cleans up expired sessions that are older than the specified date
	DeleteExpiredSessions(ctx context.Context, expiresAt time.Time) error
	// Deletes a session by its ID
	DeleteSession(ctx context.Context, id uuid.UUID) error
	// Deletes all sessions for a specific user
	DeleteSessionsByUserID(ctx context.Context, userID uuid.UUID) error
	// Permanently deletes a transaction record
	DeleteTransaction(ctx context.Context, id uuid.UUID) error
	// Deletes all transactions for a specific user
	DeleteTransactionsByUserID(ctx context.Context, userID uuid.UUID) error
	// Permanently deletes a user record
	DeleteUser(ctx context.Context, id uuid.UUID) error
	// Permanently deletes a waitlist entry
	DeleteWaitlistEntry(ctx context.Context, id uuid.UUID) error
	// Retrieves all waitlist entries for export
	ExportWaitlistEntries(ctx context.Context) ([]Waitlist, error)
	GetActiveDeviceTokensForUser(ctx context.Context, userID uuid.UUID) ([]UserDeviceTokens, error)
	// Retrieves active (non-expired, non-blocked) sessions with pagination
	GetActiveSessions(ctx context.Context, arg GetActiveSessionsParams) ([]Sessions, error)
	// Retrieves active sessions for a specific user
	GetActiveSessionsByUserID(ctx context.Context, userID uuid.UUID) ([]Sessions, error)
	GetDeviceTokensByPlatform(ctx context.Context, arg GetDeviceTokensByPlatformParams) ([]UserDeviceTokens, error)
	GetOTPVerificationByID(ctx context.Context, id uuid.UUID) (OtpVerifications, error)
	GetOTPVerificationByUserAndPurpose(ctx context.Context, arg GetOTPVerificationByUserAndPurposeParams) (OtpVerifications, error)
	// Retrieves a session by its ID
	GetSessionByID(ctx context.Context, id uuid.UUID) (Sessions, error)
	// Retrieves a session by refresh token
	GetSessionByRefreshToken(ctx context.Context, refreshToken string) (Sessions, error)
	// Retrieves all sessions for a specific user
	GetSessionsByUserID(ctx context.Context, userID uuid.UUID) ([]Sessions, error)
	// Retrieves a single transaction by its ID
	GetTransactionByID(ctx context.Context, id uuid.UUID) (Transactions, error)
	// Retrieves a single transaction by its transaction hash
	GetTransactionByTxHash(ctx context.Context, txHash string) (Transactions, error)
	// Retrieves transactions by status
	GetTransactionsByStatus(ctx context.Context, arg GetTransactionsByStatusParams) ([]Transactions, error)
	// Retrieves all transactions for a specific user
	GetTransactionsByUserID(ctx context.Context, userID uuid.UUID) ([]Transactions, error)
	// Retrieves transactions for a specific user with a specific status
	GetTransactionsByUserIDAndStatus(ctx context.Context, arg GetTransactionsByUserIDAndStatusParams) ([]Transactions, error)
	GetUnverifiedOTPsForUser(ctx context.Context, userID pgtype.UUID) ([]OtpVerifications, error)
	GetUser(ctx context.Context, id uuid.UUID) (Users, error)
	// Retrieves a single user by their email address
	GetUserByEmail(ctx context.Context, email string) (Users, error)
	GetUserDeviceTokenByDeviceToken(ctx context.Context, deviceToken string) (UserDeviceTokens, error)
	GetUserDeviceTokenByID(ctx context.Context, id uuid.UUID) (UserDeviceTokens, error)
	// Retrieves a single waitlist entry by email address
	GetWaitlistEntryByEmail(ctx context.Context, email string) (Waitlist, error)
	// Retrieves a single waitlist entry by ID
	GetWaitlistEntryByID(ctx context.Context, id uuid.UUID) (Waitlist, error)
	// Retrieves a single waitlist entry by referral code
	GetWaitlistEntryByReferralCode(ctx context.Context, referralCode string) (Waitlist, error)
	// Gets the position of an entry in the waitlist (by signup date)
	GetWaitlistPosition(ctx context.Context, id uuid.UUID) (int64, error)
	// Gets waitlist statistics grouped by referral source
	GetWaitlistStatsBySource(ctx context.Context) ([]GetWaitlistStatsBySourceRow, error)
	// Gets waitlist statistics grouped by status
	GetWaitlistStatsByStatus(ctx context.Context) ([]GetWaitlistStatsByStatusRow, error)
	InValidateOTP(ctx context.Context, id uuid.UUID) error
	// Lists users with pagination support
	ListUsers(ctx context.Context, arg ListUsersParams) ([]Users, error)
	// Lists users filtered by account type with pagination
	ListUsersByAccountType(ctx context.Context, arg ListUsersByAccountTypeParams) ([]Users, error)
	// Lists waitlist entries with pagination and filtering support
	ListWaitlistEntries(ctx context.Context, arg ListWaitlistEntriesParams) ([]Waitlist, error)
	RevokeDeviceToken(ctx context.Context, id uuid.UUID) (UserDeviceTokens, error)
	SearchDeviceTokens(ctx context.Context, arg SearchDeviceTokensParams) ([]UserDeviceTokens, error)
	// Searches for users by name, email, or nationality with pagination
	SearchUsers(ctx context.Context, arg SearchUsersParams) ([]Users, error)
	// Searches for waitlist entries by email or name with pagination
	SearchWaitlist(ctx context.Context, arg SearchWaitlistParams) ([]Waitlist, error)
	UpdateDeviceTokenDetails(ctx context.Context, arg UpdateDeviceTokenDetailsParams) (UserDeviceTokens, error)
	UpdateDeviceTokenLastUsed(ctx context.Context, arg UpdateDeviceTokenLastUsedParams) (UserDeviceTokens, error)
	UpdateDeviceTokenPushNotificationToken(ctx context.Context, arg UpdateDeviceTokenPushNotificationTokenParams) (UserDeviceTokens, error)
	UpdateOTPAttempts(ctx context.Context, id uuid.UUID) (OtpVerifications, error)
	// Updates just the refresh token of a session
	UpdateRefreshToken(ctx context.Context, arg UpdateRefreshTokenParams) (Sessions, error)
	// Updates session details
	UpdateSession(ctx context.Context, arg UpdateSessionParams) (Sessions, error)
	// Updates transaction details and returns the updated transaction
	UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transactions, error)
	// Updates the status of a transaction and returns the updated transaction
	UpdateTransactionStatus(ctx context.Context, arg UpdateTransactionStatusParams) (Transactions, error)
	// Updates user details and returns the updated user
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
	// Updates a user's email address with validation that the new email is unique
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (Users, error)
	// Updates a user's password
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	// Updates the status of a waitlist entry
	UpdateWaitlistEntryStatus(ctx context.Context, arg UpdateWaitlistEntryStatusParams) error
	UpsertUserDeviceToken(ctx context.Context, arg UpsertUserDeviceTokenParams) (UserDeviceTokens, error)
	VerifyOTP(ctx context.Context, arg VerifyOTPParams) (OtpVerifications, error)
}

var _ Querier = (*Queries)(nil)
