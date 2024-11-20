package policy_test

import (
	"testing"

	"github.com/qba73/policy"
)

func TestValidatePolicy_PassesOnValidInput(t *testing.T) {
	t.Parallel()
}

func TestValidateAPIKeyPolicy_PassOnValidInput(t *testing.T) {
	t.Parallel()

	tt := []struct {
		apiKey *policy.APIKey
		name   string
	}{
		{
			apiKey: &policy.APIKey{
				SuppliedIn: &policy.SuppliedIn{
					Header: []string{"X-API-Key"},
					Query:  []string{"api_key"},
				},
				ClientSecret: "secret",
			},
			name: "policy with all values provided",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := policy.ValidateAPIKey(tc.apiKey)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestValidateAPIKeyPolicy_FailsOnInvalidInput(t *testing.T) {
	t.Parallel()
	tests := []struct {
		apiKey *policy.APIKey
		msg    string
	}{
		{
			apiKey: &policy.APIKey{
				SuppliedIn: &policy.SuppliedIn{
					Query: []string{
						"api_key",
					},
				},
			},
			msg: "missing secret",
		},
		{
			apiKey: &policy.APIKey{
				SuppliedIn:   &policy.SuppliedIn{},
				ClientSecret: "secret",
			},
			msg: "both header and query are missing",
		},
		{
			apiKey: &policy.APIKey{
				SuppliedIn: &policy.SuppliedIn{
					Header: []string{
						`api.key"`,
					},
				},
				ClientSecret: "secret",
			},
			msg: "invalid header",
		},
		{
			apiKey: &policy.APIKey{
				SuppliedIn: &policy.SuppliedIn{
					Query: []string{
						`api_key\`,
					},
				},
				ClientSecret: "secret",
			},
			msg: "invalid query",
		},
		{
			apiKey: &policy.APIKey{
				SuppliedIn: &policy.SuppliedIn{
					Query: []string{
						`api_key`,
					},
				},
				ClientSecret: "secret_1",
			},
			msg: "invalid secret name",
		},
		// {
		// 	apiKey: &policy.APIKey{
		// 		SuppliedIn:   nil,
		// 		ClientSecret: "secret",
		// 	},
		// },
		// {
		// 	apiKey: &policy.APIKey{
		// 		ClientSecret: "secret_1",
		// 	},
		// 	msg: "no suppliedIn provided",
		// },
		// {
		// 	apiKey: nil, msg: "no apikey provided",
		// },
	}

	for _, tc := range tests {
		t.Run(tc.msg, func(t *testing.T) {
			err := policy.ValidateAPIKey(tc.apiKey)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
