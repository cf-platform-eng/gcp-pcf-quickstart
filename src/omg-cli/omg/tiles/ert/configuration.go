/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ert

import (
	"encoding/json"
	"fmt"
	"omg-cli/config"
	"omg-cli/omg/tiles"
	"omg-cli/ops_manager"

	"github.com/imdario/mergo"
)

type Properties struct {
	// Domains
	AppsDomain                tiles.Value              `json:".cloud_controller.apps_domain"`
	SysDomain                 tiles.Value              `json:".cloud_controller.system_domain"`
	TcpRouting                tiles.Value              `json:".properties.tcp_routing"`
	TcpRoutingReservablePorts tiles.Value              `json:".properties.tcp_routing.enable.reservable_ports"`
	GoRouterSSLCiphers        tiles.Value              `json:".properties.gorouter_ssl_ciphers"`
	HAProxySSLCiphers         tiles.Value              `json:".properties.haproxy_ssl_ciphers"`
	SkipSSLVerification       tiles.BooleanValue       `json:".ha_proxy.skip_cert_verify"`
	HAProxyForwardTLS         tiles.Value              `json:".properties.haproxy_forward_tls"`
	IngressCertificates       tiles.CertificateValue   `json:".properties.networking_poe_ssl_certs"`
	CredhubEncryptionKey      tiles.EncryptionKeyValue `json:".properties.credhub_key_encryption_passwords"`
	// Application Security Groups
	SecurityAcknowledgement tiles.Value `json:".properties.security_acknowledgement"`
	// UAA
	ServiceProviderCredentials tiles.OldCertificateValue `json:".uaa.service_provider_key_credentials"`

	UaaDbChoice   *tiles.Value        `json:".properties.uaa_database,omitempty"`
	UaaDbIp       *tiles.Value        `json:".properties.uaa_database.external.host,omitempty"`
	UaaDbPort     *tiles.IntegerValue `json:".properties.uaa_database.external.port,omitempty"`
	UaaDbUsername *tiles.Value        `json:".properties.uaa_database.external.uaa_username,omitempty"`
	UaaDbPassword *tiles.SecretValue  `json:".properties.uaa_database.external.uaa_password,omitempty"`

	// Databases
	ErtDbChoice tiles.Value         `json:".properties.system_database"`
	ErtDbIp     *tiles.Value        `json:".properties.system_database.external.host,omitempty"`
	ErtDbPort   *tiles.IntegerValue `json:".properties.system_database.external.port,omitempty"`

	ErtDbAppUsageUsername            *tiles.Value       `json:".properties.system_database.external.app_usage_service_username,omitempty"`
	ErtDbAppUsagePassword            *tiles.SecretValue `json:".properties.system_database.external.app_usage_service_password,omitempty"`
	ErtDbAutoscaleUsername           *tiles.Value       `json:".properties.system_database.external.autoscale_username,omitempty"`
	ErtDbAutoscalePassword           *tiles.SecretValue `json:".properties.system_database.external.autoscale_password,omitempty"`
	ErtDbCloudControllerUsername     *tiles.Value       `json:".properties.system_database.external.ccdb_username,omitempty"`
	ErtDbCloudControllerPassword     *tiles.SecretValue `json:".properties.system_database.external.ccdb_password,omitempty"`
	ErtDbDiegoUsername               *tiles.Value       `json:".properties.system_database.external.diego_username,omitempty"`
	ErtDbDiegoPassword               *tiles.SecretValue `json:".properties.system_database.external.diego_password,omitempty"`
	ErtDbLocketUsername              *tiles.Value       `json:".properties.system_database.external.locket_username,omitempty"`
	ErtDbLocketPassword              *tiles.SecretValue `json:".properties.system_database.external.locket_password,omitempty"`
	ErtDbNetworkPolicyServerUsername *tiles.Value       `json:".properties.system_database.external.networkpolicyserver_username,omitempty"`
	ErtDbNetworkPolicyServerPassword *tiles.SecretValue `json:".properties.system_database.external.networkpolicyserver_password,omitempty"`
	ErtDbNfsUsername                 *tiles.Value       `json:".properties.system_database.external.nfsvolume_username,omitempty"`
	ErtDbNfsPassword                 *tiles.SecretValue `json:".properties.system_database.external.nfsvolume_password,omitempty"`
	ErtDbNotificationsUsername       *tiles.Value       `json:".properties.system_database.external.notifications_username,omitempty"`
	ErtDbNotificationsPassword       *tiles.SecretValue `json:".properties.system_database.external.notifications_password,omitempty"`
	ErtDbAccountUsername             *tiles.Value       `json:".properties.system_database.external.account_username,omitempty"`
	ErtDbAccountPassword             *tiles.SecretValue `json:".properties.system_database.external.account_password,omitempty"`
	ErtDbRoutingUsername             *tiles.Value       `json:".properties.system_database.external.routing_username,omitempty"`
	ErtDbRoutingPassword             *tiles.SecretValue `json:".properties.system_database.external.routing_password,omitempty"`
	ErtDbSilkUsername                *tiles.Value       `json:".properties.system_database.external.silk_username,omitempty"`
	ErtDbSilkPassword                *tiles.SecretValue `json:".properties.system_database.external.silk_password,omitempty"`

	// MySQL
	MySqlMonitorRecipientEmail tiles.Value `json:".mysql_monitor.recipient_email"`
}

