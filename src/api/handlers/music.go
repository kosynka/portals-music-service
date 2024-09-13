package handlers

import (
	"musicservice/src/api/models"
	s3service "musicservice/src/api/services"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func GetMusicList(c *gin.Context) {
	prefix := "music/"
	s3Client := s3service.GetS3Client()
	bucketName := s3service.GetBucketName()
	bucketURL := s3service.GetBucketURL()

	bucketObjects, err := s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(prefix),
	})

	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": err.Error(),
			},
		)

		return
	}

	contents := filterFolder(bucketObjects.Contents, prefix)

	c.JSON(
		http.StatusOK, gin.H{
			"status": "success",
			"context": gin.H{
				"music": buildMusicList(contents, bucketURL),
			},
		},
	)
}

// First element is the folder name
func filterFolder(objects []*s3.Object, prefix string) []*s3.Object {
	if len(objects) > 0 && *objects[0].Key == prefix {
		return objects[1:]
	}

	return objects
}

func buildMusicList(objects []*s3.Object, bucketURL string) []models.Music {
	var musics []models.Music

	for _, obj := range objects {
		musics = append(musics, models.Music{
			TITLE: *obj.Key,
			URL:   bucketURL + "/" + *obj.Key,
			SIZE:  *obj.Size,
			ETAG:  *obj.ETag,
		})
	}

	return musics
}
