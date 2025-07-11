// Code generated by BobGen psql v0.38.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	models "github.com/sandonemaki/my_read_memo_Go_API/backend/cmd/core/pkg/dbmodels"
	"github.com/stephenafamo/bob"
)

type ReadingHistoryMod interface {
	Apply(context.Context, *ReadingHistoryTemplate)
}

type ReadingHistoryModFunc func(context.Context, *ReadingHistoryTemplate)

func (f ReadingHistoryModFunc) Apply(ctx context.Context, n *ReadingHistoryTemplate) {
	f(ctx, n)
}

type ReadingHistoryModSlice []ReadingHistoryMod

func (mods ReadingHistoryModSlice) Apply(ctx context.Context, n *ReadingHistoryTemplate) {
	for _, f := range mods {
		f.Apply(ctx, n)
	}
}

// ReadingHistoryTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type ReadingHistoryTemplate struct {
	ID         func() int64
	UserUlid   func() string
	ContentURL func() string
	RecordedAt func() time.Time
	CreatedAt  func() time.Time

	r readingHistoryR
	f *Factory
}

type readingHistoryR struct {
	UserUlidUser *readingHistoryRUserUlidUserR
}

type readingHistoryRUserUlidUserR struct {
	o *UserTemplate
}

// Apply mods to the ReadingHistoryTemplate
func (o *ReadingHistoryTemplate) Apply(ctx context.Context, mods ...ReadingHistoryMod) {
	for _, mod := range mods {
		mod.Apply(ctx, o)
	}
}

// setModelRels creates and sets the relationships on *models.ReadingHistory
// according to the relationships in the template. Nothing is inserted into the db
func (t ReadingHistoryTemplate) setModelRels(o *models.ReadingHistory) {
	if t.r.UserUlidUser != nil {
		rel := t.r.UserUlidUser.o.Build()
		rel.R.UserUlidReadingHistories = append(rel.R.UserUlidReadingHistories, o)
		o.UserUlid = rel.Ulid // h2
		o.R.UserUlidUser = rel
	}
}

// BuildSetter returns an *models.ReadingHistorySetter
// this does nothing with the relationship templates
func (o ReadingHistoryTemplate) BuildSetter() *models.ReadingHistorySetter {
	m := &models.ReadingHistorySetter{}

	if o.ID != nil {
		val := o.ID()
		m.ID = &val
	}
	if o.UserUlid != nil {
		val := o.UserUlid()
		m.UserUlid = &val
	}
	if o.ContentURL != nil {
		val := o.ContentURL()
		m.ContentURL = &val
	}
	if o.RecordedAt != nil {
		val := o.RecordedAt()
		m.RecordedAt = &val
	}
	if o.CreatedAt != nil {
		val := o.CreatedAt()
		m.CreatedAt = &val
	}

	return m
}

