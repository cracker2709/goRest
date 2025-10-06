package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// InitGoogleCloudAuth initializes Google Cloud authentication using a service account key file.
func InitGoogleCloudAuth(keyFilePath string) (context.Context, option.ClientOption) {
	ctx := context.Background()
	clientOption := option.WithCredentialsFile(keyFilePath)
	// You can now use ctx and clientOption with Google Cloud clients, e.g. storage.NewClient(ctx, clientOption)
	log.Println("Google Cloud authentication initialized")
	return ctx, clientOption
}

func listFilesInBucket(ctx context.Context, clientOption option.ClientOption, bucketName string) ([]string, error) {
	cctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	client, err := storage.NewClient(cctx, clientOption)
	if err != nil {
		return nil, fmt.Errorf("create storage client: %w", err)
	}
	defer func() {
		if cerr := client.Close(); cerr != nil {
			log.Printf("storage client close: %v", cerr)
		}
	}()

	it := client.Bucket(bucketName).Objects(cctx, nil)
	var names []string
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("list objects in bucket %q: %w", bucketName, err)
		}
		names = append(names, attrs.Name)
	}
	return names, nil
}

func main() {
	ctx, clientOpts := InitGoogleCloudAuth("path/to/creds.json")
	bucket, err := listFilesInBucket(ctx, clientOpts, "bucket_name")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bucket)
}
