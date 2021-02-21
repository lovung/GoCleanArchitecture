package gormutil

import (
	"time"

	"github.com/lovung/GoCleanArchitecture/pkg/jwtutil"
	"gorm.io/gorm"
)

const (
	maxPageSize = 1000
)

// Paging defines paging struct
type Paging struct{ Size, Number uint64 }

// WithID where id = <input value>
func WithID(id interface{}) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where("id = (?)", id)
	}
}

// WithDeletedAtNull where code = <input value>
func WithDeletedAtNull(table string) func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		return d.Where(table + ".deleted_at IS NULL")
	}
}

// WithUserID where office_id = office_id from context
func WithUserID() func(d *gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		claim := d.Statement.Context.Value(jwtutil.JWTClaimsKey)
		table := d.Statement.Table
		if claim != nil {
			if claimInfo, ok := claim.(jwtutil.JWTClaims); ok {
				if table != "" {
					return d.Where("`"+table+"`.`user_id` = (?)", claimInfo.UserID)
				}
				return d.Where("`user_id` = (?)", claimInfo.UserID)
			}
		}
		return d
	}
}

// UpdateWithClaimInfo add update office_id and updated_by from context
func UpdateWithClaimInfo(d *gorm.DB, data map[string]interface{}) map[string]interface{} {
	claim := d.Statement.Context.Value(jwtutil.JWTClaimsKey)
	if claim != nil {
		if claimInfo, ok := claim.(jwtutil.JWTClaims); ok {
			if _, ok := data["updated_by"]; ok {
				data["updated_by"] = claimInfo.UserID
			}
		}
	}
	return data
}

// DeletedByFromClaim for the hook after deleting record.
func DeletedByFromClaim(d *gorm.DB) (map[string]interface{}, bool) {
	claim := d.Statement.Context.Value(jwtutil.JWTClaimsKey)

	clause := make(map[string]interface{}, 1)
	if claimInfo, ok := claim.(jwtutil.JWTClaims); ok {
		clause["deleted_by"] = claimInfo.UserID
		return clause, true
	}

	// This situation is not expected to happen ecause AuthMiddleware ensures JWTClaims exists
	return clause, false
}

// SoftDeleteClauseFromClaim for Bulk SoftDelete
func SoftDeleteClauseFromClaim(d *gorm.DB, now time.Time) (map[string]interface{}, bool) {
	claim := d.Statement.Context.Value(jwtutil.JWTClaimsKey)

	clause := make(map[string]interface{}, 2)
	if claimInfo, ok := claim.(jwtutil.JWTClaims); ok {
		clause["deleted_by"] = claimInfo.UserID
		clause["deleted_at"] = now
		return clause, true
	}

	// This situation is not expected to happen ecause AuthMiddleware ensures JWTClaims exists
	return clause, false
}

// SoftDeleteClauseFromClaimWithDefaultTime for Bulk SoftDelete with Now
func SoftDeleteClauseFromClaimWithDefaultTime(d *gorm.DB) (map[string]interface{}, bool) {
	now := time.Now()
	return SoftDeleteClauseFromClaim(d, now)
}

// Paginate page with limit and offset
func Paginate(paging Paging) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case paging.Size == 0: // used for current have no paging, will remove in the future
			paging.Size = maxPageSize
		case paging.Size > maxPageSize:
			paging.Size = maxPageSize
		}
		if paging.Number == 0 {
			paging.Number = 1
		}

		offset := (paging.Number - 1) * paging.Size
		return db.Offset(int(offset)).Limit(int(paging.Size))
	}
}
