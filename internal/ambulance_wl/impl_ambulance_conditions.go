package ambulance_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Nasledujúci kód je kópiou vygenerovaného a zakomentovaného kódu zo súboru api_ambulance_conditions.go
func (this *implAmbulanceConditionsAPI) GetConditions(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		result := ambulance.PredefinedConditions
		if result == nil {
			result = []Condition{}
		}
		return nil, result, http.StatusOK
	})
}

// GetCondition - Provides details about condition
func (this *implAmbulanceConditionsAPI) GetCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		conditionCode := ctx.Param("conditionCode")
		for _, condition := range ambulance.PredefinedConditions {
			if condition.Code == conditionCode {
				return nil, condition, http.StatusOK
			}
		}
		return nil, nil, http.StatusNotFound
	})
}

// CreateCondition - Saves new condition definition
func (this *implAmbulanceConditionsAPI) CreateCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		condition := Condition{}
		err := ctx.BindJSON(&condition)
		if err != nil {
			return nil, err.Error(), http.StatusBadRequest
		}
		ambulance.PredefinedConditions = append(ambulance.PredefinedConditions, condition)
		return ambulance, condition, http.StatusCreated
	})
}

// UpdateCondition - Updates specific condition
func (this *implAmbulanceConditionsAPI) UpdateCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		conditionCode := ctx.Param("conditionCode")
		condition := Condition{}
		err := ctx.BindJSON(&condition)
		if err != nil {
			return nil, err.Error(), http.StatusBadRequest
		}
		for i, c := range ambulance.PredefinedConditions {
			if c.Code == conditionCode {
				ambulance.PredefinedConditions[i] = condition
				return ambulance, condition, http.StatusOK
			}
		}
		return nil, nil, http.StatusNotFound
	})
}

// DeleteCondition - Deletes specific condition
func (this *implAmbulanceConditionsAPI) DeleteCondition(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(
		ctx *gin.Context,
		ambulance *Ambulance,
	) (updatedAmbulance *Ambulance, responseContent interface{}, status int) {
		conditionCode := ctx.Param("conditionCode")
		for i, c := range ambulance.PredefinedConditions {
			if c.Code == conditionCode {
				ambulance.PredefinedConditions = append(ambulance.PredefinedConditions[:i], ambulance.PredefinedConditions[i+1:]...)
				return ambulance, nil, http.StatusNoContent
			}
		}
		return nil, nil, http.StatusNotFound
	})
}
