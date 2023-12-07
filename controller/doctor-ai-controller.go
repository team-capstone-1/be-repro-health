// ai-controller.go

package controller

import (
	"capstone-project/dto"
	m "capstone-project/middleware"
	"capstone-project/model"
	"capstone-project/repository"
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DoctorAIController struct {
	DoctorAIRepo repository.DoctorAIRepository
}

func NewDoctorAIController(doctorAiRepo repository.DoctorAIRepository) *DoctorAIController {
	return &DoctorAIController{
		DoctorAIRepo: doctorAiRepo,
	}
}

func (ac *DoctorAIController) storeResultInDatabase(c echo.Context, doctorStoreDB model.DoctorHealthRecommendation) error {
	ac.DoctorAIRepo.DoctorStoreChatToDB(doctorStoreDB)
	return nil
}

func (ac *DoctorAIController) DoctorGetHealthRecommendation(c echo.Context) error {
	doctorID := m.ExtractTokenUserId(c)
	if doctorID == uuid.Nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":  "unauthorized",
			"response": "Permission Denied: Doctor is not valid.",
		})
	}

	req := dto.HealthRecommendationDoctorRequest{}
	errBind := c.Bind(&req)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "error bind data",
			"response": errBind.Error(),
		})
	}

	language := "id"

	if false {
		language = "en"
	}

	var sessionID uuid.UUID = req.SessionID
	var sessionExists bool
	fmt.Printf("sessionID: %v\n", sessionID)

	if isFirstQuestion(req.Message, sessionID) {
		fmt.Printf("sessionID: %v\n", sessionID)
		sessionExists = false
		sessionID = uuid.New()
	} else {
		sessionID, sessionExists = ac.getSessionIDFromDatabase(c, doctorID)
		fmt.Printf("sessionID: %v\n", sessionID)
	}

	if isNonReproductiveHealthDoctorQuestion(req.Message) {
		response := "Saya Emilia tidak bisa menjawab seputar hal diluar kesehatan reproduksi. Apakah ada pertanyaan lain yang berkaitan dengan kesehatan reproduksi? Atau mungkin anda bisa memakai kalimat dengan satu atau lebih kata kunci yang membuat saya bisa memahami pertanyaan anda hehehe..."
		resp := dto.DoctorHealthRecommendationResponse{
			SessionID: sessionID,
			Status:    "success",
			Data:      response,
		}
		return c.JSON(http.StatusOK, resp)
	} else {
		// Continue with the rest of the logic
		if !sessionExists {
			sessionID = uuid.New()
		} else {
			previousQuestion, err := ac.DoctorAIRepo.GetPreviousQuestion(context.Background(), sessionID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message":  "error getting previous question",
					"response": err.Error(),
				})
			}

			result, err := ac.DoctorAIRepo.DoctorGetHealthRecommendationWithContext(context.Background(), req.Message, previousQuestion, req.Message, language)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message":  "error getting health recommendations from AI",
					"response": err.Error(),
				})
			}

			doctorStoreDB := model.DoctorHealthRecommendation{
				ID:        uuid.New(),
				SessionID: sessionID,
				DoctorID:  doctorID,
				Question:  req.Message,
				Answer:    result,
			}
			ac.DoctorAIRepo.DoctorStoreChatToDB(doctorStoreDB)

			resp := dto.DoctorHealthRecommendationResponse{
				SessionID: sessionID,
				Status:    "success",
				Data:      result,
			}
			return c.JSON(http.StatusOK, resp)
		}
	}

	result, err := ac.DoctorAIRepo.DoctorGetHealthRecommendation(context.Background(), req.Message, language)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":  "error getting health recommendations from AI",
			"response": err.Error(),
		})
	}

	doctorStoreDB := model.DoctorHealthRecommendation{
		ID:        uuid.New(),
		SessionID: sessionID,
		DoctorID:  doctorID,
		Question:  req.Message,
		Answer:    result,
	}
	ac.DoctorAIRepo.DoctorStoreChatToDB(doctorStoreDB)

	resp := dto.DoctorHealthRecommendationResponse{
		SessionID: sessionID,
		Status:    "success",
		Data:      result,
	}

	return c.JSON(http.StatusOK, resp)
}

