package constant

// Success
const (
	SUCCESS_LOGIN               = "logged in successfully"
	SUCCESS_REGISTER            = "register successfully"
	SUCCESS_CREATED             = "data created successfully"
	SUCCESS_RETRIEVED           = "data retrieved successfully"
	SUCCESS_UPDATED             = "data updated successfully"
)

// Error
const (
	ERROR_ID_NOT_FOUND         = "id not found"
	ERROR_ID_INVALID           = "invalid id"
	ERROR_LOGIN                = "login failed"
	ERROR_PASSWORD_INVALID     = "invalid password"
	ERROR_PASSWORD_HASH        = "error hashing password"
	ERROR_PASSWORD_CONFIRM     = "password do not match"
	ERROR_DATA_NOTFOUND        = "data not found"
	ERROR_DATA_EMPTY           = "data is empty"
	ERROR_DATA_EXIST           = "data already exists"
	ERROR_DATA_TYPE            = "data type unsupported"
	ERROR_DATA_RETRIEVED       = "failed to retrieve data"
	ERROR_DATA_INVALID         = "invalid data. allowed data: "
	ERROR_MIN_LENGTH           = "minimum length is %d characters"
	ERROR_MAX_LENGTH           = "maximum length is %d characters"
	ERROR_TOKEN_INVALID        = "invalid token"
	ERROR_TOKEN_GENERATE       = "generate token failed"
	ERROR_TOKEN_NOTFOUND       = "token not found"
)
