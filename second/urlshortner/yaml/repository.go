package yaml

type Repository struct {
	data []byte
}

// NewRepository returns a new YAML urlshortner Repository
// satisfying the urlshortner Repository and handling url shortening
// via YAML format
func NewRepository(data []byte) Repository {
	return Repository{
		data: data,
	}
}