func (ac *DoctorAIController) getSessionIDFromDatabase(c echo.Context, doctorID uuid.UUID) (uuid.UUID, bool) {
	sessionID, err := ac.DoctorAIRepo.GetSessionIDFromDatabase(context.Background(), doctorID)
	if err == nil {
		return sessionID, true
	}
	return uuid.UUID{}, false
}

func isFirstQuestion(message string, sessionID uuid.UUID) bool {
	if sessionID == uuid.Nil {
		return true
	}

	return strings.Contains(strings.ToLower(message), "start")
}

func isNonReproductiveHealthDoctorQuestion(question string) bool {
	lowerQuestion := strings.ToLower(question)

	nonReproductiveKeywords := []string{
		"cuaca", "olahraga", "keuangan", "kesehatan umum", "teknologi", "hiburan", "politik", "sejarah", "wisata",
		"makanan", "musik", "film", "fashion", "pendidikan", "ilmu pengetahuan", "bisnis", "budaya", "seni", "agama",
		"lingkungan", "kebugaran", "gayahidup", "hewan peliharaan", "karier", "selebritas", "buku", "game", "rumah", "berkebun",
		"keluarga", "hubungan", "pengembangan diri", "motivasi", "teknologi", "media sosial", "peristiwa terkini", "proyek DIY", "masak", "fotografi",
		"perkawinan", "mobil", "aktivitas outdoor", "hobi", "keuangan pribadi", "kesehatan mental", "kesehatan jiwa", "yoga", "meditasi", "kesadaran diri",
		"travel", "kuliner", "pengetahuan umum", "teater", "seni rupa", "arsitektur", "mode", "literatur", "perjalanan", "teknologi terbaru",
		"penelitian sains", "strategi bisnis", "startup", "kepemimpinan", "strategi komunikasi", "destinasi wisata", "petualangan alam", "hewan liar",
		"lingkungan hidup", "keberlanjutan", "olahraga ekstrem", "pelatihan kebugaran", "olahraga air", "olahraga tim", "fotografi alam", "fotografi potret",
		"petualangan musik", "festival musik", "konser", "kehidupan keluarga", "parenting", "pendidikan anak", "hubungan orangtua-anak", "mainan edukatif",
		"pembangunan karir", "manajemen waktu", "strategi pekerjaan", "wawancara pekerjaan", "kehidupan kantoran", "kewirausahaan", "manajemen proyek",
		"bisnis online", "investasi", "keuangan keluarga", "pajak", "perencanaan pensiun", "investasi saham", "properti", "gadget terbaru", "aplikasi mobile",
		"keamanan siber", "privasi online", "kehidupan media sosial", "tren berita", "kebijakan pemerintah", "politik global", "sejarah dunia", "tokoh sejarah",
		"budaya dunia", "agama-agama dunia", "ritual keagamaan", "sosialisme", "pengembangan masyarakat", "pemberdayaan masyarakat", "aktivisme sosial",
		"volunteerism", "kegiatan amal", "gizi seimbang", "diet sehat", "hidup aktif", "olahraga indoor", "olahraga luar ruangan", "resep masakan",
		"tren kuliner", "kebugaran fisik", "kesehatan jantung", "penyakit kronis", "perawatan kulit", "tren fashion", "gayahidup minimalis", "desain interior",
		"rumah pintar", "teknologi hijau", "seni bela diri", "olahraga bertarung", "permainan papan", "permainan video", "film dokumenter", "film independen",
		"animasi", "seni pertunjukan", "seni patung", "seni lukis", "sastra", "puisi", "penulisan kreatif", "sosiologi", "masalah sosial", "lingkungan kerja",
		"keberlanjutan bisnis", "pemanfaatan teknologi", "analisis data", "perkembangan teknologi terbaru", "inovasi bisnis", "kewirausahaan teknologi",
		"teknologi pendidikan", "e-learning", "keamanan teknologi", "virtual reality", "teknologi terkini", "inovasi digital", "pemrograman komputer",
		"game development", "desain UX/UI", "pemrograman web", "teknologi cloud", "blockchain", "cryptocurrency", "trading saham", "pemasaran digital",
		"strategi konten", "strategi pemasaran", "branding", "manajemen merek", "strategi penjualan", "manajemen proyek konstruksi", "arsitektur modern",
		"desain arsitektur", "pemikiran filosofis", "etika", "hukum", "bahasa dan linguistik", "sastra klasik", "sastra kontemporer", "puisi modern",
		"penelitian sosial", "teori psikologi", "psikologi perkembangan", "motivasi belajar", "pendidikan inklusif", "inovasi pendidikan", "teknologi dalam pendidikan",
		"pendidikan STEM", "penelitian ilmiah", "penemuan ilmiah", "perjalanan ruang angkasa", "teknologi penerbangan", "petualangan alam", "petualangan laut",
		"petualangan gunung", "outdoor survival", "hobi memancing", "hobi memanjat", "hobi berkemah", "hobi bersepeda", "hobi fotografi", "hobi melukis",
		"hobi memasak", "hobi merajut", "hobi membuat kerajinan tangan", "perjalanan keluarga", "kegiatan keluarga", "permainan edukatif untuk anak-anak",
		"perkembangan anak-anak", "psikologi anak-anak", "pengasuhan anak-anak", "metode belajar anak-anak", "buku anak-anak", "teknologi pendidikan anak-anak",
		"teknologi pengembangan karir", "keterampilan profesional", "manajemen waktu", "pemecahan masalah", "komunikasi efektif", "kepemimpinan tim",
		"pengembangan tim", "pemahaman diri", "keseimbangan kehidupan kerja", "strategi manajemen stres", "kebahagiaan dan kepuasan kerja",
		"wisata kuliner", "wisata budaya", "wisata petualangan", "wisata alam", "pengalaman travel yang unik", "seni fotografi perjalanan", "travel hacking",
		"konsep keberlanjutan", "pengurangan jejak karbon", "pemanfaatan energi terbarukan", "perlindungan satwa liar", "penanaman pohon",
		"perubahan iklim dan dampaknya", "solusi ramah lingkungan", "masakan sehat", "resept makanan organik", "diet pescetarian", "diet pola makan",
		"manajemen berat badan", "rutinitas olahraga", "olahraga kekuatan", "olahraga kardiovaskular", "olahraga fleksibilitas", "olahraga mental",
		"teknik meditasi", "teknik mindfulness", "praktik yoga", "spiritualitas modern", "kehidupan pikiran dan tubuh", "kesehatan holistik", "pengobatan holistik",
		"pengobatan alternatif", "akupunktur", "refleksiologi", "pengobatan herbal", "penyembuhan kristal", "meditasi transcendental", "teknik pernapasan",
		"perjalanan spiritual", "kajian agama", "perbandingan agama", "teologi", "cerita keberhasilan startup", "strategi wirausaha", "panduan wirausaha",
		"peluang bisnis", "tren pasar", "analisis investasi", "manajemen keuangan", "perencanaan keuangan", "investasi properti", "reksadana", "asuransi jiwa",
		"asuransi kesehatan", "asuransi kendaraan", "manajemen risiko", "panduan perencanaan pensiun", "kesadaran kesehatan mental", "teknik manajemen stres",
		"teknik pereda kecemasan", "dukungan depresi", "praktik perawatan diri", "rutinitas kebugaran", "kebiasaan makan sehat", "resept superfood",
		"meditasi dan kesadaran", "spiritualitas", "keterkaitan pikiran dan tubuh", "keseluruhan kesehatan", "ramuan alami", "pendekatan pengobatan alternatif",
		"praktik penyembuhan tradisional", "proyek DIY rumah", "ide organisasi rumah", "tips decluttering", "gaya hidup minimalis", "teknologi rumah pintar",
		"pengasuhan anak-anak", "kegiatan ikatan keluarga", "permainan edukatif untuk anak-anak", "sumber belajar", "bantuan pekerjaan rumah",
		"nasihat hubungan", "komunikasi dalam hubungan", "tips cinta dan romansa", "nasihat kencan", "tips pernikahan", "buku pengembangan diri",
		"strategi pertumbuhan pribadi", "kutipan motivasi", "kisah sukses", "menetapkan tujuan", "teknologi dan dampak masyarakat", "pertimbangan etika dalam teknologi",
		"masalah privasi data", "tindakan keamanan siber", "tren budaya internet", "dampak media sosial", "analisis berita terkini", "pembaruan berita dunia",
		"politik", "kebijakan pemerintah", "insight peristiwa sejarah", "eksploitasi tokoh sejarah", "pemahaman warisan budaya", "wawasan agama-agama dunia",
		"eksplorasi praktik keagamaan", "pelestarian lingkungan", "tips gaya hidup berkelanjutan", "eksplorasi inisiatif hijau", "pelestarian satwa liar",
		"tips hidup sehat", "tips gizi", "tren diet", "nasihat manajemen berat badan", "eksplorasi rutinitas olahraga", "ide dekorasi rumah",
		"tren desain interior", "inspirasi proyek rumah DIY", "tips berkebun", "eksplorasi aktivitas outdoor", "nasihat pengasuhan",
		"pemahaman milestone perkembangan anak", "eksplorasi tips pendidikan", "penemuan sumber belajar", "nasihat pengembangan karir", "advice pertumbuhan profesional",
		"peningkatan produktivitas di tempat kerja", "keterampilan kepemimpinan", "kegiatan membangun tim", "peningkatan strategi komunikasi", "eksplorasi destinasi wisata",
		"pengalaman petualangan travel", "penemuan pengalaman budaya", "petualangan eksplorasi alam", "inspirasi fotografi satwa liar", "penyadaran masalah lingkungan",
		"tips gaya hidup berkelanjutan", "penerapan praktik hidup hijau", "promosi kesadaran perubahan iklim", "tren teknologi terkini", "eksplorasi inovasi",
		"inspirasi cerita sukses startup", "penerapan tips wirausaha", "pemahaman strategi bisnis", "eksplorasi tren pasar", "penemuan peluang investasi",
		"penasihat perencanaan keuangan", "advice perencanaan pensiun", "tips manajemen kredit",
	}

	for _, keyword := range nonReproductiveKeywords {
		if strings.Contains(lowerQuestion, keyword) {
			return true
		}
	}

	reproductivePatterns := []string{
		"hamil", "kelahiran", "menstruasi", "ovulasi", "fertilisasi", "sel telur", "sperma", "kondom", "kb", "ovarium",
		"testis", "rahim", "serviks", "vagina", "penis", "kehamilan", "fetus", "plasenta", "melahirkan", "menopause",
		"menstruasi", "keputihan", "endometriosis", "polip", "mioma", "pcos", "infertilitas", "seksualitas", "disfungsi ereksi",
		"kanker payudara", "kanker serviks", "kanker ovarium", "kanker prostat", "kondisi menular seksual", "HIV", "AIDS",
		"kutil kelamin", "herpes genital", "sifilis", "gonore", "klamidia", "hepatitis B", "kondisi urologi", "disfungsi seksual",
		"penyakit menular seksual", "hiv", "kesehatan reproduksi", "konseling reproduksi", "anatomi reproduksi", "hubungan seksual",
		"kebugaran reproduksi", "penyakit menular seksual", "perawatan kehamilan", "kesehatan seksual", "konseling reproduksi",
		"keamanan reproduksi", "pencegahan kehamilan", "pemahaman seksual", "kesuburan", "pemahaman perawatan kesuburan",
		"pemahaman kehamilan", "pemahaman infertilitas", "pemahaman kondisi menular seksual", "perawatan prenatal", "perawatan postnatal",
		"kebutuhan kesehatan reproduksi", "pemahaman kontrasepsi", "kesehatan dan kebersihan reproduksi", "kesehatan reproduksi remaja",
		"pemahaman menopause", "perawatan menopause", "pemahaman endometriosis", "pemahaman mioma", "pemahaman polip", "pemahaman disfungsi seksual",
		"pemahaman kanker payudara", "pemahaman kanker serviks", "pemahaman kanker ovarium", "pemahaman kanker prostat",
		"pemahaman penyakit menular seksual", "pemahaman HIV", "pemahaman AIDS", "pemahaman kutil kelamin", "pemahaman herpes genital",
		"pemahaman sifilis", "pemahaman gonore", "pemahaman klamidia", "pemahaman hepatitis B", "pemahaman kondisi urologi",
		"kesehatan reproduksi pria", "kesehatan reproduksi wanita", "pemeriksaan kesehatan reproduksi", "faktor risiko infertilitas",
		"kesehatan rahim", "kesehatan testis", "kesehatan ovarium", "penyakit menular seksual pada remaja", "pengobatan infertilitas",
		"pola makan untuk kesuburan", "bahaya rokok terhadap kesuburan", "peran hormon dalam reproduksi", "kesehatan reproduksi remaja perempuan",
		"kesehatan reproduksi remaja laki-laki", "bahaya alkohol terhadap reproduksi", "merawat kesehatan reproduksi", "kesehatan seksual LGBTQ+",
		"pertanyaan seputar kesehatan reproduksi", "merencanakan kehamilan", "tanda-tanda kehamilan", "pemantauan kesehatan reproduksi",
		"bahaya merokok pada kehamilan", "nutrisi untuk kesuburan", "pemahaman reproduksi manusia", "faktor risiko kanker prostat",
		"kehamilan pada usia muda", "kehamilan di usia tua", "faktor risiko kanker payudara", "tanda-tanda penyakit menular seksual",
		"kehamilan trimester pertama", "kehamilan trimester kedua", "kehamilan trimester ketiga", "manfaat menyusui", "penyebab keguguran",
		"faktor risiko endometriosis", "pengobatan mioma", "risiko kanker serviks", "pengobatan kanker ovarium", "pengobatan kanker prostat",
		"mencegah infeksi saluran reproduksi", "pemahaman aspek psikologis reproduksi", "pemahaman perubahan hormonal pada reproduksi",
		"penyakit menular seksual pada ibu hamil", "kesehatan reproduksi masyarakat", "tanda-tanda gangguan reproduksi", "pengaruh olahraga pada kesuburan",
		"pemahaman tentang konsepsi", "tanda-tanda gangguan hormonal pada reproduksi", "pemahaman tentang kesuburan",
	}

	reproductiveRegex := regexp.MustCompile(strings.Join(reproductivePatterns, "|"))

	return !reproductiveRegex.MatchString(lowerQuestion)
}

