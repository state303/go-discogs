package batch

import (
	"context"
	"github.com/state303/go-discogs/internal/testutils"
	"github.com/state303/go-discogs/model"
	"github.com/state303/go-discogs/src/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
)

func TestBatch(t *testing.T) {
	pg := testutils.GetDatabase(testutils.Postgres)
	dsn := testutils.GetDsn(testutils.Postgres, pg)
	db, err := database.GetConnect(dsn)
	require.NoError(t, err)

	var (
		ctx   = context.Background()
		chunk = 5
	)
	order := NewOrder(ctx, chunk, "testdata/artist.xml.gz", db)

	res := newBatch().UpdateArtist(order)()
	require.NoError(t, res.Err())
	require.NotZero(t, res.Count())

	order = NewOrder(ctx, chunk, "testdata/label.xml.gz", db)

	res = newBatch().UpdateLabel(order)()
	require.NoError(t, res.Err())
	require.NotZero(t, res.Count())

	order = NewOrder(ctx, chunk, "testdata/master.xml.gz", db)
	res = newBatch().UpdateMaster(order)()
	require.NoError(t, res.Err())
	require.NotZero(t, res.Count())

	var masters []*model.Master
	db.Session(&gorm.Session{}).Find(&masters)
	require.NotEmpty(t, masters)
	for _, master := range masters {
		var count int64
		db.Session(&gorm.Session{}).Model(&model.MasterStyle{}).Where("master_id = ?", master.ID).Count(&count)
		require.NotZero(t, count)
		db.Session(&gorm.Session{}).Model(&model.MasterGenre{}).Where("master_id = ?", master.ID)
		require.NotZero(t, count)
	}

	order = NewOrder(ctx, chunk, "testdata/release.xml.gz", db)
	res = newBatch().UpdateRelease(order)()
	require.NoError(t, res.Err())
	var count int64
	db.Session(&gorm.Session{}).Model(&model.Release{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.ReleaseContract{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.ReleaseIdentifier{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.ReleaseFormat{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.ReleaseTrack{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.ReleaseCreditedArtist{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.ReleaseVideo{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.LabelRelease{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.ReleaseArtist{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.ReleaseGenre{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.ReleaseStyle{}).Count(&count)
	require.NotZero(t, count)
	db.Session(&gorm.Session{}).Model(&model.MasterMainRelease{}).Count(&count)
	require.NotZero(t, count)
}

func Test_batch_UpdateLabel(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		order Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Step
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &batch{
				db: tt.fields.db,
			}
			assert.Equalf(t, tt.want, b.UpdateLabel(tt.args.order), "UpdateLabel(%v)", tt.args.order)
		})
	}
}

func Test_batch_UpdateMaster(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		order Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Step
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &batch{
				db: tt.fields.db,
			}
			assert.Equalf(t, tt.want, b.UpdateMaster(tt.args.order), "UpdateMaster(%v)", tt.args.order)
		})
	}
}

func Test_batch_UpdateRelease(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		order Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Step
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &batch{
				db: tt.fields.db,
			}
			assert.Equalf(t, tt.want, b.UpdateRelease(tt.args.order), "UpdateRelease(%v)", tt.args.order)
		})
	}
}
