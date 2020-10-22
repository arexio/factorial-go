package factorial

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	webhookURL = "/api/v1/webhooks"
)

// Webhook contains all the webhook information
type Webhook struct {
	SubscriptionType string `json:"subscription_type"`
}

// CreateWebhookRequest keeps the information needed
// for create a new webhook
type CreateWebhookRequest struct {
	SubscriptionType string `json:"subscription_type"`
	TargetURL        string `json:"target_url"`
}

// DeleteWebhookRequest keeps the information needed
// for delete a new webhook
type DeleteWebhookRequest struct {
	SubscriptionType string `json:"subscription_type"`
}

// CreateWebhook creates a subscription for a determined webhook type.
// If webhook already exists, it just changes the target_url.
func (c Client) CreateWebhook(w CreateWebhookRequest) (Webhook, error) {
	var webhook Webhook

	bytes, err := json.Marshal(w)
	if err != nil {
		return webhook, err
	}

	resp, err := c.post(webhookURL, bytes)
	if err != nil {
		return webhook, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&webhook); err != nil {
		return webhook, err
	}

	return webhook, nil
}

// DeleteWebhook deletes a subscription to a webhook.
func (c Client) DeleteWebhook(w DeleteWebhookRequest) (Webhook, error) {
	var webhook Webhook

	body, err := json.Marshal(w)
	if err != nil {
		return webhook, err
	}

	// Not used the Client because this endpoint is not following the REST definition
	// and we don't want to break our pattern for it
	req, err := http.NewRequest(http.MethodDelete, c.apiURL+webhookURL, bytes.NewReader(body))
	if err != nil {
		return webhook, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return webhook, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&webhook); err != nil {
		return webhook, err
	}

	return webhook, nil
}

// ListWebhooks gets a list of all subscribed webhooks for current user.
func (c Client) ListWebhooks() ([]Webhook, error) {
	var webhooks []Webhook

	resp, err := c.get(webhookURL, nil)
	if err != nil {
		return webhooks, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&webhooks); err != nil {
		return webhooks, err
	}

	return webhooks, nil
}
