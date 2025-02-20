package request

type SendEmailRequest struct {
    Recipient string `json:"recipient" binding:"required,email"`
    Subject   string `json:"subject" binding:"required"`
    Body      string `json:"body" binding:"required"`
}

type SaveDraftRequest struct {
    Recipient string `json:"recipient" binding:"required,email"`
    Subject   string `json:"subject" binding:"required"`
    Body      string `json:"body" binding:"required"`
}
