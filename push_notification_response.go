package apns

// The maximum number of seconds we're willing to wait for a response
// from the Apple Push Notification Service.
const TimeoutSeconds = 5

const (
	NO_ERRORS            uint8 = 0
	PROCESSING_ERROR     uint8 = 1
	MISSING_DEVICE_TOKEN uint8 = 2
	MISSING_TOPIC        uint8 = 3
	MISSING_PAYLOAD      uint8 = 4
	INVALID_TOKEN_SIZE   uint8 = 5
	INVALID_TOPIC_SIZE   uint8 = 6
	INVALID_PAYLOAD_SIZE uint8 = 7
	INVALID_TOKEN        uint8 = 8
	SHUTDOWN             uint8 = 10
	PROTOCOL_ERROR       uint8 = 128
	UNKNOWN              uint8 = 255
	CERT_ERROR           uint8 = 101
	INTERNET_ERROR       uint8 = 103
	INTERNET_WRITE_ERROR uint8 = 107
)

// This enumerates the response codes that Apple defines
// for push notification attempts.
var ApplePushResponses = map[uint8]string{
	0:   "NO_ERRORS",
	1:   "PROCESSING_ERROR",
	2:   "MISSING_DEVICE_TOKEN",
	3:   "MISSING_TOPIC",
	4:   "MISSING_PAYLOAD",
	5:   "INVALID_TOKEN_SIZE",
	6:   "INVALID_TOPIC_SIZE",
	7:   "INVALID_PAYLOAD_SIZE",
	8:   "INVALID_TOKEN",
	10:  "SHUTDOWN",
	128: "PROTOCOL_ERROR",
	255: "UNKNOWN",
}

// Not Apple response code
var PushResponses = map[uint8]string{
	101: "CERT_ERROR",
	103: "INTERNET_ERROR",
}

// PushNotificationResponse details what Apple had to say, if anything.
type PushNotificationResponse struct {
	Success       bool
	AppleResponse string
	ResponseCode  uint8
	Error         error
	Identifier    int32
}

// NewPushNotificationResponse creates and returns a new PushNotificationResponse
// structure; it defaults to being unsuccessful at first.
func NewPushNotificationResponse() (resp *PushNotificationResponse) {
	resp = new(PushNotificationResponse)
	resp.Success = false
	return
}
