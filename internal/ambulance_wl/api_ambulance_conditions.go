/*
 * Waiting List Api
 *
 * Ambulance Waiting List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xkappel@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package ambulance_wl

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

type AmbulanceConditionsAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // CreateCondition - Saves new condition definition
   CreateCondition(ctx *gin.Context)

    // DeleteCondition - Deletes specific condition
   DeleteCondition(ctx *gin.Context)

    // GetCondition - Provides details about condition
   GetCondition(ctx *gin.Context)

    // GetConditions - Provides the list of conditions associated with ambulance
   GetConditions(ctx *gin.Context)

    // UpdateCondition - Updates specific condition
   UpdateCondition(ctx *gin.Context)

 }

// partial implementation of AmbulanceConditionsAPI - all functions must be implemented in add on files
type implAmbulanceConditionsAPI struct {

}

func newAmbulanceConditionsAPI() AmbulanceConditionsAPI {
  return &implAmbulanceConditionsAPI{}
}

func (this *implAmbulanceConditionsAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/waiting-list/:ambulanceId/condition", this.CreateCondition)
  routerGroup.Handle( http.MethodDelete, "/waiting-list/:ambulanceId/condition/:conditionCode", this.DeleteCondition)
  routerGroup.Handle( http.MethodGet, "/waiting-list/:ambulanceId/condition/:conditionCode", this.GetCondition)
  routerGroup.Handle( http.MethodGet, "/waiting-list/:ambulanceId/condition", this.GetConditions)
  routerGroup.Handle( http.MethodPut, "/waiting-list/:ambulanceId/condition/:conditionCode", this.UpdateCondition)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // CreateCondition - Saves new condition definition
// func (this *implAmbulanceConditionsAPI) CreateCondition(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteCondition - Deletes specific condition
// func (this *implAmbulanceConditionsAPI) DeleteCondition(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetCondition - Provides details about condition
// func (this *implAmbulanceConditionsAPI) GetCondition(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetConditions - Provides the list of conditions associated with ambulance
// func (this *implAmbulanceConditionsAPI) GetConditions(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateCondition - Updates specific condition
// func (this *implAmbulanceConditionsAPI) UpdateCondition(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//

