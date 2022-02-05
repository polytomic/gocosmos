package gocosmos

import (
	"net/http"
	"testing"
)

func TestOverLimitErr(t *testing.T) {
	t.Run("returns the error message", func(t *testing.T) {
		e := NewOverLimitError(http.StatusTooManyRequests, []byte(overLimitResponse))
		if http.StatusTooManyRequests != e.StatusCode {
			t.Errorf("expected status code to be %d; got %d", http.StatusTooManyRequests, e.StatusCode)
		}

		if e.Error() != "Request rate is large. More Request Units may be needed, so no changes were made. Please retry this request later. Learn more: http://aka.ms/cosmosdb-error-429" {
			t.Error("over limit error did not return message from response")
		}
	})
}

const overLimitResponse = `{
    "code": "429",
    "message": "Message: {\"Errors\":[\"Request rate is large. More Request Units may be needed, so no changes were made. Please retry this request later. Learn more: http://aka.ms/cosmosdb-error-429\"]}\r\nActivityId: ecc73276-17fa-40b3-b3bc-ba2cc0999bd9, Request URI: /apps/06f97901-5530-4c65-a094-ceac7aeba20b/services/a1cf8c5b-1d8c-46cc-b081-474d4571e673/partitions/ada22d39-6f76-4d07-8716-4c2801c6ecdc/replicas/132883176128386148s/, RequestStats: \r\nRequestStartTime: 2022-02-03T00:42:42.9998655Z, RequestEndTime: 2022-02-03T00:42:43.0098845Z,  Number of regions attempted:1\r\n{\"systemHistory\":[{\"dateUtc\":\"2022-02-03T00:41:45.5688049Z\",\"cpu\":1.852,\"memory\":466450700.000,\"threadInfo\":{\"isThreadStarving\":\"False\",\"threadWaitIntervalInMs\":0.0133,\"availableThreads\":32765,\"minThreads\":52,\"maxThreads\":32767}},{\"dateUtc\":\"2022-02-03T00:41:55.5790555Z\",\"cpu\":1.324,\"memory\":465989604.000,\"threadInfo\":{\"isThreadStarving\":\"False\",\"threadWaitIntervalInMs\":0.0092,\"availableThreads\":32764,\"minThreads\":52,\"maxThreads\":32767}},{\"dateUtc\":\"2022-02-03T00:42:05.5891643Z\",\"cpu\":1.506,\"memory\":465654180.000,\"threadInfo\":{\"isThreadStarving\":\"False\",\"threadWaitIntervalInMs\":0.0147,\"availableThreads\":32764,\"minThreads\":52,\"maxThreads\":32767}},{\"dateUtc\":\"2022-02-03T00:42:15.5993989Z\",\"cpu\":1.213,\"memory\":465313416.000,\"threadInfo\":{\"isThreadStarving\":\"False\",\"threadWaitIntervalInMs\":0.0123,\"availableThreads\":32763,\"minThreads\":52,\"maxThreads\":32767}},{\"dateUtc\":\"2022-02-03T00:42:25.6095266Z\",\"cpu\":1.771,\"memory\":464642872.000,\"threadInfo\":{\"isThreadStarving\":\"False\",\"threadWaitIntervalInMs\":0.0133,\"availableThreads\":32764,\"minThreads\":52,\"maxThreads\":32767}},{\"dateUtc\":\"2022-02-03T00:42:35.6197589Z\",\"cpu\":1.926,\"memory\":466263304.000,\"threadInfo\":{\"isThreadStarving\":\"False\",\"threadWaitIntervalInMs\":0.0289,\"availableThreads\":32764,\"minThreads\":52,\"maxThreads\":32767}}]}\r\nRequestStart: 2022-02-03T00:42:42.9998655Z; ResponseTime: 2022-02-03T00:42:43.0098845Z; StoreResult: StorePhysicalAddress: rntbd://cdb-ms-prod-westus2-fd37.documents.azure.com:14141/apps/06f97901-5530-4c65-a094-ceac7aeba20b/services/a1cf8c5b-1d8c-46cc-b081-474d4571e673/partitions/ada22d39-6f76-4d07-8716-4c2801c6ecdc/replicas/132883176128386148s/, LSN: 5460, GlobalCommittedLsn: 5459, PartitionKeyRangeId: , IsValid: True, StatusCode: 429, SubStatusCode: 3200, RequestCharge: 0.38, ItemLSN: -1, SessionToken: , UsingLocalLSN: True, TransportException: null, BELatencyMs: , ActivityId: ecc73276-17fa-40b3-b3bc-ba2cc0999bd9, RetryAfterInMs: 299937, TransportRequestTimeline: {\"requestTimeline\":[{\"event\": \"Created\", \"startTimeUtc\": \"2022-02-03T00:42:42.9998655Z\", \"durationInMs\": 0.0092},{\"event\": \"ChannelAcquisitionStarted\", \"startTimeUtc\": \"2022-02-03T00:42:42.9998747Z\", \"durationInMs\": 13.2949},{\"event\": \"Pipelined\", \"startTimeUtc\": \"2022-02-03T00:42:43.0131696Z\", \"durationInMs\": 0.0865},{\"event\": \"Transit Time\", \"startTimeUtc\": \"2022-02-03T00:42:43.0132561Z\", \"durationInMs\": 0.4604},{\"event\": \"Received\", \"startTimeUtc\": \"2022-02-03T00:42:43.0137165Z\", \"durationInMs\": 0.1261},{\"event\": \"Completed\", \"startTimeUtc\": \"2022-02-03T00:42:43.0138426Z\", \"durationInMs\": 0}],\"requestSizeInBytes\":546,\"responseMetadataSizeInBytes\":91,\"responseBodySizeInBytes\":174};\r\n ResourceType: Document, OperationType: ReadFeed\r\n, SDK: Microsoft.Azure.Documents.Common/2.14.0"
}`
