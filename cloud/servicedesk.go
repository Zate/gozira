package cloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/google/go-querystring/query"
)

// ServiceDeskService handles ServiceDesk for the Jira instance / API.
type ServiceDeskService service

// ServiceDeskOrganizationDTO is a DTO for ServiceDesk organizations
type ServiceDeskOrganizationDTO struct {
	OrganizationID int `json:"organizationId,omitempty" structs:"organizationId,omitempty"`
}

// GetOrganizations returns a list of
// all organizations associated with a service desk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-servicedesk-servicedeskid-organization-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) GetOrganizations(ctx context.Context, serviceDeskID interface{}, start int, limit int, accountID string) (*PagedDTO, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization?start=%d&limit=%d", serviceDeskID, start, limit)
	if accountID != "" {
		apiEndPoint += fmt.Sprintf("&accountId=%s", accountID)
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndPoint, nil)
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return nil, nil, err
	}

	orgs := new(PagedDTO)
	resp, err := s.client.Do(req, &orgs)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return orgs, resp, nil
}

// AddOrganization adds an organization to
// a service desk. If the organization ID is already
// associated with the service desk, no change is made
// and the resource returns a 204 success code.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-servicedesk-servicedeskid-organization-post
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) AddOrganization(ctx context.Context, serviceDeskID interface{}, organizationID int) (*Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", serviceDeskID)

	organization := ServiceDeskOrganizationDTO{
		OrganizationID: organizationID,
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndPoint, organization)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return resp, jerr
	}

	return resp, nil
}

// RemoveOrganization removes an organization
// from a service desk. If the organization ID does not
// match an organization associated with the service desk,
// no change is made and the resource returns a 204 success code.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-organization/#api-rest-servicedeskapi-servicedesk-servicedeskid-organization-delete
// Caller must close resp.Body
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) RemoveOrganization(ctx context.Context, serviceDeskID interface{}, organizationID int) (*Response, error) {
	apiEndPoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", serviceDeskID)

	organization := ServiceDeskOrganizationDTO{
		OrganizationID: organizationID,
	}

	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndPoint, organization)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return resp, jerr
	}

	return resp, nil
}

// AddCustomers adds customers to the given service desk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-customer-post
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) AddCustomers(ctx context.Context, serviceDeskID interface{}, acountIDs ...string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	payload := struct {
		AccountIDs []string `json:"accountIds"`
	}{
		AccountIDs: acountIDs,
	}
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, payload)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, NewJiraError(resp, err)
	}

	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp, nil
}

// RemoveCustomers removes customers to the given service desk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-customer-delete
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) RemoveCustomers(ctx context.Context, serviceDeskID interface{}, acountIDs ...string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	payload := struct {
		AccountIDs []string `json:"accountIDs"`
	}{
		AccountIDs: acountIDs,
	}
	req, err := s.client.NewRequest(ctx, http.MethodDelete, apiEndpoint, payload)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, NewJiraError(resp, err)
	}

	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp, nil
}

// ListCustomers lists customers for a ServiceDesk.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-customer-get
//
// TODO Double check this method if this works as expected, is using the latest API and the response is complete
// This double check effort is done for v2 - Remove this two lines if this is completed.
func (s *ServiceDeskService) ListCustomers(ctx context.Context, serviceDeskID interface{}, options *CustomerListOptions) (*CustomerList, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	// this is an experiemntal endpoint
	req.Header.Set("X-ExperimentalApi", "opt-in")

	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, err
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	customerList := new(CustomerList)
	if err := json.NewDecoder(resp.Body).Decode(customerList); err != nil {
		return nil, resp, fmt.Errorf("could not unmarshall the data into struct")
	}

	return customerList, resp, nil
}

