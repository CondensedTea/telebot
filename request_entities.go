package telebot

type KeyboardButtonRequestUser struct {
	// RequestID is identifier of the request, which will be received back in the UserShared object.
	// Must be unique within the message.
	RequestID int32 `json:"request_id"`

	UserIsBot     bool `json:"user_is_bot,omitempty"`
	UserIsPremium bool `json:"user_is_premium,omitempty"`
}

type KeyboardButtonRequestChat struct {
	// RequestID is identifier of the request, which will be received back in the UserShared object.
	// Must be unique within the message.
	RequestID int32 `json:"request_id"`

	ChatIsChannel bool `json:"chat_is_channel"`
}

type UserSharedUpdate struct {
	RequestID int32 `json:"request_id"`
	UserID    int   `json:"user_id"`
}

type ChatSharedUpdate struct {
	RequestID int32 `json:"request_id"`
	ChatID    int   `json:"chat_id"`
}
