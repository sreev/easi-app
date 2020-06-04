package storage

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"

	"github.com/cmsgov/easi-app/pkg/models"
	"github.com/cmsgov/easi-app/pkg/testhelpers"
)

func (s StoreTestSuite) TestCreateSystemIntake() {
	s.Run("create a new system intake", func() {
		intake := models.SystemIntake{
			EUAUserID: testhelpers.RandomEUAID(),
			Status:    models.SystemIntakeStatusDRAFT,
			Requester: null.StringFrom("Test requester"),
		}

		created, err := s.store.CreateSystemIntake(&intake)
		s.NoError(err)
		s.Equal(intake.EUAUserID, created.EUAUserID)
		s.Equal(intake.Status, created.Status)
		s.Equal(intake.Requester, created.Requester)
		s.False(created.ID == uuid.Nil)
	})

	s.Run("cannot save without EUA ID", func() {
		partialIntake := models.SystemIntake{
			Status: models.SystemIntakeStatusDRAFT,
		}

		_, err := s.store.CreateSystemIntake(&partialIntake)

		s.Error(err)
		s.Equal("pq: new row for relation \"system_intake\" violates check constraint \"eua_id_check\"", err.Error())
	})

	euaTests := []string{
		"f",
		"F",
		"5CHAR",
		"$BAD",
	}
	for _, tc := range euaTests {
		s.Run(fmt.Sprintf("cannot save with invalid EUA ID: %s", tc), func() {
			partialIntake := models.SystemIntake{
				Status: models.SystemIntakeStatusDRAFT,
			}
			partialIntake.EUAUserID = "F"

			_, err := s.store.CreateSystemIntake(&partialIntake)

			s.Error(err)
			s.Equal("pq: new row for relation \"system_intake\" violates check constraint \"eua_id_check\"", err.Error())
		})
	}

	s.Run("cannot create with invalid status", func() {
		partialIntake := models.SystemIntake{
			EUAUserID: testhelpers.RandomEUAID(),
			Status:    "fakeStatus",
			Requester: null.StringFrom("Test requester"),
		}

		created, err := s.store.CreateSystemIntake(&partialIntake)

		s.Error(err)
		s.Equal("pq: invalid input value for enum system_intake_status: \"fakeStatus\"", err.Error())
		s.Equal(&models.SystemIntake{}, created)
	})
}

func (s StoreTestSuite) TestUpdateSystemIntake() {
	s.Run("update an existing system intake", func() {
		intake, err := s.store.CreateSystemIntake(&models.SystemIntake{
			EUAUserID: testhelpers.RandomEUAID(),
			Status:    models.SystemIntakeStatusDRAFT,
			Requester: null.StringFrom("Test requester"),
		})
		s.NoError(err)
		now := time.Now()

		intake.UpdatedAt = &now
		intake.ISSO = null.StringFrom("test isso")

		updated, err := s.store.UpdateSystemIntake(intake)
		s.NoError(err, "failed to update system intake")
		s.Equal(intake, updated)
	})

	s.Run("EUA ID will not update", func() {
		originalIntake := models.SystemIntake{
			EUAUserID: testhelpers.RandomEUAID(),
			Status:    models.SystemIntakeStatusDRAFT,
			Requester: null.StringFrom("Test requester"),
		}
		_, err := s.store.CreateSystemIntake(&originalIntake)
		s.NoError(err)
		originalEUA := originalIntake.EUAUserID
		partialIntake := models.SystemIntake{
			ID:        originalIntake.ID,
			EUAUserID: testhelpers.RandomEUAID(),
			Status:    models.SystemIntakeStatusDRAFT,
			Requester: null.StringFrom("Test requester"),
		}
		partialIntake.EUAUserID = "NEWS"

		_, err = s.store.UpdateSystemIntake(&partialIntake)
		s.NoError(err, "failed to update system intake")

		updated, err := s.store.FetchSystemIntakeByID(originalIntake.ID)
		s.NoError(err)

		s.Equal(originalEUA, updated.EUAUserID)
	})
}