type LargeFootprintResources struct {
	TcpRouter                    tiles.Resource `json:"tcp_router"`
	Router                       tiles.Resource `json:"router"`
	DiegoBrain                   tiles.Resource `json:"diego_brain"`
	ConsulServer                 tiles.Resource `json:"consul_server"`
	Nats                         tiles.Resource `json:"nats"`
	NfsServer                    tiles.Resource `json:"nfs_server"`
	MysqlProxy                   tiles.Resource `json:"mysql_proxy"`
	Mysql                        tiles.Resource `json:"mysql"`
	BackupPrepare                tiles.Resource `json:"backup_restore"`
	DiegoDatabase                tiles.Resource `json:"diego_database"`
	Uaa                          tiles.Resource `json:"uaa"`
	CloudController              tiles.Resource `json:"cloud_controller"`
	HaProxy                      tiles.Resource `json:"ha_proxy"`
	MysqlMonitor                 tiles.Resource `json:"mysql_monitor"`
	ClockGlobal                  tiles.Resource `json:"clock_global"`
	CloudControllerWorker        tiles.Resource `json:"cloud_controller_worker"`
	DiegoCell                    tiles.Resource `json:"diego_cell"`
	LoggregatorTrafficcontroller tiles.Resource `json:"loggregator_trafficcontroller"`
	SyslogAdapter                tiles.Resource `json:"syslog_adapter"`
	SyslogScheduler              tiles.Resource `json:"syslog_scheduler"`
	Doppler                      tiles.Resource `json:"doppler"`
}

type SmallFootprintResources struct {
	TcpRouter tiles.Resource `json:"tcp_router"`
	Router    tiles.Resource `json:"router"`

	Database    tiles.Resource `json:"database"`
	Control     tiles.Resource `json:"control"`
	Compute     tiles.Resource `json:"compute"`
	FileStorage tiles.Resource `json:"blobstore"`

	HaProxy       tiles.Resource `json:"ha_proxy"`
	BackupPrepare tiles.Resource `json:"backup_restore"`
	MysqlMonitor  tiles.Resource `json:"mysql_monitor"`
}

