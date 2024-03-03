package test

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kushanD95/traval-diary/app/controller/handler"
	"github.com/kushanD95/traval-diary/app/services"
	"github.com/kushanD95/traval-diary/package/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TestSuite struct {
	suite.Suite
	DB       *gorm.DB
	Mock     sqlmock.Sqlmock
	DbConn   *services.DBConn
	DiarySrv *config.AppConfig
}

var (
	TestTime     time.Time
	testDataPath = "testData/testScenarios.json"
)

var handlerMethodsMap = map[string]func(c *fiber.Ctx) (err error){
	"Register": handler.Register,
}

func (s *TestSuite) SetupSuite() {

	var (
		db  *sql.DB
		err error
	)

	db, s.Mock, _ = sqlmock.New()

	s.DB, err = gorm.Open(
		postgres.New(postgres.Config{Conn: db}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

	require.NoError(s.T(), err)

	dbConn := &services.DBConn{
		Db: s.DB,
	}
	s.DbConn = dbConn

	s.DbConn.SetupDB()

	s.DiarySrv = &config.AppConfig{}
	s.DiarySrv.InitConfig()
	s.DiarySrv.InitLogger()

	config.AppConfigutarion = s.DiarySrv

	str := "2024-01-01T00:00:00.000Z"
	TestTime, _ = time.Parse(time.RFC3339, str)
}

func (s *TestSuite) AfterTest(_, _ string) {
	// lg := s.DiarySrv.GetLogger()

	isExpectationsWereMet := s.Mock.ExpectationsWereMet()
	s.Suite.T().Log("AfterTest", zap.Any("s.Mock.ExpectationsWereMet: ", isExpectationsWereMet))
	require.NoError(s.T(), isExpectationsWereMet)
}

func TestHandlerInit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestHandler() {
	// lg := s.DiarySrv.GetLogger()
	var testServices []Service
	testServiceRead, err := os.ReadFile(testDataPath)
	if err != nil {
		// s.Suite.T().Error(FileReadError, zap.Error(err))
		s.Suite.T().Error(err, FileReadError)
	}

	err = json.Unmarshal(testServiceRead, &testServices)
	if err != nil {
		// s.Suite.T().Error(FileUnmarshalError, zap.Error(err))
		s.Suite.T().Error(err, FileUnmarshalError)
	}

	for _, service := range testServices {
		s.Suite.T().Log(UnitTestDesc, zap.String(ServiceName, service.ServiceName))
		var testCases []Case

		testCaseRead, err := os.ReadFile(service.FilePath + service.ScenariosFileName)
		if err != nil {
			// s.Suite.T().Error(FileReadError, zap.Error(err))
			s.Suite.T().Error(err, FileReadError)
		}

		err = json.Unmarshal(testCaseRead, &testCases)
		if err != nil {
			// s.Suite.T().Error(FileUnmarshalError, zap.Error(err))
			s.Suite.T().Error(err, FileUnmarshalError)
		}

		for _, testCase := range testCases {
			s.T().Run(testCase.Name, func(t *testing.T) {
				s.Suite.T().Log(
					UnitTestDesc,
					zap.String(TestName, testCase.Name),
					zap.String(TestDesc, testCase.Description))
				app, req := buildTestRequest(service, testCase)
				if testCase.IsDBMockRequired {
					s.MockDBDetails(testCase.DbMockSteps)
				}
				res, err := app.Test(req, -1)
				if err != nil {
					s.Suite.T().Error(TestCaseError, zap.String(UnitTestDesc, testCase.Name))
				}

				isResStatusCodeEq := assert.Equal(s.T(), res.StatusCode, testCase.ExpectedStatusCode)

				isResponseEq := true
				expectedResponse, err := os.ReadFile(filepath.Join(service.FilePath, testCase.ResponseFileName+".json"))
				if err != nil {
					s.Suite.T().Error(UnitTestDesc, zap.String(FileReadError, testCase.Name), zap.Error(err))
				}

				responseBody, err := io.ReadAll(res.Body)
				if err != nil {
					s.Suite.T().Error(UnitTestDesc, zap.String(ResponseReadError, testCase.Name), zap.Error(err))
				}

				if !testCase.IsIgnoreCompare {
					isResponseEq = assert.JSONEq(s.T(), string(expectedResponse), string(responseBody))
				}

				if !isResponseEq || !isResStatusCodeEq {
					s.Suite.T().Log(
						UnitTestDesc,
						zap.Int(ActualStatusCode, res.StatusCode),
						zap.Int(ExpectedStatusCode, testCase.ExpectedStatusCode),
						zap.String(ActualResponse, string(responseBody)),
						zap.String(ExpectedResponse, string(expectedResponse)))
					t.Fail()
				}

			})
		}
	}
}

func buildTestRequest(service Service, testCase Case) (*fiber.App, *http.Request) {
	var req *http.Request
	app := fiber.New()

	switch service.HttpMethodType {
	case Get:
		app.Get(service.ServiceUrl, handlerMethodsMap[service.WorkflowService])
		req = httptest.NewRequest(service.HttpMethodType, testCase.RequestUrl, nil)
	case Post:
		bodyDetails, _ := os.ReadFile(filepath.Join(service.FilePath, testCase.RequestFileName+".json"))
		app.Post(service.ServiceUrl, handlerMethodsMap[service.WorkflowService])
		req = httptest.NewRequest(service.HttpMethodType, testCase.RequestUrl, strings.NewReader(string(bodyDetails)))
	}
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return app, req
}

func (s *TestSuite) MockDBDetails(mockDbSteps []DbStep) {
	for _, step := range mockDbSteps {

		switch step.DbAction {
		case Begin:
			s.Mock.ExpectBegin()
		case Commit:
			s.Mock.ExpectCommit()
		case TransactionFailure:
			s.Mock.ExpectBegin().WillReturnError(fmt.Errorf(TransactionError))
		case ExecuteQuery:
			{
				if step.IsNeedToEditArgs {
					var args []driver.Value
					for _, arg := range step.SqlArgs {
						if arg == AnyArgs {
							args = append(args, sqlmock.AnyArg())
						} else {
							args = append(args, arg)
						}
					}
					step.SqlArgs = args
				}

				var rowDetails []RowsToAdd
				for _, row := range step.Rows {
					var rowValues []driver.Value
					for _, data := range row.Values {
						if data == "time" {
							rowValues = append(rowValues, TestTime)
						} else if data == "int" {
							rowValues = append(rowValues, 1)
						} else {
							rowValues = append(rowValues, data)
						}
					}

					rowDetails = append(rowDetails, RowsToAdd{Values: rowValues})
				}

				rowsAdd := []*sqlmock.Rows{}

				var isRecExist bool = false
				if !step.IsErrorRequired {
					for _, row := range rowDetails {
						isRecExist = true
						rowsAdd = append(rowsAdd, sqlmock.NewRows(step.ReturnRow).AddRow(row.Values...))
					}
					if isRecExist {
						rowsAdd = append(rowsAdd, sqlmock.NewRows(step.ReturnRow))
					}
					query := QueryMap[step.SqlQuery]
					s.Mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(step.SqlArgs...).WillReturnRows(rowsAdd...)
				} else {
					query := QueryMap[step.SqlQuery]
					s.Mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(step.SqlArgs...).WillReturnError(fmt.Errorf(TransactionError))
				}
			}
		case ExecuteQueryUpdate:
			query := QueryMap[step.SqlQuery]
			if !step.IsErrorRequired {
				s.Mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(0, 1))
			} else {
				s.Mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(fmt.Errorf(TransactionError))
			}

		}
	}
}
