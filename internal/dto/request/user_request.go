package request

import (
	"fmt"
)

type UpdateUserSettingsRequest struct {
	PreferredTone string `json:"preferred_tone"`
	AutoSend      bool   `json:"auto_send"`
}

func (r *UpdateUserSettingsRequest) Validate() error {
	validTones := map[string]bool{"formal": true, "casual": true, "neutral": true}
	if !validTones[r.PreferredTone] {
		return fmt.Errorf("preferred tone must be one of: %v", validTones)
	}
	return nil
}