func (s StoreTestSuite) TestSaveSystemIntake() {
	s.Run("save a new system intake", func() {
		intake := testhelpers.NewSystemIntake()

		err := s.store.SaveSystemIntake(&intake)

		s.NoError(err, "failed to save system intake")
		var actualIntake models.SystemIntake
		err = s.db.Get(&actualIntake, "SELECT * FROM system_intake WHERE id=$1", intake.ID)
		s.NoError(err, "failed to fetch saved intake")
		s.Equal(actualIntake, intake)
	})

	partialIntake := models.SystemIntake{}
	s.Run("cannot save without EUA ID", func() {
		id, _ := uuid.NewUUID()
		partialIntake.ID = id
		partialIntake.Status = models.SystemIntakeStatusDRAFT

		err := s.store.SaveSystemIntake(&partialIntake)

		s.Error(err)
		s.Equal("pq: new row for relation \"system_intake\" violates check constraint \"eua_id_check\"", err.Error())
	})

	euaTests := []string{
		"f",
		"F",
		"5CHAR",
		"$BAD",
	}
	for _, tc := range euaTests {
		s.Run(fmt.Sprintf("cannot save with invalid EUA ID: %s", tc), func() {
			partialIntake.EUAUserID = "F"

			err := s.store.SaveSystemIntake(&partialIntake)

			s.Error(err)
			s.Equal("pq: new row for relation \"system_intake\" violates check constraint \"eua_id_check\"", err.Error())
		})
	}

	s.Run("cannot save with invalid status", func() {
		partialIntake.EUAUserID = "FAKE"
		partialIntake.Status = "fakeStatus"

		err := s.store.SaveSystemIntake(&partialIntake)

		s.Error(err)
		s.Equal("pq: invalid input value for enum system_intake_status: \"fakeStatus\"", err.Error())
	})

	s.Run("save a partial system intake", func() {
		partialIntake.Status = models.SystemIntakeStatusDRAFT
		partialIntake.Requester = null.StringFrom("Test Requester")

		err := s.store.SaveSystemIntake(&partialIntake)

		s.NoError(err, "failed to save system intake")
		var actualIntake models.SystemIntake
		err = s.db.Get(&actualIntake, "SELECT * FROM system_intake WHERE id=$1", partialIntake.ID)
		s.NoError(err, "failed to fetch saved intake")
		s.Equal(actualIntake, partialIntake)
	})

	s.Run("update a partial system intake", func() {
		partialIntake.Requester = null.StringFrom("Fix Requester")

		err := s.store.SaveSystemIntake(&partialIntake)

		s.NoError(err, "failed to save system intake")
		var actualIntake models.SystemIntake
		err = s.db.Get(&actualIntake, "SELECT * FROM system_intake WHERE id=$1", partialIntake.ID)
		s.NoError(err, "failed to fetch saved intake")
		s.Equal(actualIntake, partialIntake)
	})

	s.Run("EUA ID will not update", func() {
		originalEUA := partialIntake.EUAUserID
		partialIntake.EUAUserID = "NEWS"

		err := s.store.SaveSystemIntake(&partialIntake)

		s.NoError(err, "failed to save system intake")
		var actualIntake models.SystemIntake
		err = s.db.Get(&actualIntake, "SELECT * FROM system_intake WHERE id=$1", partialIntake.ID)
		s.NoError(err, "failed to fetch saved intake")
		s.Equal(originalEUA, actualIntake.EUAUserID)
	})
}

func (s StoreTestSuite) TestFetchSystemIntakeByID() {
	s.Run("golden path to fetch a system intake", func() {
		intake := testhelpers.NewSystemIntake()
		id := intake.ID
		tx := s.db.MustBegin()
		_, err := tx.NamedExec("INSERT INTO system_intake (id, eua_user_id, status) VALUES (:id, :eua_user_id, :status)", &intake)
		s.NoError(err)
		err = tx.Commit()
		s.NoError(err)

		fetched, err := s.store.FetchSystemIntakeByID(id)

		s.NoError(err, "failed to fetch system intake")
		s.Equal(intake.ID, fetched.ID)
		s.Equal(intake.EUAUserID, fetched.EUAUserID)
	})

	s.Run("cannot without an ID that exists in the db", func() {
		badUUID, _ := uuid.Parse("")
		fetched, err := s.store.FetchSystemIntakeByID(badUUID)

		s.Error(err)
		s.Equal("sql: no rows in result set", err.Error())
		s.Equal(&models.SystemIntake{}, fetched)
	})
}

