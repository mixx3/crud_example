package api

type LogService interface {
	Create(schema *LogPostSchema) error
	GetAll() ([]*LogGetSchema, error)
	Get(id int) (*LogGetSchema, error)
	Delete(id int) error
}

type fakeLogService struct {
	logRepo LogRepository
}

func NewLogService(logRepo LogRepository) LogService {
	return &fakeLogService{logRepo}
}

func (s *fakeLogService) Create(schema *LogPostSchema) error {
	s.logRepo.Add(schema)
	return nil
}

func (s *fakeLogService) GetAll() ([]*LogGetSchema, error) {
	res := make([]*LogGetSchema, 0)
	a, err := s.logRepo.Get(0)
	res = append(res, a)
	return res, err
}

func (s *fakeLogService) Get(id int) (*LogGetSchema, error) {
	res, err := s.logRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *fakeLogService) Delete(id int) error {
	err := s.logRepo.Delete(id)
	return err
}
