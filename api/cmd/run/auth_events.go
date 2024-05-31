package run

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/config"
	"github.com/codelite7/momentum/api/ent"
	entuser "github.com/codelite7/momentum/api/ent/user"
	"github.com/workos/workos-go/v3/pkg/events"
	"github.com/workos/workos-go/v3/pkg/organizations"
	"github.com/workos/workos-go/v3/pkg/usermanagement"
	"time"
)

func HandleAuthEvents() {
	events.SetAPIKey(config.WorkosApiKey)
	usermanagement.SetAPIKey(config.WorkosApiKey)
	organizations.SetAPIKey(config.WorkosApiKey)
	for {
		PollForEvents()
		time.Sleep(config.WorkosPollInterval)
	}
}

func PollForEvents() error {
	eventTypes := []string{
		"user.created",
	}

	cursor, err := common.EntClient.WorkosEventCursor.Get(context.Background(), 1)
	if err != nil {
		if ent.IsNotFound(err) {
			cursor, err = common.EntClient.WorkosEventCursor.Create().Save(context.Background())
			if err != nil {
				return errors.New("error creating new workos event cursor: " + err.Error())
			}
		} else {
			return errors.New("failed to get workos event cursor: " + err.Error())
		}
	}
	opts := events.ListEventsOpts{
		Events: eventTypes,
	}
	if cursor.UserCreatedCursor != "" {
		opts.After = cursor.UserCreatedCursor
	}
	events, err := events.ListEvents(context.Background(), opts)
	if err != nil {
		return errors.New("error polling for events: " + err.Error())
	}
	for _, event := range events.Data {
		err = handleEvent(event)
		if err != nil {
			return errors.New("error handling event: " + err.Error())
		}
		cursor, err = cursor.Update().SetUserCreatedCursor(event.ID).Save(context.Background())
		if err != nil {
			return errors.New("error updating cursor: " + err.Error())
		}
	}
	return nil
}

func handleEvent(event events.Event) (err error) {
	switch event.Event {
	case "user.created":
		err = handleUserCreatedEvent(event)
	default:
		err = errors.New(fmt.Sprintf("unknown event: %s", event.Event))
	}
	return
}

func handleUserCreatedEvent(event events.Event) error {
	var userEvent UserCreatedEvent
	err := json.Unmarshal(event.Data, &userEvent)
	if err != nil {
		return errors.New(fmt.Sprintf("error unmarshalling event: %s", err.Error()))
	}
	_, err = maybeCreateUser(userEvent)
	if err != nil {
		return err
	}
	err = maybeCreateOrg(userEvent)
	if err != nil {
		return err
	}
	return nil
}

func maybeCreateOrg(userEvent UserCreatedEvent) error {
	userHasOrg, err := userHasOrganization(userEvent.Id)
	if err != nil {
		return err
	}
	if !userHasOrg {
		err = bootstrapOrganizationForNewUser(userEvent)
	}
	return nil
}

func userHasOrganization(userId string) (bool, error) {
	memberships, err := usermanagement.ListOrganizationMemberships(context.Background(), usermanagement.ListOrganizationMembershipsOpts{UserID: userId})
	if err != nil {
		return false, err
	}

	return len(memberships.Data) > 0, nil

}

func bootstrapOrganizationForNewUser(userEvent UserCreatedEvent) error {
	org, err := createOrganizationForNewUser(userEvent)
	if err != nil {
		return err
	}
	err = addUserToOrganization(userEvent.Id, org.ID)
	if err != nil {
		return err
	}

	return nil
}

func createOrganizationForNewUser(userEvent UserCreatedEvent) (organizations.Organization, error) {
	return organizations.CreateOrganization(context.Background(), organizations.CreateOrganizationOpts{Name: userEvent.Email})
}

func addUserToOrganization(userId, organizationId string) error {
	_, err := usermanagement.CreateOrganizationMembership(context.Background(), usermanagement.CreateOrganizationMembershipOpts{
		UserID:         userId,
		OrganizationID: organizationId,
	})
	if err != nil {
		return err
	}
	return nil
}

func maybeCreateUser(event UserCreatedEvent) (user *ent.User, err error) {
	user, err = common.EntClient.User.Query().Where(entuser.Email(event.Email)).First(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			user, err = common.EntClient.User.Create().SetEmail(event.Email).Save(context.Background())
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	return
}

type UserCreatedEvent struct {
	Object            string    `json:"object"`
	Id                string    `json:"id"`
	Email             string    `json:"email"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	EmailVerified     bool      `json:"email_verified"`
	ProfilePictureUrl string    `json:"profile_picture_url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
