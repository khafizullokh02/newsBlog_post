package db

import (
	"context"
	"os"
	"testing"

	"github.com/jaswdr/faker"
)

var (
	fake      faker.Faker
	testStore Store
)

const (
	dbSource = "postgres://root:secret@localhost:5432/newsBlog_post?sslmode=disable"
)

func TestMain(m *testing.M) {
	fake = faker.New()
	testStore = NewStore(context.Background(), dbSource)
	os.Exit(m.Run())
}
