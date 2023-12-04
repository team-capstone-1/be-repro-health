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
	lastMonthConsultation := time.Now().AddDate(0, -2, 0)
	lastMonthConsultationData, err := repository.GetConsultationByDoctorAndMonth(doctor, lastMonthConsultation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get last month consultation data",
			"response": err.Error(),
		})
	}

	thisMonthConsultation := time.Now().AddDate(0, -1, 0)
	thisMonthConsultationData, err := repository.GetConsultationByDoctorAndMonth(doctor, thisMonthConsultation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get this month consultation data",
			"response": err.Error(),
		})
	}

	allConsultations := append(lastMonthConsultationData, thisMonthConsultationData...)

	totalConsultationLastMonth := float64(len(lastMonthConsultationData))
	totalConsultationThisMonth := float64(len(thisMonthConsultationData))

	var consultationPercentage float64
	if totalConsultationLastMonth > 0 {
		consultationPercentage = ((totalConsultationThisMonth - totalConsultationLastMonth) / totalConsultationLastMonth) * 100
	} else if totalConsultationLastMonth == 0 && totalConsultationThisMonth > 0 {
		consultationPercentage = 100.0
	} else {
		consultationPercentage = 0.0
	}

	// Patient
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
	lastMonthArticle := time.Now().AddDate(0, -2, 0)
	lastMonthArticleData, err := repository.DoctorGetAllArticlesByMonth(doctor, lastMonthArticle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get last month article data",
			"response": err.Error(),
		})
	}

	thisMonthArticle := time.Now().AddDate(0, -1, 0)
	thisMonthArticleData, err := repository.DoctorGetAllArticlesByMonth(doctor, thisMonthArticle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get this month article data",
			"response": err.Error(),
		})
	}

	allArticles := append(lastMonthArticleData, thisMonthArticleData...)

	totalArticleLastMonth := calculateTotalArticle(lastMonthArticleData)
	totalArticleThisMonth := calculateTotalArticle(thisMonthArticleData)

	var articlePercentage float64
	if totalArticleLastMonth > 0 {
		articlePercentage = ((totalArticleThisMonth - totalArticleLastMonth) / totalArticleLastMonth) * 100
	} else if totalArticleLastMonth == 0 && totalArticleThisMonth > 0 {
		articlePercentage = 100.0
	} else {
		articlePercentage = 0.0
	}

	totalTransaction := len(transactionResponseData)
	totalConsultations := calculateTotalConsultation(allConsultations)
	totalTransactions := calculateTotalTransaction(allTransactions)
	totalPatients := calculateTotalPatient(allPatients)
	totalArticle := calculateTotalArticle(allArticles)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get dashboard data for one month",
		// Consultation
		"totalConsultations":         totalConsultations,
		"totalConsultationLastMonth": totalConsultationLastMonth,
		"totalConsultationThisMonth": totalConsultationThisMonth,
		"consultationPercentage":     consultationPercentage,
		// Patients
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
		"totalArticles":         totalArticle,
		"totalArticleLastMonth": totalArticleLastMonth,
		"totalArticleThisMonth": totalArticleThisMonth,
		"articlePercentage":     articlePercentage,
	})
}

