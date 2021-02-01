package common

type Aioss struct {
	interview/accessKeyId     string `json:"interview/accesskeyId"`
	interview/accessKeySecret string `json:"interview/accesskeySecret"`
	BucketName      string `json:"bucketName"`
	EndPoint        string `json:"endPoint"`
}
