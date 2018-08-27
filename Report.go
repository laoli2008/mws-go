package main

type Report struct {
	Client
}

/**
 * Get Report
 *
 * The GetReport operation returns the contents of a report. Reports can potentially be
 * very large (>100MB) which is why we only return one report at a time, and in a
 * streaming fashion.
 *
 * @see GetReportRequest
 * @return GetReportResponse
 *
 * @throws ReportsException
 */
func (api Report) GetReport() {
}

/**
 * Get Report Schedule Count
 *
 * returns the number of report schedules
 *
 * @see GetReportScheduleCountRequest
 * @return GetReportScheduleCountResponse
 */
func (api Report) GetReportScheduleCount() {
}

/**
 * Get Report Request List By Next Token
 *
 * retrieve the next batch of list items and if there are more items to retrieve
 *
 * @see GetReportRequestListByNextTokenRequest
 * @return GetReportRequestListByNextTokenResponse
 */
func (api Report) GetReportRequestListByNextToken() {
}

/**
 * Update Report Acknowledgements
 *
 * The UpdateReportAcknowledgements operation updates the acknowledged status of one or more reports.
 *
 * @see UpdateReportAcknowledgementsRequest
 * @return UpdateReportAcknowledgementsResponse
 */
func (api Report) UpdateReportAcknowledgements() {
}

/**
 * Get Report Count
 *
 * returns a count of reports matching your criteria;
 * by default, the number of reports generated in the last 90 days,
 * regardless of acknowledgement status
 *
 * @see GetReportCountRequest
 * @return GetReportCountResponse
 */
func (api Report) GetReportCount() {
}

/**
 * Request Report
 *
 * requests the generation of a report
 *
 * @see RequestReportRequest
 * @return RequestReportResponse
 */
func (api Report) RequestReport() {
}

/**
 * Cancel Report Requests
 *
 * cancels report requests that have not yet started processing,
 * by default all those within the last 90 days
 *
 * @see CancelReportRequestsRequest
 * @return CancelReportRequestsResponse
 */
func (api Report) CancelReportRequests() {
}

/**
 * Get Report List
 *
 * returns a list of reports; by default the most recent ten reports,
 * regardless of their acknowledgement status
 *
 * @see GetReportListRequest
 * @return GetReportListResponse
 */
func (api Report) GetReportList() {
}

/**
 * Get Report Request List
 *
 * returns a list of report requests ids and their associated metadata
 *
 * @see GetReportRequestListRequest
 * @return GetReportRequestListResponse
 */
func (api Report) GetReportRequestList() {
}

/**
 * Get Report Schedule List By Next Token
 *
 * retrieve the next batch of list items and if there are more items to retrieve
 *
 * @see GetReportScheduleListByNextTokenRequest
 * @return GetReportScheduleListByNextTokenResponse
 */
func (api Report) GetReportScheduleListByNextToken() {
}

/**
 * Get Report List By Next Token
 *
 * retrieve the next batch of list items and if there are more items to retrieve
 *
 * @see GetReportListByNextTokenRequest
 * @return GetReportListByNextTokenResponse
 */
func (api Report) GetReportListByNextToken() {
}

/**
 * Manage Report Schedule
 *
 * Creates, updates, or deletes a report schedule
 * for a given report type, such as order reports in particular.
 *
 * @see ManageReportScheduleRequest
 * @return ManageReportScheduleResponse
 */
func (api Report) ManageReportSchedule() {
}

/**
 * Get Report Request Count
 *
 * returns a count of report requests; by default all the report
 * requests in the last 90 days
 *
 * @see GetReportRequestCountRequest
 * @return GetReportRequestCountResponse
 */
func (api Report) GetReportRequestCount() {
}

/**
 * Get Report Schedule List
 *
 * returns the list of report schedules
 *
 * @see GetReportScheduleListRequest
 * @return GetReportScheduleListResponse
 */
func (api Report) GetReportScheduleList() {
}
