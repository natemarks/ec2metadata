# ec2metadata
Simple functions to get metadata on an EC2 instance
```
import nmec2 github.com/natemarks/ec2metadata
// get metadata using IMDSv2
instanceID, err := nmec2.GetV2("instance-id)


// get metadata using deprecated IMDSv1
instanceID, err := nmec2.GetV1("instance-id)
```