func (s StoreTestSuite) TestFetchSystemIntakesByEuaID() {
	s.Run("golden path to fetch system intakes", func() {
		intake := testhelpers.NewSystemIntake()
		intake2 := testhelpers.NewSystemIntake()
		intake2.EUAUserID = intake.EUAUserID
		tx := s.db.MustBegin()
		_, err := tx.NamedExec("INSERT INTO system_intake (id, eua_user_id, status) VALUES (:id, :eua_user_id, :status)", &intake)
		s.NoError(err)
		_, err = tx.NamedExec("INSERT INTO system_intake (id, eua_user_id, status) VALUES (:id, :eua_user_id, :status)", &intake2)
		s.NoError(err)
		err = tx.Commit()
		s.NoError(err)

		fetched, err := s.store.FetchSystemIntakesByEuaID(intake.EUAUserID)

		s.NoError(err, "failed to fetch system intakes")
		s.Len(fetched, 2)
		s.Equal(intake.EUAUserID, fetched[0].EUAUserID)
	})

	s.Run("fetches no results with other EUA ID", func() {
		fetched, err := s.store.FetchSystemIntakesByEuaID(testhelpers.RandomEUAID())

		s.NoError(err)
		s.Len(fetched, 0)
		s.Equal(models.SystemIntakes{}, fetched)
	})
}

func (s StoreTestSuite) TestFetchSystemIntakeMetrics() {
	// create a random year to avoid test collisions
	// uses postgres max year minus 1000000
	rand.Seed(time.Now().UnixNano())
	endYear := rand.Intn(294276)
	endDate := time.Date(endYear, 0, 0, 0, 0, 0, 0, time.UTC)
	startDate := endDate.AddDate(0, -1, 0)
	var startedTests = []struct {
		name          string
		createdAt     time.Time
		expectedCount int
	}{
		{"start time is included", startDate, 1},
		{"end time is not included", endDate, 1},
		{"mid time is included", startDate.AddDate(0, 0, 1), 2},
		{"before time is not included", startDate.AddDate(0, 0, -1), 2},
		{"after time is not included", endDate.AddDate(0, 0, 1), 2},
	}
	for _, tt := range startedTests {
		s.Run(fmt.Sprintf("%s for started count", tt.name), func() {
			intake := testhelpers.NewSystemIntake()
			intake.CreatedAt = &tt.createdAt
			err := s.store.SaveSystemIntake(&intake)
			s.NoError(err)

			metrics, err := s.store.FetchSystemIntakeMetrics(startDate, endDate)

			s.NoError(err)
			s.Equal(tt.expectedCount, metrics.Started)
		})
	}

	endYear = rand.Intn(294276)
	endDate = time.Date(endYear, 0, 0, 0, 0, 0, 0, time.UTC)
	startDate = endDate.AddDate(0, -1, 0)
	var completedTests = []struct {
		name          string
		createdAt     time.Time
		submittedAt   time.Time
		expectedCount int
	}{
		{
			"started but not finished is not included",
			startDate,
			endDate.AddDate(0, 0, 1),
			0,
		},
		{
			"started and finished is included",
			startDate,
			startDate.AddDate(0, 0, 1),
			1,
		},
		{
			"started before is not included",
			startDate.AddDate(0, 0, -1),
			startDate.AddDate(0, 0, 1),
			1,
		},
	}
	for _, tt := range completedTests {
		s.Run(fmt.Sprintf("%s for completed count", tt.name), func() {
			intake := testhelpers.NewSystemIntake()
			intake.CreatedAt = &tt.createdAt
			intake.SubmittedAt = &tt.submittedAt
			err := s.store.SaveSystemIntake(&intake)
			s.NoError(err)

			metrics, err := s.store.FetchSystemIntakeMetrics(startDate, endDate)

			s.NoError(err)
			s.Equal(tt.expectedCount, metrics.CompletedOfStarted)
		})
	}

	endYear = rand.Intn(294276)
	endDate = time.Date(endYear, 0, 0, 0, 0, 0, 0, time.UTC)
	startDate = endDate.AddDate(0, -1, 0)
	var fundedTests = []struct {
		name           string
		submittedAt    time.Time
		funded         bool
		completedCount int
		fundedCount    int
	}{
		{
			"completed out of range and funded",
			endDate.AddDate(0, 0, 1),
			true,
			0,
			0,
		},
		{
			"completed in range and funded",
			startDate,
			true,
			1,
			1,
		},
		{
			"completed in range and not funded",
			startDate,
			false,
			2,
			1,
		},
	}
	for _, tt := range fundedTests {
		s.Run(tt.name, func() {
			intake := testhelpers.NewSystemIntake()
			intake.SubmittedAt = &tt.submittedAt
			intake.ExistingFunding = null.BoolFrom(tt.funded)
			err := s.store.SaveSystemIntake(&intake)
			s.NoError(err)

			metrics, err := s.store.FetchSystemIntakeMetrics(startDate, endDate)

			s.NoError(err)
			s.Equal(tt.completedCount, metrics.Completed)
			s.Equal(tt.fundedCount, metrics.Funded)
		})
	}
}
