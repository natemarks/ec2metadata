package ec2metadata

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
)

const (
	tokenURL        = "http://169.254.169.254/latest/api/token"
	metadataURL     = "http://169.254.169.254/latest/meta-data/"
	tokenTTLSeconds = "21600"
)

func getIMDSV2Token() (token string, err error) {
	req, err := http.NewRequest("PUT", tokenURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Aws-Ec2-Metadata-Token-Ttl-Seconds", tokenTTLSeconds)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	token = string(body)
	return token, err
}

// GetV2 use IMDSv2 to get EC2 instance metadata
func GetV2(path string) (value string, err error) {
	token, err := getIMDSV2Token()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", metadataURL+path, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Aws-Ec2-Metadata-Token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	value = string(body)
	return value, err
}

// GetV1 use IMDSv1 to get EC2 instance metadata
func GetV1(path string) (value string, err error) {
	req, err := http.NewRequest("GET", metadataURL+path, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	value = string(body)
	return value, err
}

// GetAWSEc2Metadata get metadata using
// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/feature/ec2/imds#Client.GetMetadata
func GetAWSEc2Metadata(path string) (value string, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	client := imds.NewFromConfig(cfg)
	output, err := client.GetMetadata(context.TODO(), &imds.GetMetadataInput{
		Path: path,
	})
	if err != nil {
		return "", err
	}
	defer output.Content.Close()
	bytes, err := io.ReadAll(output.Content)
	if err != nil {
		return "", err
	}
	resp := string(bytes)
	return resp, err
}
