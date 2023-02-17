package api

type LogService interface {
	Create(schema *LogPostSchema) error
	GetAll() (*[]LogPostSchema, error)
	Get(id int) (*LogPostSchema, error)
	Delete(id int) error
}

type logService struct {
	logRepo LogRepository
}

func NewLogService(logRepo LogRepository) LogService {
	return &logService{logRepo}
}

func (s *logService) Create(schema *LogPostSchema) error {
	return nil
}

func (s *logService) GetAll() (*[]LogPostSchema, error) {
	return nil, nil
}

func (s *logService) Get(id int) (*LogPostSchema, error) {
	res, err := s.logRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *logService) Delete(id int) error {
	err := s.logRepo.Delete(id)
	return err
}
