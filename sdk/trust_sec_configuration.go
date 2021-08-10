package isegosdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type TrustSecConfigurationService service

//GetSecurityGroups 🚧 getSecurityGroups
/* 🚧 getSecurityGroups

 */
func (s *TrustSecConfigurationService) GetSecurityGroups() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/config/trustsec/getSecurityGroups"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSecurityGroups")
	}

	return response, err

}

//GetSecurityGroupACLs 🚧 getSecurityGroupAcls
/* 🚧 getSecurityGroupAcls

 */
func (s *TrustSecConfigurationService) GetSecurityGroupACLs() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/config/trustsec/getSecurityGroupAcls"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetSecurityGroupAcls")
	}

	return response, err

}

//GetEgressPolicies 🚧 getEgressPolicies
/* 🚧 getEgressPolicies

 */
func (s *TrustSecConfigurationService) GetEgressPolicies() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/config/trustsec/getEgressPolicies"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetEgressPolicies")
	}

	return response, err

}

//GetEgressMatrices 🚧 getEgressMatrices
/* 🚧 getEgressMatrices

 */
func (s *TrustSecConfigurationService) GetEgressMatrices() (*resty.Response, error) {
	setHost(s.client, "_px_grid")
	path := "/pxgrid/ise/radius/ise/config/trustsec/getEgressMatrices"

	setCSRFToken(s.client)
	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	getCSFRToken(response.Header())
	if response.IsError() {
		return response, fmt.Errorf("error with operation GetEgressMatrices")
	}

	return response, err

}
