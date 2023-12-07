// // ai-controller.go

package controller

// import (
// 	"capstone-project/dto"
// 	m "capstone-project/middleware"
// 	"capstone-project/model"
// 	"capstone-project/repository"
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"regexp"
// 	"strings"

// 	"github.com/google/uuid"
// 	"github.com/labstack/echo/v4"
// )

// type PatientAIController struct {
// 	PatientAIRepo repository.PatientAIRepository
// }

// func NewPatientAIController(patientAIRepo repository.PatientAIRepository) *PatientAIController {
// 	return &PatientAIController{
// 		PatientAIRepo: patientAIRepo,
// 	}
// }

// func getCategoryFromQuestion(question string) string {
// 	if strings.Contains(strings.ToLower(question), "janji temu") {
// 		return "janji temu"
// 	} else if strings.Contains(strings.ToLower(question), "artikel") {
// 		return "artikel"
// 	} else if strings.Contains(strings.ToLower(question), "forum") {
// 		return "forum"
// 	} else if strings.Contains(strings.ToLower(question), "riwayat") {
// 		return "riwayat"
// 	} else if strings.Contains(strings.ToLower(question), "profile") {
// 		return "profile"
// 	}
// 	return "lainnya"
// }

// func (ac *PatientAIController) GetHealthRecommendation(c echo.Context) error {
// 	userID := m.ExtractTokenUserId(c)
// 	if userID == uuid.Nil {
// 		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
// 			"message":  "unauthorized",
// 			"response": "Permission Denied: User is not valid.",
// 		})
// 	}

// 	req := dto.HealthRecommendationRequest{}
// 	errBind := c.Bind(&req)
// 	if errBind != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message":  "error bind data",
// 			"response": errBind.Error(),
// 		})
// 	}

// 	language := "id"

// 	if false {
// 		language = "en"
// 	}

// 	var PatientSessionID uuid.UUID = req.PatientSessionID
// 	var sessionExists bool
// 	fmt.Printf("sessionID: %v\n", PatientSessionID)

// 	if isFirstQuestion(req.Message, PatientSessionID) {
// 		fmt.Printf("sessionID: %v\n", PatientSessionID)
// 		sessionExists = false
// 		PatientSessionID = uuid.New()
// 	} else {
// 		PatientSessionID, sessionExists = ac.getPatientSessionIDFromDatabase(c, userID)
// 		fmt.Printf("sessionID: %v\n", PatientSessionID)
// 	}

// 	if isNonReproductiveHealthQuestion(req.Message) {
// 		response := "Saya Emilia tidak bisa menjawab seputar hal diluar kesehatan reproduksi. Apakah ada pertanyaan lain yang berkaitan dengan kesehatan reproduksi?"
// 		resp := dto.HealthRecommendationResponse{
// 			Status: "success",
// 			Data:   response,
// 		}
// 		return c.JSON(http.StatusOK, resp)
// 	} else {
// 		// Continue with the rest of the logic
// 		if !sessionExists {
// 			PatientSessionID = uuid.New()
// 		} else {
// 			previousQuestion, err := ac.PatientAIRepo.GetPreviousQuestion(context.Background(), sessionID)
// 			if err != nil {
// 				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 					"message":  "error getting previous question",
// 					"response": err.Error(),
// 				})
// 			}

// 			result, err := ac.PatientAIRepo.PatientGetHealthRecommendationWithContext(context.Background(), req.Message, previousQuestion, req.Message, language)
// 			if err != nil {
// 				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 					"message":  "error getting health recommendations from AI",
// 					"response": err.Error(),
// 				})
// 			}

// 			doctorStoreDB := model.DoctorHealthRecommendation{
// 				ID:        uuid.New(),
// 				SessionID: sessionID,
// 				DoctorID:  doctorID,
// 				Question:  req.Message,
// 				Answer:    result,
// 			}
// 			ac.PatientAIRepo.DoctorStoreChatToDB(doctorStoreDB)

// 			resp := dto.DoctorHealthRecommendationResponse{
// 				SessionID: sessionID,
// 				Status:    "success",
// 				Data:      result,
// 			}
// 			return c.JSON(http.StatusOK, resp)
// 		}
// 	}

