package main

import (
	"github.com/newrelic/infra-integrations-sdk/data/metric"
)

type metricDefinition struct {
	MetricName  string
	SourceType  metric.SourceType
	IsAttribute bool
}

// HAProxyFrontendStats holds the metric definitions for a frontend
var HAProxyFrontendStats = map[string]metricDefinition{
	"pxname":        {MetricName: "proxyName", IsAttribute: true},
	"svname":        {MetricName: "serviceName", IsAttribute: true},
	"scur":          {MetricName: "currentSessions", SourceType: metric.GAUGE},
	"smax":          {MetricName: "maxSessions", SourceType: metric.GAUGE},
	"stot":          {MetricName: "sessionsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"bin":           {MetricName: "bytesInPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"bout":          {MetricName: "bytesOutPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"dreq":          {MetricName: "requestsDenied.securityConcernsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"dresp":         {MetricName: "responsesDenied.securityConcernsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"ereq":          {MetricName: "requestErrorsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"status":        {MetricName: "status", IsAttribute: true},
	"rate_max":      {MetricName: "maxSessionsPerSecond", SourceType: metric.GAUGE},
	"hrsp_1xx":      {MetricName: "http100ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_2xx":      {MetricName: "http200ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_3xx":      {MetricName: "http300ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_4xx":      {MetricName: "http400ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_5xx":      {MetricName: "http500ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_other":    {MetricName: "httpOtherResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"req_rate_max":  {MetricName: "httpRequests.maxPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"req_tot":       {MetricName: "httpRequestsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"mode":          {MetricName: "mode", IsAttribute: true},
	"conn_rate_max": {MetricName: "maxConnectionsPerSecond", SourceType: metric.GAUGE},
	"conn_tot":      {MetricName: "connectionsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"intercepted":   {MetricName: "interceptedRequestsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"dcon":          {MetricName: "requestsDenied.tcpRequestConnectionRulesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"dses":          {MetricName: "requestsDenied.tcpRequestSessionRulesPerSecond", SourceType: metric.CUMULATIVE_RATE},
}

// HAProxyBackendStats holds the metric definitions for a backend
var HAProxyBackendStats = map[string]metricDefinition{
	"pxname":      {MetricName: "proxyName", IsAttribute: true},
	"qcur":        {MetricName: "currentQueuedRequestsWithoutServer", SourceType: metric.GAUGE},
	"qmax":        {MetricName: "maxQueuedRequestsWithoutServer", SourceType: metric.GAUGE},
	"scur":        {MetricName: "currentSessions", SourceType: metric.GAUGE},
	"smax":        {MetricName: "maxSessions", SourceType: metric.GAUGE},
	"stot":        {MetricName: "sessionsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"bin":         {MetricName: "bytesInPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"bout":        {MetricName: "bytesOutPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"dreq":        {MetricName: "requestsDenied.securityConcernsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"dresp":       {MetricName: "responsesDenied.securityConcernsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"econ":        {MetricName: "connectingRequestErrorsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"eresp":       {MetricName: "responseErrorsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"wretr":       {MetricName: "connectionRetriesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"wredis":      {MetricName: "requestRedispatchPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"status":      {MetricName: "status", IsAttribute: true},
	"weight":      {MetricName: "totalWeight", SourceType: metric.GAUGE},
	"act":         {MetricName: "activeServers", SourceType: metric.GAUGE},
	"bck":         {MetricName: "backupServers", SourceType: metric.GAUGE},
	"chkdown":     {MetricName: "upToDownTransitionsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"lastchg":     {MetricName: "timeSinceLastUpDownTransitionInSeconds", SourceType: metric.GAUGE},
	"downtime":    {MetricName: "downtimeInSeconds", SourceType: metric.GAUGE},
	"lbtot":       {MetricName: "serverSelectedPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"rate_max":    {MetricName: "maxSessionsPerSecond", SourceType: metric.GAUGE},
	"hrsp_1xx":    {MetricName: "http100ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_2xx":    {MetricName: "http200ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_3xx":    {MetricName: "http300ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_4xx":    {MetricName: "http400ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_5xx":    {MetricName: "http500ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_other":  {MetricName: "httpOtherResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"req_tot":     {MetricName: "httpRequestsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"cli_abrt":    {MetricName: "dataTransfersAbortedByClientPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"srv_abrt":    {MetricName: "dataTransfersAbortedByServerPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"comp_in":     {MetricName: "httpResponseBytesFedToCompressorPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"comp_out":    {MetricName: "httpResponseBytesEmittedByCompressorPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"comp_byp":    {MetricName: "bytesThatBypassedCompressorPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"comp_rsp":    {MetricName: "httpResponsesCompressedPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"lastsess":    {MetricName: "timeSinceLastSessionAssignedInSeconds", SourceType: metric.GAUGE},
	"qtime":       {MetricName: "averageQueueTimeInSeconds", SourceType: metric.GAUGE},
	"ctime":       {MetricName: "averageConnectTimeInSeconds", SourceType: metric.GAUGE},
	"rtime":       {MetricName: "averageResponseTimeInSeconds", SourceType: metric.GAUGE},
	"ttime":       {MetricName: "averageTotalSessionTimeInSeconds", SourceType: metric.GAUGE},
	"cookie":      {MetricName: "cookieName", IsAttribute: true},
	"mode":        {MetricName: "mode", IsAttribute: true},
	"intercepted": {MetricName: "interceptedRequestsPerSecond", SourceType: metric.CUMULATIVE_RATE},
}

// HAProxyServerStats holds the metric definitions for a server
var HAProxyServerStats = map[string]metricDefinition{
	"pxname":         {MetricName: "proxyName", IsAttribute: true},
	"svname":         {MetricName: "serviceName", IsAttribute: true},
	"qcur":           {MetricName: "currentQueuedRequestsWithoutServer", SourceType: metric.GAUGE},
	"qmax":           {MetricName: "maxQueuedRequestsWithoutServer", SourceType: metric.GAUGE},
	"scur":           {MetricName: "currentSessions", SourceType: metric.GAUGE},
	"smax":           {MetricName: "maxSessions", SourceType: metric.GAUGE},
	"stot":           {MetricName: "sessionsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"bin":            {MetricName: "bytesInPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"bout":           {MetricName: "bytesOutPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"dreq":           {MetricName: "requestsDenied.securityConcernsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"dresp":          {MetricName: "responsesDenied.securityConcernsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"econ":           {MetricName: "connectingRequestErrorsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"eresp":          {MetricName: "responseErrorsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"wretr":          {MetricName: "connectionRetriesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"wredis":         {MetricName: "requestRedispatchPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"status":         {MetricName: "status", IsAttribute: true},
	"weight":         {MetricName: "serverWeight", SourceType: metric.GAUGE},
	"act":            {MetricName: "isActive", SourceType: metric.GAUGE},
	"bck":            {MetricName: "isBackup", SourceType: metric.GAUGE},
	"chkfail":        {MetricName: "failedChecksPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"chkdown":        {MetricName: "upToDownTransitionsPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"lastchg":        {MetricName: "timeSinceLastUpDownTransitionInSeconds", SourceType: metric.GAUGE},
	"downtime":       {MetricName: "downtimeInSeconds", SourceType: metric.GAUGE},
	"sid":            {MetricName: "serverID", IsAttribute: true},
	"throttle":       {MetricName: "throttlePercentage", SourceType: metric.GAUGE},
	"lbtot":          {MetricName: "serverSelectedPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"rate_max":       {MetricName: "maxSessionsPerSecond", SourceType: metric.GAUGE},
	"check_status":   {MetricName: "healthCheckStatus", IsAttribute: true},
	"check_code":     {MetricName: "layerCode", IsAttribute: true},
	"check_duration": {MetricName: "healthCheckDurationInMilliseconds", SourceType: metric.GAUGE},
	"hrsp_1xx":       {MetricName: "http100ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_2xx":       {MetricName: "http200ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_3xx":       {MetricName: "http300ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_4xx":       {MetricName: "http400ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_5xx":       {MetricName: "http500ResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hrsp_other":     {MetricName: "httpOtherResponsesPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"hanafail":       {MetricName: "failedHealthCheckDetails", IsAttribute: true},
	"cli_abrt":       {MetricName: "dataTransfersAbortedByClientPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"srv_abrt":       {MetricName: "dataTransfersAbortedByServerPerSecond", SourceType: metric.CUMULATIVE_RATE},
	"lastsess":       {MetricName: "timeSinceLastSessionAssignedInSeconds", SourceType: metric.GAUGE},
	"last_chk":       {MetricName: "healthCheckContents", IsAttribute: true},
	"last_agt":       {MetricName: "agentCheckContents", IsAttribute: true},
	"qtime":          {MetricName: "averageQueueTimeInSeconds", SourceType: metric.GAUGE},
	"ctime":          {MetricName: "averageConnectTimeInSeconds", SourceType: metric.GAUGE},
	"rtime":          {MetricName: "averageResponseTimeInSeconds", SourceType: metric.GAUGE},
	"ttime":          {MetricName: "averageTotalSessionTimeInSeconds", SourceType: metric.GAUGE},
	"agent_status":   {MetricName: "agentStatus", IsAttribute: true},
	"agent_duration": {MetricName: "agentDurationSeconds", SourceType: metric.GAUGE},
	"check_desc":     {MetricName: "checkStatusDescription", IsAttribute: true},
	"agent_desc":     {MetricName: "agentStatusDescription", IsAttribute: true},
	"cookie":         {MetricName: "cookieValue", IsAttribute: true},
	"mode":           {MetricName: "mode", IsAttribute: true},
}

// HAProxyInventory holds the list of inventory items to be collected from the stats request
var HAProxyInventory = map[string]struct{}{
	"slim":     {},
	"pid":      {},
	"sid":      {},
	"iid":      {},
	"rate_lim": {},
	"qmax":     {},
}
