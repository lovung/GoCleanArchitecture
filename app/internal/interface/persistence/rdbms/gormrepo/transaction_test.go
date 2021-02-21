package gormrepo

import (
	"reflect"
	"testing"

	"github.com/lovung/GoCleanArchitecture/pkg/testhelper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewTxnDataSQL(t *testing.T) {
	t.Parallel()
	gDB, _, err := testhelper.OpenDBConnection()
	assert.NoError(t, err)
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *TxnDataSQL
	}{
		{
			args: args{
				db: gDB,
			},
			want: &TxnDataSQL{
				db: gDB,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTxnDataSQL(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTxnDataSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
