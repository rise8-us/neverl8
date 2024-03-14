package repository_test

import (
	"os"
	"testing"

	testutil "github.com/rise8-us/neverl8/test/integration/testcontainers"
	"gorm.io/gorm"
)

var db *gorm.DB

// TestMain sets up the test database
func TestMain(m *testing.M) {
	testDB := testutil.SetupTestDB()
	db = testDB.DB

	code := runTests(m, testDB.TearDown)
	os.Exit(code)
}

// While the m.Run() function already runs all the tests, os.Exit() does not respect the defer method.
// This function is used to ensure that the test database is torn down after all tests are run.
func runTests(m *testing.M, tearDown func()) int {
	defer tearDown()
	return m.Run()
}
