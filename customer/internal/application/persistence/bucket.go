package persistence

import "context"

type Bucket interface {
	UploadAvatar(ctx context.Context, name string, fileBuffer []byte) error
}