func (ac *DoctorAIController) GetHealthRecommendationDoctorHistory(c echo.Context) error {
	doctorIDParam := c.Param("doctor_id")
	fmt.Println("Doctor ID from URL:", doctorIDParam)
	doctorID, err := uuid.Parse(doctorIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "error parse doctor id",
			"response": err.Error(),
		})
	}

	responseData, err := ac.DoctorAIRepo.DoctorGetAllHealthRecommendationsByDoctorID(doctorID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "failed get healthRecommendations",
			"response": err.Error(),
		})
	}

	var healthRecommendationResponse []dto.HealthRecommendationHistoryDoctorResponse
	for _, healthRecommendation := range responseData {
		healthRecommendationResponse = append(healthRecommendationResponse, dto.ConvertToHealthRecommendationHistoryDoctorResponse(healthRecommendation))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get healthRecommendations",
		"response": healthRecommendationResponse,
	})
}

func (ac *DoctorAIController) GetHealthRecommendationDoctorHistoryFromSession(c echo.Context) error {
	sessionIDParam := c.Param("session_id")
	fmt.Println("Session ID from URL:", sessionIDParam)
	sessionID, err := uuid.Parse(sessionIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "error parse session id",
			"response": err.Error(),
		})
	}

	responseData, err := ac.DoctorAIRepo.DoctorGetAllHealthRecommendationsBySession(sessionID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":  "failed get healthRecommendations",
			"response": err.Error(),
		})
	}

	var healthRecommendationResponse []dto.HealthRecommendationHistoryDoctorResponse
	for _, healthRecommendation := range responseData {
		healthRecommendationResponse = append(healthRecommendationResponse, dto.ConvertToHealthRecommendationHistoryDoctorResponse(healthRecommendation))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get healthRecommendations",
		"response": healthRecommendationResponse,
	})
}