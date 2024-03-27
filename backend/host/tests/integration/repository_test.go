package integration

import (
	"os"
	"testing"
	"time"

	"github.com/rise8-us/neverl8/host"
	"github.com/rise8-us/neverl8/model"
	testutil "github.com/rise8-us/neverl8/test_config"
	"github.com/stretchr/testify/assert"
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

func GetSampleHosts() []model.Host {
	hosts := &[]model.Host{
		{HostName: "Host 1", ID: 1, LastMeetingTime: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{HostName: "Host 2", ID: 2, LastMeetingTime: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)},
	}
	return *hosts
}

func TestCreateHost(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := host.NewHostRepository(db)

	hosts := GetSampleHosts()

	// Create Hosts 1 and 2
	createdHost, err := repo.CreateHost(&hosts[0])
	assert.NoError(t, err, "expected no error")

	createdHost2, err := repo.CreateHost(&hosts[1])
	assert.NoError(t, err, "expected no error")

	assert.NotNil(t, createdHost, "expected host to be created")
	assert.NotNil(t, createdHost2, "expected second host to be created")
	assert.Equal(t, uint(1), createdHost.ID, "expected host id to be 1")
	assert.Equal(t, uint(2), createdHost2.ID, "expected host id to be 2")
}

func TestGetHostByID(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := host.NewHostRepository(db)

	hosts := GetSampleHosts()

	// Create host
	createdHost, err := repo.CreateHost(&hosts[0])
	assert.NoError(t, err, "expected no error")

	// Get Host 1
	retrievedHost, err := repo.GetHostByID(createdHost.ID)
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, retrievedHost, "expected host to be retrieved")
	assert.Equal(t, createdHost.ID, retrievedHost.ID, "expected host to equal retrieved host")
}

func TestGetAllHosts(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := host.NewHostRepository(db)

	hosts := GetSampleHosts()

	// Create Hosts 1 and 2
	_, err := repo.CreateHost(&hosts[0])
	assert.NoError(t, err, "expected no error")

	_, err = repo.CreateHost(&hosts[1])
	assert.NoError(t, err, "expected no error")

	// Get all Hosts
	hosts, err = repo.GetAllHosts()
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, hosts, "expected hosts to be retrieved")
	assert.Equal(t, 2, len(hosts), "expected 2 hosts to be retrieved")
}

func TestCreateHostTimePreferences(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.TimePreference{})
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := host.NewHostRepository(db)

	hosts := GetSampleHosts()

	// Create host
	createdHost, err := repo.CreateHost(&hosts[0])
	assert.NoError(t, err, "expected no error")

	// Create TimePreference for Host 1
	layout := "15:04"
	startTime, _ := time.Parse(layout, "09:00")
	endTime, _ := time.Parse(layout, "17:00")

	timePreference := model.TimePreference{HostID: createdHost.ID, StartTime: startTime, EndTime: endTime}
	// Test Create TimePreferences for Host 1
	createdTimePreference, err := repo.CreateTimePreference(&timePreference)
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, createdTimePreference, "expected time preferences to be created")
}

func TestGetHostTimePreferences(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.TimePreference{})
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := host.NewHostRepository(db)

	hosts := GetSampleHosts()

	// Create host
	createdHost, err := repo.CreateHost(&hosts[0])
	assert.NoError(t, err, "expected no error")

	// Create TimePreference for Host 1
	layout := "15:04"
	startTime, _ := time.Parse(layout, "09:00")
	endTime, _ := time.Parse(layout, "17:00")

	timePreference := model.TimePreference{HostID: createdHost.ID, StartTime: startTime, EndTime: endTime}

	createdTimePreference, err := repo.CreateTimePreference(&timePreference)
	assert.NoError(t, err, "expected no error when creating time preference")
	assert.NotNil(t, createdTimePreference, "expected time preferences to be created")

	// Get all Hosts
	newHost, _ := repo.GetAllHosts()
	assert.NoError(t, err, "expected no error")

	// Get TimePreferences for Host 1
	assert.NotNil(t, newHost[0].TimePreferences, "expected time preferences to be retrieved")
	assert.Equal(t, &timePreference, createdTimePreference, "expected time preference to equal retrieved time preference")
}

func TestCreateHostCalendar(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Host{})
		db.Where("1 = 1").Delete(&model.Calendar{})
	})

	repo := host.NewHostRepository(db)

	hosts := GetSampleHosts()
	calendar := model.Calendar{GoogleCalendarID: "johndoe@rise8.us"}

	// Test Create Calendar for host
	createdCalendar, err := repo.CreateCalendar(&calendar, &hosts[0])
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, createdCalendar, "expected calendar to be created")
	assert.Equal(t, calendar, *createdCalendar, "expected calendar to equal created calendar")
}
