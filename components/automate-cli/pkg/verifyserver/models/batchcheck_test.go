package models

import (
	"errors"
	"reflect"
	"testing"

	"github.com/chef/automate/lib/config"
	"github.com/stretchr/testify/assert"
)

func TestAppendCertsByIpToNodeCerts(t *testing.T) {
	// Test case 1: certsByIP is not nil, ipList is not empty
	c := &Config{}
	certsByIP := []config.CertByIP{
		{
			IP:         "192.168.0.1",
			PrivateKey: "private_key1",
			PublicKey:  "public_key1",
		},
		{
			IP:         "192.168.0.2",
			PrivateKey: "private_key2",
			PublicKey:  "public_key2",
		},
	}
	ipList := []string{"192.168.0.3", "192.168.0.4"}
	privateKey := "private_key"
	publicKey := "public_key"
	adminKey := "admin_key"
	adminCert := "admin_cert"

	c.appendCertsByIpToNodeCerts(&certsByIP, ipList, privateKey, publicKey, adminKey, adminCert)

	expected := []NodeCert{
		{
			IP:        "192.168.0.1",
			Key:       "private_key1",
			Cert:      "public_key1",
			AdminKey:  "admin_key",
			AdminCert: "admin_cert",
		},
		{
			IP:        "192.168.0.2",
			Key:       "private_key2",
			Cert:      "public_key2",
			AdminKey:  "admin_key",
			AdminCert: "admin_cert",
		},
		{
			IP:        "192.168.0.3",
			Key:       "private_key",
			Cert:      "public_key",
			AdminKey:  "admin_key",
			AdminCert: "admin_cert",
		},
		{
			IP:        "192.168.0.4",
			Key:       "private_key",
			Cert:      "public_key",
			AdminKey:  "admin_key",
			AdminCert: "admin_cert",
		},
	}

	// Check if each expected element exists in the actual result
	for _, expectedNode := range expected {
		found := containsElement(c.Certificate.Nodes, expectedNode)
		if !found {
			t.Errorf("Test case 1 failed: Expected node %+v not found in the result.", expectedNode)
		}
	}

	// Test case 2: certsByIP is nil, ipList is empty
	c = &Config{}
	certsByIP = nil
	ipList = []string{""}
	privateKey = "private_key"
	publicKey = "public_key"
	adminKey = "admin_key"
	adminCert = "admin_cert"

	c.appendCertsByIpToNodeCerts(&certsByIP, ipList, privateKey, publicKey, adminKey, adminCert)

	expected = []NodeCert{
		{
			IP:        "",
			Key:       "private_key",
			Cert:      "public_key",
			AdminKey:  "admin_key",
			AdminCert: "admin_cert",
		},
	}

	// Check if each expected element exists in the actual result
	for _, expectedNode := range expected {
		found := containsElement(c.Certificate.Nodes, expectedNode)
		if !found {
			t.Errorf("Test case 2 failed: Expected node %+v not found in the result.", expectedNode)
		}
	}
}

func containsElement(nodes []NodeCert, targetNode NodeCert) bool {
	for _, node := range nodes {
		if reflect.DeepEqual(node, targetNode) {
			return true
		}
	}
	return false
}

