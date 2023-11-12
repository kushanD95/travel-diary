package utils

// Logger init
const (
	METHOD         = "Method"
	Init_error     = "Logging client creation failed. error: %v"
	ERROR          = "Error"
	STARTED        = "Started"
	END            = "End"
	END_WITH_ERROR = "End with error"
)

// method names
const (
	// handler
	CHECK    = "Check"
	LOGIN    = "Login"
	REGISTER = "Register"
	READY    = "Ready"

	//service
	FETCH_CONFIG_SERVICE  = "FetchConfigService"
	REGISTER_USER_SERVICE = "RegisterUserService"

	//repository
	GET_CONFIG    = "GetConfig"
	REGISTER_USER = "RegisterUser"

	// validator
	ENV_CONFIG_REQ = "EnvConfigReq"
	USER_REQ       = "UserReq"
)

// logger handler
const (
	CHECK_HANDLER    = CHECK + " Handler %v"
	LOGIN_HANDLER    = LOGIN + " Handler %v"
	REGISTER_HANDLER = REGISTER + " Handler %v"
	READY_HANDLER    = READY + " Handler %v"
)

// logger services
const (
	FETCH_CONFIG_SERVICE_LOG  = FETCH_CONFIG_SERVICE + " Service  %v"
	REGISTER_USER_SERVICE_LOG = REGISTER_USER_SERVICE + " Service  %v"
)

// logger repository
const (
	GET_CONFIG_REPO    = GET_CONFIG + " Repository %v"
	REGISTER_USER_REPO = REGISTER_USER + " Repository %v"
)

// logger request validators
const (
	ENV_CONFIG_REQ_LOG = ENV_CONFIG_REQ + " Validator %v"
	USER_REQ_LOG       = USER_REQ + " Validator %v"
)

// logs
const (
	RECEIVED_PAYLOAD = "received payload %v"
)
