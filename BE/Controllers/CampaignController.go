package Controllers

import (
	"net/http"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
)

func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel
	// var guid = uuid.New().String()
	var data Const.Paging
	// * Get Kols from the database based on the range of pageIndex and pageSize
	// * TODO: Implement the logic to get parameters from the request
	// ? If parameter passed in the request is not valid, return the response with HTTP Status Bad Request (400)
	// @params: pageIndex
	// @params: pageSize
	if err := context.ShouldBind(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// * Perform Logic Here
	// ! Pass the parameters to the Logic Layer
	kols, totalCount, error := Logic.GetKolLogic(Initializers.DB, data)
	if error != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = error.Error()
		KolsVM.PageIndex = data.PageIndex
		KolsVM.PageSize = data.PageSize
		// KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// * Return the response after the logic is executed
	// ? If the logic is successful, return the response with HTTP Status OK (200)
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = data.PageIndex
	KolsVM.PageSize = data.PageSize
	// KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = totalCount
	context.JSON(http.StatusOK, KolsVM)
}