func GetDataCountForDoctorControllerOneWeek(c echo.Context) error {
	doctor := m.ExtractTokenUserId(c)
	if doctor == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	// Consultation
	lastWeekConsultation := time.Now().AddDate(0, 0, -14)
	lastWeekConsultationData, err := repository.GetConsultationByDoctorAndWeek(doctor, lastWeekConsultation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get last month consultation data",
			"response": err.Error(),
		})
	}

	thisWeekConsultation := time.Now().AddDate(0, 0, -7)
	thisWeekConsultationData, err := repository.GetConsultationByDoctorAndWeek(doctor, thisWeekConsultation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get this month consultation data",
			"response": err.Error(),
		})
	}

	allConsultations := append(lastWeekConsultationData, thisWeekConsultationData...)

	totalConsultationLastWeek := float64(len(lastWeekConsultationData))
	totalConsultationThisWeek := float64(len(thisWeekConsultationData))

	var consultationPercentage float64
	if totalConsultationLastWeek > 0 {
		consultationPercentage = ((totalConsultationThisWeek - totalConsultationLastWeek) / totalConsultationLastWeek) * 100
	} else if totalConsultationLastWeek == 0 && totalConsultationThisWeek > 0 {
		consultationPercentage = 100.0
	} else {
		consultationPercentage = 0.0
	}

	// Patient
	lastWeekPatient := time.Now().AddDate(0, 0, -14)
	lastWeekPatientData, err := repository.GetPatientByDoctorAndWeek(doctor, lastWeekPatient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get last month patient data",
			"response": err.Error(),
		})
	}

	thisWeekPatient := time.Now().AddDate(0, 0, -7)
	thisWeekPatientData, err := repository.GetPatientByDoctorAndWeek(doctor, thisWeekPatient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get this month patient data",
			"response": err.Error(),
		})
	}

	allPatients := append(lastWeekPatientData, thisWeekPatientData...)

	totalPatientLastWeek := float64(len(lastWeekPatientData))
	totalPatientThisWeek := float64(len(thisWeekPatientData))

	var patientPercentage float64
	if totalPatientLastWeek > 0 {
		patientPercentage = ((totalPatientThisWeek - totalPatientLastWeek) / totalPatientLastWeek) * 100
	} else if totalPatientLastWeek == 0 && totalPatientThisWeek > 0 {
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

	lastWeekTrasaction := time.Now().AddDate(0, 0, -14)
	lastWeekTrasactionData, err := repository.GetDoneTransactionsByDoctorAndWeek(doctor, lastWeekTrasaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get last month transaction data",
			"response": err.Error(),
		})
	}

	thisWeekTrasaction := time.Now().AddDate(0, 0, -7)
	thisWeekTrasactionData, err := repository.GetDoneTransactionsByDoctorAndWeek(doctor, thisWeekTrasaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get this month transaction data",
			"response": err.Error(),
		})
	}

	allTransactions := append(lastWeekTrasactionData, thisWeekTrasactionData...)

	totalPriceLastWeekTransaction := calculateTotalTransaction(lastWeekTrasactionData)
	totalPriceThisWeekTransaction := calculateTotalTransaction(thisWeekTrasactionData)

	var transactionPercentage float64
	if totalPriceLastWeekTransaction > 0 {
		transactionPercentage = ((totalPriceThisWeekTransaction - totalPriceLastWeekTransaction) / totalPriceLastWeekTransaction) * 100
	} else if totalPriceLastWeekTransaction > 0 {
		transactionPercentage = 100.0
	} else {
		transactionPercentage = 0.0
	}

	// Article

	lastWeekArticle := time.Now().AddDate(0, 0, -14)
	lastWeekArticleData, err := repository.DoctorGetAllArticlesByWeek(doctor, lastWeekArticle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get last week article data",
			"response": err.Error(),
		})
	}

	thisWeekArticle := time.Now().AddDate(0, 0, -7)
	thisWeekArticleData, err := repository.DoctorGetAllArticlesByWeek(doctor, thisWeekArticle)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message":  "failed get this week article data",
			"response": err.Error(),
		})
	}

	allArticles := append(lastWeekArticleData, thisWeekArticleData...)

	totalArticleLastWeek := calculateTotalArticle(lastWeekArticleData)
	totalArticleThisWeek := calculateTotalArticle(thisWeekArticleData)

	var articlePercentage float64
	if totalArticleLastWeek > 0 {
		articlePercentage = ((totalArticleThisWeek - totalArticleLastWeek) / totalArticleLastWeek) * 100
	} else if totalArticleLastWeek == 0 && totalArticleThisWeek > 0 {
		articlePercentage = 100.0
	} else {
		articlePercentage = 0.0
	}

	totalTransaction := len(transactionResponseData)
	totalConsultations := calculateTotalConsultation(allConsultations)
	totalTransactions := calculateTotalTransaction(allTransactions)
	totalPatients := calculateTotalPatient(allPatients)
	totalArticle := calculateTotalArticle(allArticles)

	return c.JSON(http.StatusOK, map[string]any{
		"message": "success get dashboard data for one week",
		// Consultation
		"totalConsultations":        totalConsultations,
		"totalConsultationLastWeek": totalConsultationLastWeek,
		"totalConsultationThisWeek": totalConsultationThisWeek,
		"consultationPercentage":    consultationPercentage,
		// Patients
		"totalPatients":        totalPatients,
		"totalPatientLastWeek": totalPatientLastWeek,
		"totalPatientThisWeek": totalPatientThisWeek,
		"patientPercentage":    patientPercentage,
		// Transaction
		"totalTransaction":      totalTransaction,
		"totalTransactions":     totalTransactions,
		"totalPriceLastWeek":    totalPriceLastWeekTransaction,
		"totalPriceThisWeek":    totalPriceThisWeekTransaction,
		"transactionPercentage": transactionPercentage,
		// Article
		"totalArticles":        totalArticle,
		"totalArticleLastWeek": totalArticleLastWeek,
		"totalArticleThisWeek": totalArticleThisWeek,
		"articlePercentage":    articlePercentage,
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

func calculateTotalArticle(articles []model.Article) float64 {
	totalArticle := 0.0
	for range articles {
		totalArticle += 1
	}
	return totalArticle
}
