// Package awss3 handles all functions related to AWS S3 Service
/*
Copyright © 2022 m.luqman077@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package awss3

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadFile uploads file specified using the filePath argument into the
// specified bucket under the said objectKey
func UploadFile(ctx context.Context, cfg aws.Config, bucketName, objectKey, filePath string) error {
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	uploadFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer uploadFile.Close()
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   uploadFile,
	})
	return err
}
