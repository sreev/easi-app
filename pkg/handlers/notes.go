package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/cmsgov/easi-app/pkg/apperrors"
	"github.com/cmsgov/easi-app/pkg/models"
)

type fetchNotes func(context.Context, uuid.UUID) ([]*models.Note, error)
type createNote func(context.Context, *models.Note) (*models.Note, error)

// NewNotesHandler is a constructor for SystemListHandler
func NewNotesHandler(base HandlerBase, fetch fetchNotes, create createNote) NotesHandler {
	return NotesHandler{
		HandlerBase: base,
		FetchNotes:  fetch,
		CreateNote:  create,
	}
}

// NotesHandler is the handler for interacting with admin Notes
// associated with a SystemIntake
type NotesHandler struct {
	HandlerBase
	FetchNotes fetchNotes
	CreateNote createNote
}

// Handle handles a web request and returns a list of systems
func (h NotesHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["intake_id"]
		valErr := apperrors.NewValidationError(
			errors.New("system intake failed validation"),
			models.SystemIntake{},
			"",
		)
		if id == "" {
			valErr.WithValidation("path.intakeID", "is required")
			h.WriteErrorResponse(r.Context(), w, &valErr)
			return
		}
		uuid, err := uuid.Parse(id)
		if err != nil {
			valErr.WithValidation("path.intakeID", "must be UUID")
			h.WriteErrorResponse(r.Context(), w, &valErr)
			return
		}

		switch r.Method {
		case "GET":
			notes, err := h.FetchNotes(r.Context(), uuid)
			if err != nil {
				h.WriteErrorResponse(r.Context(), w, err)
				return
			}

			js, err := json.Marshal(notes)
			if err != nil {
				h.WriteErrorResponse(r.Context(), w, err)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			_, err = w.Write(js)
			if err != nil {
				h.WriteErrorResponse(r.Context(), w, err)
				return
			}

		case "POST":
			if r.Body == nil {
				h.WriteErrorResponse(
					r.Context(),
					w,
					&apperrors.BadRequestError{Err: errors.New("empty request not allowed")},
				)
				return
			}
			defer r.Body.Close()

			decoder := json.NewDecoder(r.Body)
			note := models.Note{}

			err := decoder.Decode(&note)
			if err != nil {
				h.WriteErrorResponse(r.Context(), w, &apperrors.BadRequestError{Err: err})
				return
			}

			note.SystemIntakeID = uuid

			createdNote, err := h.CreateNote(r.Context(), &note)
			if err != nil {
				h.WriteErrorResponse(r.Context(), w, err)
				return
			}

			responseBody, err := json.Marshal(createdNote)
			if err != nil {
				h.WriteErrorResponse(r.Context(), w, err)
				return
			}

			w.WriteHeader(http.StatusCreated)
			_, err = w.Write(responseBody)
			if err != nil {
				h.WriteErrorResponse(r.Context(), w, err)
				return
			}
			return
		default:
			h.WriteErrorResponse(r.Context(), w, &apperrors.MethodNotAllowedError{Method: r.Method})
			return
		}

	}
}