// 	questionCategory := getCategoryFromQuestion(req.Message)
// 	var aiResponse string

// 	if questionCategory == "janji temu" {
// 		if strings.Contains(strings.ToLower(req.Message), "batalkan") {
// 			aiResponse = "Untuk membatalkan janji temu, silakan ikuti langkah berikut:\n1. Buka aplikasi kami.\n2. Pilih menu 'Janji Temu'.\n3. Temukan janji temu yang ingin dibatalkan.\n4. Pilih opsi 'Batalkan Janji Temu'.\nTerima kasih."
// 		} else {
// 			aiResponse = "Anda belum membuat jadwal dengan dokter. Silakan membuat jadwal terlebih dahulu."
// 		}
// 	} else if questionCategory == "artikel" {
// 		// Check for specific questions related to the "artikel" category
// 		if strings.Contains(strings.ToLower(req.Message), "buka menu artikel") {
// 			aiResponse = "Untuk membuka menu artikel, Anda dapat:\n1. Masuk ke aplikasi kami.\n2. Pilih menu 'Artikel'.\n3. Temukan artikel yang ingin Anda baca.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "menjawab artikel") {
// 			aiResponse = "Untuk menjawab artikel, Anda dapat:\n1. Buka artikel yang ingin Anda jawab.\n2. Temukan bagian komentar atau jawaban.\n3. Tulis komentar atau jawaban Anda.\n4. Tekan tombol 'Kirim'.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "siapa yang bisa membuat artikel") {
// 			aiResponse = "Hanya dokter yang dapat membuat artikel. Jika Anda memiliki pertanyaan atau topik yang ingin diangkat, Anda dapat berbicara langsung dengan dokter Anda.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "apa yang tidak bisa dilakukan pengguna pada artikel") {
// 			aiResponse = "Pengguna dapat membaca artikel, memberikan komentar, dan bertanya kepada dokter terkait artikel. Namun, pengguna tidak dapat membuat artikel sendiri. Jika Anda memiliki konten yang ingin dibagikan, sampaikan kepada dokter Anda.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "pertanyaan artikel lainnya") {
// 			aiResponse = "Tentu, apa lagi yang ingin Anda ketahui tentang artikel? Saya siap membantu."
// 		} else {
// 			aiResponse = "Maaf, saya tidak mengerti pertanyaan Anda terkait artikel. Bisakah Anda memberikan informasi lebih lanjut?"
// 		}
// 	} else if questionCategory == "forum" {
// 		// Check for specific questions related to the "forum" category
// 		if strings.Contains(strings.ToLower(req.Message), "cara membuat forum") {
// 			aiResponse = "Untuk membuat forum, Anda dapat:\n1. Masuk ke aplikasi kami.\n2. Pilih menu 'Forum'.\n3. Pilih opsi 'Buat Forum'.\n4. Tulis pertanyaan atau topik Anda.\n5. Tekan tombol 'Kirim'.\nDokter kami akan segera merespons. Terima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "siapa yang bisa membuat forum") {
// 			aiResponse = "Hanya pengguna yang dapat membuat forum. Jika Anda memiliki pertanyaan atau topik yang ingin dibahas, silakan membuat forum dan dokter kami akan meresponsnya.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "cara dokter menjawab forum") {
// 			aiResponse = "Dokter dapat merespons forum yang dibuat oleh pengguna. Jika Anda ingin mendapatkan jawaban dari dokter, buatlah forum dan mereka akan segera merespons.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "pertanyaan forum lainnya") {
// 			aiResponse = "Tentu, apa lagi yang ingin Anda ketahui tentang forum? Saya siap membantu."
// 		} else {
// 			aiResponse = "Maaf, saya tidak mengerti pertanyaan Anda terkait forum. Bisakah Anda memberikan informasi lebih lanjut?"
// 		}
// 	} else if questionCategory == "riwayat" {
// 		// Check for specific questions related to the "riwayat" (transaction history) category
// 		if strings.Contains(strings.ToLower(req.Message), "bagaimana cara melihat riwayat transaksi saya") {
// 			aiResponse = "Untuk melihat riwayat transaksi Anda, Anda dapat:\n1. Masuk ke akun Anda.\n2. Pilih menu 'Riwayat Transaksi'.\n3. Anda akan melihat daftar transaksi yang telah Anda lakukan.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "apa yang termasuk dalam riwayat transaksi") {
// 			aiResponse = "Riwayat transaksi mencakup semua transaksi keuangan yang telah Anda lakukan, seperti pembayaran layanan atau produk kesehatan. Ini membantu Anda melacak pengeluaran dan layanan yang telah digunakan.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "bisa dihapus tidak riwayat transaksi saya") {
// 			aiResponse = "Sayangnya, riwayat transaksi tidak dapat dihapus. Ini penting untuk memastikan transparansi dan akurasi informasi keuangan Anda.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "berapa lama riwayat transaksi disimpan") {
// 			aiResponse = "Riwayat transaksi Anda akan disimpan dalam sistem dengan aman. Kami menjaga privasi data dan memastikan bahwa informasi keuangan Anda terlindungi.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "pertanyaan riwayat transaksi lainnya") {
// 			aiResponse = "Jika Anda memiliki pertanyaan lebih lanjut atau ingin informasi tambahan tentang riwayat transaksi Anda, beri tahu saya. Saya siap membantu.\nTerima kasih."
// 		} else {
// 			aiResponse = "Maaf, saya tidak mengerti pertanyaan Anda terkait riwayat transaksi. Bisakah Anda memberikan informasi lebih lanjut?"
// 		}
// 	} else if questionCategory == "profile" {
// 		// Check for specific questions related to the "profile" category
// 		if strings.Contains(strings.ToLower(req.Message), "bagaimana cara melihat profil saya") {
// 			aiResponse = "Untuk melihat profil Anda, Anda dapat:\n1. Masuk ke akun Anda.\n2. Pilih menu 'Profil' atau 'Akun Saya'.\n3. Anda akan melihat informasi profil Anda, termasuk detail pribadi dan pengaturan akun lainnya.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "apa yang dapat saya ubah di profil saya") {
// 			aiResponse = "Anda dapat mengubah beberapa informasi di profil Anda, seperti foto profil, alamat, nomor telepon, dan preferensi lainnya. Pastikan untuk memeriksa opsi pengaturan di menu profil untuk melakukan perubahan.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "bisa dihapus tidak foto profil saya") {
// 			aiResponse = "Ya, Anda dapat menghapus atau mengganti foto profil Anda. Lakukan langkah-langkah berikut:\n1. Masuk ke akun Anda.\n2. Buka menu 'Profil' atau 'Akun Saya'.\n3. Temukan opsi untuk mengganti atau menghapus foto profil Anda.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "apa informasi wajib di profil") {
// 			aiResponse = "Informasi wajib di profil biasanya melibatkan data dasar seperti nama lengkap, tanggal lahir, dan alamat. Pastikan untuk melengkapi informasi tersebut agar kami dapat memberikan layanan yang lebih baik.\nTerima kasih."
// 		} else if strings.Contains(strings.ToLower(req.Message), "pertanyaan profil lainnya") {
// 			aiResponse = "Jika Anda memiliki pertanyaan lebih lanjut atau butuh bantuan terkait profil Anda, beri tahu saya. Saya siap membantu.\nTerima kasih."
// 		} else {
// 			aiResponse = "Maaf, saya tidak mengerti pertanyaan Anda terkait profil. Bisakah Anda memberikan informasi lebih lanjut?"
// 		}
// 	} else {
// 		aiResponse = "Maaf, saya tidak mengerti pertanyaan Anda. Bisakah Anda memberikan informasi lebih lanjut?"
// 	}

