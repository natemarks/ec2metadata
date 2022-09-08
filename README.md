# ec2metadata
Simple functions to get metadata on an EC2 instance
```
import github.com/natemarks/ec2metadata
// get metadata using IMDSv2
instanceID, err := ec2metadata.GetV2("instance-id)


// get metadata using deprecated IMDSv1
instanceID, err := ec2metadata.GetV1("instance-id)
```
