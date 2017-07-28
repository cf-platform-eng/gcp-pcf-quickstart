package ert

import (
	"encoding/json"
	"fmt"
	"omg-cli/config"
	"omg-cli/omg/tiles"
	"omg-cli/ops_manager"
)

type Properties struct {
	// Domains
	AppsDomain Value `json:".cloud_controller.apps_domain"`
	SysDomain  Value `json:".cloud_controller.system_domain"`
	// Networking
	NetworkingPointOfEntry    Value `json:".properties.networking_point_of_entry"`
	TcpRouting                Value `json:".properties.tcp_routing"`
	TcpRoutingReservablePorts Value `json:".properties.tcp_routing.enable.reservable_ports"`
	//SkipSSLVerification       BooleanValue `json:".properties.route_services.enable.ignore_ssl_cert_verification"`
	SkipSSLVerification BooleanValue `json:".ha_proxy.skip_cert_verify"`
	// Application Security Groups
	SecurityAcknowledgement Value `json:".properties.security_acknowledgement"`
	// UAA
	ServiceProviderCredentials CertificateValue `json:".uaa.service_provider_key_credentials"`
	// MySQL
	MySqlMonitorRecipientEmail Value `json:".mysql_monitor.recipient_email"`
}

type Value struct {
	Value string `json:"value"`
}

type BooleanValue struct {
	Value bool `json:"value"`
}

type Certificate struct {
	PublicKey  string `json:"cert_pem"`
	PrivateKey string `json:"private_key_pem"`
}

type CertificateValue struct {
	Value Certificate `json:"value"`
}

type Resources struct {
	TcpRouter                    Resource `json:"tcp_router"`
	Router                       Resource `json:"router"`
	DiegoBrain                   Resource `json:"diego_brain"`
	ConsulServer                 Resource `json:"consul_server"`
	Nats                         Resource `json:"nats"`
	EtcdTlsServer                Resource `json:"etcd_tls_server"`
	NfsServer                    Resource `json:"nfs_server"`
	MysqlProxy                   Resource `json:"mysql_proxy"`
	Mysql                        Resource `json:"mysql"`
	BackupPrepare                Resource `json:"backup-prepare"`
	Ccdb                         Resource `json:"ccdb"`
	DiegoDatabase                Resource `json:"diego_database"`
	Uaadb                        Resource `json:"uaadb"`
	Uaa                          Resource `json:"uaa"`
	CloudController              Resource `json:"cloud_controller"`
	HaProxy                      Resource `json:"ha_proxy"`
	MysqlMonitor                 Resource `json:"mysql_monitor"`
	ClockGlobal                  Resource `json:"clock_global"`
	CloudControllerWorker        Resource `json:"cloud_controller_worker"`
	DiegoCell                    Resource `json:"diego_cell"`
	LoggregatorTrafficcontroller Resource `json:"loggregator_trafficcontroller"`
	SyslogAdapter                Resource `json:"syslog_adapter"`
	SyslogScheduler              Resource `json:"syslog_scheduler"`
	Doppler                      Resource `json:"doppler"`
	SmokeTests                   Resource `json:"smoke-tests"`
	PushAppsManager              Resource `json:"push-apps-manager"`
	Notifications                Resource `json:"notifications"`
	NotificationsUi              Resource `json:"notifications-ui"`
	PushPivotalAccount           Resource `json:"push-pivotal-account"`
	Autoscaling                  Resource `json:"autoscaling"`
	AutoscalingRegisterBroker    Resource `json:"autoscaling-register-broker"`
	Nfsbrokerpush                Resource `json:"nfsbrokerpush"`
	Bootstrap                    Resource `json:"bootstrap"`
	MysqlRejoinUnsafe            Resource `json:"mysql-rejoin-unsafe"`
}

type Resource struct {
	RouterNames       []string `json:"elb_names,omitempty"`
	Instances         int      `json:"instances,omitempty"`
	InternetConnected bool     `json:"internet_connected"`
}

func (Tile) Configure(cfg *config.Config, om *ops_manager.Sdk) error {
	if err := om.StageProduct(tile.Product); err != nil {
		return err
	}

	network := tiles.NetworkConfig(cfg.ErtSubnetName, cfg)

	networkBytes, err := json.Marshal(&network)
	if err != nil {
		return err
	}

	properties := Properties{
		AppsDomain:                 Value{fmt.Sprintf("apps.%s", cfg.DnsSuffix)},
		SysDomain:                  Value{fmt.Sprintf("sys.%s", cfg.DnsSuffix)},
		NetworkingPointOfEntry:     Value{"external_non_ssl"},
		SkipSSLVerification:        BooleanValue{true},
		TcpRouting:                 Value{"enable"},
		TcpRoutingReservablePorts:  Value{cfg.TcpPortRange},
		SecurityAcknowledgement:    Value{"X"},
		ServiceProviderCredentials: CertificateValue{Certificate{cfg.SslCertificate, cfg.SslPrivateKey}},
		MySqlMonitorRecipientEmail: Value{"admin@example.org"},
	}

	propertiesBytes, err := json.Marshal(&properties)
	if err != nil {
		return err
	}

	resoruces := Resources{
		TcpRouter: Resource{
			RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.TcpTargetPoolName)},
			InternetConnected: false,
		},
		Router: Resource{
			RouterNames:       []string{fmt.Sprintf("http:%s", cfg.HttpBackendServiceName)},
			InternetConnected: false,
		},
		DiegoBrain: Resource{
			RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.SshTargetPoolName)},
			InternetConnected: false,
		},
	}
	resorucesBytes, err := json.Marshal(&resoruces)
	if err != nil {
		return err
	}

	return om.ConfigureProduct(tile.Product.Name, string(networkBytes), string(propertiesBytes), string(resorucesBytes))
}