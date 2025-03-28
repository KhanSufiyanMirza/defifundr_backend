package commons

import (
	"testing"

	"github.com/demola234/defifundr/pkg/random"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := random.RandomString(6)

	hashPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword)

	err = CheckPassword(password, hashPassword)
	require.NoError(t, err)

	wrongPassword := random.RandomString(6)
	err = CheckPassword(wrongPassword, hashPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
	require.NotEmpty(t, wrongPassword)

	hashPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword2)
	require.NotEmpty(t, hashPassword, hashPassword2)
}
