/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs. 
 *
 * API version: 3.0
 * Contact: support@haproxy.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// HAProxy default bind configuration
type DefaultBind struct {
	AcceptNetscalerCip int32 `json:"accept_netscaler_cip,omitempty"`
	AcceptProxy bool `json:"accept_proxy,omitempty"`
	Allow0rtt bool `json:"allow_0rtt,omitempty"`
	Alpn string `json:"alpn,omitempty"`
	Backlog string `json:"backlog,omitempty"`
	CaIgnoreErr string `json:"ca_ignore_err,omitempty"`
	CaSignFile string `json:"ca_sign_file,omitempty"`
	CaSignPass string `json:"ca_sign_pass,omitempty"`
	CaVerifyFile string `json:"ca_verify_file,omitempty"`
	Ciphers string `json:"ciphers,omitempty"`
	Ciphersuites string `json:"ciphersuites,omitempty"`
	ClientSigalgs string `json:"client_sigalgs,omitempty"`
	CrlFile string `json:"crl_file,omitempty"`
	CrtIgnoreErr string `json:"crt_ignore_err,omitempty"`
	CrtList string `json:"crt_list,omitempty"`
	Curves string `json:"curves,omitempty"`
	DefaultCrtList []string `json:"default_crt_list,omitempty"`
	DeferAccept bool `json:"defer_accept,omitempty"`
	Ecdhe string `json:"ecdhe,omitempty"`
	ExposeFdListeners bool `json:"expose_fd_listeners,omitempty"`
	ForceSslv3 bool `json:"force_sslv3,omitempty"`
	ForceTlsv10 bool `json:"force_tlsv10,omitempty"`
	ForceTlsv11 bool `json:"force_tlsv11,omitempty"`
	ForceTlsv12 bool `json:"force_tlsv12,omitempty"`
	ForceTlsv13 bool `json:"force_tlsv13,omitempty"`
	GenerateCertificates bool `json:"generate_certificates,omitempty"`
	Gid int32 `json:"gid,omitempty"`
	Group string `json:"group,omitempty"`
	GuidPrefix string `json:"guid_prefix,omitempty"`
	Id string `json:"id,omitempty"`
	Interface_ string `json:"interface,omitempty"`
	Level string `json:"level,omitempty"`
	Maxconn int32 `json:"maxconn,omitempty"`
	Mode string `json:"mode,omitempty"`
	Mss string `json:"mss,omitempty"`
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Nbconn int32 `json:"nbconn,omitempty"`
	Nice int32 `json:"nice,omitempty"`
	NoAlpn bool `json:"no_alpn,omitempty"`
	NoCaNames bool `json:"no_ca_names,omitempty"`
	NoSslv3 bool `json:"no_sslv3,omitempty"`
	NoTlsTickets bool `json:"no_tls_tickets,omitempty"`
	NoTlsv10 bool `json:"no_tlsv10,omitempty"`
	NoTlsv11 bool `json:"no_tlsv11,omitempty"`
	NoTlsv12 bool `json:"no_tlsv12,omitempty"`
	NoTlsv13 bool `json:"no_tlsv13,omitempty"`
	Npn string `json:"npn,omitempty"`
	PreferClientCiphers bool `json:"prefer_client_ciphers,omitempty"`
	Proto string `json:"proto,omitempty"`
	QuicCcAlgo string `json:"quic-cc-algo,omitempty"`
	QuicForceRetry bool `json:"quic-force-retry,omitempty"`
	QuicSocket string `json:"quic-socket,omitempty"`
	SeverityOutput string `json:"severity_output,omitempty"`
	Sigalgs string `json:"sigalgs,omitempty"`
	Ssl bool `json:"ssl,omitempty"`
	SslCafile string `json:"ssl_cafile,omitempty"`
	SslCertificate string `json:"ssl_certificate,omitempty"`
	SslMaxVer string `json:"ssl_max_ver,omitempty"`
	SslMinVer string `json:"ssl_min_ver,omitempty"`
	StrictSni bool `json:"strict_sni,omitempty"`
	TcpUserTimeout int32 `json:"tcp_user_timeout,omitempty"`
	Tfo bool `json:"tfo,omitempty"`
	Thread string `json:"thread,omitempty"`
	TlsTicketKeys string `json:"tls_ticket_keys,omitempty"`
	Transparent bool `json:"transparent,omitempty"`
	Uid string `json:"uid,omitempty"`
	User string `json:"user,omitempty"`
	V4v6 bool `json:"v4v6,omitempty"`
	V6only bool `json:"v6only,omitempty"`
	Verify string `json:"verify,omitempty"`
}