// BuildManySetter returns an []*models.ReadingHistorySetter
// this does nothing with the relationship templates
func (o ReadingHistoryTemplate) BuildManySetter(number int) []*models.ReadingHistorySetter {
	m := make([]*models.ReadingHistorySetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.ReadingHistory
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use ReadingHistoryTemplate.Create
func (o ReadingHistoryTemplate) Build() *models.ReadingHistory {
	m := &models.ReadingHistory{}

	if o.ID != nil {
		m.ID = o.ID()
	}
	if o.UserUlid != nil {
		m.UserUlid = o.UserUlid()
	}
	if o.ContentURL != nil {
		m.ContentURL = o.ContentURL()
	}
	if o.RecordedAt != nil {
		m.RecordedAt = o.RecordedAt()
	}
	if o.CreatedAt != nil {
		m.CreatedAt = o.CreatedAt()
	}

	o.setModelRels(m)

	return m
}

// BuildMany returns an models.ReadingHistorySlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use ReadingHistoryTemplate.CreateMany
func (o ReadingHistoryTemplate) BuildMany(number int) models.ReadingHistorySlice {
	m := make(models.ReadingHistorySlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatableReadingHistory(m *models.ReadingHistorySetter) {
	if m.UserUlid == nil {
		val := random_string(nil)
		m.UserUlid = &val
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.ReadingHistory
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *ReadingHistoryTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.ReadingHistory) (context.Context, error) {
	var err error

	return ctx, err
}

// Create builds a readingHistory and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *ReadingHistoryTemplate) Create(ctx context.Context, exec bob.Executor) (*models.ReadingHistory, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// MustCreate builds a readingHistory and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// panics if an error occurs
func (o *ReadingHistoryTemplate) MustCreate(ctx context.Context, exec bob.Executor) *models.ReadingHistory {
	_, m, err := o.create(ctx, exec)
	if err != nil {
		panic(err)
	}
	return m
}

// CreateOrFail builds a readingHistory and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// It calls `tb.Fatal(err)` on the test/benchmark if an error occurs
func (o *ReadingHistoryTemplate) CreateOrFail(ctx context.Context, tb testing.TB, exec bob.Executor) *models.ReadingHistory {
	tb.Helper()
	_, m, err := o.create(ctx, exec)
	if err != nil {
		tb.Fatal(err)
		return nil
	}
	return m
}

// create builds a readingHistory and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *ReadingHistoryTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.ReadingHistory, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatableReadingHistory(opt)

	if o.r.UserUlidUser == nil {
		ReadingHistoryMods.WithNewUserUlidUser().Apply(ctx, o)
	}

	rel0, ok := userCtx.Value(ctx)
	if !ok {
		ctx, rel0, err = o.r.UserUlidUser.o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	opt.UserUlid = &rel0.Ulid

	m, err := models.ReadingHistories.Insert(opt).One(ctx, exec)
	if err != nil {
		return ctx, nil, err
	}
	ctx = readingHistoryCtx.WithValue(ctx, m)

	m.R.UserUlidUser = rel0

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple readingHistories and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o ReadingHistoryTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.ReadingHistorySlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// MustCreateMany builds multiple readingHistories and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// panics if an error occurs
func (o ReadingHistoryTemplate) MustCreateMany(ctx context.Context, exec bob.Executor, number int) models.ReadingHistorySlice {
	_, m, err := o.createMany(ctx, exec, number)
	if err != nil {
		panic(err)
	}
	return m
}

// CreateManyOrFail builds multiple readingHistories and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// It calls `tb.Fatal(err)` on the test/benchmark if an error occurs
func (o ReadingHistoryTemplate) CreateManyOrFail(ctx context.Context, tb testing.TB, exec bob.Executor, number int) models.ReadingHistorySlice {
	tb.Helper()
	_, m, err := o.createMany(ctx, exec, number)
	if err != nil {
		tb.Fatal(err)
		return nil
	}
	return m
}

// createMany builds multiple readingHistories and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o ReadingHistoryTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.ReadingHistorySlice, error) {
	var err error
	m := make(models.ReadingHistorySlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// ReadingHistory has methods that act as mods for the ReadingHistoryTemplate
var ReadingHistoryMods readingHistoryMods

type readingHistoryMods struct{}

func (m readingHistoryMods) RandomizeAllColumns(f *faker.Faker) ReadingHistoryMod {
	return ReadingHistoryModSlice{
		ReadingHistoryMods.RandomID(f),
		ReadingHistoryMods.RandomUserUlid(f),
		ReadingHistoryMods.RandomContentURL(f),
		ReadingHistoryMods.RandomRecordedAt(f),
		ReadingHistoryMods.RandomCreatedAt(f),
	}
}

// Set the model columns to this value
func (m readingHistoryMods) ID(val int64) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.ID = func() int64 { return val }
	})
}

// Set the Column from the function
func (m readingHistoryMods) IDFunc(f func() int64) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.ID = f
	})
}

// Clear any values for the column
func (m readingHistoryMods) UnsetID() ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.ID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m readingHistoryMods) RandomID(f *faker.Faker) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.ID = func() int64 {
			return random_int64(f)
		}
	})
}