func TestPopulateWith(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     *Config
		wantErr  bool
		err      error
	}{
		{
			name:     "PopulateWith Invalid OnPrem Config",
			filePath: "./testdata/InvalidHaOnPrem.toml",
			want:     &Config{},
			wantErr:  true,
			err:      errors.New("invalid or empty: ssh_user\ninvalid or empty: ssh_key_file\ninvalid S3 endpoint format. Endpoint should end with '.amazonaws.com'\nautomate private ip 1324.2534.1is not valid\ninvalid or empty: chef_server_private_ips\ninvalid or empty: opensearch_private_ips\ninvalid or empty: postgresql_private_ips\nurl should not include the protocol (http:// or https://): automate fqdn\nempty value: automate instance_count\ninvalid value 'automate.toml' for field 'config_file'. Expected values are: configs/automate.toml\ninvalid format. Failed to decode root_ca for automate\ninvalid format. Failed to decode private_key for automate\ninvalid format. Failed to decode public_key for automate\ninvalid format. Failed to decode private_key for automate ip\ninvalid format. Failed to decode public_key for automate ip\ninvalid value 'chef server instance_count' for field 'two'\ninvalid format. Failed to decode private_key for chef-server\ninvalid format. Failed to decode public_key for chef-server\ninvalid format. Failed to decode private_key for chef server ip\ninvalid format. Failed to decode public_key for chef server ip\nempty value: opensearch instance_count\nopensearch root_ca and/or admin_key and/or admin_cert and/or public_key and/or private_key are missing. Otherwise set enable_custom_certs to false\nopensearch ippublic_key and/or private_key are missing in certs_by_ip. Otherwise set enable_custom_certs to false\nempty value: postgresql instance_count\ninvalid format. Failed to decode root_ca for postgresql\ninvalid format. Failed to decode private_key for postgresql\ninvalid format. Failed to decode public_key for postgresql\npostgresql ip 0.0.1 for certs is not valid\ninvalid format. Failed to decode private_key for postgresql ip\ninvalid format. Failed to decode public_key for postgresql ip"),
		},
		{
			name:     "PopulateWith OnPrem Db Aws Managed Config",
			filePath: "./testdata/HaOnPremDbAwsManaged.toml",
			want: &Config{
				SSHUser: SSHUser{
					Username:     "ubuntu",
					PrivateKey:   "./testdata/A2HA.pem",
					SudoPassword: "",
				},
				Arch: "existing_nodes",
				Backup: Backup{
					FileSystem: FileSystem{
						MountLocation: "/mnt/automate_backups",
					},
					ObjectStorage: ObjectStorage{
						Endpoint:   "",
						BucketName: "",
						BasePath:   "",
						AccessKey:  "",
						SecretKey:  "",
						AWSRegion:  "us-west-1",
					},
				},
				Hardware: Hardware{
					AutomateNodeCount: 2,
					AutomateNodeIps: []string{
						"192.0.0.11", "192.0.0.12",
					},
					ChefInfraServerNodeCount: 2,
					ChefInfraServerNodeIps: []string{
						"192.0.0.11", "192.0.0.12",
					},
					PostgresqlNodeCount: 3,
					PostgresqlNodeIps:   nil,
					OpenSearchNodeCount: 3,
					OpenSearchNodeIps:   nil,
				},
				Certificate: Certificate{
					AutomateFqdn:   "chefautomate.example.com",
					ChefServerFqdn: "chefautomate.example.com",
					RootCert:       "",
					Nodes:          nil,
				},
				ExternalOS: ExternalOS{
					OSDomainName:   "managed-services-os",
					OSDomainURL:    "search-managed-services-os.us-east-1.es.amazonaws.com",
					OSUsername:     "admin",
					OSUserPassword: "Progress@123",
					OSCert:         "<cert_content>",
					OSRoleArn:      "arn:aws:iam::1127583934333:role/managed-services",
				},
				ExternalPG: ExternalPG{
					PGInstanceURL:       "managed-rds-db.c5gkx.ap-northeast-1.rds.amazonaws.com:5432",
					PGSuperuserName:     "postgres",
					PGSuperuserPassword: "Progress123",
					PGDbUserName:        "postgres",
					PGDbUserPassword:    "Progress123",
					PGRootCert:          "<cert_content>",
				},
				DeploymentState: "",
				APIToken:        "",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:     "PopulateWith AWS Managed Config",
			filePath: "./testdata/HaAwsManaged.toml",
			want: &Config{
				SSHUser: SSHUser{
					Username:     "ubuntu",
					PrivateKey:   "./testdata/A2HA.pem",
					SudoPassword: "",
				},
				Arch: "aws",
				Backup: Backup{
					FileSystem: FileSystem{
						MountLocation: "/mnt/automate_backups",
					},
					ObjectStorage: ObjectStorage{
						Endpoint:   "",
						BucketName: "automate-test",
						BasePath:   "",
						AccessKey:  "",
						SecretKey:  "",
						AWSRegion:  "",
					},
				},
				Hardware: Hardware{
					AutomateNodeCount:        2,
					AutomateNodeIps:          nil,
					ChefInfraServerNodeCount: 2,
					ChefInfraServerNodeIps:   nil,
					PostgresqlNodeCount:      3,
					PostgresqlNodeIps:        nil,
					OpenSearchNodeCount:      3,
					OpenSearchNodeIps:        nil,
				},
				Certificate: Certificate{
					AutomateFqdn:   "chefautomate.example.com",
					ChefServerFqdn: "chefautomate.example.com",
					RootCert:       "-----BEGIN CERTIFICATE-----\nMIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl\nMCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp\nU3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw\nNjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE\nChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp\nZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3\nDQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf\n8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN\n+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0\nX9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa\nK4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA\n1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G\nA1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR\nzt0fhvRbVazc1xDCDqmI56FspGowaDELMAkGA1UEBhMCVVMxJTAjBgNVBAoTHFN0\nYXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD\nbGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w\nDQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3\nL7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D\neruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl\nxy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp\nVSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY\nWQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8fF5Q=\n-----END CERTIFICATE-----",
					Nodes: []NodeCert{
						{
							IP:        "",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgzCCAmugAwIBAgIJAPMNo6eG0UBgMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIwWhcNMjYwNTIxMTEzOTIwWjBlMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEzARBgNVBAMMCmNoZWZjbGllbnQwggEiMA0GCSqGSIb3\nDQEBAQUAA4IBDwAwggEKAoIBAQCyJbvhtNLZrgz7aa1W1gIGFrwtI5h23S99A8VZ\nlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6lSopXzmyzxgbuLSnAmmicNrsF\nI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/npzXIEa3HENUghVHOyR1nY8d\nvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFovC2z+HikVhkt8NB1pD+7ZVTq\nS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2jpMvFSVBdUi3HgdK+wO5pD0Jt\ngWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTjs04t/R1BAgMBAAGjODA2MB0G\nA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAVBgNVHREEDjAMggpjaGVmY2xp\nZW50MA0GCSqGSIb3DQEBCwUAA4IBAQAOHPBbbNCFCQ+Fg5YSvgXCglNUXfKoeCla\nLjruho1re6mJQmeW0H5OTEL7CssEHM6yzR6rlVlVnnu0+RAFlb6vTXo8eK4HaXS2\nA+OY7oP7UpMcxj3fxuWWgg1lnW/8a9Z+9JNSm9M59MpU17whvq/i8EJppECK4pD/\nZ9J12XYepWxARbVUcgWyHFIjnSoNpFwgrm1xybnyJ2WnjRRjpq6JvkP5+eklESEm\nZ4wjUCokbRJ8gbC7tuVgQk6DNM123j9djBm+A+5WlbvHUrD5trfbkp8kUJY4jaJF\ncsKeX402wz9P7XM5eGsToNpAZq41Q7mFzz14DfqFNttaCMHMYi4k\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCyJbvhtNLZrgz7\naa1W1gIGFrwtI5h23S99A8VZlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6l\nSopXzmyzxgbuLSnAmmicNrsFI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/\nnpzXIEa3HENUghVHOyR1nY8dvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFo\nvC2z+HikVhkt8NB1pD+7ZVTqS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2j\npMvFSVBdUi3HgdK+wO5pD0JtgWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTj\ns04t/R1BAgMBAAECggEATFvmM3Gpawqyn9UFKpJina+DCyoVwOU/5KsIHUTP0XkC\n3ASgWDPb1tozy36RmmjMcbFz9lOboZhiwoC32bkuWMRJ1i9flSHaMxM2iJaHckhh\nSaaubvFptw9of7XllG7m23CRyJJMuEXT4YCgI4m/Jd+kcIWtjzGniA53+LvxUM/S\n+pR2j5cc7yHYcsc4tz842UelVQJWnMwjHwYDsuZC4ssRq1dffq7ifgfhcrxcwAXg\nMMBEtmJsTeaBYjbiM4qdVTwc8BpX9lafgUi3asapR0X2u2k1hOs5fS0qdEf5IarN\nFqwQDj5Xurcs0xvUoTPC7eEqFMh1SRVw8mS0hZG0EQKBgQDZ9GHhIPjDd3NwT38R\nhHYd3i1EKUtJgoOlqN/WUSlcFBUKxEsQTF+5pZipPYFUjb8+DwzwQMmeSXAs9A+Z\nPCFQDO1EHlE1Hv6A0Fv7RPTt3lnkZpo7SdSf/VqxKVCLOIkeU9tEMr0d/8mJ8hbj\nlDIUmfZtWIo5eZnM6fQF8rM8LQKBgQDRPoIf1d5M6eMDz6Ss0dPF/UkhEhsXUQgh\nd0puiyKuk0e2v6q8sESKV4CnGDFd08XLMcSnOw3SS/neLhZ7lixjtewUoDAVwxlB\nNos3W6FBIgyQRFp5SBjuHhKjPiK+BPlGsc7+39I12aGxE0r5w7Pp4X+4Z7hDF8ID\nqlZtDkIN5QKBgQCe8aIjnHjtiwHraH3hF3lP5MOcDoUx8XTx7Up3L6760EZcGLQp\nCZlReFrxKMJVGB3cMvubhZPC1AlzLvTlKb2ddB/fakCMfbLZ25kIj8wSX/GsJ8rX\n68qcdhWaVue+75bHQB4KCPpzkyK1b4+TnXI8Jd9Y9JWwvmYT0pU7dTeSbQKBgQC1\nko6MXaQoDhWG6xq1NOeWOXLKFdIYa6Kol8GpJ2eTIg7rEGtyjWsMuV3UofPEvc43\nwxopG9+ki3VqTYgI+onOhME2LMNNPx2dL12jTgoiYQ+R6R6xe9TWXJZDvdmcFujR\nZd5/4W2ieRYMePdowWBQJfQU6zxETEt5rsiMngDH2QKBgHB8r2LQvvQX4L0w2wY+\nOJZOGKk4AMXIFM/J0qY60DLnw0B4RuzSXcKw2EGYRJihbGJAI8+KczU7LpVlOsw5\nFzQabRrE33NPYWFv0bqzas9/p9mPwFPzhq13mkrLOHwf71T63OEalaOh9WzaE2Tx\nhTp1M8sBDwrDLxCEk0X3NEL+\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgzCCAmugAwIBAgIJAPMNo6eG0UBiMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIxWhcNMjYwNTIxMTEzOTIxWjBlMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEzARBgNVBAMMCmNoZWZjbGllbnQwggEiMA0GCSqGSIb3\nDQEBAQUAA4IBDwAwggEKAoIBAQDQlkc2SQkhzEbDBluJDr/9UvP4cObMbYf7r3Yy\n9vxds5AikR9gK0tkamAEK9e9uPMi2xSIvDPxNkwew2XlWmMxOHsUpUPt2gaaT2DE\nueiTae+TnVcLtFBIYuo3D9udDW2XAwf0rnjDWifVrbOsr7++hr3eqcg1k46+f+KU\nkHwtaYG7noMAq/vwAGkN0hN+Pfa8ILZhXgtoSrLk5vJ5KavKmD0lFMSzg5AHTT8U\nNfarUGdH7bCUeZGwZ6MHGdFtUDBPvJKdSgxShoV+03DmGx4TdbCYaS0bSVWAOxHW\nfJ3oh3tizjp8n8u/REqSZQpsWx0us8kjf9JICBYSafjxMBiZAgMBAAGjODA2MB0G\nA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAVBgNVHREEDjAMggpjaGVmY2xp\nZW50MA0GCSqGSIb3DQEBCwUAA4IBAQABwtwMsJlif2kAa9cYdSwaG0Zn5gHlqX0u\n7GW8RFcTwkU+5ZpzbdcgX5XtKq7a7LDdGTYZpvmiKAvsSY4L5vDV0tfZRsJSIyOS\n4kWG8P4LdtsXld2Px9V/fLGxTB//aldM3K5NDOG441KyBskaZz406oL1SJP0mEh6\nJM+I3uOVqxG7o36ntLn6feSWeL6sgv1CpgzQH2kiQpgWK/T7raGxuSvngfuFoSyr\nv82fAJ2GD/Cw/0E5IFj/AVPbCWY9EzE9m2AkkFhfpXTP1qLXfIaIlFZiiZZrW+fL\nEM5hpYgcYA7V2hP++Im5U7MEndhRASbHW+XvCR7WKMm5V+Rt3wjy\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDQlkc2SQkhzEbD\nBluJDr/9UvP4cObMbYf7r3Yy9vxds5AikR9gK0tkamAEK9e9uPMi2xSIvDPxNkwe\nw2XlWmMxOHsUpUPt2gaaT2DEueiTae+TnVcLtFBIYuo3D9udDW2XAwf0rnjDWifV\nrbOsr7++hr3eqcg1k46+f+KUkHwtaYG7noMAq/vwAGkN0hN+Pfa8ILZhXgtoSrLk\n5vJ5KavKmD0lFMSzg5AHTT8UNfarUGdH7bCUeZGwZ6MHGdFtUDBPvJKdSgxShoV+\n03DmGx4TdbCYaS0bSVWAOxHWfJ3oh3tizjp8n8u/REqSZQpsWx0us8kjf9JICBYS\nafjxMBiZAgMBAAECggEAH07/P2G1EjKkSG+y256wKKkD372qlvK134xVtI6oELR8\nJQQdboTxGxBwew/NoTeanxe9PEzriwA4asGvkL6BdwjgSJgJ2zDHnu/dOYoiI8ZZ\n5JFQWHxnNIZMW5lwwGEcmvckgZCSpdfpdMRO3NPAdyuoYjyfxZLxcRym+N+7E+Gl\n3izIlRNz4fHfdsH6rttYMJpn/T2MMyJ3BRZmGc8JPraDzHCMoMnVv2bAhJTpbglr\ndFUgzE+eIa/b0ExE/sd3+WVG2hLqrrphwlFczmPlb/7nPlnDL42tVBmZnlexJm5x\npk1GePV2OmlT98rW9ItX0cwVxz15LBpXLIIXG9F7pQKBgQDw5DvUlIsqreJ61jQH\nXd5QRVKmjTkninL4oA5+0YxlzkU3kDhMTvNlmmACcLqOmJ4uOTB107auVnv/XEo2\nnYtWMglx6YtdO0q2z8w2ywpnoI9PfIZtAUKNqLBK1BcD5/c9lRE05o1LFylcvL/Y\nP9NZYhKkWYGOt0aFIbjNphyW1wKBgQDdq1ySbIst64pw16DssYkvA681Lr8Sbp++\n3F+d1aXIKkoOYPxmdA+/lwv/47LWeBB4QDItyBAxx74pesSkEy+BbdkDNSSnMEsS\nJW9mCn7gqbKwZeX9TWbdcqXtjXz1XrQZxEslVGOQJPRt7ClGfSVA+NjYtjv7iqCI\nxcVzseiODwKBgQCItW5DCX4lXYN/pNroJ1yIf58VSGZcS1VORj+Tt0aPbE2Z5+4b\nWF8HlWHRYLpvPKvgnbIj3F/7drduR6kSb7xo0YLMs/bUlVakgy9pFTe1cciDGq+L\nY0Cq9kX+YXkiTV3iBBw8wm8DY4SkzbWueyJtwpvDy8wb+2U5HtcrVo85BwKBgQC9\n+Sx++LNXCXQ+PS5Xa9esCTZRF9z9CP2y7t6rP/yyTTvmksv9ah5NDkBkb1pHX+KN\njEb04W6vmwWoOuTn0OF3xRKlIxhkiIjt5lNQWlJebFENyGaQ7ZLo2mbF7epXx3AG\nXSohte8WC/XHdwvwszQIOLxvDc7eRvJKBWSxQJmTlQKBgDf7Cx8/ym9TzCIfWx3+\nAfCtltTglZAWgt0KjaOlT94lFmv+rJrDbp+rGSp4zvTW+U1/Xxt5+YdZGpMD8OND\n0Wy6xd38bTp7jOzYdxTebg+KC+Gum0asXZK9+ra6vBRYsqXX7sET7rEsFMrJiata\ndXO0rPR4kQZSuTqBSIxAf4Jb\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDfzCCAmegAwIBAgIJAPMNo6eG0UBdMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIwWhcNMjYwNTIxMTEzOTIwWjBjMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxETAPBgNVBAMMCGNoZWZub2RlMIIBIjANBgkqhkiG9w0B\nAQEFAAOCAQ8AMIIBCgKCAQEAxcr/RHZymJ7x6V3fVHNs+ppzWOSj+WX0JDm3P0/K\nR1dV5Drzicl8tJVpiR3fniCQOYo/wEA6uRplvmB1unPTTRa712X+WrzLX29sXJoL\n128thRpp6XI7R3fOXyNPEAGZO6gHnsSpHqdPZCCL6LmIS9wh4/+AqQ+KdoChxJ6D\n7RjVmivCU4p5tyLtbNhO47IEP/pTXXUmnX29zlY6dC1wcEsRt3zvwKJYwgXUgMsh\nVzB2h+7I8OBuPBXTjUFC8o4+yr7xjtIB1AKQYHeZaxWoy0z8YFUwCY4Zrvy2P7Yw\nE2dAVipnSUkNNtmRCseLqdchmhxwVKWCP1GokFZqfdksVwIDAQABozYwNDAdBgNV\nHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwEwYDVR0RBAwwCoIIY2hlZm5vZGUw\nDQYJKoZIhvcNAQELBQADggEBAK3O2VEYJQ4byjMYxuxmVEwHaNrl1XrwDZ7pKciq\nx9GyLFwYva4svKZfawUbZUQsIBLxRAaUDPVdfb8MYn1lTI4q2y9XGT1osM0SuL6f\ngllp4Yg4rTc94G5yXzhnWYCqrX9XZK0muFKCuJnniC0VmrP9zfojUsa0x6qvIfKf\nEB5u6SKRbQGS1ECryGQfiwXzhwy17/Qw7Ab44ufDpWmiw6fdNY/KhBKIwCRSHoJJ\nQQSO1eGYc/f3j7Kve5LaTOxADqBhGcxyItqdEJlmhKwAp2aOmO6pZLPCuZvCmvDL\nDv6bUUXSsZF4fb1diLIBpmD1hh8OGNY65LUPpzAxJeZvo5w=\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDFyv9EdnKYnvHp\nXd9Uc2z6mnNY5KP5ZfQkObc/T8pHV1XkOvOJyXy0lWmJHd+eIJA5ij/AQDq5GmW+\nYHW6c9NNFrvXZf5avMtfb2xcmgvXby2FGmnpcjtHd85fI08QAZk7qAeexKkep09k\nIIvouYhL3CHj/4CpD4p2gKHEnoPtGNWaK8JTinm3Iu1s2E7jsgQ/+lNddSadfb3O\nVjp0LXBwSxG3fO/AoljCBdSAyyFXMHaH7sjw4G48FdONQULyjj7KvvGO0gHUApBg\nd5lrFajLTPxgVTAJjhmu/LY/tjATZ0BWKmdJSQ022ZEKx4up1yGaHHBUpYI/UaiQ\nVmp92SxXAgMBAAECggEBALClRSk9p8bKXT6QCb6Af5mois+fExrPhSU9Ln0qo3rn\nctwsEgjCm88jiWdd+LJeXrAk2h62vjtGaguGVl44x0OXxBbxDiK3beJDvsFNCrpS\nnpK7Lk/BJ1QCmZq6DAg9hT6UKIoRFQE9Z1gDATDNUf5+EP5w19Uk/gIri03wS95Y\nt6xnA11cdP1tpEnGgqG+0wgjxgsYz6EdrBOj/UQeIGE9tCM/G13+cZFNIDhuVCMm\nIRkMqbdBVjlXmqi1f+bxYYPV8+2A6yCHN4x9fq9LZzAYXfReaG+wzGLIhXwvKEtq\nTOBCFkgnYu7wYZYwNEtmTO49xaZQfWCpNAbREupLD7ECgYEA4a10ye11NUGy7Cym\nFzM8XgTI0gF4NCg+UMoGR7D1uoVOI6f0Y4JTwauYRZaR7K5IN4k4i27GgEerYfck\nLGWpkA6gpkclunUQ8Ubxm0YBHZnfmheqMoayP/SzoGQdknfdPkiXhaI37IWejKcO\n5ILfxSGK8vzaXPZpKomljk9hJ40CgYEA4F5nqkWYVzbWKUPboap8v8Ct03y6DzP+\nS1vgQRsm01JQnNfklETXof8+YCJqfMuSrykzRGAe7TwGWmLgwLvlxiWgRsTAfeo2\ns4tTZmMcBVa2MvtqdJD2njOlfFMfCejkV/8KKXy9DJ1yQ9euWtblesU9CBJRr5SC\n9svOPGHYCHMCgYA9samHuj6cfIVpQxt0pDEQksZDgttVhtriQxhMaPgEMYUXAkcx\nHOPAwiQygeMKjOp5JC4tD+98Chu0AFgHOxOLqjQIwNJzkqU7EGXkSNLtQK979JQ2\nk9QO39prMnNTIyl8aWPiyGH5at3ZHaJYnd6GiZDutGkNmN9PHaoAqXqp0QKBgCav\noGg3f8Dp75tF3ATQBJp7en1QsDQW3u3XdZ9EMzmUo9mnT/5QsG16OSMSTBIgd7ZE\nAFb1y99TzjSff+k7fK7hpfUNz7LmQ3BJwaORyy8QeHHp770RkbRNa2c4Xc2znkud\n6f6lR2N5ck5ITgPTsdWtVIyju/nuPXaYRYMby8gJAoGBANoLWRp5Vyic8S8hBS/C\nDILTDvolupQTrcUfMN0shdTdjxcHnU8Wr0XhMTrMdKWR4hVPuvv/Y58oG3w709Bo\nUghRv7EU457jiqqUtoI/R+NV+WGK3v1Yf57dd7qCH+7KVsnckmwwLzx2JXqzcpqe\nPuFdbAb/aYtEr88nS44l8FXv\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAPMNo6eG0UBaMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTE5WhcNMjYwNTIxMTEzOTE5WjBkMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEjAQBgNVBAMMCWNoZWZub2RlMTCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAKP+6Y2O2wAUGqWBBLDseChAoIu+Kr/Nc/RKvRky\nWKLMZj8Hmknqu7HjZsAXQVwquov8oHOSjifd9OA0EdxklQUeIv3VmrNbGNzavGZJ\n31g6E8fvHfUT6jOMEYoKU05H6nJ/TAusqizcexWBkoUzSUqZv/Eh+ssxpofOpQTa\nV8/80hxuoCAsxiDB07kuY497NaUf6gqE+TqrFfcqCJ1UB9BDruk1gP5K4w2BPz85\nizPP9sTeXNLuNGu3r/xQq9Xz6NL9vMFuzPnUURhwKJLdd3my8OkLQhxMeJy+nnIe\n14/JjPHEIS38/RPt0hShVdDvzQ0Avxj8IYrpNREejuiyMz0CAwEAAaM3MDUwHQYD\nVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBQGA1UdEQQNMAuCCWNoZWZub2Rl\nMTANBgkqhkiG9w0BAQsFAAOCAQEAI9B0IJgmeqgLgFvBiHhGsNzxAKo3z+YmU2Ta\nbqmcBO8gTGnLsnaQPs25sDPvM7YEkcjazhj74+f+L70+rAXl45TLkLQnIGpa4Bbg\nuBpYonPRzK3aSiDcnTeTH7LivuTJJQZptaT9jrcAcpK6AzWCopWR/E1rQ/oRCfiu\n4/PGV2nllNHC8rZm4YB3uftmjaWiwISf/gRSuD5yGu1TcCnYr5w3PhvkVAKHYLdj\ntIUlDTuTFXO/92ZCuW74YqBay+dAIB4ThU0jvUNYWTFV8MHmBgQ9CfG9WW4pjORI\nqalq5zXuKm1t+lrxCFND6cXu9Uk8HtrnSS5IqF0DprdepkyYhA==\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCj/umNjtsAFBql\ngQSw7HgoQKCLviq/zXP0Sr0ZMliizGY/B5pJ6rux42bAF0FcKrqL/KBzko4n3fTg\nNBHcZJUFHiL91ZqzWxjc2rxmSd9YOhPH7x31E+ozjBGKClNOR+pyf0wLrKos3HsV\ngZKFM0lKmb/xIfrLMaaHzqUE2lfP/NIcbqAgLMYgwdO5LmOPezWlH+oKhPk6qxX3\nKgidVAfQQ67pNYD+SuMNgT8/OYszz/bE3lzS7jRrt6/8UKvV8+jS/bzBbsz51FEY\ncCiS3Xd5svDpC0IcTHicvp5yHtePyYzxxCEt/P0T7dIUoVXQ780NAL8Y/CGK6TUR\nHo7osjM9AgMBAAECggEASiB0AxdaaDuuG7conqwUV+V2bBPmENJWIksSFGyMYfHQ\nGZdfJyAh/PNTw2n/kiCCN7pV8EeDWAPcpucCV8NjFHAd0uyVQ5Letx1r4TRs7t05\nibrMqLV6vBgI6YNnSk/5ag2eGvzN4v8552utBeY7r6u1ddItIWFs65/9OSdUX98m\n9iKC9n4D7pFvJsRoUfeD5qf5tF2cmAGS8z+y3502LPx0rNoJchHh06bkEwQIdA7Y\nTUsXIj6RXZqHgcyD0CGVI0gsT6lSywSzfUQvgLEV6Py4VD+1t2jo97bgScbLaFRN\nupC41HJFloYBvin0jtjgo/x8OUTGgW5IkBmrhw5roQKBgQDRoQiWiNtVYh+3AX2g\nDkDXqUbJLdWEElbkjxqzrC19jHCOI12S67MvNZGZ3ET//5agF5pWo/ZE1ZUy7WoJ\n2C9IdbKauGFlA0Drl1xrjHXU22/w8iDFp/F9lhbvu0vGIlCGGZlarl97Ulrha33a\nEHtxNtX7jlu3yrZdReqrUWEhWQKBgQDIRbyaVbVw8nBAKj49cqi8AayD/aAt1sYE\nKuDG0pv1ucXkOYtO2kvTxVUdBIPozdK4nVo/XT+mZ7PYSculukegjkO/Y51F6Tg+\n4jlQxMFv30UAhslacFcSNy5KCePZX/5sUbHYa6Amp5dYXMf1hud02PdIMoNygNZ4\nKKeQhL7ghQKBgF1XtTlCi1fDr5ePpF6muhzNlWVzcUWz3Nk9F4i1vDPRWzUPblVD\nerAkzEaUnGzZZDq5B9JYhAo2iI76xGLJzpQXRIY8X7HY9wlwhoilLLqxU3EYf5tD\novZm5KOu5Ji/ItfzgiOszXteOnVxpcJ54F2TK0kuJIz8SKPTxCCwxe1RAoGAPA3N\nVGpHEitgxZzlNP/g4R+PX7T6B0TT9AP3iyc0ZSbj1F/9ChQjkMknkJ/9/h1aBsoI\ned+4amnGYCEg0/1b5SVD42w3iPM6ToD/ttyJNMa6pkHEtz3gnjG1y7XTgSdr34dP\n0RnU2EKA+5o2y8U8Oqmk3R1olTlVFor6VDe6FRECgYAsgRurQPvNyWTjtXp8QLZ/\njYa8jqXkS7a26JVGKG6gq5hoaFq/fRksNYyC4H88lD/zxRmzkFkrEDIyhjVh5OMN\nPAJTLOnebC6W2xbisRJqDB2LfxmbaH2FA4kO6UzZkhYjudThu4y8uOQlEcVQO7GU\nVPVlMVkJXQmXIl6v3hCN+Q==\n-----END PRIVATE KEY-----",
							AdminKey:  "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQD0aBYmZyhNjsbi\n16ybvxYSBYrIKTdCrZ2g/FeAnfxVY4/85J6KWvvfoWFDqp0nAShlmQb7Fe5GQuN1\nIMTDbpCgYxAVX6Kmk6E6iKcmpko5p5K2vEo2A0/0V9vh2W5mdJjjGOtjDgtYAYR9\naDMdLdyagKfb3AnaOx+TotQuij0nFyYaTKGwK7sh5/y9Pyt0floeB0apndiJbh+e\nlDzG2t8QLBivsuIl0pI510PZSVVcWIBkJ2RVAPyy8gy/wOInsE08o2+0SrIxz2f7\nj+4ZFKfXEUiBopiwy7k5aFeq2gXF2r/Ndpi96uNsP4CXxHkUd+7jcFZQVHPiXNIX\nP6MCLbLXAgMBAAECggEACD7WuHb0eiFd/lsuXJbGxNbhBr21Oo+m6L56qUErOSpB\nulNwMdS9+J52LJU99gno9fyCqsfjoQUyrUnsuXcqc+7DpSTz1NDYOKRRl1E24dkQ\nbw/NJSNZeDHanjT6r4QxgD/f+RiJM2/hq2VvjAV3EtNSVm2G+5DREOcGZ4eMZpwl\ncPg1xrfEOf4Q7mLmiGxBW7tdGzilTsQSwzYifMWzdRmuxNSB8+jIu7Zy53mdUNyc\nI8nDxX8nfLND83GSPlzYuu2S5+cIHumKwtfBcde06Mn72bCLlkHF/AQEWMzGSqv6\nBQAeK3Dr1wI0JZ2Bil6L5Jld3xuRmu++i0DMF8pnwQKBgQD834ChWyoNu1YbXijS\nf53TEfuRZEBgyiOtfgrW723L9uMM975OCBJrhZ4/+WlUGIiZvr2COES3ikxYCzQ6\nkcukAjj9f0JupIOxABT+dwcJIJ4JDdFD8tUwAvY9ys2/WoOs5TdcCxRntrjL0xjG\nZUtDnMI90aOmihY1LB6yEVwh4QKBgQD3bchTW6jC2UdS52Kosdy7YhsgTdiZ+Mqs\nkL52SzBgd0lK4K35nHMUd+C3O90icKEfrpJX1cSbyFfSuimUQdK0I1bl0ijL8+uJ\nFkhOC0iy6um8xO/dvP66j7cuvNXYfOiJAqFJfGiSLe/vuT+JerbFFEidIRXqWSlh\nLIHtvxbbtwKBgBxWL2Plg2DmjU+jzY9JHbZ5XWd9hHlULYtThINxcSxaDjd1y62S\n2f2Si5k/qb3ywdv4s+PTyl+G7+ct2jx1+gv288v0Zs1fQiKjj7a0P+WV8h+xnLGw\nlJM8wbtK7qNy0S6ewQVfeHnmz+6HSU9yKmz5NAsZYu1WrAZpW0c5CsoBAoGBAL1P\naumUhM/obLDaxtqpk2hvjK+vwB02hON5r7BUoQP94L8AnzwPXuF3QyEPFYfHQxA5\nglDgBxjmNYPO2gdMQYmATHl0zbAWxczSlqnX6lyybfn3eEtg0kktsot5Aeks0MIb\nmAngvSWzLhRt2VY35OVvOou2h80RQR7Pbe3YugWLAoGATLlLM3U80bexIiyvwmot\n0gFgd8NIIDF48jTYEpZ9lsfieb7QfkvVVlvm3YsWaDtSd3OswWN6tkyFnz2GJxEB\nj8yVCBhAOxBfDRpMRzOpBe2v8u/o2Pxtw1leHmHbWYqbfwJnupnx3wnrzCVm+oLP\nUHB/aximuVLK+eRz/ZIZ3KU=\n-----END PRIVATE KEY-----",
							AdminCert: "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAPMNo6eG0UBZMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTE5WhcNMjYwNTIxMTEzOTE5WjBkMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEjAQBgNVBAMMCWNoZWZhZG1pbjCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAPRoFiZnKE2OxuLXrJu/FhIFisgpN0KtnaD8V4Cd\n/FVjj/zknopa+9+hYUOqnScBKGWZBvsV7kZC43UgxMNukKBjEBVfoqaToTqIpyam\nSjmnkra8SjYDT/RX2+HZbmZ0mOMY62MOC1gBhH1oMx0t3JqAp9vcCdo7H5Oi1C6K\nPScXJhpMobAruyHn/L0/K3R+Wh4HRqmd2IluH56UPMba3xAsGK+y4iXSkjnXQ9lJ\nVVxYgGQnZFUA/LLyDL/A4iewTTyjb7RKsjHPZ/uP7hkUp9cRSIGimLDLuTloV6ra\nBcXav812mL3q42w/gJfEeRR37uNwVlBUc+Jc0hc/owItstcCAwEAAaM3MDUwHQYD\nVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBQGA1UdEQQNMAuCCWNoZWZhZG1p\nbjANBgkqhkiG9w0BAQsFAAOCAQEAD4uod0aFfmwHuXshXFPcvu5gkkvd2Prd8pMz\nTiOu0A1IYgVHvAg0Bqs8bReZLsijUKhhjmlfwaUrW1i2r5jYp9Acd2K/rKwl7rov\ngaP60IgkszFCjfZFU1Zb0+6OlPuG8PquMPSJuJDFm5mQkw52vb81guvF/cGlKCuI\nfPwEr1CSOyHmQMQLLJsdzfB2RQ6MfG1pCH1FNPCfWjK8Rrd0Hic/sQ4Yd6q/Bmek\nb9KA5lUyk1Ox/YBBfm4RfmeKRqQKtASF8UJG3eCut0h3+uqyQpCxAx/8MlipkF7g\n1EvU3J0+e99VMZJ/wNlrAla08VsV/hjXd+NIZYkl31cghLXmTA==\n-----END CERTIFICATE-----",
						},
					},
				},
				ExternalOS: ExternalOS{
					OSDomainName:   "managed-services-os",
					OSDomainURL:    "search-managed-services-os-eckom3msrwqlmjlgbdu.us-east-1.es.amazonaws.com",
					OSUsername:     "admin",
					OSUserPassword: "Progress@123",
					OSCert:         "",
					OSRoleArn:      "arn:aws:iam::1127583934333:role/managed-services",
				},
				ExternalPG: ExternalPG{
					PGInstanceURL:       "managed-rds-db.cww4poze5gkx.ap-northeast-1.rds.amazonaws.com:5432",
					PGSuperuserName:     "postgres",
					PGSuperuserPassword: "chefautomate",
					PGDbUserName:        "postgres",
					PGDbUserPassword:    "chefautomate",
					PGRootCert:          "",
				},
				DeploymentState: "",
				APIToken:        "",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:     "PopulateWith OnPrem Config",
			filePath: "./testdata/HaOnPrem.toml",
			want: &Config{
				SSHUser: SSHUser{
					Username:     "ubuntu",
					PrivateKey:   "./testdata/A2HA.pem",
					SudoPassword: "",
				},
				Arch: "existing_nodes",
				Backup: Backup{
					FileSystem: FileSystem{
						MountLocation: "automate_backups",
					},
					ObjectStorage: ObjectStorage{
						Endpoint:   "s3.amazonaws.com",
						BucketName: "test",
						BasePath:   "",
						AccessKey:  "test_access_key",
						SecretKey:  "test_secret_key",
						AWSRegion:  "us-west-1",
					},
				},
				Hardware: Hardware{
					AutomateNodeCount: 2,
					AutomateNodeIps: []string{
						"192.0.0.1", "192.0.0.2",
					},
					ChefInfraServerNodeCount: 2,
					ChefInfraServerNodeIps: []string{
						"192.0.1.1", "192.0.1.2",
					},
					PostgresqlNodeCount: 3,
					PostgresqlNodeIps: []string{
						"192.0.3.1", "192.0.3.2", "192.0.3.3",
					},
					OpenSearchNodeCount: 3,
					OpenSearchNodeIps: []string{
						"192.0.2.1", "192.0.2.2", "192.0.2.3",
					},
				},
				Certificate: Certificate{
					AutomateFqdn:   "chefautomate.example.com",
					ChefServerFqdn: "chefautomate.example.com",
					RootCert:       "-----BEGIN CERTIFICATE-----\nMIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl\nMCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp\nU3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw\nNjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE\nChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp\nZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3\nDQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf\n8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN\n+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0\nX9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa\nK4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA\n1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G\nA1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR\nzt0fhvRbVazc1xDCDqmI56FspGowaDELMAkGA1UEBhMCVVMxJTAjBgNVBAoTHFN0\nYXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD\nbGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w\nDQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3\nL7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D\neruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl\nxy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp\nVSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY\nWQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8fF5Q=\n-----END CERTIFICATE-----",
					Nodes: []NodeCert{
						{
							IP:        "192.0.0.1",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgzCCAmugAwIBAgIJAPMNo6eG0UBgMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIwWhcNMjYwNTIxMTEzOTIwWjBlMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEzARBgNVBAMMCmNoZWZjbGllbnQwggEiMA0GCSqGSIb3\nDQEBAQUAA4IBDwAwggEKAoIBAQCyJbvhtNLZrgz7aa1W1gIGFrwtI5h23S99A8VZ\nlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6lSopXzmyzxgbuLSnAmmicNrsF\nI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/npzXIEa3HENUghVHOyR1nY8d\nvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFovC2z+HikVhkt8NB1pD+7ZVTq\nS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2jpMvFSVBdUi3HgdK+wO5pD0Jt\ngWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTjs04t/R1BAgMBAAGjODA2MB0G\nA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAVBgNVHREEDjAMggpjaGVmY2xp\nZW50MA0GCSqGSIb3DQEBCwUAA4IBAQAOHPBbbNCFCQ+Fg5YSvgXCglNUXfKoeCla\nLjruho1re6mJQmeW0H5OTEL7CssEHM6yzR6rlVlVnnu0+RAFlb6vTXo8eK4HaXS2\nA+OY7oP7UpMcxj3fxuWWgg1lnW/8a9Z+9JNSm9M59MpU17whvq/i8EJppECK4pD/\nZ9J12XYepWxARbVUcgWyHFIjnSoNpFwgrm1xybnyJ2WnjRRjpq6JvkP5+eklESEm\nZ4wjUCokbRJ8gbC7tuVgQk6DNM123j9djBm+A+5WlbvHUrD5trfbkp8kUJY4jaJF\ncsKeX402wz9P7XM5eGsToNpAZq41Q7mFzz14DfqFNttaCMHMYi4k\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCyJbvhtNLZrgz7\naa1W1gIGFrwtI5h23S99A8VZlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6l\nSopXzmyzxgbuLSnAmmicNrsFI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/\nnpzXIEa3HENUghVHOyR1nY8dvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFo\nvC2z+HikVhkt8NB1pD+7ZVTqS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2j\npMvFSVBdUi3HgdK+wO5pD0JtgWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTj\ns04t/R1BAgMBAAECggEATFvmM3Gpawqyn9UFKpJina+DCyoVwOU/5KsIHUTP0XkC\n3ASgWDPb1tozy36RmmjMcbFz9lOboZhiwoC32bkuWMRJ1i9flSHaMxM2iJaHckhh\nSaaubvFptw9of7XllG7m23CRyJJMuEXT4YCgI4m/Jd+kcIWtjzGniA53+LvxUM/S\n+pR2j5cc7yHYcsc4tz842UelVQJWnMwjHwYDsuZC4ssRq1dffq7ifgfhcrxcwAXg\nMMBEtmJsTeaBYjbiM4qdVTwc8BpX9lafgUi3asapR0X2u2k1hOs5fS0qdEf5IarN\nFqwQDj5Xurcs0xvUoTPC7eEqFMh1SRVw8mS0hZG0EQKBgQDZ9GHhIPjDd3NwT38R\nhHYd3i1EKUtJgoOlqN/WUSlcFBUKxEsQTF+5pZipPYFUjb8+DwzwQMmeSXAs9A+Z\nPCFQDO1EHlE1Hv6A0Fv7RPTt3lnkZpo7SdSf/VqxKVCLOIkeU9tEMr0d/8mJ8hbj\nlDIUmfZtWIo5eZnM6fQF8rM8LQKBgQDRPoIf1d5M6eMDz6Ss0dPF/UkhEhsXUQgh\nd0puiyKuk0e2v6q8sESKV4CnGDFd08XLMcSnOw3SS/neLhZ7lixjtewUoDAVwxlB\nNos3W6FBIgyQRFp5SBjuHhKjPiK+BPlGsc7+39I12aGxE0r5w7Pp4X+4Z7hDF8ID\nqlZtDkIN5QKBgQCe8aIjnHjtiwHraH3hF3lP5MOcDoUx8XTx7Up3L6760EZcGLQp\nCZlReFrxKMJVGB3cMvubhZPC1AlzLvTlKb2ddB/fakCMfbLZ25kIj8wSX/GsJ8rX\n68qcdhWaVue+75bHQB4KCPpzkyK1b4+TnXI8Jd9Y9JWwvmYT0pU7dTeSbQKBgQC1\nko6MXaQoDhWG6xq1NOeWOXLKFdIYa6Kol8GpJ2eTIg7rEGtyjWsMuV3UofPEvc43\nwxopG9+ki3VqTYgI+onOhME2LMNNPx2dL12jTgoiYQ+R6R6xe9TWXJZDvdmcFujR\nZd5/4W2ieRYMePdowWBQJfQU6zxETEt5rsiMngDH2QKBgHB8r2LQvvQX4L0w2wY+\nOJZOGKk4AMXIFM/J0qY60DLnw0B4RuzSXcKw2EGYRJihbGJAI8+KczU7LpVlOsw5\nFzQabRrE33NPYWFv0bqzas9/p9mPwFPzhq13mkrLOHwf71T63OEalaOh9WzaE2Tx\nhTp1M8sBDwrDLxCEk0X3NEL+\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "192.0.0.2",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgzCCAmugAwIBAgIJAPMNo6eG0UBgMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIwWhcNMjYwNTIxMTEzOTIwWjBlMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEzARBgNVBAMMCmNoZWZjbGllbnQwggEiMA0GCSqGSIb3\nDQEBAQUAA4IBDwAwggEKAoIBAQCyJbvhtNLZrgz7aa1W1gIGFrwtI5h23S99A8VZ\nlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6lSopXzmyzxgbuLSnAmmicNrsF\nI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/npzXIEa3HENUghVHOyR1nY8d\nvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFovC2z+HikVhkt8NB1pD+7ZVTq\nS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2jpMvFSVBdUi3HgdK+wO5pD0Jt\ngWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTjs04t/R1BAgMBAAGjODA2MB0G\nA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAVBgNVHREEDjAMggpjaGVmY2xp\nZW50MA0GCSqGSIb3DQEBCwUAA4IBAQAOHPBbbNCFCQ+Fg5YSvgXCglNUXfKoeCla\nLjruho1re6mJQmeW0H5OTEL7CssEHM6yzR6rlVlVnnu0+RAFlb6vTXo8eK4HaXS2\nA+OY7oP7UpMcxj3fxuWWgg1lnW/8a9Z+9JNSm9M59MpU17whvq/i8EJppECK4pD/\nZ9J12XYepWxARbVUcgWyHFIjnSoNpFwgrm1xybnyJ2WnjRRjpq6JvkP5+eklESEm\nZ4wjUCokbRJ8gbC7tuVgQk6DNM123j9djBm+A+5WlbvHUrD5trfbkp8kUJY4jaJF\ncsKeX402wz9P7XM5eGsToNpAZq41Q7mFzz14DfqFNttaCMHMYi4k\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCyJbvhtNLZrgz7\naa1W1gIGFrwtI5h23S99A8VZlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6l\nSopXzmyzxgbuLSnAmmicNrsFI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/\nnpzXIEa3HENUghVHOyR1nY8dvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFo\nvC2z+HikVhkt8NB1pD+7ZVTqS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2j\npMvFSVBdUi3HgdK+wO5pD0JtgWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTj\ns04t/R1BAgMBAAECggEATFvmM3Gpawqyn9UFKpJina+DCyoVwOU/5KsIHUTP0XkC\n3ASgWDPb1tozy36RmmjMcbFz9lOboZhiwoC32bkuWMRJ1i9flSHaMxM2iJaHckhh\nSaaubvFptw9of7XllG7m23CRyJJMuEXT4YCgI4m/Jd+kcIWtjzGniA53+LvxUM/S\n+pR2j5cc7yHYcsc4tz842UelVQJWnMwjHwYDsuZC4ssRq1dffq7ifgfhcrxcwAXg\nMMBEtmJsTeaBYjbiM4qdVTwc8BpX9lafgUi3asapR0X2u2k1hOs5fS0qdEf5IarN\nFqwQDj5Xurcs0xvUoTPC7eEqFMh1SRVw8mS0hZG0EQKBgQDZ9GHhIPjDd3NwT38R\nhHYd3i1EKUtJgoOlqN/WUSlcFBUKxEsQTF+5pZipPYFUjb8+DwzwQMmeSXAs9A+Z\nPCFQDO1EHlE1Hv6A0Fv7RPTt3lnkZpo7SdSf/VqxKVCLOIkeU9tEMr0d/8mJ8hbj\nlDIUmfZtWIo5eZnM6fQF8rM8LQKBgQDRPoIf1d5M6eMDz6Ss0dPF/UkhEhsXUQgh\nd0puiyKuk0e2v6q8sESKV4CnGDFd08XLMcSnOw3SS/neLhZ7lixjtewUoDAVwxlB\nNos3W6FBIgyQRFp5SBjuHhKjPiK+BPlGsc7+39I12aGxE0r5w7Pp4X+4Z7hDF8ID\nqlZtDkIN5QKBgQCe8aIjnHjtiwHraH3hF3lP5MOcDoUx8XTx7Up3L6760EZcGLQp\nCZlReFrxKMJVGB3cMvubhZPC1AlzLvTlKb2ddB/fakCMfbLZ25kIj8wSX/GsJ8rX\n68qcdhWaVue+75bHQB4KCPpzkyK1b4+TnXI8Jd9Y9JWwvmYT0pU7dTeSbQKBgQC1\nko6MXaQoDhWG6xq1NOeWOXLKFdIYa6Kol8GpJ2eTIg7rEGtyjWsMuV3UofPEvc43\nwxopG9+ki3VqTYgI+onOhME2LMNNPx2dL12jTgoiYQ+R6R6xe9TWXJZDvdmcFujR\nZd5/4W2ieRYMePdowWBQJfQU6zxETEt5rsiMngDH2QKBgHB8r2LQvvQX4L0w2wY+\nOJZOGKk4AMXIFM/J0qY60DLnw0B4RuzSXcKw2EGYRJihbGJAI8+KczU7LpVlOsw5\nFzQabRrE33NPYWFv0bqzas9/p9mPwFPzhq13mkrLOHwf71T63OEalaOh9WzaE2Tx\nhTp1M8sBDwrDLxCEk0X3NEL+\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "192.0.1.1",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgzCCAmugAwIBAgIJAPMNo6eG0UBgMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIwWhcNMjYwNTIxMTEzOTIwWjBlMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEzARBgNVBAMMCmNoZWZjbGllbnQwggEiMA0GCSqGSIb3\nDQEBAQUAA4IBDwAwggEKAoIBAQCyJbvhtNLZrgz7aa1W1gIGFrwtI5h23S99A8VZ\nlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6lSopXzmyzxgbuLSnAmmicNrsF\nI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/npzXIEa3HENUghVHOyR1nY8d\nvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFovC2z+HikVhkt8NB1pD+7ZVTq\nS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2jpMvFSVBdUi3HgdK+wO5pD0Jt\ngWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTjs04t/R1BAgMBAAGjODA2MB0G\nA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAVBgNVHREEDjAMggpjaGVmY2xp\nZW50MA0GCSqGSIb3DQEBCwUAA4IBAQAOHPBbbNCFCQ+Fg5YSvgXCglNUXfKoeCla\nLjruho1re6mJQmeW0H5OTEL7CssEHM6yzR6rlVlVnnu0+RAFlb6vTXo8eK4HaXS2\nA+OY7oP7UpMcxj3fxuWWgg1lnW/8a9Z+9JNSm9M59MpU17whvq/i8EJppECK4pD/\nZ9J12XYepWxARbVUcgWyHFIjnSoNpFwgrm1xybnyJ2WnjRRjpq6JvkP5+eklESEm\nZ4wjUCokbRJ8gbC7tuVgQk6DNM123j9djBm+A+5WlbvHUrD5trfbkp8kUJY4jaJF\ncsKeX402wz9P7XM5eGsToNpAZq41Q7mFzz14DfqFNttaCMHMYi4k\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCyJbvhtNLZrgz7\naa1W1gIGFrwtI5h23S99A8VZlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6l\nSopXzmyzxgbuLSnAmmicNrsFI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/\nnpzXIEa3HENUghVHOyR1nY8dvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFo\nvC2z+HikVhkt8NB1pD+7ZVTqS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2j\npMvFSVBdUi3HgdK+wO5pD0JtgWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTj\ns04t/R1BAgMBAAECggEATFvmM3Gpawqyn9UFKpJina+DCyoVwOU/5KsIHUTP0XkC\n3ASgWDPb1tozy36RmmjMcbFz9lOboZhiwoC32bkuWMRJ1i9flSHaMxM2iJaHckhh\nSaaubvFptw9of7XllG7m23CRyJJMuEXT4YCgI4m/Jd+kcIWtjzGniA53+LvxUM/S\n+pR2j5cc7yHYcsc4tz842UelVQJWnMwjHwYDsuZC4ssRq1dffq7ifgfhcrxcwAXg\nMMBEtmJsTeaBYjbiM4qdVTwc8BpX9lafgUi3asapR0X2u2k1hOs5fS0qdEf5IarN\nFqwQDj5Xurcs0xvUoTPC7eEqFMh1SRVw8mS0hZG0EQKBgQDZ9GHhIPjDd3NwT38R\nhHYd3i1EKUtJgoOlqN/WUSlcFBUKxEsQTF+5pZipPYFUjb8+DwzwQMmeSXAs9A+Z\nPCFQDO1EHlE1Hv6A0Fv7RPTt3lnkZpo7SdSf/VqxKVCLOIkeU9tEMr0d/8mJ8hbj\nlDIUmfZtWIo5eZnM6fQF8rM8LQKBgQDRPoIf1d5M6eMDz6Ss0dPF/UkhEhsXUQgh\nd0puiyKuk0e2v6q8sESKV4CnGDFd08XLMcSnOw3SS/neLhZ7lixjtewUoDAVwxlB\nNos3W6FBIgyQRFp5SBjuHhKjPiK+BPlGsc7+39I12aGxE0r5w7Pp4X+4Z7hDF8ID\nqlZtDkIN5QKBgQCe8aIjnHjtiwHraH3hF3lP5MOcDoUx8XTx7Up3L6760EZcGLQp\nCZlReFrxKMJVGB3cMvubhZPC1AlzLvTlKb2ddB/fakCMfbLZ25kIj8wSX/GsJ8rX\n68qcdhWaVue+75bHQB4KCPpzkyK1b4+TnXI8Jd9Y9JWwvmYT0pU7dTeSbQKBgQC1\nko6MXaQoDhWG6xq1NOeWOXLKFdIYa6Kol8GpJ2eTIg7rEGtyjWsMuV3UofPEvc43\nwxopG9+ki3VqTYgI+onOhME2LMNNPx2dL12jTgoiYQ+R6R6xe9TWXJZDvdmcFujR\nZd5/4W2ieRYMePdowWBQJfQU6zxETEt5rsiMngDH2QKBgHB8r2LQvvQX4L0w2wY+\nOJZOGKk4AMXIFM/J0qY60DLnw0B4RuzSXcKw2EGYRJihbGJAI8+KczU7LpVlOsw5\nFzQabRrE33NPYWFv0bqzas9/p9mPwFPzhq13mkrLOHwf71T63OEalaOh9WzaE2Tx\nhTp1M8sBDwrDLxCEk0X3NEL+\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "192.0.1.2",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgzCCAmugAwIBAgIJAPMNo6eG0UBgMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIwWhcNMjYwNTIxMTEzOTIwWjBlMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEzARBgNVBAMMCmNoZWZjbGllbnQwggEiMA0GCSqGSIb3\nDQEBAQUAA4IBDwAwggEKAoIBAQCyJbvhtNLZrgz7aa1W1gIGFrwtI5h23S99A8VZ\nlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6lSopXzmyzxgbuLSnAmmicNrsF\nI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/npzXIEa3HENUghVHOyR1nY8d\nvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFovC2z+HikVhkt8NB1pD+7ZVTq\nS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2jpMvFSVBdUi3HgdK+wO5pD0Jt\ngWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTjs04t/R1BAgMBAAGjODA2MB0G\nA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAVBgNVHREEDjAMggpjaGVmY2xp\nZW50MA0GCSqGSIb3DQEBCwUAA4IBAQAOHPBbbNCFCQ+Fg5YSvgXCglNUXfKoeCla\nLjruho1re6mJQmeW0H5OTEL7CssEHM6yzR6rlVlVnnu0+RAFlb6vTXo8eK4HaXS2\nA+OY7oP7UpMcxj3fxuWWgg1lnW/8a9Z+9JNSm9M59MpU17whvq/i8EJppECK4pD/\nZ9J12XYepWxARbVUcgWyHFIjnSoNpFwgrm1xybnyJ2WnjRRjpq6JvkP5+eklESEm\nZ4wjUCokbRJ8gbC7tuVgQk6DNM123j9djBm+A+5WlbvHUrD5trfbkp8kUJY4jaJF\ncsKeX402wz9P7XM5eGsToNpAZq41Q7mFzz14DfqFNttaCMHMYi4k\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCyJbvhtNLZrgz7\naa1W1gIGFrwtI5h23S99A8VZlRySlbvJpNpJsQZpwBqW9HtUSsErL8u5dakgFn6l\nSopXzmyzxgbuLSnAmmicNrsFI/O5BkZtacvebOjBpfm0LnPsZIfhqnlqvn8nM3a/\nnpzXIEa3HENUghVHOyR1nY8dvFgBHElfMbIGnLUkBZaHoQcXonG8nAw89IzoiBFo\nvC2z+HikVhkt8NB1pD+7ZVTqS/oPYO5yoHLBFc7V8X8DyS9MoPyoFBCYu7V+/d2j\npMvFSVBdUi3HgdK+wO5pD0JtgWNWWfdP6XpiL5VmvwEeOx8yJ6pPIrTPZtcFinTj\ns04t/R1BAgMBAAECggEATFvmM3Gpawqyn9UFKpJina+DCyoVwOU/5KsIHUTP0XkC\n3ASgWDPb1tozy36RmmjMcbFz9lOboZhiwoC32bkuWMRJ1i9flSHaMxM2iJaHckhh\nSaaubvFptw9of7XllG7m23CRyJJMuEXT4YCgI4m/Jd+kcIWtjzGniA53+LvxUM/S\n+pR2j5cc7yHYcsc4tz842UelVQJWnMwjHwYDsuZC4ssRq1dffq7ifgfhcrxcwAXg\nMMBEtmJsTeaBYjbiM4qdVTwc8BpX9lafgUi3asapR0X2u2k1hOs5fS0qdEf5IarN\nFqwQDj5Xurcs0xvUoTPC7eEqFMh1SRVw8mS0hZG0EQKBgQDZ9GHhIPjDd3NwT38R\nhHYd3i1EKUtJgoOlqN/WUSlcFBUKxEsQTF+5pZipPYFUjb8+DwzwQMmeSXAs9A+Z\nPCFQDO1EHlE1Hv6A0Fv7RPTt3lnkZpo7SdSf/VqxKVCLOIkeU9tEMr0d/8mJ8hbj\nlDIUmfZtWIo5eZnM6fQF8rM8LQKBgQDRPoIf1d5M6eMDz6Ss0dPF/UkhEhsXUQgh\nd0puiyKuk0e2v6q8sESKV4CnGDFd08XLMcSnOw3SS/neLhZ7lixjtewUoDAVwxlB\nNos3W6FBIgyQRFp5SBjuHhKjPiK+BPlGsc7+39I12aGxE0r5w7Pp4X+4Z7hDF8ID\nqlZtDkIN5QKBgQCe8aIjnHjtiwHraH3hF3lP5MOcDoUx8XTx7Up3L6760EZcGLQp\nCZlReFrxKMJVGB3cMvubhZPC1AlzLvTlKb2ddB/fakCMfbLZ25kIj8wSX/GsJ8rX\n68qcdhWaVue+75bHQB4KCPpzkyK1b4+TnXI8Jd9Y9JWwvmYT0pU7dTeSbQKBgQC1\nko6MXaQoDhWG6xq1NOeWOXLKFdIYa6Kol8GpJ2eTIg7rEGtyjWsMuV3UofPEvc43\nwxopG9+ki3VqTYgI+onOhME2LMNNPx2dL12jTgoiYQ+R6R6xe9TWXJZDvdmcFujR\nZd5/4W2ieRYMePdowWBQJfQU6zxETEt5rsiMngDH2QKBgHB8r2LQvvQX4L0w2wY+\nOJZOGKk4AMXIFM/J0qY60DLnw0B4RuzSXcKw2EGYRJihbGJAI8+KczU7LpVlOsw5\nFzQabRrE33NPYWFv0bqzas9/p9mPwFPzhq13mkrLOHwf71T63OEalaOh9WzaE2Tx\nhTp1M8sBDwrDLxCEk0X3NEL+\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "192.0.3.1",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDfzCCAmegAwIBAgIJAPMNo6eG0UBdMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIwWhcNMjYwNTIxMTEzOTIwWjBjMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxETAPBgNVBAMMCGNoZWZub2RlMIIBIjANBgkqhkiG9w0B\nAQEFAAOCAQ8AMIIBCgKCAQEAxcr/RHZymJ7x6V3fVHNs+ppzWOSj+WX0JDm3P0/K\nR1dV5Drzicl8tJVpiR3fniCQOYo/wEA6uRplvmB1unPTTRa712X+WrzLX29sXJoL\n128thRpp6XI7R3fOXyNPEAGZO6gHnsSpHqdPZCCL6LmIS9wh4/+AqQ+KdoChxJ6D\n7RjVmivCU4p5tyLtbNhO47IEP/pTXXUmnX29zlY6dC1wcEsRt3zvwKJYwgXUgMsh\nVzB2h+7I8OBuPBXTjUFC8o4+yr7xjtIB1AKQYHeZaxWoy0z8YFUwCY4Zrvy2P7Yw\nE2dAVipnSUkNNtmRCseLqdchmhxwVKWCP1GokFZqfdksVwIDAQABozYwNDAdBgNV\nHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwEwYDVR0RBAwwCoIIY2hlZm5vZGUw\nDQYJKoZIhvcNAQELBQADggEBAK3O2VEYJQ4byjMYxuxmVEwHaNrl1XrwDZ7pKciq\nx9GyLFwYva4svKZfawUbZUQsIBLxRAaUDPVdfb8MYn1lTI4q2y9XGT1osM0SuL6f\ngllp4Yg4rTc94G5yXzhnWYCqrX9XZK0muFKCuJnniC0VmrP9zfojUsa0x6qvIfKf\nEB5u6SKRbQGS1ECryGQfiwXzhwy17/Qw7Ab44ufDpWmiw6fdNY/KhBKIwCRSHoJJ\nQQSO1eGYc/f3j7Kve5LaTOxADqBhGcxyItqdEJlmhKwAp2aOmO6pZLPCuZvCmvDL\nDv6bUUXSsZF4fb1diLIBpmD1hh8OGNY65LUPpzAxJeZvo5w=\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDFyv9EdnKYnvHp\nXd9Uc2z6mnNY5KP5ZfQkObc/T8pHV1XkOvOJyXy0lWmJHd+eIJA5ij/AQDq5GmW+\nYHW6c9NNFrvXZf5avMtfb2xcmgvXby2FGmnpcjtHd85fI08QAZk7qAeexKkep09k\nIIvouYhL3CHj/4CpD4p2gKHEnoPtGNWaK8JTinm3Iu1s2E7jsgQ/+lNddSadfb3O\nVjp0LXBwSxG3fO/AoljCBdSAyyFXMHaH7sjw4G48FdONQULyjj7KvvGO0gHUApBg\nd5lrFajLTPxgVTAJjhmu/LY/tjATZ0BWKmdJSQ022ZEKx4up1yGaHHBUpYI/UaiQ\nVmp92SxXAgMBAAECggEBALClRSk9p8bKXT6QCb6Af5mois+fExrPhSU9Ln0qo3rn\nctwsEgjCm88jiWdd+LJeXrAk2h62vjtGaguGVl44x0OXxBbxDiK3beJDvsFNCrpS\nnpK7Lk/BJ1QCmZq6DAg9hT6UKIoRFQE9Z1gDATDNUf5+EP5w19Uk/gIri03wS95Y\nt6xnA11cdP1tpEnGgqG+0wgjxgsYz6EdrBOj/UQeIGE9tCM/G13+cZFNIDhuVCMm\nIRkMqbdBVjlXmqi1f+bxYYPV8+2A6yCHN4x9fq9LZzAYXfReaG+wzGLIhXwvKEtq\nTOBCFkgnYu7wYZYwNEtmTO49xaZQfWCpNAbREupLD7ECgYEA4a10ye11NUGy7Cym\nFzM8XgTI0gF4NCg+UMoGR7D1uoVOI6f0Y4JTwauYRZaR7K5IN4k4i27GgEerYfck\nLGWpkA6gpkclunUQ8Ubxm0YBHZnfmheqMoayP/SzoGQdknfdPkiXhaI37IWejKcO\n5ILfxSGK8vzaXPZpKomljk9hJ40CgYEA4F5nqkWYVzbWKUPboap8v8Ct03y6DzP+\nS1vgQRsm01JQnNfklETXof8+YCJqfMuSrykzRGAe7TwGWmLgwLvlxiWgRsTAfeo2\ns4tTZmMcBVa2MvtqdJD2njOlfFMfCejkV/8KKXy9DJ1yQ9euWtblesU9CBJRr5SC\n9svOPGHYCHMCgYA9samHuj6cfIVpQxt0pDEQksZDgttVhtriQxhMaPgEMYUXAkcx\nHOPAwiQygeMKjOp5JC4tD+98Chu0AFgHOxOLqjQIwNJzkqU7EGXkSNLtQK979JQ2\nk9QO39prMnNTIyl8aWPiyGH5at3ZHaJYnd6GiZDutGkNmN9PHaoAqXqp0QKBgCav\noGg3f8Dp75tF3ATQBJp7en1QsDQW3u3XdZ9EMzmUo9mnT/5QsG16OSMSTBIgd7ZE\nAFb1y99TzjSff+k7fK7hpfUNz7LmQ3BJwaORyy8QeHHp770RkbRNa2c4Xc2znkud\n6f6lR2N5ck5ITgPTsdWtVIyju/nuPXaYRYMby8gJAoGBANoLWRp5Vyic8S8hBS/C\nDILTDvolupQTrcUfMN0shdTdjxcHnU8Wr0XhMTrMdKWR4hVPuvv/Y58oG3w709Bo\nUghRv7EU457jiqqUtoI/R+NV+WGK3v1Yf57dd7qCH+7KVsnckmwwLzx2JXqzcpqe\nPuFdbAb/aYtEr88nS44l8FXv\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "192.0.3.2",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDfzCCAmegAwIBAgIJAPMNo6eG0UBdMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIwWhcNMjYwNTIxMTEzOTIwWjBjMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxETAPBgNVBAMMCGNoZWZub2RlMIIBIjANBgkqhkiG9w0B\nAQEFAAOCAQ8AMIIBCgKCAQEAxcr/RHZymJ7x6V3fVHNs+ppzWOSj+WX0JDm3P0/K\nR1dV5Drzicl8tJVpiR3fniCQOYo/wEA6uRplvmB1unPTTRa712X+WrzLX29sXJoL\n128thRpp6XI7R3fOXyNPEAGZO6gHnsSpHqdPZCCL6LmIS9wh4/+AqQ+KdoChxJ6D\n7RjVmivCU4p5tyLtbNhO47IEP/pTXXUmnX29zlY6dC1wcEsRt3zvwKJYwgXUgMsh\nVzB2h+7I8OBuPBXTjUFC8o4+yr7xjtIB1AKQYHeZaxWoy0z8YFUwCY4Zrvy2P7Yw\nE2dAVipnSUkNNtmRCseLqdchmhxwVKWCP1GokFZqfdksVwIDAQABozYwNDAdBgNV\nHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwEwYDVR0RBAwwCoIIY2hlZm5vZGUw\nDQYJKoZIhvcNAQELBQADggEBAK3O2VEYJQ4byjMYxuxmVEwHaNrl1XrwDZ7pKciq\nx9GyLFwYva4svKZfawUbZUQsIBLxRAaUDPVdfb8MYn1lTI4q2y9XGT1osM0SuL6f\ngllp4Yg4rTc94G5yXzhnWYCqrX9XZK0muFKCuJnniC0VmrP9zfojUsa0x6qvIfKf\nEB5u6SKRbQGS1ECryGQfiwXzhwy17/Qw7Ab44ufDpWmiw6fdNY/KhBKIwCRSHoJJ\nQQSO1eGYc/f3j7Kve5LaTOxADqBhGcxyItqdEJlmhKwAp2aOmO6pZLPCuZvCmvDL\nDv6bUUXSsZF4fb1diLIBpmD1hh8OGNY65LUPpzAxJeZvo5w=\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDFyv9EdnKYnvHp\nXd9Uc2z6mnNY5KP5ZfQkObc/T8pHV1XkOvOJyXy0lWmJHd+eIJA5ij/AQDq5GmW+\nYHW6c9NNFrvXZf5avMtfb2xcmgvXby2FGmnpcjtHd85fI08QAZk7qAeexKkep09k\nIIvouYhL3CHj/4CpD4p2gKHEnoPtGNWaK8JTinm3Iu1s2E7jsgQ/+lNddSadfb3O\nVjp0LXBwSxG3fO/AoljCBdSAyyFXMHaH7sjw4G48FdONQULyjj7KvvGO0gHUApBg\nd5lrFajLTPxgVTAJjhmu/LY/tjATZ0BWKmdJSQ022ZEKx4up1yGaHHBUpYI/UaiQ\nVmp92SxXAgMBAAECggEBALClRSk9p8bKXT6QCb6Af5mois+fExrPhSU9Ln0qo3rn\nctwsEgjCm88jiWdd+LJeXrAk2h62vjtGaguGVl44x0OXxBbxDiK3beJDvsFNCrpS\nnpK7Lk/BJ1QCmZq6DAg9hT6UKIoRFQE9Z1gDATDNUf5+EP5w19Uk/gIri03wS95Y\nt6xnA11cdP1tpEnGgqG+0wgjxgsYz6EdrBOj/UQeIGE9tCM/G13+cZFNIDhuVCMm\nIRkMqbdBVjlXmqi1f+bxYYPV8+2A6yCHN4x9fq9LZzAYXfReaG+wzGLIhXwvKEtq\nTOBCFkgnYu7wYZYwNEtmTO49xaZQfWCpNAbREupLD7ECgYEA4a10ye11NUGy7Cym\nFzM8XgTI0gF4NCg+UMoGR7D1uoVOI6f0Y4JTwauYRZaR7K5IN4k4i27GgEerYfck\nLGWpkA6gpkclunUQ8Ubxm0YBHZnfmheqMoayP/SzoGQdknfdPkiXhaI37IWejKcO\n5ILfxSGK8vzaXPZpKomljk9hJ40CgYEA4F5nqkWYVzbWKUPboap8v8Ct03y6DzP+\nS1vgQRsm01JQnNfklETXof8+YCJqfMuSrykzRGAe7TwGWmLgwLvlxiWgRsTAfeo2\ns4tTZmMcBVa2MvtqdJD2njOlfFMfCejkV/8KKXy9DJ1yQ9euWtblesU9CBJRr5SC\n9svOPGHYCHMCgYA9samHuj6cfIVpQxt0pDEQksZDgttVhtriQxhMaPgEMYUXAkcx\nHOPAwiQygeMKjOp5JC4tD+98Chu0AFgHOxOLqjQIwNJzkqU7EGXkSNLtQK979JQ2\nk9QO39prMnNTIyl8aWPiyGH5at3ZHaJYnd6GiZDutGkNmN9PHaoAqXqp0QKBgCav\noGg3f8Dp75tF3ATQBJp7en1QsDQW3u3XdZ9EMzmUo9mnT/5QsG16OSMSTBIgd7ZE\nAFb1y99TzjSff+k7fK7hpfUNz7LmQ3BJwaORyy8QeHHp770RkbRNa2c4Xc2znkud\n6f6lR2N5ck5ITgPTsdWtVIyju/nuPXaYRYMby8gJAoGBANoLWRp5Vyic8S8hBS/C\nDILTDvolupQTrcUfMN0shdTdjxcHnU8Wr0XhMTrMdKWR4hVPuvv/Y58oG3w709Bo\nUghRv7EU457jiqqUtoI/R+NV+WGK3v1Yf57dd7qCH+7KVsnckmwwLzx2JXqzcpqe\nPuFdbAb/aYtEr88nS44l8FXv\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "192.0.3.3",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDfzCCAmegAwIBAgIJAPMNo6eG0UBdMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTIwWhcNMjYwNTIxMTEzOTIwWjBjMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxETAPBgNVBAMMCGNoZWZub2RlMIIBIjANBgkqhkiG9w0B\nAQEFAAOCAQ8AMIIBCgKCAQEAxcr/RHZymJ7x6V3fVHNs+ppzWOSj+WX0JDm3P0/K\nR1dV5Drzicl8tJVpiR3fniCQOYo/wEA6uRplvmB1unPTTRa712X+WrzLX29sXJoL\n128thRpp6XI7R3fOXyNPEAGZO6gHnsSpHqdPZCCL6LmIS9wh4/+AqQ+KdoChxJ6D\n7RjVmivCU4p5tyLtbNhO47IEP/pTXXUmnX29zlY6dC1wcEsRt3zvwKJYwgXUgMsh\nVzB2h+7I8OBuPBXTjUFC8o4+yr7xjtIB1AKQYHeZaxWoy0z8YFUwCY4Zrvy2P7Yw\nE2dAVipnSUkNNtmRCseLqdchmhxwVKWCP1GokFZqfdksVwIDAQABozYwNDAdBgNV\nHSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwEwYDVR0RBAwwCoIIY2hlZm5vZGUw\nDQYJKoZIhvcNAQELBQADggEBAK3O2VEYJQ4byjMYxuxmVEwHaNrl1XrwDZ7pKciq\nx9GyLFwYva4svKZfawUbZUQsIBLxRAaUDPVdfb8MYn1lTI4q2y9XGT1osM0SuL6f\ngllp4Yg4rTc94G5yXzhnWYCqrX9XZK0muFKCuJnniC0VmrP9zfojUsa0x6qvIfKf\nEB5u6SKRbQGS1ECryGQfiwXzhwy17/Qw7Ab44ufDpWmiw6fdNY/KhBKIwCRSHoJJ\nQQSO1eGYc/f3j7Kve5LaTOxADqBhGcxyItqdEJlmhKwAp2aOmO6pZLPCuZvCmvDL\nDv6bUUXSsZF4fb1diLIBpmD1hh8OGNY65LUPpzAxJeZvo5w=\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDFyv9EdnKYnvHp\nXd9Uc2z6mnNY5KP5ZfQkObc/T8pHV1XkOvOJyXy0lWmJHd+eIJA5ij/AQDq5GmW+\nYHW6c9NNFrvXZf5avMtfb2xcmgvXby2FGmnpcjtHd85fI08QAZk7qAeexKkep09k\nIIvouYhL3CHj/4CpD4p2gKHEnoPtGNWaK8JTinm3Iu1s2E7jsgQ/+lNddSadfb3O\nVjp0LXBwSxG3fO/AoljCBdSAyyFXMHaH7sjw4G48FdONQULyjj7KvvGO0gHUApBg\nd5lrFajLTPxgVTAJjhmu/LY/tjATZ0BWKmdJSQ022ZEKx4up1yGaHHBUpYI/UaiQ\nVmp92SxXAgMBAAECggEBALClRSk9p8bKXT6QCb6Af5mois+fExrPhSU9Ln0qo3rn\nctwsEgjCm88jiWdd+LJeXrAk2h62vjtGaguGVl44x0OXxBbxDiK3beJDvsFNCrpS\nnpK7Lk/BJ1QCmZq6DAg9hT6UKIoRFQE9Z1gDATDNUf5+EP5w19Uk/gIri03wS95Y\nt6xnA11cdP1tpEnGgqG+0wgjxgsYz6EdrBOj/UQeIGE9tCM/G13+cZFNIDhuVCMm\nIRkMqbdBVjlXmqi1f+bxYYPV8+2A6yCHN4x9fq9LZzAYXfReaG+wzGLIhXwvKEtq\nTOBCFkgnYu7wYZYwNEtmTO49xaZQfWCpNAbREupLD7ECgYEA4a10ye11NUGy7Cym\nFzM8XgTI0gF4NCg+UMoGR7D1uoVOI6f0Y4JTwauYRZaR7K5IN4k4i27GgEerYfck\nLGWpkA6gpkclunUQ8Ubxm0YBHZnfmheqMoayP/SzoGQdknfdPkiXhaI37IWejKcO\n5ILfxSGK8vzaXPZpKomljk9hJ40CgYEA4F5nqkWYVzbWKUPboap8v8Ct03y6DzP+\nS1vgQRsm01JQnNfklETXof8+YCJqfMuSrykzRGAe7TwGWmLgwLvlxiWgRsTAfeo2\ns4tTZmMcBVa2MvtqdJD2njOlfFMfCejkV/8KKXy9DJ1yQ9euWtblesU9CBJRr5SC\n9svOPGHYCHMCgYA9samHuj6cfIVpQxt0pDEQksZDgttVhtriQxhMaPgEMYUXAkcx\nHOPAwiQygeMKjOp5JC4tD+98Chu0AFgHOxOLqjQIwNJzkqU7EGXkSNLtQK979JQ2\nk9QO39prMnNTIyl8aWPiyGH5at3ZHaJYnd6GiZDutGkNmN9PHaoAqXqp0QKBgCav\noGg3f8Dp75tF3ATQBJp7en1QsDQW3u3XdZ9EMzmUo9mnT/5QsG16OSMSTBIgd7ZE\nAFb1y99TzjSff+k7fK7hpfUNz7LmQ3BJwaORyy8QeHHp770RkbRNa2c4Xc2znkud\n6f6lR2N5ck5ITgPTsdWtVIyju/nuPXaYRYMby8gJAoGBANoLWRp5Vyic8S8hBS/C\nDILTDvolupQTrcUfMN0shdTdjxcHnU8Wr0XhMTrMdKWR4hVPuvv/Y58oG3w709Bo\nUghRv7EU457jiqqUtoI/R+NV+WGK3v1Yf57dd7qCH+7KVsnckmwwLzx2JXqzcpqe\nPuFdbAb/aYtEr88nS44l8FXv\n-----END PRIVATE KEY-----",
							AdminKey:  "",
							AdminCert: "",
						}, {
							IP:        "192.0.2.1",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAPMNo6eG0UBaMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTE5WhcNMjYwNTIxMTEzOTE5WjBkMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEjAQBgNVBAMMCWNoZWZub2RlMTCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAKP+6Y2O2wAUGqWBBLDseChAoIu+Kr/Nc/RKvRky\nWKLMZj8Hmknqu7HjZsAXQVwquov8oHOSjifd9OA0EdxklQUeIv3VmrNbGNzavGZJ\n31g6E8fvHfUT6jOMEYoKU05H6nJ/TAusqizcexWBkoUzSUqZv/Eh+ssxpofOpQTa\nV8/80hxuoCAsxiDB07kuY497NaUf6gqE+TqrFfcqCJ1UB9BDruk1gP5K4w2BPz85\nizPP9sTeXNLuNGu3r/xQq9Xz6NL9vMFuzPnUURhwKJLdd3my8OkLQhxMeJy+nnIe\n14/JjPHEIS38/RPt0hShVdDvzQ0Avxj8IYrpNREejuiyMz0CAwEAAaM3MDUwHQYD\nVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBQGA1UdEQQNMAuCCWNoZWZub2Rl\nMTANBgkqhkiG9w0BAQsFAAOCAQEAI9B0IJgmeqgLgFvBiHhGsNzxAKo3z+YmU2Ta\nbqmcBO8gTGnLsnaQPs25sDPvM7YEkcjazhj74+f+L70+rAXl45TLkLQnIGpa4Bbg\nuBpYonPRzK3aSiDcnTeTH7LivuTJJQZptaT9jrcAcpK6AzWCopWR/E1rQ/oRCfiu\n4/PGV2nllNHC8rZm4YB3uftmjaWiwISf/gRSuD5yGu1TcCnYr5w3PhvkVAKHYLdj\ntIUlDTuTFXO/92ZCuW74YqBay+dAIB4ThU0jvUNYWTFV8MHmBgQ9CfG9WW4pjORI\nqalq5zXuKm1t+lrxCFND6cXu9Uk8HtrnSS5IqF0DprdepkyYhA==\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCj/umNjtsAFBql\ngQSw7HgoQKCLviq/zXP0Sr0ZMliizGY/B5pJ6rux42bAF0FcKrqL/KBzko4n3fTg\nNBHcZJUFHiL91ZqzWxjc2rxmSd9YOhPH7x31E+ozjBGKClNOR+pyf0wLrKos3HsV\ngZKFM0lKmb/xIfrLMaaHzqUE2lfP/NIcbqAgLMYgwdO5LmOPezWlH+oKhPk6qxX3\nKgidVAfQQ67pNYD+SuMNgT8/OYszz/bE3lzS7jRrt6/8UKvV8+jS/bzBbsz51FEY\ncCiS3Xd5svDpC0IcTHicvp5yHtePyYzxxCEt/P0T7dIUoVXQ780NAL8Y/CGK6TUR\nHo7osjM9AgMBAAECggEASiB0AxdaaDuuG7conqwUV+V2bBPmENJWIksSFGyMYfHQ\nGZdfJyAh/PNTw2n/kiCCN7pV8EeDWAPcpucCV8NjFHAd0uyVQ5Letx1r4TRs7t05\nibrMqLV6vBgI6YNnSk/5ag2eGvzN4v8552utBeY7r6u1ddItIWFs65/9OSdUX98m\n9iKC9n4D7pFvJsRoUfeD5qf5tF2cmAGS8z+y3502LPx0rNoJchHh06bkEwQIdA7Y\nTUsXIj6RXZqHgcyD0CGVI0gsT6lSywSzfUQvgLEV6Py4VD+1t2jo97bgScbLaFRN\nupC41HJFloYBvin0jtjgo/x8OUTGgW5IkBmrhw5roQKBgQDRoQiWiNtVYh+3AX2g\nDkDXqUbJLdWEElbkjxqzrC19jHCOI12S67MvNZGZ3ET//5agF5pWo/ZE1ZUy7WoJ\n2C9IdbKauGFlA0Drl1xrjHXU22/w8iDFp/F9lhbvu0vGIlCGGZlarl97Ulrha33a\nEHtxNtX7jlu3yrZdReqrUWEhWQKBgQDIRbyaVbVw8nBAKj49cqi8AayD/aAt1sYE\nKuDG0pv1ucXkOYtO2kvTxVUdBIPozdK4nVo/XT+mZ7PYSculukegjkO/Y51F6Tg+\n4jlQxMFv30UAhslacFcSNy5KCePZX/5sUbHYa6Amp5dYXMf1hud02PdIMoNygNZ4\nKKeQhL7ghQKBgF1XtTlCi1fDr5ePpF6muhzNlWVzcUWz3Nk9F4i1vDPRWzUPblVD\nerAkzEaUnGzZZDq5B9JYhAo2iI76xGLJzpQXRIY8X7HY9wlwhoilLLqxU3EYf5tD\novZm5KOu5Ji/ItfzgiOszXteOnVxpcJ54F2TK0kuJIz8SKPTxCCwxe1RAoGAPA3N\nVGpHEitgxZzlNP/g4R+PX7T6B0TT9AP3iyc0ZSbj1F/9ChQjkMknkJ/9/h1aBsoI\ned+4amnGYCEg0/1b5SVD42w3iPM6ToD/ttyJNMa6pkHEtz3gnjG1y7XTgSdr34dP\n0RnU2EKA+5o2y8U8Oqmk3R1olTlVFor6VDe6FRECgYAsgRurQPvNyWTjtXp8QLZ/\njYa8jqXkS7a26JVGKG6gq5hoaFq/fRksNYyC4H88lD/zxRmzkFkrEDIyhjVh5OMN\nPAJTLOnebC6W2xbisRJqDB2LfxmbaH2FA4kO6UzZkhYjudThu4y8uOQlEcVQO7GU\nVPVlMVkJXQmXIl6v3hCN+Q==\n-----END PRIVATE KEY-----",
							AdminKey:  "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQD0aBYmZyhNjsbi\n16ybvxYSBYrIKTdCrZ2g/FeAnfxVY4/85J6KWvvfoWFDqp0nAShlmQb7Fe5GQuN1\nIMTDbpCgYxAVX6Kmk6E6iKcmpko5p5K2vEo2A0/0V9vh2W5mdJjjGOtjDgtYAYR9\naDMdLdyagKfb3AnaOx+TotQuij0nFyYaTKGwK7sh5/y9Pyt0floeB0apndiJbh+e\nlDzG2t8QLBivsuIl0pI510PZSVVcWIBkJ2RVAPyy8gy/wOInsE08o2+0SrIxz2f7\nj+4ZFKfXEUiBopiwy7k5aFeq2gXF2r/Ndpi96uNsP4CXxHkUd+7jcFZQVHPiXNIX\nP6MCLbLXAgMBAAECggEACD7WuHb0eiFd/lsuXJbGxNbhBr21Oo+m6L56qUErOSpB\nulNwMdS9+J52LJU99gno9fyCqsfjoQUyrUnsuXcqc+7DpSTz1NDYOKRRl1E24dkQ\nbw/NJSNZeDHanjT6r4QxgD/f+RiJM2/hq2VvjAV3EtNSVm2G+5DREOcGZ4eMZpwl\ncPg1xrfEOf4Q7mLmiGxBW7tdGzilTsQSwzYifMWzdRmuxNSB8+jIu7Zy53mdUNyc\nI8nDxX8nfLND83GSPlzYuu2S5+cIHumKwtfBcde06Mn72bCLlkHF/AQEWMzGSqv6\nBQAeK3Dr1wI0JZ2Bil6L5Jld3xuRmu++i0DMF8pnwQKBgQD834ChWyoNu1YbXijS\nf53TEfuRZEBgyiOtfgrW723L9uMM975OCBJrhZ4/+WlUGIiZvr2COES3ikxYCzQ6\nkcukAjj9f0JupIOxABT+dwcJIJ4JDdFD8tUwAvY9ys2/WoOs5TdcCxRntrjL0xjG\nZUtDnMI90aOmihY1LB6yEVwh4QKBgQD3bchTW6jC2UdS52Kosdy7YhsgTdiZ+Mqs\nkL52SzBgd0lK4K35nHMUd+C3O90icKEfrpJX1cSbyFfSuimUQdK0I1bl0ijL8+uJ\nFkhOC0iy6um8xO/dvP66j7cuvNXYfOiJAqFJfGiSLe/vuT+JerbFFEidIRXqWSlh\nLIHtvxbbtwKBgBxWL2Plg2DmjU+jzY9JHbZ5XWd9hHlULYtThINxcSxaDjd1y62S\n2f2Si5k/qb3ywdv4s+PTyl+G7+ct2jx1+gv288v0Zs1fQiKjj7a0P+WV8h+xnLGw\nlJM8wbtK7qNy0S6ewQVfeHnmz+6HSU9yKmz5NAsZYu1WrAZpW0c5CsoBAoGBAL1P\naumUhM/obLDaxtqpk2hvjK+vwB02hON5r7BUoQP94L8AnzwPXuF3QyEPFYfHQxA5\nglDgBxjmNYPO2gdMQYmATHl0zbAWxczSlqnX6lyybfn3eEtg0kktsot5Aeks0MIb\nmAngvSWzLhRt2VY35OVvOou2h80RQR7Pbe3YugWLAoGATLlLM3U80bexIiyvwmot\n0gFgd8NIIDF48jTYEpZ9lsfieb7QfkvVVlvm3YsWaDtSd3OswWN6tkyFnz2GJxEB\nj8yVCBhAOxBfDRpMRzOpBe2v8u/o2Pxtw1leHmHbWYqbfwJnupnx3wnrzCVm+oLP\nUHB/aximuVLK+eRz/ZIZ3KU=\n-----END PRIVATE KEY-----",
							AdminCert: "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAPMNo6eG0UBZMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTE5WhcNMjYwNTIxMTEzOTE5WjBkMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEjAQBgNVBAMMCWNoZWZhZG1pbjCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAPRoFiZnKE2OxuLXrJu/FhIFisgpN0KtnaD8V4Cd\n/FVjj/zknopa+9+hYUOqnScBKGWZBvsV7kZC43UgxMNukKBjEBVfoqaToTqIpyam\nSjmnkra8SjYDT/RX2+HZbmZ0mOMY62MOC1gBhH1oMx0t3JqAp9vcCdo7H5Oi1C6K\nPScXJhpMobAruyHn/L0/K3R+Wh4HRqmd2IluH56UPMba3xAsGK+y4iXSkjnXQ9lJ\nVVxYgGQnZFUA/LLyDL/A4iewTTyjb7RKsjHPZ/uP7hkUp9cRSIGimLDLuTloV6ra\nBcXav812mL3q42w/gJfEeRR37uNwVlBUc+Jc0hc/owItstcCAwEAAaM3MDUwHQYD\nVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBQGA1UdEQQNMAuCCWNoZWZhZG1p\nbjANBgkqhkiG9w0BAQsFAAOCAQEAD4uod0aFfmwHuXshXFPcvu5gkkvd2Prd8pMz\nTiOu0A1IYgVHvAg0Bqs8bReZLsijUKhhjmlfwaUrW1i2r5jYp9Acd2K/rKwl7rov\ngaP60IgkszFCjfZFU1Zb0+6OlPuG8PquMPSJuJDFm5mQkw52vb81guvF/cGlKCuI\nfPwEr1CSOyHmQMQLLJsdzfB2RQ6MfG1pCH1FNPCfWjK8Rrd0Hic/sQ4Yd6q/Bmek\nb9KA5lUyk1Ox/YBBfm4RfmeKRqQKtASF8UJG3eCut0h3+uqyQpCxAx/8MlipkF7g\n1EvU3J0+e99VMZJ/wNlrAla08VsV/hjXd+NIZYkl31cghLXmTA==\n-----END CERTIFICATE-----",
						}, {
							IP:        "192.0.2.2",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAPMNo6eG0UBaMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTE5WhcNMjYwNTIxMTEzOTE5WjBkMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEjAQBgNVBAMMCWNoZWZub2RlMTCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAKP+6Y2O2wAUGqWBBLDseChAoIu+Kr/Nc/RKvRky\nWKLMZj8Hmknqu7HjZsAXQVwquov8oHOSjifd9OA0EdxklQUeIv3VmrNbGNzavGZJ\n31g6E8fvHfUT6jOMEYoKU05H6nJ/TAusqizcexWBkoUzSUqZv/Eh+ssxpofOpQTa\nV8/80hxuoCAsxiDB07kuY497NaUf6gqE+TqrFfcqCJ1UB9BDruk1gP5K4w2BPz85\nizPP9sTeXNLuNGu3r/xQq9Xz6NL9vMFuzPnUURhwKJLdd3my8OkLQhxMeJy+nnIe\n14/JjPHEIS38/RPt0hShVdDvzQ0Avxj8IYrpNREejuiyMz0CAwEAAaM3MDUwHQYD\nVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBQGA1UdEQQNMAuCCWNoZWZub2Rl\nMTANBgkqhkiG9w0BAQsFAAOCAQEAI9B0IJgmeqgLgFvBiHhGsNzxAKo3z+YmU2Ta\nbqmcBO8gTGnLsnaQPs25sDPvM7YEkcjazhj74+f+L70+rAXl45TLkLQnIGpa4Bbg\nuBpYonPRzK3aSiDcnTeTH7LivuTJJQZptaT9jrcAcpK6AzWCopWR/E1rQ/oRCfiu\n4/PGV2nllNHC8rZm4YB3uftmjaWiwISf/gRSuD5yGu1TcCnYr5w3PhvkVAKHYLdj\ntIUlDTuTFXO/92ZCuW74YqBay+dAIB4ThU0jvUNYWTFV8MHmBgQ9CfG9WW4pjORI\nqalq5zXuKm1t+lrxCFND6cXu9Uk8HtrnSS5IqF0DprdepkyYhA==\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCj/umNjtsAFBql\ngQSw7HgoQKCLviq/zXP0Sr0ZMliizGY/B5pJ6rux42bAF0FcKrqL/KBzko4n3fTg\nNBHcZJUFHiL91ZqzWxjc2rxmSd9YOhPH7x31E+ozjBGKClNOR+pyf0wLrKos3HsV\ngZKFM0lKmb/xIfrLMaaHzqUE2lfP/NIcbqAgLMYgwdO5LmOPezWlH+oKhPk6qxX3\nKgidVAfQQ67pNYD+SuMNgT8/OYszz/bE3lzS7jRrt6/8UKvV8+jS/bzBbsz51FEY\ncCiS3Xd5svDpC0IcTHicvp5yHtePyYzxxCEt/P0T7dIUoVXQ780NAL8Y/CGK6TUR\nHo7osjM9AgMBAAECggEASiB0AxdaaDuuG7conqwUV+V2bBPmENJWIksSFGyMYfHQ\nGZdfJyAh/PNTw2n/kiCCN7pV8EeDWAPcpucCV8NjFHAd0uyVQ5Letx1r4TRs7t05\nibrMqLV6vBgI6YNnSk/5ag2eGvzN4v8552utBeY7r6u1ddItIWFs65/9OSdUX98m\n9iKC9n4D7pFvJsRoUfeD5qf5tF2cmAGS8z+y3502LPx0rNoJchHh06bkEwQIdA7Y\nTUsXIj6RXZqHgcyD0CGVI0gsT6lSywSzfUQvgLEV6Py4VD+1t2jo97bgScbLaFRN\nupC41HJFloYBvin0jtjgo/x8OUTGgW5IkBmrhw5roQKBgQDRoQiWiNtVYh+3AX2g\nDkDXqUbJLdWEElbkjxqzrC19jHCOI12S67MvNZGZ3ET//5agF5pWo/ZE1ZUy7WoJ\n2C9IdbKauGFlA0Drl1xrjHXU22/w8iDFp/F9lhbvu0vGIlCGGZlarl97Ulrha33a\nEHtxNtX7jlu3yrZdReqrUWEhWQKBgQDIRbyaVbVw8nBAKj49cqi8AayD/aAt1sYE\nKuDG0pv1ucXkOYtO2kvTxVUdBIPozdK4nVo/XT+mZ7PYSculukegjkO/Y51F6Tg+\n4jlQxMFv30UAhslacFcSNy5KCePZX/5sUbHYa6Amp5dYXMf1hud02PdIMoNygNZ4\nKKeQhL7ghQKBgF1XtTlCi1fDr5ePpF6muhzNlWVzcUWz3Nk9F4i1vDPRWzUPblVD\nerAkzEaUnGzZZDq5B9JYhAo2iI76xGLJzpQXRIY8X7HY9wlwhoilLLqxU3EYf5tD\novZm5KOu5Ji/ItfzgiOszXteOnVxpcJ54F2TK0kuJIz8SKPTxCCwxe1RAoGAPA3N\nVGpHEitgxZzlNP/g4R+PX7T6B0TT9AP3iyc0ZSbj1F/9ChQjkMknkJ/9/h1aBsoI\ned+4amnGYCEg0/1b5SVD42w3iPM6ToD/ttyJNMa6pkHEtz3gnjG1y7XTgSdr34dP\n0RnU2EKA+5o2y8U8Oqmk3R1olTlVFor6VDe6FRECgYAsgRurQPvNyWTjtXp8QLZ/\njYa8jqXkS7a26JVGKG6gq5hoaFq/fRksNYyC4H88lD/zxRmzkFkrEDIyhjVh5OMN\nPAJTLOnebC6W2xbisRJqDB2LfxmbaH2FA4kO6UzZkhYjudThu4y8uOQlEcVQO7GU\nVPVlMVkJXQmXIl6v3hCN+Q==\n-----END PRIVATE KEY-----",
							AdminKey:  "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQD0aBYmZyhNjsbi\n16ybvxYSBYrIKTdCrZ2g/FeAnfxVY4/85J6KWvvfoWFDqp0nAShlmQb7Fe5GQuN1\nIMTDbpCgYxAVX6Kmk6E6iKcmpko5p5K2vEo2A0/0V9vh2W5mdJjjGOtjDgtYAYR9\naDMdLdyagKfb3AnaOx+TotQuij0nFyYaTKGwK7sh5/y9Pyt0floeB0apndiJbh+e\nlDzG2t8QLBivsuIl0pI510PZSVVcWIBkJ2RVAPyy8gy/wOInsE08o2+0SrIxz2f7\nj+4ZFKfXEUiBopiwy7k5aFeq2gXF2r/Ndpi96uNsP4CXxHkUd+7jcFZQVHPiXNIX\nP6MCLbLXAgMBAAECggEACD7WuHb0eiFd/lsuXJbGxNbhBr21Oo+m6L56qUErOSpB\nulNwMdS9+J52LJU99gno9fyCqsfjoQUyrUnsuXcqc+7DpSTz1NDYOKRRl1E24dkQ\nbw/NJSNZeDHanjT6r4QxgD/f+RiJM2/hq2VvjAV3EtNSVm2G+5DREOcGZ4eMZpwl\ncPg1xrfEOf4Q7mLmiGxBW7tdGzilTsQSwzYifMWzdRmuxNSB8+jIu7Zy53mdUNyc\nI8nDxX8nfLND83GSPlzYuu2S5+cIHumKwtfBcde06Mn72bCLlkHF/AQEWMzGSqv6\nBQAeK3Dr1wI0JZ2Bil6L5Jld3xuRmu++i0DMF8pnwQKBgQD834ChWyoNu1YbXijS\nf53TEfuRZEBgyiOtfgrW723L9uMM975OCBJrhZ4/+WlUGIiZvr2COES3ikxYCzQ6\nkcukAjj9f0JupIOxABT+dwcJIJ4JDdFD8tUwAvY9ys2/WoOs5TdcCxRntrjL0xjG\nZUtDnMI90aOmihY1LB6yEVwh4QKBgQD3bchTW6jC2UdS52Kosdy7YhsgTdiZ+Mqs\nkL52SzBgd0lK4K35nHMUd+C3O90icKEfrpJX1cSbyFfSuimUQdK0I1bl0ijL8+uJ\nFkhOC0iy6um8xO/dvP66j7cuvNXYfOiJAqFJfGiSLe/vuT+JerbFFEidIRXqWSlh\nLIHtvxbbtwKBgBxWL2Plg2DmjU+jzY9JHbZ5XWd9hHlULYtThINxcSxaDjd1y62S\n2f2Si5k/qb3ywdv4s+PTyl+G7+ct2jx1+gv288v0Zs1fQiKjj7a0P+WV8h+xnLGw\nlJM8wbtK7qNy0S6ewQVfeHnmz+6HSU9yKmz5NAsZYu1WrAZpW0c5CsoBAoGBAL1P\naumUhM/obLDaxtqpk2hvjK+vwB02hON5r7BUoQP94L8AnzwPXuF3QyEPFYfHQxA5\nglDgBxjmNYPO2gdMQYmATHl0zbAWxczSlqnX6lyybfn3eEtg0kktsot5Aeks0MIb\nmAngvSWzLhRt2VY35OVvOou2h80RQR7Pbe3YugWLAoGATLlLM3U80bexIiyvwmot\n0gFgd8NIIDF48jTYEpZ9lsfieb7QfkvVVlvm3YsWaDtSd3OswWN6tkyFnz2GJxEB\nj8yVCBhAOxBfDRpMRzOpBe2v8u/o2Pxtw1leHmHbWYqbfwJnupnx3wnrzCVm+oLP\nUHB/aximuVLK+eRz/ZIZ3KU=\n-----END PRIVATE KEY-----",
							AdminCert: "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAPMNo6eG0UBZMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTE5WhcNMjYwNTIxMTEzOTE5WjBkMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEjAQBgNVBAMMCWNoZWZhZG1pbjCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAPRoFiZnKE2OxuLXrJu/FhIFisgpN0KtnaD8V4Cd\n/FVjj/zknopa+9+hYUOqnScBKGWZBvsV7kZC43UgxMNukKBjEBVfoqaToTqIpyam\nSjmnkra8SjYDT/RX2+HZbmZ0mOMY62MOC1gBhH1oMx0t3JqAp9vcCdo7H5Oi1C6K\nPScXJhpMobAruyHn/L0/K3R+Wh4HRqmd2IluH56UPMba3xAsGK+y4iXSkjnXQ9lJ\nVVxYgGQnZFUA/LLyDL/A4iewTTyjb7RKsjHPZ/uP7hkUp9cRSIGimLDLuTloV6ra\nBcXav812mL3q42w/gJfEeRR37uNwVlBUc+Jc0hc/owItstcCAwEAAaM3MDUwHQYD\nVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBQGA1UdEQQNMAuCCWNoZWZhZG1p\nbjANBgkqhkiG9w0BAQsFAAOCAQEAD4uod0aFfmwHuXshXFPcvu5gkkvd2Prd8pMz\nTiOu0A1IYgVHvAg0Bqs8bReZLsijUKhhjmlfwaUrW1i2r5jYp9Acd2K/rKwl7rov\ngaP60IgkszFCjfZFU1Zb0+6OlPuG8PquMPSJuJDFm5mQkw52vb81guvF/cGlKCuI\nfPwEr1CSOyHmQMQLLJsdzfB2RQ6MfG1pCH1FNPCfWjK8Rrd0Hic/sQ4Yd6q/Bmek\nb9KA5lUyk1Ox/YBBfm4RfmeKRqQKtASF8UJG3eCut0h3+uqyQpCxAx/8MlipkF7g\n1EvU3J0+e99VMZJ/wNlrAla08VsV/hjXd+NIZYkl31cghLXmTA==\n-----END CERTIFICATE-----",
						}, {
							IP:        "192.0.2.3",
							Cert:      "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAPMNo6eG0UBaMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTE5WhcNMjYwNTIxMTEzOTE5WjBkMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEjAQBgNVBAMMCWNoZWZub2RlMTCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAKP+6Y2O2wAUGqWBBLDseChAoIu+Kr/Nc/RKvRky\nWKLMZj8Hmknqu7HjZsAXQVwquov8oHOSjifd9OA0EdxklQUeIv3VmrNbGNzavGZJ\n31g6E8fvHfUT6jOMEYoKU05H6nJ/TAusqizcexWBkoUzSUqZv/Eh+ssxpofOpQTa\nV8/80hxuoCAsxiDB07kuY497NaUf6gqE+TqrFfcqCJ1UB9BDruk1gP5K4w2BPz85\nizPP9sTeXNLuNGu3r/xQq9Xz6NL9vMFuzPnUURhwKJLdd3my8OkLQhxMeJy+nnIe\n14/JjPHEIS38/RPt0hShVdDvzQ0Avxj8IYrpNREejuiyMz0CAwEAAaM3MDUwHQYD\nVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBQGA1UdEQQNMAuCCWNoZWZub2Rl\nMTANBgkqhkiG9w0BAQsFAAOCAQEAI9B0IJgmeqgLgFvBiHhGsNzxAKo3z+YmU2Ta\nbqmcBO8gTGnLsnaQPs25sDPvM7YEkcjazhj74+f+L70+rAXl45TLkLQnIGpa4Bbg\nuBpYonPRzK3aSiDcnTeTH7LivuTJJQZptaT9jrcAcpK6AzWCopWR/E1rQ/oRCfiu\n4/PGV2nllNHC8rZm4YB3uftmjaWiwISf/gRSuD5yGu1TcCnYr5w3PhvkVAKHYLdj\ntIUlDTuTFXO/92ZCuW74YqBay+dAIB4ThU0jvUNYWTFV8MHmBgQ9CfG9WW4pjORI\nqalq5zXuKm1t+lrxCFND6cXu9Uk8HtrnSS5IqF0DprdepkyYhA==\n-----END CERTIFICATE-----",
							Key:       "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCj/umNjtsAFBql\ngQSw7HgoQKCLviq/zXP0Sr0ZMliizGY/B5pJ6rux42bAF0FcKrqL/KBzko4n3fTg\nNBHcZJUFHiL91ZqzWxjc2rxmSd9YOhPH7x31E+ozjBGKClNOR+pyf0wLrKos3HsV\ngZKFM0lKmb/xIfrLMaaHzqUE2lfP/NIcbqAgLMYgwdO5LmOPezWlH+oKhPk6qxX3\nKgidVAfQQ67pNYD+SuMNgT8/OYszz/bE3lzS7jRrt6/8UKvV8+jS/bzBbsz51FEY\ncCiS3Xd5svDpC0IcTHicvp5yHtePyYzxxCEt/P0T7dIUoVXQ780NAL8Y/CGK6TUR\nHo7osjM9AgMBAAECggEASiB0AxdaaDuuG7conqwUV+V2bBPmENJWIksSFGyMYfHQ\nGZdfJyAh/PNTw2n/kiCCN7pV8EeDWAPcpucCV8NjFHAd0uyVQ5Letx1r4TRs7t05\nibrMqLV6vBgI6YNnSk/5ag2eGvzN4v8552utBeY7r6u1ddItIWFs65/9OSdUX98m\n9iKC9n4D7pFvJsRoUfeD5qf5tF2cmAGS8z+y3502LPx0rNoJchHh06bkEwQIdA7Y\nTUsXIj6RXZqHgcyD0CGVI0gsT6lSywSzfUQvgLEV6Py4VD+1t2jo97bgScbLaFRN\nupC41HJFloYBvin0jtjgo/x8OUTGgW5IkBmrhw5roQKBgQDRoQiWiNtVYh+3AX2g\nDkDXqUbJLdWEElbkjxqzrC19jHCOI12S67MvNZGZ3ET//5agF5pWo/ZE1ZUy7WoJ\n2C9IdbKauGFlA0Drl1xrjHXU22/w8iDFp/F9lhbvu0vGIlCGGZlarl97Ulrha33a\nEHtxNtX7jlu3yrZdReqrUWEhWQKBgQDIRbyaVbVw8nBAKj49cqi8AayD/aAt1sYE\nKuDG0pv1ucXkOYtO2kvTxVUdBIPozdK4nVo/XT+mZ7PYSculukegjkO/Y51F6Tg+\n4jlQxMFv30UAhslacFcSNy5KCePZX/5sUbHYa6Amp5dYXMf1hud02PdIMoNygNZ4\nKKeQhL7ghQKBgF1XtTlCi1fDr5ePpF6muhzNlWVzcUWz3Nk9F4i1vDPRWzUPblVD\nerAkzEaUnGzZZDq5B9JYhAo2iI76xGLJzpQXRIY8X7HY9wlwhoilLLqxU3EYf5tD\novZm5KOu5Ji/ItfzgiOszXteOnVxpcJ54F2TK0kuJIz8SKPTxCCwxe1RAoGAPA3N\nVGpHEitgxZzlNP/g4R+PX7T6B0TT9AP3iyc0ZSbj1F/9ChQjkMknkJ/9/h1aBsoI\ned+4amnGYCEg0/1b5SVD42w3iPM6ToD/ttyJNMa6pkHEtz3gnjG1y7XTgSdr34dP\n0RnU2EKA+5o2y8U8Oqmk3R1olTlVFor6VDe6FRECgYAsgRurQPvNyWTjtXp8QLZ/\njYa8jqXkS7a26JVGKG6gq5hoaFq/fRksNYyC4H88lD/zxRmzkFkrEDIyhjVh5OMN\nPAJTLOnebC6W2xbisRJqDB2LfxmbaH2FA4kO6UzZkhYjudThu4y8uOQlEcVQO7GU\nVPVlMVkJXQmXIl6v3hCN+Q==\n-----END PRIVATE KEY-----",
							AdminKey:  "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQD0aBYmZyhNjsbi\n16ybvxYSBYrIKTdCrZ2g/FeAnfxVY4/85J6KWvvfoWFDqp0nAShlmQb7Fe5GQuN1\nIMTDbpCgYxAVX6Kmk6E6iKcmpko5p5K2vEo2A0/0V9vh2W5mdJjjGOtjDgtYAYR9\naDMdLdyagKfb3AnaOx+TotQuij0nFyYaTKGwK7sh5/y9Pyt0floeB0apndiJbh+e\nlDzG2t8QLBivsuIl0pI510PZSVVcWIBkJ2RVAPyy8gy/wOInsE08o2+0SrIxz2f7\nj+4ZFKfXEUiBopiwy7k5aFeq2gXF2r/Ndpi96uNsP4CXxHkUd+7jcFZQVHPiXNIX\nP6MCLbLXAgMBAAECggEACD7WuHb0eiFd/lsuXJbGxNbhBr21Oo+m6L56qUErOSpB\nulNwMdS9+J52LJU99gno9fyCqsfjoQUyrUnsuXcqc+7DpSTz1NDYOKRRl1E24dkQ\nbw/NJSNZeDHanjT6r4QxgD/f+RiJM2/hq2VvjAV3EtNSVm2G+5DREOcGZ4eMZpwl\ncPg1xrfEOf4Q7mLmiGxBW7tdGzilTsQSwzYifMWzdRmuxNSB8+jIu7Zy53mdUNyc\nI8nDxX8nfLND83GSPlzYuu2S5+cIHumKwtfBcde06Mn72bCLlkHF/AQEWMzGSqv6\nBQAeK3Dr1wI0JZ2Bil6L5Jld3xuRmu++i0DMF8pnwQKBgQD834ChWyoNu1YbXijS\nf53TEfuRZEBgyiOtfgrW723L9uMM975OCBJrhZ4/+WlUGIiZvr2COES3ikxYCzQ6\nkcukAjj9f0JupIOxABT+dwcJIJ4JDdFD8tUwAvY9ys2/WoOs5TdcCxRntrjL0xjG\nZUtDnMI90aOmihY1LB6yEVwh4QKBgQD3bchTW6jC2UdS52Kosdy7YhsgTdiZ+Mqs\nkL52SzBgd0lK4K35nHMUd+C3O90icKEfrpJX1cSbyFfSuimUQdK0I1bl0ijL8+uJ\nFkhOC0iy6um8xO/dvP66j7cuvNXYfOiJAqFJfGiSLe/vuT+JerbFFEidIRXqWSlh\nLIHtvxbbtwKBgBxWL2Plg2DmjU+jzY9JHbZ5XWd9hHlULYtThINxcSxaDjd1y62S\n2f2Si5k/qb3ywdv4s+PTyl+G7+ct2jx1+gv288v0Zs1fQiKjj7a0P+WV8h+xnLGw\nlJM8wbtK7qNy0S6ewQVfeHnmz+6HSU9yKmz5NAsZYu1WrAZpW0c5CsoBAoGBAL1P\naumUhM/obLDaxtqpk2hvjK+vwB02hON5r7BUoQP94L8AnzwPXuF3QyEPFYfHQxA5\nglDgBxjmNYPO2gdMQYmATHl0zbAWxczSlqnX6lyybfn3eEtg0kktsot5Aeks0MIb\nmAngvSWzLhRt2VY35OVvOou2h80RQR7Pbe3YugWLAoGATLlLM3U80bexIiyvwmot\n0gFgd8NIIDF48jTYEpZ9lsfieb7QfkvVVlvm3YsWaDtSd3OswWN6tkyFnz2GJxEB\nj8yVCBhAOxBfDRpMRzOpBe2v8u/o2Pxtw1leHmHbWYqbfwJnupnx3wnrzCVm+oLP\nUHB/aximuVLK+eRz/ZIZ3KU=\n-----END PRIVATE KEY-----",
							AdminCert: "-----BEGIN CERTIFICATE-----\nMIIDgTCCAmmgAwIBAgIJAPMNo6eG0UBZMA0GCSqGSIb3DQEBCwUAMGMxCzAJBgNV\nBAYTAlVTMRMwEQYDVQQIDApXYXNoaW5ndG9uMRAwDgYDVQQHDAdTZWF0dGxlMRow\nGAYDVQQKDBFDaGVmIFNvZnR3YXJlIEluYzERMA8GA1UEAwwIcHJvZ3Jlc3MwHhcN\nMjMwNTIyMTEzOTE5WhcNMjYwNTIxMTEzOTE5WjBkMQswCQYDVQQGEwJVUzETMBEG\nA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEaMBgGA1UECgwRQ2hl\nZiBTb2Z0d2FyZSBJbmMxEjAQBgNVBAMMCWNoZWZhZG1pbjCCASIwDQYJKoZIhvcN\nAQEBBQADggEPADCCAQoCggEBAPRoFiZnKE2OxuLXrJu/FhIFisgpN0KtnaD8V4Cd\n/FVjj/zknopa+9+hYUOqnScBKGWZBvsV7kZC43UgxMNukKBjEBVfoqaToTqIpyam\nSjmnkra8SjYDT/RX2+HZbmZ0mOMY62MOC1gBhH1oMx0t3JqAp9vcCdo7H5Oi1C6K\nPScXJhpMobAruyHn/L0/K3R+Wh4HRqmd2IluH56UPMba3xAsGK+y4iXSkjnXQ9lJ\nVVxYgGQnZFUA/LLyDL/A4iewTTyjb7RKsjHPZ/uP7hkUp9cRSIGimLDLuTloV6ra\nBcXav812mL3q42w/gJfEeRR37uNwVlBUc+Jc0hc/owItstcCAwEAAaM3MDUwHQYD\nVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMBQGA1UdEQQNMAuCCWNoZWZhZG1p\nbjANBgkqhkiG9w0BAQsFAAOCAQEAD4uod0aFfmwHuXshXFPcvu5gkkvd2Prd8pMz\nTiOu0A1IYgVHvAg0Bqs8bReZLsijUKhhjmlfwaUrW1i2r5jYp9Acd2K/rKwl7rov\ngaP60IgkszFCjfZFU1Zb0+6OlPuG8PquMPSJuJDFm5mQkw52vb81guvF/cGlKCuI\nfPwEr1CSOyHmQMQLLJsdzfB2RQ6MfG1pCH1FNPCfWjK8Rrd0Hic/sQ4Yd6q/Bmek\nb9KA5lUyk1Ox/YBBfm4RfmeKRqQKtASF8UJG3eCut0h3+uqyQpCxAx/8MlipkF7g\n1EvU3J0+e99VMZJ/wNlrAla08VsV/hjXd+NIZYkl31cghLXmTA==\n-----END CERTIFICATE-----",
						},
					},
				},
				ExternalOS: ExternalOS{
					OSDomainName:   "",
					OSDomainURL:    "",
					OSUsername:     "",
					OSUserPassword: "",
					OSCert:         "",
					OSRoleArn:      "",
				},
				ExternalPG: ExternalPG{
					PGInstanceURL:       "",
					PGSuperuserName:     "",
					PGSuperuserPassword: "",
					PGDbUserName:        "",
					PGDbUserPassword:    "",
					PGRootCert:          "",
				},
				DeploymentState: "",
				APIToken:        "",
			},
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			haConfig := &config.HaDeployConfig{}
			err := haConfig.Parse(tt.filePath)
			assert.NoErrorf(t, err, "Error parsing HaDeployConfig: %v", err)
			c := &Config{}
			err = c.PopulateWith(haConfig)
			if tt.wantErr {
				assert.Equal(t, tt.err.Error(), err.Error())
			}

			assert.Equal(t, tt.want, c)
		})
	}
}
