package forms

// CommentaryForm form for creating Commentaries
type CommentaryForm struct {
	Author  string `json:"author" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Content string `json:"content" binding:"required"`
}
