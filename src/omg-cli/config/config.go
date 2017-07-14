package config

type Config struct {
	OpsManagerIp            string
	NetworkName             string
	DeploymentTargetTag     string
	MgmtSubnetName          string
	MgmtSubnetGateway       string
	MgmtSubnetCIDR          string
	ServicesSubnetName      string
	ServicesSubnetGateway   string
	ServicesSubnetCIDR      string
	ErtSubnetName           string
	ErtSubnetGateway        string
	ErtSubnetCIDR           string
	HttpBackendServiceName  string
	SshTargetPoolName       string
	TcpTargetPoolName       string
	TcpPortRange            string
	WebSocketTargetPoolName string
	BuildpacksBucket        string
	DropletsBucket          string
	PackagesBucket          string
	ResourcesBucket         string
	DirectorBucket          string
	RootDomain              string
	SslCertificate          string
	SslPrivateKey           string
	OpsManServiceAccount    string

	Region      string
	Zone1       string
	Zone2       string
	Zone3       string
	ProjectName string

	OpsManUsername         string
	OpsManPassword         string
	OpsManDecryptionPhrase string
}

func fillInDefaults(cfg *Config) {
	cfg.OpsManUsername = "foo"
	cfg.OpsManPassword = "foobar"
	cfg.OpsManDecryptionPhrase = "foobar"
}