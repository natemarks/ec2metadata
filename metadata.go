package puppers

import (
	"io/ioutil"
	"net/http"
)

const (
	tokenUrl        = "http://169.254.169.254/latest/api/token"
	metadataUrl     = "http://169.254.169.254/latest/meta-data/"
	tokenTtlSeconds = "21600"
)

func getIMDSV2Token() (token string, err error) {
	req, err := http.NewRequest("PUT", "http://169.254.169.254/latest/api/token", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-Aws-Ec2-Metadata-Token-Ttl-Seconds", tokenTtlSeconds)

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
		panic(err)
	}

	req, err := http.NewRequest("GET", metadataUrl+path, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Aws-Ec2-Metadata-Token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	value = string(body)
	return value, err
}

// GetV1 use IMDSv1 to get EC2 instance metadata
func GetV1(path string) (value string, err error) {
	req, err := http.NewRequest("GET", metadataUrl+path, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	value = string(body)
	return value, err
}