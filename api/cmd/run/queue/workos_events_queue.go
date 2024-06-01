package queue

import (
	"context"
	"errors"
	"fmt"
	"github.com/acaloiaro/neoq"
	"github.com/acaloiaro/neoq/handler"
	"github.com/acaloiaro/neoq/jobs"
	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/ent/tenant"
	entuser "github.com/codelite7/momentum/api/ent/user"
	"github.com/workos/workos-go/v4/pkg/events"
	"github.com/workos/workos-go/v4/pkg/organizations"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"strings"
	"time"
)

const workosUserCreatedEventQueueName = "workos_user_created"

func startWorkosUserCreatedEventHandler(queue neoq.Neoq) error {
	h := handler.New(workosUserCreatedEventQueueName, workosUserCreatedEventHandler)
	err := queue.Start(context.Background(), h)
	if err != nil {
		return err
	}

	return nil
}

func workosUserCreatedEventHandler(ctx context.Context) error {
	err := handleWorkosUserCreatedEvent(ctx)
	if err != nil {
		fmt.Println(fmt.Sprintf("error handling workos user created event: %v", err))
		return err
	}

	return nil
}

func handleWorkosUserCreatedEvent(ctx context.Context) error {
	job, err := jobs.FromContext(ctx)
	if err != nil {
		return err
	}
	var event events.Event
	err = convertType(job.Payload, &event)
	if err != nil {
		return err
	}
	err = handleUserCreatedEvent(event)
	if err != nil {
		return err
	}
	return nil
}

func EnqueueWorkosUserCreatedEvent(event events.Event) error {
	var payload map[string]interface{}
	err := convertType(event, &payload)
	if err != nil {
		return err
	}
	job := &jobs.Job{
		Queue:   workosUserCreatedEventQueueName,
		Payload: payload,
	}
	_, err = queue.Enqueue(context.Background(), job)
	if err != nil {
		return err
	}

	return nil
}

func handleUserCreatedEvent(event events.Event) error {
	var eventUser User
	err := convertType(event.Data, &eventUser)
	if err != nil {
		return errors.New(fmt.Sprintf("error unmarshalling event: %s", err.Error()))
	}
	workosUser, err := getWorkosUser(eventUser.Id)
	if err != nil {
		if strings.Contains(err.Error(), "User not found") {
			// if the user doesn't exist in workos anymore, there's nothing to do
			return nil
		}
		return err
	}
	// check user memberships
	memberships, err := getUserMemberships(workosUser)
	if len(memberships) > 1 {
		return errors.New(fmt.Sprintf("multiple memberships found for new user: %s", workosUser.ID))
	}
	// if no memberships, create an org for the user and add the user to it
	var org organizations.Organization
	if len(memberships) == 0 {
		org, err = createOrganizationForNewUser(workosUser)
		if err != nil {
			return err
		}
		_, err = addUserToOrganization(workosUser, org)
		if err != nil {
			return err
		}
	} else {
		org, err = getOrganization(memberships[0])
		if err != nil {
			return err
		}
	}
	// look up tenant by org id
	tenant, err := common.EntClient.Tenant.Query().Where(tenant.WorkosOrgID(org.ID)).First(context.Background())
	if err != nil {
		// if no tenant, create one
		if ent.IsNotFound(err) {
			tenant, err = common.EntClient.Tenant.Create().SetWorkosOrgID(org.ID).Save(context.Background())
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	// look up our user by workos user id
	user, err := common.EntClient.User.Query().Where(entuser.WorkosUserID(workosUser.ID)).First(context.Background())
	if err != nil {
		// if user doesn't exist, create one
		if ent.IsNotFound(err) {
			user, err = common.EntClient.User.Create().SetEmail(workosUser.Email).SetWorkosUserID(workosUser.ID).SetActiveTenant(tenant).AddTenants(tenant).Save(context.Background())
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	// set the eventUser's active tenant
	_, err = user.Update().SetActiveTenant(tenant).Save(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func getWorkosUser(userId string) (usermanagement.User, error) {
	return usermanagement.GetUser(context.Background(), usermanagement.GetUserOpts{User: userId})
}

func getUserMemberships(user usermanagement.User) ([]usermanagement.OrganizationMembership, error) {
	response, err := usermanagement.ListOrganizationMemberships(context.Background(), usermanagement.ListOrganizationMembershipsOpts{UserID: user.ID})
	if err != nil {
		return nil, err
	}

	return response.Data, nil

}

func createOrganizationForNewUser(user usermanagement.User) (organizations.Organization, error) {
	return organizations.CreateOrganization(context.Background(), organizations.CreateOrganizationOpts{Name: user.Email})
}

func addUserToOrganization(user usermanagement.User, organization organizations.Organization) (usermanagement.OrganizationMembership, error) {
	return usermanagement.CreateOrganizationMembership(context.Background(), usermanagement.CreateOrganizationMembershipOpts{
		UserID:         user.ID,
		OrganizationID: organization.ID,
	})
}

func getOrganization(membership usermanagement.OrganizationMembership) (organizations.Organization, error) {
	return organizations.GetOrganization(context.Background(), organizations.GetOrganizationOpts{Organization: membership.OrganizationID})
}

type User struct {
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
