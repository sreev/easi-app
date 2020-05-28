package services

import (
	"context"
	"errors"

	"github.com/facebookgo/clock"
	"github.com/google/uuid"
	"github.com/guregu/null"
	"go.uber.org/zap"

	"github.com/cmsgov/easi-app/pkg/appcontext"
	"github.com/cmsgov/easi-app/pkg/apperrors"
	"github.com/cmsgov/easi-app/pkg/models"
)

func (s ServicesTestSuite) TestSystemIntakesByEuaIDFetcher() {
	logger := zap.NewNop()
	fakeEuaID := "FAKE"
	serviceConfig := NewConfig(logger)
	serviceConfig.clock = clock.NewMock()

	s.Run("successfully fetches System Intakes by EUA ID without an error", func() {
		fetch := func(euaID string) (models.SystemIntakes, error) {
			return models.SystemIntakes{
				models.SystemIntake{
					EUAUserID: fakeEuaID,
				},
			}, nil
		}
		fetchSystemIntakesByEuaID := NewFetchSystemIntakesByEuaID(serviceConfig, fetch)
		intakes, err := fetchSystemIntakesByEuaID(fakeEuaID)
		s.NoError(err)
		s.Equal(fakeEuaID, intakes[0].EUAUserID)
	})

	s.Run("returns query error when fetch fails", func() {
		fetch := func(euaID string) (models.SystemIntakes, error) {
			return models.SystemIntakes{}, errors.New("fetch failed")
		}
		fetchSystemIntakesByEuaID := NewFetchSystemIntakesByEuaID(serviceConfig, fetch)
		intakes, err := fetchSystemIntakesByEuaID("FAKE")

		s.IsType(&apperrors.QueryError{}, err)
		s.Equal(models.SystemIntakes{}, intakes)
	})
}

func (s ServicesTestSuite) TestAuthorizeSaveSystemIntake() {
	logger := zap.NewNop()
	authorizeSaveSystemIntake := NewAuthorizeSaveSystemIntake(logger)

	s.Run("No EUA ID fails auth", func() {
		ctx := context.Background()

		ok, err := authorizeSaveSystemIntake(ctx, &models.SystemIntake{})

		s.False(ok)
		s.IsType(&apperrors.ContextError{}, err)
	})

	s.Run("Mismatched EUA ID fails auth", func() {
		ctx := context.Background()
		ctx = appcontext.WithUser(ctx, models.User{EUAUserID: "ZYXW"})
		intake := models.SystemIntake{
			EUAUserID: "ABCD",
		}

		ok, err := authorizeSaveSystemIntake(ctx, &intake)

		s.False(ok)
		s.NoError(err)
	})

	s.Run("Matched EUA ID passes auth", func() {
		ctx := context.Background()
		ctx = appcontext.WithUser(ctx, models.User{EUAUserID: "ABCD"})
		intake := models.SystemIntake{
			EUAUserID: "ABCD",
		}

		ok, err := authorizeSaveSystemIntake(ctx, &intake)

		s.True(ok)
		s.NoError(err)
	})
}

