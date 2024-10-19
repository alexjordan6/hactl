/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.
 *
 * API version: 3.0
 * Contact: support@haproxy.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package dataplaneclient

// One entry in stick table
type StickTableEntry struct {
	BytesInCnt   int32  `json:"bytes_in_cnt,omitempty"`
	BytesInRate  int32  `json:"bytes_in_rate,omitempty"`
	BytesOutCnt  int32  `json:"bytes_out_cnt,omitempty"`
	BytesOutRate int32  `json:"bytes_out_rate,omitempty"`
	ConnCnt      int32  `json:"conn_cnt,omitempty"`
	ConnCur      int32  `json:"conn_cur,omitempty"`
	ConnRate     int32  `json:"conn_rate,omitempty"`
	Exp          int32  `json:"exp,omitempty"`
	Gpc0         int32  `json:"gpc0,omitempty"`
	Gpc0Rate     int32  `json:"gpc0_rate,omitempty"`
	Gpc1         int32  `json:"gpc1,omitempty"`
	Gpc1Rate     int32  `json:"gpc1_rate,omitempty"`
	Gpt0         int32  `json:"gpt0,omitempty"`
	HttpErrCnt   int32  `json:"http_err_cnt,omitempty"`
	HttpErrRate  int32  `json:"http_err_rate,omitempty"`
	HttpReqCnt   int32  `json:"http_req_cnt,omitempty"`
	HttpReqRate  int32  `json:"http_req_rate,omitempty"`
	Id           string `json:"id,omitempty"`
	Key          string `json:"key,omitempty"`
	ServerId     int32  `json:"server_id,omitempty"`
	SessCnt      int32  `json:"sess_cnt,omitempty"`
	SessRate     int32  `json:"sess_rate,omitempty"`
	Use          bool   `json:"use,omitempty"`
}
