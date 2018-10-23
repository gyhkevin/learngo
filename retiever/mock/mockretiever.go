package mock

type Retriever struct {
	Contents string
}

func (Retriever) Get(url string) string {
	panic("implement me")
}

