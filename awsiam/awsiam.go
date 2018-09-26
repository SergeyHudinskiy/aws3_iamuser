package awsiam

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

//NewIAMUser returns
func NewIAMUser(accessKeyID string, secretAccessKey string, region string, bindingID string) (string, error) {

	os.Setenv("AWS_ACCESS_KEY_ID", accessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", secretAccessKey)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svc := iam.New(sess)

	if awserr, ok := err.(awserr.Error); ok && awserr.Code() == iam.ErrCodeNoSuchEntityException {
		result, err := svc.CreateUser(&iam.CreateUserInput{
			UserName: &bindingID,
		})

		if err != nil {
			fmt.Println("CreateUser Error", err)
			return "", err
		}

		return aws.StringValue(result.User.Arn), nil
	} else {
		return "", err
	}
}