// Set the model columns to this value
func (m readingHistoryMods) UserUlid(val string) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.UserUlid = func() string { return val }
	})
}

// Set the Column from the function
func (m readingHistoryMods) UserUlidFunc(f func() string) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.UserUlid = f
	})
}

// Clear any values for the column
func (m readingHistoryMods) UnsetUserUlid() ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.UserUlid = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m readingHistoryMods) RandomUserUlid(f *faker.Faker) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.UserUlid = func() string {
			return random_string(f)
		}
	})
}

// Set the model columns to this value
func (m readingHistoryMods) ContentURL(val string) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.ContentURL = func() string { return val }
	})
}

// Set the Column from the function
func (m readingHistoryMods) ContentURLFunc(f func() string) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.ContentURL = f
	})
}

// Clear any values for the column
func (m readingHistoryMods) UnsetContentURL() ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.ContentURL = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m readingHistoryMods) RandomContentURL(f *faker.Faker) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.ContentURL = func() string {
			return random_string(f, "255")
		}
	})
}

// Set the model columns to this value
func (m readingHistoryMods) RecordedAt(val time.Time) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.RecordedAt = func() time.Time { return val }
	})
}

// Set the Column from the function
func (m readingHistoryMods) RecordedAtFunc(f func() time.Time) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.RecordedAt = f
	})
}

// Clear any values for the column
func (m readingHistoryMods) UnsetRecordedAt() ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.RecordedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m readingHistoryMods) RandomRecordedAt(f *faker.Faker) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.RecordedAt = func() time.Time {
			return random_time_Time(f)
		}
	})
}

// Set the model columns to this value
func (m readingHistoryMods) CreatedAt(val time.Time) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.CreatedAt = func() time.Time { return val }
	})
}

// Set the Column from the function
func (m readingHistoryMods) CreatedAtFunc(f func() time.Time) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.CreatedAt = f
	})
}

// Clear any values for the column
func (m readingHistoryMods) UnsetCreatedAt() ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.CreatedAt = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m readingHistoryMods) RandomCreatedAt(f *faker.Faker) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(_ context.Context, o *ReadingHistoryTemplate) {
		o.CreatedAt = func() time.Time {
			return random_time_Time(f)
		}
	})
}

func (m readingHistoryMods) WithParentsCascading() ReadingHistoryMod {
	return ReadingHistoryModFunc(func(ctx context.Context, o *ReadingHistoryTemplate) {
		if isDone, _ := readingHistoryWithParentsCascadingCtx.Value(ctx); isDone {
			return
		}
		ctx = readingHistoryWithParentsCascadingCtx.WithValue(ctx, true)
		{

			related := o.f.NewUser(ctx, UserMods.WithParentsCascading())
			m.WithUserUlidUser(related).Apply(ctx, o)
		}
	})
}

func (m readingHistoryMods) WithUserUlidUser(rel *UserTemplate) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(ctx context.Context, o *ReadingHistoryTemplate) {
		o.r.UserUlidUser = &readingHistoryRUserUlidUserR{
			o: rel,
		}
	})
}

func (m readingHistoryMods) WithNewUserUlidUser(mods ...UserMod) ReadingHistoryMod {
	return ReadingHistoryModFunc(func(ctx context.Context, o *ReadingHistoryTemplate) {
		related := o.f.NewUser(ctx, mods...)

		m.WithUserUlidUser(related).Apply(ctx, o)
	})
}

func (m readingHistoryMods) WithoutUserUlidUser() ReadingHistoryMod {
	return ReadingHistoryModFunc(func(ctx context.Context, o *ReadingHistoryTemplate) {
		o.r.UserUlidUser = nil
	})
}
