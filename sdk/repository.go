package isegosdk

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type RepositoryService service

type ResponseRepositoryGetRepositories struct {
	Response *[]ResponseRepositoryGetRepositoriesResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}

type ResponseRepositoryGetRepositoriesResponse struct {
	Name       string `json:"name,omitempty"`       // Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Protocol   string `json:"protocol,omitempty"`   //
	Path       string `json:"path,omitempty"`       // Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Password   string `json:"password,omitempty"`   // Password can contain alphanumeric and/or special characters.
	ServerName string `json:"serverName,omitempty"` //
	UserName   string `json:"userName,omitempty"`   // Username can contain alphanumeric characters.
	EnablePki  *bool  `json:"enablePki,omitempty"`  //
}

type ResponseRepositoryCreateRepository struct {
	Success *ResponseRepositoryCreateRepositorySuccess `json:"success,omitempty"` //
	Version string                                     `json:"version,omitempty"` //
}

type ResponseRepositoryCreateRepositorySuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseRepositoryGetRepository struct {
	Response *ResponseRepositoryGetRepositoryResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}

type ResponseRepositoryGetRepositoryResponse struct {
	Name       string `json:"name,omitempty"`       // Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Protocol   string `json:"protocol,omitempty"`   //
	Path       string `json:"path,omitempty"`       // Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Password   string `json:"password,omitempty"`   // Password can contain alphanumeric and/or special characters.
	ServerName string `json:"serverName,omitempty"` //
	UserName   string `json:"userName,omitempty"`   // Username can contain alphanumeric characters.
	EnablePki  *bool  `json:"enablePki,omitempty"`  //
}

type ResponseRepositoryUpdateRepository struct {
	Success *ResponseRepositoryUpdateRepositorySuccess `json:"success,omitempty"` //
	Version string                                     `json:"version,omitempty"` //
}

type ResponseRepositoryUpdateRepositorySuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseRepositoryDeleteRepository struct {
	Success *ResponseRepositoryDeleteRepositorySuccess `json:"success,omitempty"` //
	Version string                                     `json:"version,omitempty"` //
}

type ResponseRepositoryDeleteRepositorySuccess struct {
	Message string `json:"message,omitempty"` //
}

type ResponseRepositoryGetRepositoryFiles struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

type RequestRepositoryCreateRepository struct {
	Name       string `json:"name,omitempty"`       // Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Protocol   string `json:"protocol,omitempty"`   //
	Path       string `json:"path,omitempty"`       // Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Password   string `json:"password,omitempty"`   // Password can contain alphanumeric and/or special characters.
	ServerName string `json:"serverName,omitempty"` //
	UserName   string `json:"userName,omitempty"`   // Username can contain alphanumeric characters.
	EnablePki  *bool  `json:"enablePki,omitempty"`  //
}

type RequestRepositoryUpdateRepository struct {
	Name       string `json:"name,omitempty"`       // Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Protocol   string `json:"protocol,omitempty"`   //
	Path       string `json:"path,omitempty"`       // Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Password   string `json:"password,omitempty"`   // Password can contain alphanumeric and/or special characters.
	ServerName string `json:"serverName,omitempty"` //
	UserName   string `json:"userName,omitempty"`   // Username can contain alphanumeric characters.
	EnablePki  *bool  `json:"enablePki,omitempty"`  //
}

//GetRepositories Get list of repositories
/* This will get the full list of repository definitions on the system.


 */
func (s *RepositoryService) GetRepositories() (*ResponseRepositoryGetRepositories, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/repository"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseRepositoryGetRepositories{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRepositories")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRepositoryGetRepositories)
	return result, response, err

}

//GetRepository Get a specific repository
/* Get a specific repository identified by the name passed in the URL.


@param name name path parameter. Unique name for a repository
*/
func (s *RepositoryService) GetRepository(name string) (*ResponseRepositoryGetRepository, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/repository/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseRepositoryGetRepository{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRepository")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRepositoryGetRepository)
	return result, response, err

}

//GetRepositoryFiles Get list of files in a repository
/* This will get the full list of files present in the named repository.


@param name name path parameter. Unique name for a repository
*/
func (s *RepositoryService) GetRepositoryFiles(name string) (*ResponseRepositoryGetRepositoryFiles, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/repository/{name}/files"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseRepositoryGetRepositoryFiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRepositoryFiles")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRepositoryGetRepositoryFiles)
	return result, response, err

}

//CreateRepository Create a new repository
/* Create a new repository in the system. The name provided for the
repository must be unique.


*/
func (s *RepositoryService) CreateRepository(requestRepositoryCreateRepository *RequestRepositoryCreateRepository) (*ResponseRepositoryCreateRepository, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/repository"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestRepositoryCreateRepository).
		SetResult(&ResponseRepositoryCreateRepository{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateRepository")
	}

	result := response.Result().(*ResponseRepositoryCreateRepository)
	return result, response, err

}

//UpdateRepository Update the definition of a specific repository
/* Update the definition of a specific repository, providing ALL parameters for the repository.


@param name name path parameter. Unique name for a repository
*/
func (s *RepositoryService) UpdateRepository(name string, requestRepositoryUpdateRepository *RequestRepositoryUpdateRepository) (*ResponseRepositoryUpdateRepository, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/repository/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestRepositoryUpdateRepository).
		SetResult(&ResponseRepositoryUpdateRepository{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateRepository")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRepositoryUpdateRepository)
	return result, response, err

}

//DeleteRepository Delete a specific repository
/* Long description TBD


@param name name path parameter. Unique name for a repository
*/
func (s *RepositoryService) DeleteRepository(name string) (*ResponseRepositoryDeleteRepository, *resty.Response, error) {
	setHost(s.client, "_ui")
	path := "/api/v1/repository/{name}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseRepositoryDeleteRepository{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteRepository")
	}

	getCSFRToken(response.Header())

	result := response.Result().(*ResponseRepositoryDeleteRepository)
	return result, response, err

}
