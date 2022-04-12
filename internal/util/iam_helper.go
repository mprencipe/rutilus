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
type CredentialUser struct {
	UserName                  string
	Arn                       string
	UserCreationTime          *time.Time
	PasswordEnabled           *bool
	PasswordLastUsed          *time.Time
	PasswordLastChanged       *time.Time
	PasswordNextRotation      *time.Time
	MfaActive                 *bool
	AccessKey1Active          *bool
	AccessKey1LastRotated     *time.Time
	AccessKey1LastUsed        *time.Time
	AccessKey1LastUsedRegion  string
	AccessKey1LastUsedService string
	AccessKey2Active          *bool
	AccessKey2LastRotated     *time.Time
	AccessKey2LastUsed        *time.Time
	AccessKey2LastUsedRegion  string
	AccessKey2LastUsedService string
	Cert1Active               *bool
	Cert1LastRotated          *time.Time
	Cert2Active               *bool
	Cert2LastRotated          *time.Time
}

type CredentialReport struct {
	Users []CredentialUser
}

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
	if str == "N/A" || str == "no_information" || str == "not_supported" {
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
		user.Arn = columns[1]
		user.UserCreationTime = parseTime(columns[2], "UserCreationTime")
		user.PasswordEnabled = parseBool(columns[3], "PasswordEnabled")
		user.PasswordLastUsed = parseTime(columns[4], "PasswordLastUsed")
		user.PasswordNextRotation = parseTime(columns[6], "PasswordNextRotation")
		user.PasswordLastUsed = parseTime(columns[6], "PasswordLastUsed")
		user.MfaActive = parseBool(columns[7], "MfaActive")
		user.AccessKey1Active = parseBool(columns[8], "AccessKey1Active")
		user.AccessKey1LastRotated = parseTime(columns[9], "AccessKey1LastRotated")
		user.AccessKey1LastUsed = parseTime(columns[10], "AccessKey1LastUsed")
		user.AccessKey1LastUsedRegion = columns[11]
		user.AccessKey1LastUsedService = columns[12]
		user.AccessKey2Active = parseBool(columns[13], "AccessKey2Active")
		user.AccessKey2LastRotated = parseTime(columns[14], "AccessKey2LastRotated")
		user.AccessKey2LastUsed = parseTime(columns[15], "AccessKey2LastUsed")
		user.AccessKey2LastUsedRegion = columns[16]
		user.AccessKey2LastUsedService = columns[17]
		user.Cert1Active = parseBool(columns[18], "Cert1Active")
		user.Cert1LastRotated = parseTime(columns[19], "Cert1LastRotated")
		user.Cert2Active = parseBool(columns[20], "Cert2Active")
		user.Cert2LastRotated = parseTime(columns[21], "Cert2LastRotated")
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
