package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
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
