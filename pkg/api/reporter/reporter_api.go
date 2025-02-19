package reporter

import (
	"net/http"
	"reflect"
	"soarca/internal/controller/informer"
	"soarca/internal/logger"
	"soarca/pkg/api/error"
	"soarca/pkg/models/api"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var log *logger.Log

type Empty struct{}

func init() {
	log = logger.Logger(reflect.TypeOf(Empty{}).PkgPath(), logger.Info, "", logger.Json)
}

// reportHandler implements the handler functions that can be called by the gin api is dependent on a database.
type reportHandler struct {
	informer informer.IExecutionInformer
}

// NewReportHandler makes a new instance of playbookControler
func NewReportHandler(informer informer.IExecutionInformer) *reportHandler {
	return &reportHandler{informer: informer}
}

// GetExecutions GET handler for obtaining all the executions that can be retrieved.
// Returns this to the gin context as a list if execution IDs in json format
//
//	@Summary	gets all the UUIDs for the executions that can be retireved
//	@Schemes
//	@Description	return all stored executions
//	@Tags			reporter
//	@Produce		json
//	@success		200	{array}		api.PlaybookExecutionReport
//	@failure		400	{object}	api.Error
//	@Router			/reporter [GET]
func (reportHandler *reportHandler) GetExecutions(g *gin.Context) {
	executions, err := reportHandler.informer.GetExecutions()
	if err != nil {
		log.Debug("Could not get executions from informer")
		error.SendErrorResponse(g, http.StatusInternalServerError, "Could not get executions from informer", "GET /reporter/", "")
		return
	}

	executionsParsed := []api.PlaybookExecutionReport{}
	for _, executionEntry := range executions {
		executionEntryParsed, err := parseCachePlaybookEntry(executionEntry)
		if err != nil {
			log.Debug("Could not parse entry to reporter result model")
			log.Error(err)
			error.SendErrorResponse(g, http.StatusInternalServerError, "Could not parse execution report", "GET /reporter/", "")
			return
		}
		executionsParsed = append(executionsParsed, executionEntryParsed)
	}

	g.JSON(http.StatusOK, executionsParsed)
}

// GetExecutionReport GET handler for obtaining the information about an execution.
// Returns this to the gin context as a PlaybookExecutionReport object at soarca/model/api/reporter
//
//	@Summary	gets information about an ongoing playbook execution
//	@Schemes
//	@Description	return execution information
//	@Tags			reporter
//	@Produce		json
//	@Param			id	path		string	true	"execution identifier"
//	@success		200	{object}	api.PlaybookExecutionReport
//	@failure		400	{object}	api.Error
//	@Router			/reporter/{id} [GET]
func (handler *reportHandler) GetExecutionReport(g *gin.Context) {
	id := g.Param("id")
	log.Trace("Trying to obtain execution for id: ", id)
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Debug("Could not parse id parameter for request")
		error.SendErrorResponse(g, http.StatusBadRequest, "Could not parse id parameter for request", "GET /reporter/"+id, err.Error())
		return
	}

	executionEntry, err := handler.informer.GetExecutionReport(uuid)
	if err != nil {
		log.Debug("Could not find execution for given id")
		log.Error(err)
		error.SendErrorResponse(g, http.StatusBadRequest, "Could not find execution for given ID", "GET /reporter/"+id, "")
		return
	}

	executionEntryParsed, err := parseCachePlaybookEntry(executionEntry)
	if err != nil {
		log.Debug("Could not parse entry to reporter result model")
		log.Error(err)
		error.SendErrorResponse(g, http.StatusInternalServerError, "Could not parse execution report", "GET /reporter/"+id, "")
		return
	}
	g.JSON(http.StatusOK, executionEntryParsed)
}
