package request

type GenerateReplyRequest struct {
    EmailID string `json:"email_id" binding:"required"`
    Content string `json:"content" binding:"required"`
}

type SummarizeEmailRequest struct {
    EmailID string `json:"email_id" binding:"required"`
}
