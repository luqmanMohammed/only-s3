// Package cmd initializes the required cli flags using cobra
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
package cmd

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/luqmanMohammed/only-s3/pkg/awss3"
	"github.com/spf13/cobra"
)

var (
	bucketName *string
	objectKey  *string
	filePath   *string
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file into AWS S3",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatalf("Failed to load aws config : %v", err)
		}
		if err := awss3.UploadFile(context.TODO(), config, *bucketName, *objectKey, *filePath); err != nil {
			log.Fatalf("Failed to upload file : %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	bucketName = uploadCmd.Flags().StringP("bucket-name", "b", "", "Name of the bucket to upload the file to")
	objectKey = uploadCmd.Flags().StringP("object-key", "o", "", "Object key in bucket")
	filePath = uploadCmd.Flags().StringP("file-path", "p", "", "Path to the file to upload")

	uploadCmd.MarkFlagRequired("bucket-name")
	uploadCmd.MarkFlagRequired("object-key")
	uploadCmd.MarkFlagRequired("file-path")
}
