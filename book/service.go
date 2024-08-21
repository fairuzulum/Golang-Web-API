package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(BookRequest BookRequest) (Book, error) {
	price, _ := BookRequest.Price.Int64()
	discount, _ := BookRequest.Discount.Int64()
	rating, _ := BookRequest.Rating.Int64()

	book := Book{
		Title:       BookRequest.Title,
		Description: BookRequest.Description,
		Price:       int(price),
		Rating:      int(rating),
		Discount:    int(discount),
	}
	return s.repository.Create(book)
}

func (s *service) Update(ID int, BookRequest BookRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)

	price, _ := BookRequest.Price.Int64()
	discount, _ := BookRequest.Discount.Int64()
	rating, _ := BookRequest.Rating.Int64()

	book.Title = BookRequest.Title
	book.Description = BookRequest.Description
	book.Price = int(price)
	book.Rating = int(rating)
	book.Discount = int(discount)

	
	return s.repository.Update(book)
}

func (s *service) Delete(ID int) (Book, error) {
	book, _ := s.repository.FindByID(ID)
	return s.repository.Delete(book)
}