// 	result, err := ac.AIRepo.GetHealthRecommendation(context.Background(), req.Message, language)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"message":  "error getting health recommendations from AI",
// 			"response": err.Error(),
// 		})
// 	}

// 	storeDB := model.HealthRecommendation{
// 		ID:        uuid.New(),
// 		PatientID: req.PatientID,
// 		Question:  req.Message,
// 		Answer:    result,
// 	}
// 	ac.AIRepo.StoreChatToDB(storeDB)

// 	resp := dto.HealthRecommendationResponse{
// 		Status: "success",
// 		Data:   aiResponse,
// 	}

// 	return c.JSON(http.StatusOK, resp)
// }

// func (ac *PatientAIController) getPatientSessionIDFromDatabase(c echo.Context, userID uuid.UUID) (uuid.UUID, bool) {
// 	sessionID, err := ac.PatientAIRepo.GetPatientSessionIDFromDatabase(context.Background(), userID)
// 	if err == nil {
// 		return sessionID, true
// 	}
// 	return uuid.UUID{}, false
// }

// func isNonReproductiveHealthQuestion(question string) bool {
// 	lowerQuestion := strings.ToLower(question)

// 	nonReproductiveKeywords := []string{
// 		"cuaca", "olahraga", "keuangan", "kesehatan umum", "teknologi", "hiburan", "politik", "sejarah", "wisata",
// 		"makanan", "musik", "film", "fashion", "pendidikan", "ilmu pengetahuan", "bisnis", "budaya", "seni", "agama",
// 		"lingkungan", "kebugaran", "gayahidup", "hewan peliharaan", "karier", "selebritas", "buku", "game", "rumah", "berkebun",
// 		"keluarga", "hubungan", "pengembangan diri", "motivasi", "teknologi", "media sosial", "peristiwa terkini", "proyek DIY", "masak", "fotografi",
// 		"perkawinan", "mobil", "aktivitas outdoor", "hobi", "keuangan pribadi", "kesehatan mental", "kesehatan jiwa", "yoga", "meditasi", "kesadaran diri",
// 		"travel", "kuliner", "pengetahuan umum", "teater", "seni rupa", "arsitektur", "mode", "literatur", "perjalanan", "teknologi terbaru",
// 		"penelitian sains", "strategi bisnis", "startup", "kepemimpinan", "strategi komunikasi", "destinasi wisata", "petualangan alam", "hewan liar",
// 		"lingkungan hidup", "keberlanjutan", "olahraga ekstrem", "pelatihan kebugaran", "olahraga air", "olahraga tim", "fotografi alam", "fotografi potret",
// 		"petualangan musik", "festival musik", "konser", "kehidupan keluarga", "parenting", "pendidikan anak", "hubungan orangtua-anak", "mainan edukatif",
// 		"pembangunan karir", "manajemen waktu", "strategi pekerjaan", "wawancara pekerjaan", "kehidupan kantoran", "kewirausahaan", "manajemen proyek",
// 		"bisnis online", "investasi", "keuangan keluarga", "pajak", "perencanaan pensiun", "investasi saham", "properti", "gadget terbaru", "aplikasi mobile",
// 		"keamanan siber", "privasi online", "kehidupan media sosial", "tren berita", "kebijakan pemerintah", "politik global", "sejarah dunia", "tokoh sejarah",
// 		"budaya dunia", "agama-agama dunia", "ritual keagamaan", "sosialisme", "pengembangan masyarakat", "pemberdayaan masyarakat", "aktivisme sosial",
// 		"volunteerism", "kegiatan amal", "gizi seimbang", "diet sehat", "hidup aktif", "olahraga indoor", "olahraga luar ruangan", "resep masakan",
// 		"tren kuliner", "kebugaran fisik", "kesehatan jantung", "penyakit kronis", "perawatan kulit", "tren fashion", "gayahidup minimalis", "desain interior",
// 		"rumah pintar", "teknologi hijau", "seni bela diri", "olahraga bertarung", "permainan papan", "permainan video", "film dokumenter", "film independen",
// 		"animasi", "seni pertunjukan", "seni patung", "seni lukis", "sastra", "puisi", "penulisan kreatif", "sosiologi", "masalah sosial", "lingkungan kerja",
// 		"keberlanjutan bisnis", "pemanfaatan teknologi", "analisis data", "perkembangan teknologi terbaru", "inovasi bisnis", "kewirausahaan teknologi",
// 		"teknologi pendidikan", "e-learning", "keamanan teknologi", "virtual reality", "teknologi terkini", "inovasi digital", "pemrograman komputer",
// 		"game development", "desain UX/UI", "pemrograman web", "teknologi cloud", "blockchain", "cryptocurrency", "trading saham", "pemasaran digital",
// 		"strategi konten", "strategi pemasaran", "branding", "manajemen merek", "strategi penjualan", "manajemen proyek konstruksi", "arsitektur modern",
// 		"desain arsitektur", "pemikiran filosofis", "etika", "hukum", "bahasa dan linguistik", "sastra klasik", "sastra kontemporer", "puisi modern",
// 		"penelitian sosial", "teori psikologi", "psikologi perkembangan", "motivasi belajar", "pendidikan inklusif", "inovasi pendidikan", "teknologi dalam pendidikan",
// 		"pendidikan STEM", "penelitian ilmiah", "penemuan ilmiah", "perjalanan ruang angkasa", "teknologi penerbangan", "petualangan alam", "petualangan laut",
// 		"petualangan gunung", "outdoor survival", "hobi memancing", "hobi memanjat", "hobi berkemah", "hobi bersepeda", "hobi fotografi", "hobi melukis",
// 		"hobi memasak", "hobi merajut", "hobi membuat kerajinan tangan", "perjalanan keluarga", "kegiatan keluarga", "permainan edukatif untuk anak-anak",
// 		"perkembangan anak-anak", "psikologi anak-anak", "pengasuhan anak-anak", "metode belajar anak-anak", "buku anak-anak", "teknologi pendidikan anak-anak",
// 		"teknologi pengembangan karir", "keterampilan profesional", "manajemen waktu", "pemecahan masalah", "komunikasi efektif", "kepemimpinan tim",
// 		"pengembangan tim", "pemahaman diri", "keseimbangan kehidupan kerja", "strategi manajemen stres", "kebahagiaan dan kepuasan kerja",
// 		"wisata kuliner", "wisata budaya", "wisata petualangan", "wisata alam", "pengalaman travel yang unik", "seni fotografi perjalanan", "travel hacking",
// 		"konsep keberlanjutan", "pengurangan jejak karbon", "pemanfaatan energi terbarukan", "perlindungan satwa liar", "penanaman pohon",
// 		"perubahan iklim dan dampaknya", "solusi ramah lingkungan", "masakan sehat", "resept makanan organik", "diet pescetarian", "diet pola makan",
// 		"manajemen berat badan", "rutinitas olahraga", "olahraga kekuatan", "olahraga kardiovaskular", "olahraga fleksibilitas", "olahraga mental",
// 		"teknik meditasi", "teknik mindfulness", "praktik yoga", "spiritualitas modern", "kehidupan pikiran dan tubuh", "kesehatan holistik", "pengobatan holistik",
// 		"pengobatan alternatif", "akupunktur", "refleksiologi", "pengobatan herbal", "penyembuhan kristal", "meditasi transcendental", "teknik pernapasan",
// 		"perjalanan spiritual", "kajian agama", "perbandingan agama", "teologi", "cerita keberhasilan startup", "strategi wirausaha", "panduan wirausaha",
// 		"peluang bisnis", "tren pasar", "analisis investasi", "manajemen keuangan", "perencanaan keuangan", "investasi properti", "reksadana", "asuransi jiwa",
// 		"asuransi kesehatan", "asuransi kendaraan", "manajemen risiko", "panduan perencanaan pensiun", "kesadaran kesehatan mental", "teknik manajemen stres",
// 		"teknik pereda kecemasan", "dukungan depresi", "praktik perawatan diri", "rutinitas kebugaran", "kebiasaan makan sehat", "resept superfood",
// 		"meditasi dan kesadaran", "spiritualitas", "keterkaitan pikiran dan tubuh", "keseluruhan kesehatan", "ramuan alami", "pendekatan pengobatan alternatif",
// 		"praktik penyembuhan tradisional", "proyek DIY rumah", "ide organisasi rumah", "tips decluttering", "gaya hidup minimalis", "teknologi rumah pintar",
// 		"pengasuhan anak-anak", "kegiatan ikatan keluarga", "permainan edukatif untuk anak-anak", "sumber belajar", "bantuan pekerjaan rumah",
// 		"nasihat hubungan", "komunikasi dalam hubungan", "tips cinta dan romansa", "nasihat kencan", "tips pernikahan", "buku pengembangan diri",
// 		"strategi pertumbuhan pribadi", "kutipan motivasi", "kisah sukses", "menetapkan tujuan", "teknologi dan dampak masyarakat", "pertimbangan etika dalam teknologi",
// 		"masalah privasi data", "tindakan keamanan siber", "tren budaya internet", "dampak media sosial", "analisis berita terkini", "pembaruan berita dunia",
// 		"politik", "kebijakan pemerintah", "insight peristiwa sejarah", "eksploitasi tokoh sejarah", "pemahaman warisan budaya", "wawasan agama-agama dunia",
// 		"eksplorasi praktik keagamaan", "pelestarian lingkungan", "tips gaya hidup berkelanjutan", "eksplorasi inisiatif hijau", "pelestarian satwa liar",
// 		"tips hidup sehat", "tips gizi", "tren diet", "nasihat manajemen berat badan", "eksplorasi rutinitas olahraga", "ide dekorasi rumah",
// 		"tren desain interior", "inspirasi proyek rumah DIY", "tips berkebun", "eksplorasi aktivitas outdoor", "nasihat pengasuhan",
// 		"pemahaman milestone perkembangan anak", "eksplorasi tips pendidikan", "penemuan sumber belajar", "nasihat pengembangan karir", "advice pertumbuhan profesional",
// 		"peningkatan produktivitas di tempat kerja", "keterampilan kepemimpinan", "kegiatan membangun tim", "peningkatan strategi komunikasi", "eksplorasi destinasi wisata",
// 		"pengalaman petualangan travel", "penemuan pengalaman budaya", "petualangan eksplorasi alam", "inspirasi fotografi satwa liar", "penyadaran masalah lingkungan",
// 		"tips gaya hidup berkelanjutan", "penerapan praktik hidup hijau", "promosi kesadaran perubahan iklim", "tren teknologi terkini", "eksplorasi inovasi",
// 		"inspirasi cerita sukses startup", "penerapan tips wirausaha", "pemahaman strategi bisnis", "eksplorasi tren pasar", "penemuan peluang investasi",
// 		"penasihat perencanaan keuangan", "advice perencanaan pensiun", "tips manajemen kredit",
// 	}

