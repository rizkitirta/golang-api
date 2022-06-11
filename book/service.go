package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookUpdateRequest BookUpdateRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository BookRepository
}

func NewService(repository BookRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	// return V1
	// return s.repository.FindAll()

	// return V2
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      4,
	}
	book, err := s.repository.Create(book)

	return book, err
}

func (s *service) Update(ID int, bookUpdateRequest BookUpdateRequest) (Book, error) {
	book, err := s.repository.FindById(ID)
	price, _ := bookUpdateRequest.Price.Int64()

	book.Title = bookUpdateRequest.Title
	book.Price = int(price)
	book.Description = bookUpdateRequest.Description
	book, err = s.repository.Update(book)

	return book, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	book, err = s.repository.Delete(book)

	return book, err
}
