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

type NativeStatStats struct {
	Act int32 `json:"act,omitempty"`
	Addr string `json:"addr,omitempty"`
	AgentCode int32 `json:"agent_code,omitempty"`
	AgentDesc string `json:"agent_desc,omitempty"`
	AgentDuration int32 `json:"agent_duration,omitempty"`
	AgentFall int32 `json:"agent_fall,omitempty"`
	AgentHealth int32 `json:"agent_health,omitempty"`
	AgentRise int32 `json:"agent_rise,omitempty"`
	AgentStatus string `json:"agent_status,omitempty"`
	Algo string `json:"algo,omitempty"`
	Bck int32 `json:"bck,omitempty"`
	Bin int32 `json:"bin,omitempty"`
	Bout int32 `json:"bout,omitempty"`
	CheckCode int32 `json:"check_code,omitempty"`
	CheckDesc string `json:"check_desc,omitempty"`
	CheckDuration int32 `json:"check_duration,omitempty"`
	CheckFall int32 `json:"check_fall,omitempty"`
	CheckHealth int32 `json:"check_health,omitempty"`
	CheckRise int32 `json:"check_rise,omitempty"`
	CheckStatus string `json:"check_status,omitempty"`
	Chkdown int32 `json:"chkdown,omitempty"`
	Chkfail int32 `json:"chkfail,omitempty"`
	CliAbrt int32 `json:"cli_abrt,omitempty"`
	CompByp int32 `json:"comp_byp,omitempty"`
	CompIn int32 `json:"comp_in,omitempty"`
	CompOut int32 `json:"comp_out,omitempty"`
	CompRsp int32 `json:"comp_rsp,omitempty"`
	ConnRate int32 `json:"conn_rate,omitempty"`
	ConnRateMax int32 `json:"conn_rate_max,omitempty"`
	ConnTot int32 `json:"conn_tot,omitempty"`
	Cookie string `json:"cookie,omitempty"`
	Ctime int32 `json:"ctime,omitempty"`
	Dcon int32 `json:"dcon,omitempty"`
	Downtime int32 `json:"downtime,omitempty"`
	Dreq int32 `json:"dreq,omitempty"`
	Dresp int32 `json:"dresp,omitempty"`
	Dses int32 `json:"dses,omitempty"`
	Econ int32 `json:"econ,omitempty"`
	Ereq int32 `json:"ereq,omitempty"`
	Eresp int32 `json:"eresp,omitempty"`
	Hanafail string `json:"hanafail,omitempty"`
	Hrsp1xx int32 `json:"hrsp_1xx,omitempty"`
	Hrsp2xx int32 `json:"hrsp_2xx,omitempty"`
	Hrsp3xx int32 `json:"hrsp_3xx,omitempty"`
	Hrsp4xx int32 `json:"hrsp_4xx,omitempty"`
	Hrsp5xx int32 `json:"hrsp_5xx,omitempty"`
	HrspOther int32 `json:"hrsp_other,omitempty"`
	Iid int32 `json:"iid,omitempty"`
	Intercepted int32 `json:"intercepted,omitempty"`
	LastAgt string `json:"last_agt,omitempty"`
	LastChk string `json:"last_chk,omitempty"`
	Lastchg int32 `json:"lastchg,omitempty"`
	Lastsess int32 `json:"lastsess,omitempty"`
	Lbtot int32 `json:"lbtot,omitempty"`
	Mode string `json:"mode,omitempty"`
	Pid int32 `json:"pid,omitempty"`
	Qcur int32 `json:"qcur,omitempty"`
	Qlimit int32 `json:"qlimit,omitempty"`
	Qmax int32 `json:"qmax,omitempty"`
	Qtime int32 `json:"qtime,omitempty"`
	Rate int32 `json:"rate,omitempty"`
	RateLim int32 `json:"rate_lim,omitempty"`
	RateMax int32 `json:"rate_max,omitempty"`
	ReqRate int32 `json:"req_rate,omitempty"`
	ReqRateMax int32 `json:"req_rate_max,omitempty"`
	ReqTot int32 `json:"req_tot,omitempty"`
	Rtime int32 `json:"rtime,omitempty"`
	Scur int32 `json:"scur,omitempty"`
	Sid int32 `json:"sid,omitempty"`
	Slim int32 `json:"slim,omitempty"`
	Smax int32 `json:"smax,omitempty"`
	SrvAbrt int32 `json:"srv_abrt,omitempty"`
	Status string `json:"status,omitempty"`
	Stot int32 `json:"stot,omitempty"`
	Throttle int32 `json:"throttle,omitempty"`
	Tracked string `json:"tracked,omitempty"`
	Ttime int32 `json:"ttime,omitempty"`
	Weight int32 `json:"weight,omitempty"`
	Wredis int32 `json:"wredis,omitempty"`
	Wretr int32 `json:"wretr,omitempty"`
}
