# rutilus
An AWS vulnerability scanner. The goal is to implement at least CIS controls [https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-cis-controls.html](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-cis-controls.html).

Pull requests are welcome!

## Build
```
go get
go build
```

## Run
```
# export AWS_PROFILE=my-profile
# export AWS_DEFAULT_REGION=us-east-1
./rutilus
```
