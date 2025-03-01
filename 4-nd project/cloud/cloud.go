package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (cdb *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (cdb *CloudDb) Write() (content []byte) {
	return nil
}
