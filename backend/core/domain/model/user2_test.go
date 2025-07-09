package model

import (
	"database/sql"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	// Create a new User instance
	test := []struct {
		name         string
		ulid         string
		uid          string
		displayName  string
		deletedAt    sql.Null[time.Time]
		want         *User
	}{
		{
			name: "正常なユーザー",
			ulid:  "01ARZ3NDEKTSV4RRFFQ69G5FAV",
			uid:   "firebase_uid_123",
			displayName: "正常なテストユーザー",
			deletedAt: sql.Null[time.Time]{},
			want: &User{
				Ulid:        "01ARZ3NDEKTSV4RRFFQ69G5FAV",
				UID:         "firebase_uid_123",
				DisplayName: "正常なテストユーザー",
				DeletedAt:   sql.Null[time.Time]{},
			},
		},
		{
			name:        "削除済みユーザー",
			ulid:        "01ARZ3NDEKTSV4RRFFQ69G5FAV",
			uid:         "firebase_uid_123",
			displayName: "削除済みテストユーザー",
			deletedAt:   sql.Null[time.Time]{Time: time.Now(), Valid: true},
			want: &User{
				Ulid:        "01ARZ3NDEKTSV4RRFFQ69G5FAV",
				UID:         "firebase_uid_123",
				DisplayName: "削除済みテストユーザー",
				DeletedAt:   sql.Null[time.Time]{Time: time.Now(), Valid: true},
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUser(tt.ulid, tt.uid, tt.displayName, tt.deletedAt)
			if got.Ulid != tt.want.Ulid {
				t.Errorf("NewUser().Ulid = %v, want %v", got.Ulid, tt.want.Ulid)
			}
			if got.UID != tt.want.UID {
				t.Errorf("NewUser().UID = %v, want %v", got.UID, tt.want.UID)
			}
			if got.DisplayName != tt.want.DisplayName {
				t.Errorf("NewUser().DisplayName = %v, want %v", got.DisplayName, tt.want.DisplayName)
			}
			if got.DeletedAt != tt.want.DeletedAt {
				t.Errorf("NewUser().DeletedAt = %v, want %v", got.DeletedAt, tt.want.DeletedAt)
			}
		})
	}
}
