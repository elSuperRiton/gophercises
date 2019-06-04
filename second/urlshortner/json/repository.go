package json

// Repository is a dummy structure implementing the
// urlshortner Repository interface for parsing JSON data
type Repository struct {
	data []byte
}

// NewRepository returns a new Repository for mapping short urls
// using JSON format data
func NewRepository(data []byte) Repository {
	return Repository{
		data: data,
	}
}