func (*Tile) Configure(envConfig *config.EnvConfig, cfg *config.Config, om *ops_manager.Sdk) error {
	if err := om.StageProduct(product); err != nil {
		return err
	}

	network := tiles.NetworkConfig(cfg.ErtSubnetName, cfg)

	networkBytes, err := json.Marshal(&network)
	if err != nil {
		return err
	}

	properties := Properties{
		AppsDomain:          tiles.Value{Value: cfg.AppsDomain},
		SysDomain:           tiles.Value{Value: cfg.SysDomain},
		SkipSSLVerification: tiles.BooleanValue{Value: true},
		HAProxyForwardTLS:   tiles.Value{Value: "disable"},
		IngressCertificates: tiles.CertificateValue{Value: []tiles.CertificateConstruct{
			{Certificate: tiles.Certificate{PublicKey: cfg.SslCertificate, PrivateKey: cfg.SslPrivateKey},
				Name: "Certificate",
			},
		},
		},
		CredhubEncryptionKey: tiles.EncryptionKeyValue{Value: []tiles.EncryptionKey{
			{
				Name:    cfg.CredhubKey.Name,
				Key:     tiles.KeyStruct{Secret: cfg.CredhubKey.Key},
				Primary: true,
			},
		},
		},
		TcpRouting:                 tiles.Value{Value: "enable"},
		TcpRoutingReservablePorts:  tiles.Value{Value: cfg.TcpPortRange},
		GoRouterSSLCiphers:         tiles.Value{Value: "ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384"},
		HAProxySSLCiphers:          tiles.Value{Value: "DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384"},
		SecurityAcknowledgement:    tiles.Value{Value: "X"},
		ServiceProviderCredentials: tiles.OldCertificateValue{Value: tiles.Certificate{PublicKey: cfg.SslCertificate, PrivateKey: cfg.SslPrivateKey}},

		MySqlMonitorRecipientEmail: tiles.Value{Value: "admin@example.org"},
	}

	if envConfig.SmallFootprint {
		mergo.Merge(&properties, Properties{
			ErtDbChoice: tiles.Value{Value: "internal_pxc"},
		})
	} else {
		mergo.Merge(&properties, Properties{
			UaaDbChoice:   &tiles.Value{Value: "external"},
			UaaDbIp:       &tiles.Value{Value: cfg.ExternalSqlIp},
			UaaDbPort:     &tiles.IntegerValue{Value: cfg.ExternalSqlPort},
			UaaDbUsername: &tiles.Value{Value: cfg.ERTSqlUsername},
			UaaDbPassword: &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},

			ErtDbChoice:                      tiles.Value{Value: "external"},
			ErtDbIp:                          &tiles.Value{Value: cfg.ExternalSqlIp},
			ErtDbPort:                        &tiles.IntegerValue{Value: cfg.ExternalSqlPort},
			ErtDbAppUsageUsername:            &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbAppUsagePassword:            &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbAutoscaleUsername:           &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbAutoscalePassword:           &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbCloudControllerUsername:     &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbCloudControllerPassword:     &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbDiegoUsername:               &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbDiegoPassword:               &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbLocketUsername:              &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbLocketPassword:              &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbNetworkPolicyServerUsername: &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbNetworkPolicyServerPassword: &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbNfsUsername:                 &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbNfsPassword:                 &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbNotificationsUsername:       &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbNotificationsPassword:       &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbAccountUsername:             &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbAccountPassword:             &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbRoutingUsername:             &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbRoutingPassword:             &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
			ErtDbSilkUsername:                &tiles.Value{Value: cfg.ERTSqlUsername},
			ErtDbSilkPassword:                &tiles.SecretValue{Sec: tiles.Secret{Value: cfg.ERTSqlPassword}},
		})
	}

	propertiesBytes, err := json.Marshal(&properties)
	if err != nil {
		return err
	}

	resourcesBytes := []byte{}

	zero := 0
	one := 1
	three := 3
	if envConfig.SmallFootprint {
		resources := SmallFootprintResources{
			TcpRouter: tiles.Resource{
				RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.TcpTargetPoolName)},
				InternetConnected: false,
				Instances:         &one,
			},
			Router: tiles.Resource{
				RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.WssTargetPoolName), fmt.Sprintf("http:%s", cfg.HttpBackendServiceName)},
				InternetConnected: false,
			},
			Control: tiles.Resource{
				RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.SshTargetPoolName)},
				InternetConnected: false,
			},
			HaProxy:      tiles.Resource{Instances: &zero},
			MysqlMonitor: tiles.Resource{Instances: &zero},
		}
		// Healthwatch pushes quite a few apps, make sure we have enough compute
		if envConfig.IncludeHealthwatch {
			resources.Compute = tiles.Resource{
				Instances: &three,
			}
		}
		resourcesBytes, err = json.Marshal(&resources)
	} else {
		resources := LargeFootprintResources{
			TcpRouter: tiles.Resource{
				RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.TcpTargetPoolName)},
				InternetConnected: false,
				Instances:         &three,
			},
			Router: tiles.Resource{
				RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.WssTargetPoolName), fmt.Sprintf("http:%s", cfg.HttpBackendServiceName)},
				InternetConnected: false,
			},
			DiegoBrain: tiles.Resource{
				RouterNames:       []string{fmt.Sprintf("tcp:%s", cfg.SshTargetPoolName)},
				InternetConnected: false,
			},
			HaProxy:      tiles.Resource{Instances: &zero},
			MysqlProxy:   tiles.Resource{Instances: &zero},
			Mysql:        tiles.Resource{Instances: &zero},
			MysqlMonitor: tiles.Resource{Instances: &zero},
		}
		resourcesBytes, err = json.Marshal(&resources)
	}

	if err != nil {
		return err
	}
	return om.ConfigureProduct(product.Name, string(networkBytes), string(propertiesBytes), string(resourcesBytes))
}
