package books

type CreateBookRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Author      string `json:"author" validate:"required,min=1,max=255"`
	Publication string `json:"publication" validate:"required,min=1,max=255"`
}

type UpdateBookRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=1,max=255"`
	Author      *string `json:"author" validate:"omitempty,min=1,max=255"`
	Publication *string `json:"publication" validate:"omitempty,min=1,max=255"`
}

type BookResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func ToResponse(b *Book) BookResponse {
	return BookResponse{
		ID:          b.ID,
		Name:        b.Name,
		Author:      b.Author,
		Publication: b.Publication,
	}
}
