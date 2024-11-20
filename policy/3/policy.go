package policy

import (
	"errors"
)

type PolicySpec struct {
	APIKey *APIKey
}

type APIKey struct {
	SuppliedIn   *SuppliedIn
	ClientSecret string
}

type SuppliedIn struct {
	Header []string
	Query  []string
}

func ValidatePolicySpec(spec PolicySpec) error {
	if spec.APIKey != nil {
		err := ValidateAPIKey(spec.APIKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func ValidateAPIKey(apiKey *APIKey) error {
	if apiKey.SuppliedIn.Query == nil && apiKey.SuppliedIn.Header == nil {
		return errors.New("missing fields suppliedIn, at least one query or header name must be provided")
	}

	if apiKey.SuppliedIn.Header != nil {
		for _, header := range apiKey.SuppliedIn.Header {
			err := isHTTPHeader(header)
			if err != nil {
				return err
			}
		}
	}

	if apiKey.SuppliedIn.Query != nil {
		for _, query := range apiKey.SuppliedIn.Query {
			err := isValidQuery(query)
			if err != nil {
				return err
			}
		}
	}

	if apiKey.ClientSecret == "" {
		return errors.New("clientSecret cannot be empty")
	}

	return nil
}

func isHTTPHeader(header string) error {
	// concrete implementation goes here
	return nil
}

func isValidQuery(query string) error {
	// concrete implementation goes here
	return nil
}
