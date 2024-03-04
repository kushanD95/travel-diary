package test

var QueryMap = map[string]string{
	"createUser":        "INSERT INTO \"USER\" (\"USER_NAME\",\"PWD\",\"CREATED_AT\") VALUES ($1,$2,$3) RETURNING \"ID\"",
	"createUserDetails": "INSERT INTO \"USER_DETAILS\" (\"USER_ID\",\"F_NAME\",\"L_NAME\",\"COUNTRY\",\"CREATED_AT\") VALUES ($1,$2,$3,$4,$5) RETURNING \"ID\"",
}
