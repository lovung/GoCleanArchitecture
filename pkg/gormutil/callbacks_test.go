package gormutil

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/lovung/GoCleanArchitecture/pkg/jwtutil"

	"gorm.io/gorm"
)

func TestDeletedByFromClaim(t *testing.T) {
	db, _ := gorm.Open(nil, nil)
	userID := uint64(1)
	jwtClaims := jwtutil.JWTClaims{
		UserID: userID,
	}
	ctx := context.WithValue(context.Background(), jwtutil.JWTClaimsKey, jwtClaims) // nolint
	db = db.WithContext(ctx)

	type args struct {
		d *gorm.DB
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]interface{}
		want1 bool
	}{
		{
			name: "Test case1: success",
			args: args{
				d: db,
			},
			want:  map[string]interface{}{"deleted_by": userID},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DeletedByFromClaim(tt.args.d)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeletedByFromClaim() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DeletedByFromClaim() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSoftDeleteClauseFromClaim(t *testing.T) {
	db, _ := gorm.Open(nil, nil)
	userID := uint64(1)
	jwtClaims := jwtutil.JWTClaims{
		UserID: userID,
	}
	ctx := context.WithValue(context.Background(), jwtutil.JWTClaimsKey, jwtClaims) // nolint
	db = db.WithContext(ctx)

	now := time.Now()

	type args struct {
		d   *gorm.DB
		now time.Time
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]interface{}
		want1 bool
	}{
		{
			name: "Test case1: success",
			args: args{
				d:   db,
				now: now,
			},
			want: map[string]interface{}{
				"deleted_by": userID,
				"deleted_at": now,
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SoftDeleteClauseFromClaim(tt.args.d, tt.args.now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SoftDeleteClauseFromClaim() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SoftDeleteClauseFromClaim() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
