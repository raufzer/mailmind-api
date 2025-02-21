package request

import "errors"

type UpdateUserSettingsRequest struct {
    PreferredTone string `json:"preferred_tone"`
    AutoSend      bool   `json:"auto_send"`
}

func (r *UpdateUserSettingsRequest) Validate() error {
    validTones := map[string]bool{"formal": true, "casual": true, "neutral": true}
    if !validTones[r.PreferredTone] {
        return errors.New("preferred_tone must be one of: formal, casual, neutral")
    }
    return nil
}