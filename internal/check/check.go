package check

type CheckResult = int64

const (
	Warning CheckResult = iota
	Failure
	Success
)

type Check interface {
	Check() (CheckResult, error)
	Describe() string
}