func (s ServicesTestSuite) TestNewSaveSystemIntake() {
	logger := zap.NewNop()
	fetch := func(id uuid.UUID) (*models.SystemIntake, error) {
		return nil, nil
	}
	save := func(intake *models.SystemIntake) error {
		return nil
	}
	authorize := func(ctx context.Context, intake *models.SystemIntake) (bool, error) {
		return true, nil
	}
	submit := func(intake *models.SystemIntake, logger2 *zap.Logger) (string, error) {
		return "ALFABET-ID", nil
	}
	emailCount := 0
	sendEmail := func(requester string, intakeID uuid.UUID) error {
		emailCount++
		return nil
	}
	serviceConfig := NewConfig(logger)
	serviceConfig.clock = clock.NewMock()

	s.Run("returns no error when successful on save", func() {
		ctx := context.Background()
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, fetch, authorize, submit, sendEmail)

		err := saveSystemIntake(ctx, &models.SystemIntake{})

		s.NoError(err)
	})

	s.Run("returns no error when successful on submit and save", func() {
		ctx := context.Background()
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, fetch, authorize, submit, sendEmail)
		s.Equal(0, emailCount)

		err := saveSystemIntake(ctx, &models.SystemIntake{Status: models.SystemIntakeStatusSUBMITTED})

		s.NoError(err)
		s.Equal(1, emailCount)

		emailCount = 0
	})

	s.Run("returns query error when fetch fails", func() {
		ctx := context.Background()
		failFetch := func(uuid uuid.UUID) (*models.SystemIntake, error) {
			return nil, errors.New("failed to fetch system intake")
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, failFetch, authorize, submit, sendEmail)

		err := saveSystemIntake(ctx, &models.SystemIntake{})

		s.IsType(&apperrors.QueryError{}, err)
	})

	s.Run("returns query error when save fails", func() {
		ctx := context.Background()
		failSave := func(intake *models.SystemIntake) error {
			return errors.New("save failed")
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, failSave, fetch, authorize, submit, sendEmail)

		err := saveSystemIntake(ctx, &models.SystemIntake{})

		s.IsType(&apperrors.QueryError{}, err)
	})

	s.Run("saves when fetch existing fails with no results", func() {
		ctx := context.Background()
		failFetch := func(uuid uuid.UUID) (*models.SystemIntake, error) {
			return nil, errors.New("sql: no rows in result set")
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, failFetch, authorize, submit, sendEmail)

		err := saveSystemIntake(ctx, &models.SystemIntake{})

		s.NoError(err)
	})

	s.Run("returns error when authorization errors", func() {
		ctx := context.Background()
		err := errors.New("authorization failed")
		failAuthorize := func(ctx context.Context, intake *models.SystemIntake) (bool, error) {
			return false, err
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, fetch, failAuthorize, submit, sendEmail)

		actualError := saveSystemIntake(ctx, &models.SystemIntake{})

		s.Error(err)
		s.Equal(err, actualError)
	})

	s.Run("returns unauthorized error when authorization not ok", func() {
		ctx := context.Background()
		notOKAuthorize := func(ctx context.Context, intake *models.SystemIntake) (bool, error) {
			return false, nil
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, fetch, notOKAuthorize, submit, sendEmail)

		err := saveSystemIntake(ctx, &models.SystemIntake{})

		s.IsType(&apperrors.UnauthorizedError{}, err)
	})

	s.Run("returns query error when fetch fails", func() {
		ctx := context.Background()
		failFetch := func(id uuid.UUID) (*models.SystemIntake, error) {
			return &models.SystemIntake{}, errors.New("fetch failed")
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, failFetch, authorize, submit, sendEmail)

		err := saveSystemIntake(ctx, &models.SystemIntake{})

		s.IsType(&apperrors.QueryError{}, err)
	})

	s.Run("returns error when validation fails", func() {
		ctx := context.Background()
		failValidationSubmit := func(intake *models.SystemIntake, logger *zap.Logger) (string, error) {
			return "", &apperrors.ValidationError{
				Err:     errors.New("validation failed on these fields: ID"),
				ModelID: intake.ID.String(),
				Model:   intake,
			}
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, fetch, authorize, failValidationSubmit, sendEmail)

		err := saveSystemIntake(ctx, &models.SystemIntake{Status: models.SystemIntakeStatusSUBMITTED})

		s.IsType(&apperrors.ValidationError{}, err)
		s.Equal(0, emailCount)
	})

	s.Run("returns error when submission fails", func() {
		ctx := context.Background()
		failValidationSubmit := func(intake *models.SystemIntake, logger *zap.Logger) (string, error) {
			return "", &apperrors.ExternalAPIError{
				Err:       errors.New("CEDAR return result: unexpected failure"),
				ModelID:   intake.ID.String(),
				Model:     intake,
				Operation: apperrors.Submit,
				Source:    "CEDAR",
			}
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, fetch, authorize, failValidationSubmit, sendEmail)

		err := saveSystemIntake(ctx, &models.SystemIntake{Status: models.SystemIntakeStatusSUBMITTED})

		s.IsType(&apperrors.ExternalAPIError{}, err)
		s.Equal(0, emailCount)
	})

	s.Run("returns error when intake has already been submitted", func() {
		ctx := context.Background()
		alreadySubmittedIntake := models.SystemIntake{
			AlfabetID: null.StringFrom("394-141-0"),
			ID:        uuid.New(),
			Status:    models.SystemIntakeStatusSUBMITTED,
			EUAUserID: "EUAI",
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, fetch, authorize, submit, sendEmail)

		err := saveSystemIntake(ctx, &alreadySubmittedIntake)

		s.IsType(&apperrors.ResourceConflictError{}, err)
		s.Equal(0, emailCount)
	})

	s.Run("returns notification error when email fails", func() {
		ctx := context.Background()
		failSendEmail := func(requester string, intakeID uuid.UUID) error {
			return &apperrors.NotificationError{
				Err:             errors.New("failed to send Email"),
				DestinationType: apperrors.DestinationTypeEmail,
			}
		}
		saveSystemIntake := NewSaveSystemIntake(serviceConfig, save, fetch, authorize, submit, failSendEmail)

		err := saveSystemIntake(ctx, &models.SystemIntake{Status: models.SystemIntakeStatusSUBMITTED})

		s.IsType(&apperrors.NotificationError{}, err)
	})
}

func (s ServicesTestSuite) TestSystemIntakeByIDFetcher() {
	logger := zap.NewNop()
	fakeID := uuid.New()
	serviceConfig := NewConfig(logger)
	serviceConfig.clock = clock.NewMock()

	s.Run("successfully fetches System Intake by ID without an error", func() {
		fetch := func(id uuid.UUID) (*models.SystemIntake, error) {
			return &models.SystemIntake{
				ID: fakeID,
			}, nil
		}
		fetchSystemIntakeByID := NewFetchSystemIntakeByID(serviceConfig, fetch)
		intake, err := fetchSystemIntakeByID(fakeID)
		s.NoError(err)

		s.Equal(fakeID, intake.ID)
	})

	s.Run("returns query error when save fails", func() {
		fetch := func(id uuid.UUID) (*models.SystemIntake, error) {
			return &models.SystemIntake{}, errors.New("save failed")
		}
		fetchSystemIntakeByID := NewFetchSystemIntakeByID(serviceConfig, fetch)

		intake, err := fetchSystemIntakeByID(uuid.New())

		s.IsType(&apperrors.QueryError{}, err)
		s.Equal(&models.SystemIntake{}, intake)
	})
}
