package test

import "database/sql/driver"

type Service struct {
	ServiceName       string `json:"serviceName"`
	FilePath          string `json:"filePath"`
	ScenariosFileName string `json:"scenariosFileName"`
	IsGetService      bool   `json:"isGetService"`
	ServiceUrl        string `json:"serviceUrl"`
	HttpMethodType    string `json:"httpMethodType"`
	WorkflowService   string `json:"workflowService"`
}

type Case struct {
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	IsDBMockRequired   bool     `json:"isDBMockRequired"`
	RequestUrl         string   `json:"requestUrl"`
	RequestFileName    string   `json:"requestFileName"`
	ResponseFileName   string   `json:"responseFileName"`
	ExpectedStatusCode int      `json:"expectedStatusCode"`
	IsIgnoreCompare    bool     `json:"isIgnoreCompare"`
	DbMockSteps        []DbStep `json:"dbMockSteps"`
}

type DbStep struct {
	DbAction              string         `json:"dbAction"`
	SqlArgs               []driver.Value `json:"sqlArgs"`
	ReturnRow             []string       `json:"returnRow"`
	ReturnError           string         `json:"returnError"`
	Rows                  []RowsToAdd    `json:"rows"`
	IsNeedToEditRowValues bool           `json:"isNeedToEditRowValues"`
	IsNeedToEditArgs      bool           `json:"isNeedToEditArgs"`
	SqlQuery              string         `json:"sqlQuery"`
	IsErrorRequired       bool           `json:"isErrorRequired"`
}

type RowsToAdd struct {
	Values []driver.Value `json:"values"`
}
