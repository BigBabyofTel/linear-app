package data

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToken_Insert_Valid(t *testing.T) {
	user := createRandomUser(t)
	createRandomUserToken(t, user)
}

func TestToken_DeleteAllForUser(t *testing.T) {
	user := createRandomUser(t)
	createRandomUserToken(t, user)

	err := testDB.Token.DeleteAllForUser(ScopeAuthentication, user.ID)
	require.NoError(t, err)
}
