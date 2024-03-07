package repository_test

import (
	"testing"

	"github.com/rise8-us/neverl8/model"
	"github.com/rise8-us/neverl8/repository"
	"github.com/stretchr/testify/assert"
)

func GetSampleHosts() []model.Host {
	// Create new Hosts
	hosts := &[]model.Host{
		{HostName: "Host 1", TimePreferences: []model.TimePreference{}, Calendars: []model.Calendar{}, Meetings: []model.Meetings{}},
		{HostName: "Host 2", TimePreferences: []model.TimePreference{}, Calendars: []model.Calendar{}, Meetings: []model.Meetings{}},
	}
	return *hosts
}

func TestCreateHost(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := repository.NewHostRepository(db)

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

	repo := repository.NewHostRepository(db)

	hosts := GetSampleHosts()

	// Create host
	createdHost, err := repo.CreateHost(&hosts[0])
	assert.NoError(t, err, "expected no error")

	// Get Host 1
	host, err := repo.GetHostByID(createdHost.ID)
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, host, "expected host to be retrieved")
	assert.Equal(t, *createdHost, *host, "expected host to equal retrieved host")
}

func TestGetAllHosts(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Host{})
	})

	repo := repository.NewHostRepository(db)

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

	repo := repository.NewHostRepository(db)

	hosts := GetSampleHosts()

	// Create host
	createdHost, err := repo.CreateHost(&hosts[0])
	assert.NoError(t, err, "expected no error")

	// Create TimePreference for Host 1
	startTime := "14:00" // 0900 EST
	endTime := "17:00"   // 1200 EST

	timePreference := model.TimePreference{HostID: createdHost.ID, StartWindow: startTime, EndWindow: endTime}

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

	repo := repository.NewHostRepository(db)

	hosts := GetSampleHosts()

	// Create host
	createdHost, err := repo.CreateHost(&hosts[0])
	assert.NoError(t, err, "expected no error")

	// Create TimePreference for Host 1
	startTime := "14:00" // 0900 EST
	endTime := "17:00"   // 1200 EST

	timePreference := model.TimePreference{HostID: createdHost.ID, StartWindow: startTime, EndWindow: endTime}

	createdTimePreference, err := repo.CreateTimePreference(&timePreference)
	assert.NoError(t, err, "expected no error when creating time preference")
	assert.NotNil(t, createdTimePreference, "expected time preferences to be created")

	// Get all Hosts
	newHost, _ := repo.GetAllHosts()
	assert.NoError(t, err, "expected no error")

	// Get TimePreferences for Host 1
	assert.NotNil(t, newHost[0].TimePreferences, "expected time preferences to be retrieved")
	assert.Equal(t, timePreference, *createdTimePreference, "expected time preference to equal retrieved time preference")
}

func TestCreateHostCalendar(t *testing.T) {
	t.Cleanup(func() {
		db.Where("1 = 1").Delete(&model.Host{})
		db.Where("1 = 1").Delete(&model.Calendar{})
	})

	repo := repository.NewHostRepository(db)

	hosts := GetSampleHosts()
	calendar := model.Calendar{GoogleCalendarID: "johndoe@rise8.us"}

	// Test Create Calendar for host
	createdCalendar, err := repo.CreateCalendar(&calendar, &hosts[0])
	assert.NoError(t, err, "expected no error")
	assert.NotNil(t, createdCalendar, "expected calendar to be created")
	assert.Equal(t, calendar, *createdCalendar, "expected calendar to equal created calendar")
}
