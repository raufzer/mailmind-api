package request

type UpdateUserSettingsRequest struct {
    PreferredTone string `json:"preferred_tone" binding:"required,oneof=formal casual neutral"`
    AutoSend      bool   `json:"auto_send"`
}
