package data

import (
	"github.com/lucabrx/wuhu/internal/random"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUser_Insert_Valid(t *testing.T) {
	createRandomUser(t)
}

func TestUser_Insert_DuplicateEmail(t *testing.T) {
	email := random.RandString(10) + "@gmail.com"
	user1 := &User{
		Email: email,
		Name:  random.RandString(10),
		Image: random.RandString(10),
	}
	err := user1.Password.Set(random.RandString(10))
	require.NoError(t, err)
	user2 := &User{
		Email: email,
		Name:  random.RandString(10),
		Image: random.RandString(10),
	}

	err = user2.Password.Set(random.RandString(10))
	require.NoError(t, err)

	err = testDB.User.Insert(user1)
	require.NoError(t, err)
	require.NotEmpty(t, user1.ID)
	require.NotEmpty(t, user1.CreatedAt)

	err = testDB.User.Insert(user2)
	require.Error(t, err)
	require.EqualError(t, err, ErrDuplicateEmail.Error())
}

func TestUser_GetById_Valid(t *testing.T) {
	user := createRandomUser(t)

	dbUser, err := testDB.User.GetByID(user.ID)
	require.NoError(t, err)
	require.Equal(t, user, *dbUser)
}

func TestUser_GetById_NotFound(t *testing.T) {
	dbUser, err := testDB.User.GetByID(0)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Nil(t, dbUser)
}

func TestUser_GetByEmail_Valid(t *testing.T) {
	user := createRandomUser(t)

	dbUser, err := testDB.User.GetByEmail(user.Email)
	require.NoError(t, err)
	require.Equal(t, user, *dbUser)
}

func TestUser_GetByEmail_NotFound(t *testing.T) {
	dbUser, err := testDB.User.GetByEmail("hmm")
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Nil(t, dbUser)
}
