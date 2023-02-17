package api

type LogPostSchema struct {
	message string
}

type LogGetSchema struct {
	id        int
	message   string
	dtCreated int64
}

type LogGetAllSchema struct {
	items []LogGetSchema
}