// 	for _, keyword := range nonReproductiveKeywords {
// 		if strings.Contains(lowerQuestion, keyword) {
// 			return true
// 		}
// 	}

// 	reproductivePatterns := []string{
// 		"hamil", "kelahiran", "menstruasi", "ovulasi", "fertilisasi", "sel telur", "sperma", "kondom", "kb", "ovarium",
// 		"testis", "rahim", "serviks", "vagina", "penis", "kehamilan", "fetus", "plasenta", "melahirkan", "menopause",
// 		"menstruasi", "keputihan", "endometriosis", "polip", "mioma", "pcos", "infertilitas", "seksualitas", "disfungsi ereksi",
// 		"kanker payudara", "kanker serviks", "kanker ovarium", "kanker prostat", "kondisi menular seksual", "HIV", "AIDS",
// 		"kutil kelamin", "herpes genital", "sifilis", "gonore", "klamidia", "hepatitis B", "kondisi urologi", "disfungsi seksual",
// 		"penyakit menular seksual", "hiv", "kesehatan reproduksi", "konseling reproduksi", "anatomi reproduksi", "hubungan seksual",
// 		"kebugaran reproduksi", "penyakit menular seksual", "perawatan kehamilan", "kesehatan seksual", "konseling reproduksi",
// 		"keamanan reproduksi", "pencegahan kehamilan", "pemahaman seksual", "kesuburan", "pemahaman perawatan kesuburan",
// 		"pemahaman kehamilan", "pemahaman infertilitas", "pemahaman kondisi menular seksual", "perawatan prenatal", "perawatan postnatal",
// 		"kebutuhan kesehatan reproduksi", "pemahaman kontrasepsi", "kesehatan dan kebersihan reproduksi", "kesehatan reproduksi remaja",
// 		"pemahaman menopause", "perawatan menopause", "pemahaman endometriosis", "pemahaman mioma", "pemahaman polip", "pemahaman disfungsi seksual",
// 		"pemahaman kanker payudara", "pemahaman kanker serviks", "pemahaman kanker ovarium", "pemahaman kanker prostat",
// 		"pemahaman penyakit menular seksual", "pemahaman HIV", "pemahaman AIDS", "pemahaman kutil kelamin", "pemahaman herpes genital",
// 		"pemahaman sifilis", "pemahaman gonore", "pemahaman klamidia", "pemahaman hepatitis B", "pemahaman kondisi urologi",
// 		"kesehatan reproduksi pria", "kesehatan reproduksi wanita", "pemeriksaan kesehatan reproduksi", "faktor risiko infertilitas",
// 		"kesehatan rahim", "kesehatan testis", "kesehatan ovarium", "penyakit menular seksual pada remaja", "pengobatan infertilitas",
// 		"pola makan untuk kesuburan", "bahaya rokok terhadap kesuburan", "peran hormon dalam reproduksi", "kesehatan reproduksi remaja perempuan",
// 		"kesehatan reproduksi remaja laki-laki", "bahaya alkohol terhadap reproduksi", "merawat kesehatan reproduksi", "kesehatan seksual LGBTQ+",
// 		"pertanyaan seputar kesehatan reproduksi", "merencanakan kehamilan", "tanda-tanda kehamilan", "pemantauan kesehatan reproduksi",
// 		"bahaya merokok pada kehamilan", "nutrisi untuk kesuburan", "pemahaman reproduksi manusia", "faktor risiko kanker prostat",
// 		"kehamilan pada usia muda", "kehamilan di usia tua", "faktor risiko kanker payudara", "tanda-tanda penyakit menular seksual",
// 		"kehamilan trimester pertama", "kehamilan trimester kedua", "kehamilan trimester ketiga", "manfaat menyusui", "penyebab keguguran",
// 		"faktor risiko endometriosis", "pengobatan mioma", "risiko kanker serviks", "pengobatan kanker ovarium", "pengobatan kanker prostat",
// 		"mencegah infeksi saluran reproduksi", "pemahaman aspek psikologis reproduksi", "pemahaman perubahan hormonal pada reproduksi",
// 		"penyakit menular seksual pada ibu hamil", "kesehatan reproduksi masyarakat", "tanda-tanda gangguan reproduksi", "pengaruh olahraga pada kesuburan",
// 		"pemahaman tentang konsepsi", "tanda-tanda gangguan hormonal pada reproduksi", "pemahaman tentang kesuburan",
// 	}

// 	reproductiveRegex := regexp.MustCompile(strings.Join(reproductivePatterns, "|"))

// 	if reproductiveRegex.MatchString(lowerQuestion) {
// 		return false
// 	}

// 	return true
// }

// func (ac *UserPatientAIController) GetHealthRecommendationHistory(c echo.Context) error {
// 	uuid, err := uuid.Parse(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]any{
// 			"message":  "error parse id",
// 			"response": err.Error(),
// 		})
// 	}

// 	responseData, err := ac.AIRepo.GetAllHealthRecommendations(uuid)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]any{
// 			"message":  "failed get healthRecommendations",
// 			"response": err.Error(),
// 		})
// 	}

// 	var healthRecommendationResponse []dto.HealthRecommendationHistoryResponse
// 	for _, healthRecommendation := range responseData {
// 		healthRecommendationResponse = append(healthRecommendationResponse, dto.ConvertToHealthRecommendationHistoryResponse(healthRecommendation))
// 	}

// 	return c.JSON(http.StatusOK, map[string]any{
// 		"message":  "success get healthRecommendations",
// 		"response": healthRecommendationResponse,
// 	})
// }
