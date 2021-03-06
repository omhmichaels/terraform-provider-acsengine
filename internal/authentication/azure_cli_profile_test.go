package authentication

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/azure/cli"
)

func TestAzureCLIProfileFindDefaultSubscription(t *testing.T) {
	cases := []struct {
		Description            string
		Subscriptions          []cli.Subscription
		ExpectedSubscriptionID string
		ExpectError            bool
	}{
		{
			Description:   "Empty Subscriptions",
			Subscriptions: []cli.Subscription{},
			ExpectError:   true,
		},
		{
			Description: "Single Subscription",
			Subscriptions: []cli.Subscription{
				{
					ID:        "7f68fe06-9404-4db8-a5c7-29639dc4b299",
					IsDefault: true,
				},
			},
			ExpectError:            false,
			ExpectedSubscriptionID: "7f68fe06-9404-4db8-a5c7-29639dc4b299",
		},
		{
			Description: "Multiple Subscriptions with First as the Default",
			Subscriptions: []cli.Subscription{
				{
					ID:        "7f68fe06-9404-4db8-a5c7-29639dc4b299",
					IsDefault: true,
				},
				{
					ID:        "f36508bb-53b9-4aad-a2ac-2df86acf0c31",
					IsDefault: false,
				},
			},
			ExpectError:            false,
			ExpectedSubscriptionID: "7f68fe06-9404-4db8-a5c7-29639dc4b299",
		},
		{
			Description: "Multiple Subscriptions with Second as the Default",
			Subscriptions: []cli.Subscription{
				{
					ID:        "7f68fe06-9404-4db8-a5c7-29639dc4b299",
					IsDefault: false,
				},
				{
					ID:        "f36508bb-53b9-4aad-a2ac-2df86acf0c31",
					IsDefault: true,
				},
			},
			ExpectError:            false,
			ExpectedSubscriptionID: "f36508bb-53b9-4aad-a2ac-2df86acf0c31",
		},
		{
			Description: "Multiple Subscriptions with None as the Default",
			Subscriptions: []cli.Subscription{
				{
					ID:        "7f68fe06-9404-4db8-a5c7-29639dc4b299",
					IsDefault: false,
				},
				{
					ID:        "f36508bb-53b9-4aad-a2ac-2df86acf0c31",
					IsDefault: false,
				},
			},
			ExpectError: true,
		},
	}

	for _, v := range cases {
		profile := AzureCLIProfile{
			Profile: cli.Profile{
				Subscriptions: v.Subscriptions,
			},
		}
		actualSubscriptionID, err := profile.FindDefaultSubscriptionID()

		if v.ExpectError && err == nil {
			t.Fatalf("Expected an error for %q: didn't get one", v.Description)
		}

		if !v.ExpectError && err != nil {
			t.Fatalf("Expected there to be no error for %q - but got: %v", v.Description, err)
		}

		if actualSubscriptionID != v.ExpectedSubscriptionID {
			t.Fatalf("Expected Subscription ID to be %q - got %q", v.ExpectedSubscriptionID, actualSubscriptionID)
		}
	}
}

func TestAzureCLIProfileFindSubscription(t *testing.T) {
	cases := []struct {
		Description               string
		Subscriptions             []cli.Subscription
		SubscriptionIDToSearchFor string
		ExpectError               bool
	}{
		{
			Description:               "Empty Subscriptions",
			Subscriptions:             []cli.Subscription{},
			SubscriptionIDToSearchFor: "7f68fe06-9404-4db8-a5c7-29639dc4b299",
			ExpectError:               true,
		},
		{
			Description:               "Single Subscription",
			SubscriptionIDToSearchFor: "7f68fe06-9404-4db8-a5c7-29639dc4b299",
			Subscriptions: []cli.Subscription{
				{
					ID:        "7f68fe06-9404-4db8-a5c7-29639dc4b299",
					IsDefault: true,
				},
			},
			ExpectError: false,
		},
		{
			Description:               "Finding the default subscription",
			SubscriptionIDToSearchFor: "7f68fe06-9404-4db8-a5c7-29639dc4b299",
			Subscriptions: []cli.Subscription{
				{
					ID:        "7f68fe06-9404-4db8-a5c7-29639dc4b299",
					IsDefault: true,
				},
				{
					ID:        "f36508bb-53b9-4aad-a2ac-2df86acf0c31",
					IsDefault: false,
				},
			},
			ExpectError: false,
		},
		{
			Description:               "Finding a non default Subscription",
			SubscriptionIDToSearchFor: "7f68fe06-9404-4db8-a5c7-29639dc4b299",
			Subscriptions: []cli.Subscription{
				{
					ID:        "7f68fe06-9404-4db8-a5c7-29639dc4b299",
					IsDefault: false,
				},
				{
					ID:        "f36508bb-53b9-4aad-a2ac-2df86acf0c31",
					IsDefault: true,
				},
			},
			ExpectError: false,
		},
		{
			Description:               "Multiple Subscriptions with None as the Default",
			SubscriptionIDToSearchFor: "224f4ca6-117f-4928-bc0f-3df018feba3e",
			Subscriptions: []cli.Subscription{
				{
					ID:        "7f68fe06-9404-4db8-a5c7-29639dc4b299",
					IsDefault: false,
				},
				{
					ID:        "f36508bb-53b9-4aad-a2ac-2df86acf0c31",
					IsDefault: false,
				},
			},
			ExpectError: true,
		},
	}

	for _, v := range cases {
		profile := AzureCLIProfile{
			Profile: cli.Profile{
				Subscriptions: v.Subscriptions,
			},
		}

		subscription, err := profile.FindSubscription(v.SubscriptionIDToSearchFor)

		if v.ExpectError && err == nil {
			t.Fatalf("Expected an error for %q: didn't get one", v.Description)
		}

		if !v.ExpectError && err != nil {
			t.Fatalf("Expected there to be no error for %q - but got: %v", v.Description, err)
		}

		if subscription != nil && subscription.ID != v.SubscriptionIDToSearchFor {
			t.Fatalf("Expected to find Subscription ID %q - got %q", subscription.ID, v.SubscriptionIDToSearchFor)
		}
	}
}
