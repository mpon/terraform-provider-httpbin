provider "twilio" {
  # account_sid = "" // You can set env variables TWILIO_ACCOUNT_SID  # auth_token = ""  // You can set env variables TWILIO_AUTH_TOKEN
}

resource "twilio_chat_service" "test" {
  friendly_name = "test21"

  limits {
    channel_members = 1
    user_channels   = 2
  }
}

resource "twilio_chat_service" "test2" {
  friendly_name    = "test23"
  pre_webhook_url  = "https://httpbin.org/get"
  post_webhook_url = "https://httpbin.org/post"
  webhook_method   = "POST"

  webhook_filters = [
    "onMessageSend",
    "onMessageSent",
    "onChannelUpdated",
    "onMessageUpdated",
    "onMessageUpdate",
    "onUserUpdate",
  ]

  limits {
    channel_members = 3
    user_channels   = 4
  }
}
