package s3service

import (
	envHelper "musicservice/src/api/helpers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Client() *s3.S3 {
	key := envHelper.GetEnvValue("DIGITAL_OCEAN_S3_ACCESS_KEY", "")
	secret := envHelper.GetEnvValue("DIGITAL_OCEAN_S3_SECRET_KEY", "")
	url := envHelper.GetEnvValue("DIGITAL_OCEAN_S3_URL", "")
	region := envHelper.GetEnvValue("DIGITAL_OCEAN_S3_REGION", "ams3")

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String(url),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(region),
	}

	sess := session.Must(session.NewSession(s3Config))

	return s3.New(sess)
}

func GetBucketName() string {
	return envHelper.GetEnvValue("DIGITAL_OCEAN_S3_BUCKET_NAME", "")
}

func GetBucketURL() string {
	return "https://" + envHelper.GetEnvValue("DIGITAL_OCEAN_S3_URL", "") + "/" + GetBucketName()
}
