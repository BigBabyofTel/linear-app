package data

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrFieldRequired  = errors.New("field required")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	User       UserModel
	Token      TokenModel
	Issue      IssueModal
	Workspace  WorkspaceModal
	Permission PermissionModal
	Request    RequestModal
}

func NewModals(db *sql.DB) Models {
	return Models{
		User:       UserModel{DB: db},
		Token:      TokenModel{DB: db},
		Issue:      IssueModal{DB: db},
		Workspace:  WorkspaceModal{DB: db},
		Permission: PermissionModal{DB: db},
		Request:    RequestModal{DB: db},
	}
}

func newNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func newNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

func newNullBool(b bool) sql.NullBool {
	return sql.NullBool{
		Bool:  b,
		Valid: true,
	}
}

func newNullStringArray(s []string) []sql.NullString {
	var nullStrings []sql.NullString
	for _, v := range s {
		nullStrings = append(nullStrings, newNullString(v))
	}
	return nullStrings
}

func newNullDate(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func newNullPassword(b password) sql.NullString {
	if len(b.hash) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: string(b.hash),
		Valid:  true,
	}
}
