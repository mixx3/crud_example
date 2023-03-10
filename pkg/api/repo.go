package api

type LogRepository interface {
	Add(schema *LogPostSchema) error
	Get(id int) (*LogGetSchema, error)
	Delete(id int) error
}

type FakeLogRepo struct {
	container map[int]*LogPostSchema
	idIncr    int
}

func NewFakeRepo() *FakeLogRepo {
	return &FakeLogRepo{
		container: make(map[int]*LogPostSchema, 0),
		idIncr:    0,
	}
}

func (r *FakeLogRepo) Add(schema *LogPostSchema) error {
	r.container[r.idIncr] = schema
	r.idIncr += 1
	return nil
}

func (r *FakeLogRepo) Get(id int) (*LogGetSchema, error) {
	res := &LogGetSchema{Message: r.container[id].Message, ID: id}
	return res, nil
}

func (r *FakeLogRepo) Delete(id int) error {
	delete(r.container, id)
	return nil
}
