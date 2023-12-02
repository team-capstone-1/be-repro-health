package controller

import (
	m "capstone-project/middleware"
	"capstone-project/model"
	"capstone-project/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetDataCountForDoctorControllerOneMonth(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	// Consultation
	consultationResponseData, err := repository.GetConsultationsByDoctorID(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get consultation data",
			"response": err.Error(),
		})
	}

	// lastMonthConsultation := time.Now().AddDate(0, -2, 0)
	// lastMonthConsultationData, err := repository.GetConsultationByDoctorAndMonth(doctor, lastMonthConsultation)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]any{
	// 		"message":  "failed get last month consultation data",
	// 		"response": err.Error(),
	// 	})
	// }

	// thisMonthConsultation := time.Now().AddDate(0, -1, 0)
	// thisMonthConsultationData, err := repository.GetConsultationByDoctorAndMonth(doctor, thisMonthConsultation)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]any{
	// 		"message":  "failed get this month consultation data",
	// 		"response": err.Error(),
	// 	})
	// }

	// allConsultations := append(lastMonthConsultationData, thisMonthConsultationData...)

	// totalConsultationLastMonth := float64(len(lastMonthConsultationData))
	// totalConsultationThisMonth := float64(len(thisMonthConsultationData))

	// var consultationPercentage float64
	// if totalConsultationLastMonth > 0 {
	// 	consultationPercentage = ((totalConsultationThisMonth - totalConsultationLastMonth) / totalConsultationLastMonth) * 100
	// } else if totalConsultationLastMonth == 0 && totalConsultationThisMonth > 0 {
	// 	consultationPercentage = 100.0
	// } else {
	// 	consultationPercentage = 0.0
	// }

	// Patient
	patientResponseData, err := repository.GetPatientByDoctorID(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get patient data",
			"response": err.Error(),
		})
	}

	lastMonthPatient := time.Now().AddDate(0, -2, 0)
	lastMonthPatientData, err := repository.GetPatientByDoctorAndMonth(doctor, lastMonthPatient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get last month patient data",
			"response": err.Error(),
		})
	}

	thisMonthPatient := time.Now().AddDate(0, -1, 0)
	thisMonthPatientData, err := repository.GetPatientByDoctorAndMonth(doctor, thisMonthPatient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get this month patient data",
			"response": err.Error(),
		})
	}

	allPatients := append(lastMonthPatientData, thisMonthPatientData...)

	totalPatientLastMonth := float64(len(lastMonthPatientData))
	totalPatientThisMonth := float64(len(thisMonthPatientData))

	var patientPercentage float64
	if totalPatientLastMonth > 0 {
		patientPercentage = ((totalPatientThisMonth - totalPatientLastMonth) / totalPatientLastMonth) * 100
	} else if totalPatientLastMonth == 0 && totalPatientThisMonth > 0 {
		patientPercentage = 100.0
	} else {
		patientPercentage = 0.0
	}

	// Transaction
	transactionResponseData, err := repository.GetAllTransactionsByDoctorID(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get transaction data",
			"response": err.Error(),
		})
	}

	lastMonthTrasaction := time.Now().AddDate(0, -2, 0)
	lastMonthTrasactionData, err := repository.GetDoneTransactionsByDoctorAndMonth(doctor, lastMonthTrasaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get last month transaction data",
			"response": err.Error(),
		})
	}

	thisMonthTrasaction := time.Now().AddDate(0, -1, 0)
	thisMonthTrasactionData, err := repository.GetDoneTransactionsByDoctorAndMonth(doctor, thisMonthTrasaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get this month transaction data",
			"response": err.Error(),
		})
	}

	allTransactions := append(lastMonthTrasactionData, thisMonthTrasactionData...)

	totalPriceLastMonthTransaction := calculateTotalTransaction(lastMonthTrasactionData)
	totalPriceThisMonthTransaction := calculateTotalTransaction(thisMonthTrasactionData)

	var transactionPercentage float64
	if totalPriceLastMonthTransaction > 0 {
		transactionPercentage = ((totalPriceThisMonthTransaction - totalPriceLastMonthTransaction) / totalPriceLastMonthTransaction) * 100
	} else if totalPriceLastMonthTransaction > 0 {
		transactionPercentage = 100.0
	} else {
		transactionPercentage = 0.0
	}

	// Article
	articleResponseData, err := repository.DoctorGetAllArticles(doctor)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get article data",
			"response": err.Error(),
		})
	}

	totalConsultation := len(consultationResponseData)
	totalArticle := len(articleResponseData)
	totalTransaction := len(transactionResponseData)
	totalPatient := len(patientResponseData)
	// totalConsultations := calculateTotalConsultation(allConsultations)
	totalTransactions := calculateTotalTransaction(allTransactions)
	totalPatients := calculateTotalPatient(allPatients)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get dashboard data for one month",
		// Consultation
		"totalConsultation":          totalConsultation,
		// "totalConsultations":         totalConsultations,
		// "totalConsultationLastMonth": totalConsultationLastMonth,
		// "totalConsultationThisMonth": totalConsultationThisMonth,
		// "consultationPercentage":     consultationPercentage,
		// Patients
		"totalPatient":          totalPatient,
		"totalPatients":         totalPatients,
		"totalPatientLastMonth": totalPatientLastMonth,
		"totalPatientThisMonth": totalPatientThisMonth,
		"patientPercentage":     patientPercentage,
		// Transaction
		"totalTransaction":      totalTransaction,
		"totalTransactions":     totalTransactions,
		"totalPriceLastMonth":   totalPriceLastMonthTransaction,
		"totalPriceThisMonth":   totalPriceThisMonthTransaction,
		"transactionPercentage": transactionPercentage,
		// Article
		"totalArticle": totalArticle,
	})
}

func calculateTotalTransaction(transactions []model.Transaction) float64 {
	totalTransaction := 0.0
	for _, transaction := range transactions {
		totalTransaction += transaction.Total
	}
	return totalTransaction
}

func calculateTotalPatient(patients []model.Patient) float64 {
	totalPatient := 0.0
	for range patients {
		totalPatient += 1
	}
	return totalPatient
}

func calculateTotalConsultation(consultations []model.Consultation) float64 {
	totalConsultation := 0.0
	for range consultations {
		totalConsultation += 1
	}
	return totalConsultation
}