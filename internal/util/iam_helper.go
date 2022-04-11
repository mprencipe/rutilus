package util

import (
	"context"
	"rutilus/internal/config"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

var client *iam.Client

type CredentialUser struct {
	UserName           string
	PasswordEnabled    *bool
	MfaActive          *bool
	PasswordLastUsed   *time.Time
	AccessKey1LastUsed *time.Time
	AccessKey2LastUsed *time.Time
}

type CredentialReport struct {
	Users []CredentialUser
}

/*
Report CSV format
0	user
1	arn
2	user_creation_time
3	password_enabled
4	password_last_used
5	password_last_changed
6	password_next_rotation
7	mfa_active
8	access_key_1_active
9	access_key_1_last_rotated
10	access_key_1_last_used_date
11	access_key_1_last_used_region
12	access_key_1_last_used_service
13	access_key_2_active
14	access_key_2_last_rotated
15	access_key_2_last_used_date
17	access_key_2_last_used_region
18	access_key_2_last_used_service
19	cert_1_active
20	cert_1_last_rotated
21	cert_2_active
22	cert_2_last_rotated
*/
var report *CredentialReport

func GenerateCredentialReport() {
	log.Debug("Generating IAM credentials report")
	if client == nil {
		client = iam.NewFromConfig(config.Config)
		_, err := client.GenerateCredentialReport(context.TODO(), &iam.GenerateCredentialReportInput{})
		if err != nil {
			log.Fatalf("Couldn't generate credential report, %v", err)
		}
	}
}

func parseTime(str string, desc string) *time.Time {
	if str == "N/A" || str == "no_information" {
		return nil
	}
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		log.Fatalf("Error parsing time %s for column %s", str, desc)
	}
	return &parsedTime
}

func parseBool(str string, desc string) *bool {
	if str == "not_supported" {
		return nil
	}
	parsedBool, err := strconv.ParseBool(str)
	if err != nil {
		log.Fatalf("Couldn't parse bool %s for column %s", str, desc)
	}
	return &parsedBool
}

func buildReport(credReport *iam.GetCredentialReportOutput) {
	report = &CredentialReport{}
	report.Users = make([]CredentialUser, 0)
	lines := strings.Split(string(credReport.Content), "\n")
	for i, line := range lines {
		if i == 0 {
			// skip CSV header row
			continue
		}
		columns := strings.Split(line, ",")
		user := CredentialUser{}
		user.UserName = columns[0]
		user.PasswordEnabled = parseBool(columns[3], "PasswordEnabled")
		user.MfaActive = parseBool(columns[7], "MfaActive")
		user.PasswordLastUsed = parseTime(columns[4], "PasswordLastUsed")
		user.AccessKey1LastUsed = parseTime(columns[10], "AccessKey1LastUsed")
		user.AccessKey1LastUsed = parseTime(columns[15], "AccessKey1LastUsed")
		report.Users = append(report.Users, user)
	}
}

func GetCredentialReport() *CredentialReport {
	if report == nil {
		log.Debug("Credential report doesn't exist")
		credReport, err := client.GetCredentialReport(context.TODO(), &iam.GetCredentialReportInput{})
		if err != nil {
			log.Fatalf("Couldn't get credential report, %v", err)
		}
		buildReport(credReport)
	}

	return report
}
