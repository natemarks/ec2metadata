package puppers

import (
	"io/ioutil"
	"net/http"
)

const (
	tokenURL        = "http://169.254.169.254/latest/api/token"
	metadataURL     = "http://169.254.169.254/latest/meta-data/"
	tokenTTLSeconds = "21600"
)

func getIMDSV2Token() (token string, err error) {
	req, err := http.NewRequest("PUT", "http://169.254.169.254/latest/api/token", nil)
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