// ServiceDeskInfo is a DTO for ServiceDesk information
type ServiceDeskInfo struct {
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"_links,omitempty"`
	BuildChangeSet string `json:"buildChangeSet,omitempty"`
	BuildDate      struct {
		EpochMillis int64  `json:"epochMillis,omitempty"`
		Friendly    string `json:"friendly,omitempty"`
		Iso8601     string `json:"iso8601,omitempty"`
		Jira        string `json:"jira,omitempty"`
	} `json:"buildDate,omitempty"`
	IsLicensedForUse bool   `json:"isLicensedForUse,omitempty"`
	PlatformVersion  string `json:"platformVersion,omitempty"`
	Version          string `json:"version,omitempty"`
}

// ServiceDeskinfo retrieves information about the Jira Service Management instance such as software version, builds, and related links.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-info-get
func (s *ServiceDeskService) ServiceDeskinfo(ctx context.Context) (*ServiceDeskInfo, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "rest/servicedeskapi/info", nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	serviceDeskInfo := new(ServiceDeskInfo)
	if err := json.NewDecoder(resp.Body).Decode(serviceDeskInfo); err != nil {
		return nil, resp, fmt.Errorf("could not unmarshall the data into struct")
	}

	return serviceDeskInfo, resp, nil
}

// ServiceDeskListOptions specifies the optional parameters to the ServiceDeskService.ListserviceDesks method.
type ServiceDeskListOptions struct {
	Start int `url:"start,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// ServiceDeskList represents a list of service desks in the Jira Service Management instance.
type ServiceDeskList struct {
	Expands    string `json:"_expands,omitempty"`
	Size       int    `json:"size,omitempty"`
	Start      int    `json:"start,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	IsLastPage bool   `json:"isLastPage,omitempty"`
	Links      struct {
		Base    string `json:"base,omitempty"`
		Context string `json:"context,omitempty"`
		Next    string `json:"next,omitempty"`
		Prev    string `json:"prev,omitempty"`
	} `json:"_links,omitempty"`
	Values []struct {
		ID          string `json:"id,omitempty"`
		ProjectID   string `json:"projectId,omitempty"`
		ProjectName string `json:"projectName,omitempty"`
		ProjectKey  string `json:"projectKey,omitempty"`
		Links       struct {
			Self string `json:"self,omitempty"`
		} `json:"_links,omitempty"`
	} `json:"values,omitempty"`
}

// ListServiceDesks returns a list of all service desks in the Jira Service Management instance.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-get
func (s *ServiceDeskService) ListServiceDesks(ctx context.Context, options *ServiceDeskListOptions) (*ServiceDeskList, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "rest/servicedeskapi/servicedesk", nil)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create a new request: %w", err)
	}

	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, fmt.Errorf("could not create query values: %w", err)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		reqDump, _ := httputil.DumpRequestOut(resp.Request, true)
		if reqDump != nil {
			fmt.Printf("Request Dump: %s\n", reqDump)
		}

		respDump, _ := httputil.DumpResponse(resp.Response, true)
		if respDump != nil {
			fmt.Printf("Response Dump: %s\n", respDump)
		}

		return nil, resp, NewJiraError(resp, fmt.Errorf("client.do failed in ListServiceDesks: %w", err))
	}
	defer resp.Body.Close()

	if resp.Response.StatusCode != http.StatusOK {
		return nil, resp, NewJiraError(resp, fmt.Errorf("could not get service desk info"))
	}

	serviceDeskList := new(ServiceDeskList)

	if err := json.NewDecoder(resp.Body).Decode(serviceDeskList); err != nil {
		return nil, resp, fmt.Errorf("could not unmarshall the data into struct")
	}

	return serviceDeskList, resp, nil

}

// ServiceDesk is a struct that represents a JIRA Service Desk
type ServiceDesk struct {
	ID          string `json:"id,omitempty"`
	ProjectID   string `json:"projectId,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	ProjectKey  string `json:"projectKey,omitempty"`
	Links       struct {
		Self string `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

// GetServiceDeskByID returns a service desk by its ID.
//
// https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-servicedesk/#api-rest-servicedeskapi-servicedesk-servicedeskid-get
func (s *ServiceDeskService) GetServiceDeskByID(ctx context.Context, serviceDeskID interface{}) (*ServiceDesk, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v", serviceDeskID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	defer resp.Body.Close()

	serviceDesk := new(ServiceDesk)
	if err := json.NewDecoder(resp.Body).Decode(serviceDesk); err != nil {
		return nil, resp, fmt.Errorf("could not unmarshall the data into struct")
	}

	return serviceDesk, resp, nil
}
