package twilio

import (
	"testing"

	"encoding/json"
)

func TestChatServiceSerialization(t *testing.T) {
	raw := `
{
  "account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "consumption_report_interval": 100,
  "date_created": "2015-07-30T20:00:00Z",
  "date_updated": "2015-07-30T20:00:00Z",
  "default_channel_creator_role_sid": "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "default_channel_role_sid": "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "default_service_role_sid": "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "friendly_name": "friendly_name",
  "limits": {
    "channel_members": 100,
    "user_channels": 250
  },
  "links": {
    "channels": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels",
    "users": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users",
    "roles": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles",
    "bindings": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Bindings"
  },
  "notifications": {},
  "post_webhook_url": "post_webhook_url",
  "pre_webhook_url": "pre_webhook_url",
  "pre_webhook_retry_count": 2,
  "post_webhook_retry_count": 3,
  "reachability_enabled": false,
  "read_status_enabled": false,
  "sid": "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "typing_indicator_timeout": 100,
  "url": "https://chat.twilio.com/v2/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "webhook_filters": [
    "webhook_filters"
  ],
  "webhook_method": "webhook_method",
  "media": {
    "size_limit_mb": 150,
    "compatibility_message": "media compatibility message"
  }
}`
	var chatService ChatService
	err := json.Unmarshal([]byte(raw), &chatService)

	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
}
