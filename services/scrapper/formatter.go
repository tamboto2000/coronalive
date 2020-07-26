package scrapper

import (
	"math"
	"time"

	rawStruct "github.com/tamboto2000/coronalive/services/scrapper/raw"
)

//COVIDData adalah kumpulan semua data yang disatukan dalam satu object
type COVIDData struct {
	NationalSummary         Item                      `json:"nationalSummary"`
	ByDateNational          []ByDate                  `json:"byDateNational"`
	ByProvince              []ByProvince              `json:"byProvince"`
	ByGenderNational        []ByGender                `json:"byGenderNational"`
	ByAgeNational           []ByAge                   `json:"byAgeNational"`
	BySimptom               []BySimptom               `json:"bySimptom"`
	ByAccompanyingCondition []ByAccompanyingCondition `json:"byAccompanyingCondition"`
}

//ByDate data berdasarkan tanggal
type ByDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
	Item
}

//ByProvince data berdasarkan provinsi
type ByProvince struct {
	Province string     `json:"province"`
	ByGender []ByGender `json:"byGender"`
	ByAge    []ByAge    `json:"byAge"`
	Item
}

//ByGender data berdasarkan jenis kelamin
type ByGender struct {
	Gender string `json:"gender"` //MALE,FEMALE
	Item
}

//ByAge data berdasarkan umur
type ByAge struct {
	From int `json:"from"`
	To   int `json:"to"`
	Item
}

//BySimptom data berdasarkan gejala yang ditunjukan pasien
type BySimptom struct {
	//gejala
	Simptom string `json:"simptom"`
	Item
}

//ByAccompanyingCondition data berdasarkan kondisi kesehatan yang diderita pasien
type ByAccompanyingCondition struct {
	Condition string `json:"condition"`
	Item
}

//Item adalah data dasar
type Item struct {
	//positif/terkonfirmasi
	Positive int `json:"positive,omitempty"`
	//dalam perawatan
	InCare int `json:"inCare,omitempty"`
	//sembuh
	Recovered int `json:"recovered,omitempty"`
	//meninggal
	Died int `json:"died,omitempty"`
	//orang dalam pengawasan
	InMonitoring int `json:"inMonitoring,omitempty"` //ODP
	//pasien dalam pengawasan
	UnderSurveillance int `json:"underSurveillance,omitempty"` //PDP
	//total spesimen
	Specimen int `json:"specimen,omitempty"`
	//spesimen negatif
	NegativeSpecimen int `json:"negativeSpecimen,omitempty"`
}

const (
	genderMale   = "MALE"
	genderFemale = "FEMALE"
)

func fromUpdateJSON(result *COVIDData, raw *rawStruct.DataFromUpdateJSON) {
	//extract national summary data
	result.NationalSummary.Positive = raw.Update.Total.JumlahPositif
	result.NationalSummary.InCare = raw.Update.Total.JumlahDirawat
	result.NationalSummary.Recovered = raw.Update.Total.JumlahSembuh
	result.NationalSummary.Died = raw.Update.Total.JumlahMeninggal
	result.NationalSummary.InMonitoring = raw.Data.JumlahOdp
	result.NationalSummary.UnderSurveillance = raw.Data.JumlahPDP
	result.NationalSummary.Specimen = raw.Data.TotalSpesimen
	result.NationalSummary.Specimen = raw.Data.TotalSpesimenNegatif

	//extract daily national data
	for _, data := range raw.Update.Harian {
		date := time.Unix(data.Key/1000, 0)
		result.ByDateNational = append(result.ByDateNational, ByDate{
			Year:  date.Year(),
			Month: int(date.Month()),
			Day:   date.Day(),
			Item: Item{
				Died:      data.JumlahMeninggal.Value,
				Recovered: data.JumlahSembuh.Value,
				Positive:  data.JumlahPositif.Value,
				InCare:    data.JumlahDirawat.Value,
			},
		})
	}
}

func fromProvJSON(result *COVIDData, raw *rawStruct.DataFromProvJSON) {
	for _, data := range raw.ListData {
		result.ByProvince = append(result.ByProvince, ByProvince{
			Province: data.Key,
			Item: Item{
				Positive:  data.JumlahKasus,
				Recovered: data.JumlahSembuh,
				Died:      data.JumlahMeninggal,
				InCare:    data.JumlahDirawat,
			},

			ByGender: []ByGender{
				ByGender{
					Gender: genderMale,
					Item: Item{
						Positive: data.JenisKelamin[0].DocCount,
					},
				},
				ByGender{
					Gender: genderFemale,
					Item: Item{
						Positive: data.JenisKelamin[1].DocCount,
					},
				},
			},

			ByAge: []ByAge{
				ByAge{
					From: 0,
					To:   5,
					Item: Item{
						Positive: data.KelompokUmur[0].DocCount,
					},
				},

				ByAge{
					From: 6,
					To:   17,
					Item: Item{
						Positive: data.KelompokUmur[1].DocCount,
					},
				},

				ByAge{
					From: 18,
					To:   30,
					Item: Item{
						Positive: data.KelompokUmur[2].DocCount,
					},
				},

				ByAge{
					From: 31,
					To:   45,
					Item: Item{
						Positive: data.KelompokUmur[3].DocCount,
					},
				},

				ByAge{
					From: 46,
					To:   59,
					Item: Item{
						Positive: data.KelompokUmur[4].DocCount,
					},
				},

				ByAge{
					From: 60,
					To:   0,
					Item: Item{
						Positive: data.KelompokUmur[5].DocCount,
					},
				},
			},
		})
	}
}

func fromDataJSON(result *COVIDData, raw *rawStruct.DataFromDataJSON) {
	//extract data by gender national
	nationalPositivePercentil := (float64(result.NationalSummary.Positive) / 100)
	nationalRecovPercentil := (float64(result.NationalSummary.Recovered) / 100)
	nationalDiedPercentil := (float64(result.NationalSummary.Died) / 100)
	nationalInCarePercentil := (float64(result.NationalSummary.InCare) / 100)

	result.ByGenderNational = append(result.ByGenderNational, []ByGender{
		ByGender{
			Gender: genderMale,
			Item: Item{
				Positive:  int(math.RoundToEven(nationalPositivePercentil * raw.Kasus.JenisKelamin.ListData[0].DocCount)),
				Recovered: int(math.RoundToEven(nationalRecovPercentil * raw.Sembuh.JenisKelamin.ListData[0].DocCount)),
				Died:      int(math.RoundToEven(nationalDiedPercentil * raw.Meninggal.JenisKelamin.ListData[0].DocCount)),
				InCare:    int(math.RoundToEven(nationalInCarePercentil * raw.Perawatan.JenisKelamin.ListData[0].DocCount)),
			},
		},
		ByGender{
			Gender: genderFemale,
			Item: Item{
				Positive:  int(math.RoundToEven(nationalPositivePercentil * raw.Kasus.JenisKelamin.ListData[1].DocCount)),
				Recovered: int(math.RoundToEven(nationalRecovPercentil * raw.Sembuh.JenisKelamin.ListData[1].DocCount)),
				Died:      int(math.RoundToEven(nationalDiedPercentil * raw.Meninggal.JenisKelamin.ListData[1].DocCount)),
				InCare:    int(math.RoundToEven(nationalInCarePercentil * raw.Perawatan.JenisKelamin.ListData[1].DocCount)),
			},
		},
	}...)

	//extract data by age nasional
	result.ByAgeNational = append(result.ByAgeNational, []ByAge{
		ByAge{
			From: 0,
			To:   5,
			Item: Item{
				Positive:  int(math.RoundToEven(nationalPositivePercentil * raw.Kasus.KelompokUmur.ListData[0].DocCount)),
				Recovered: int(math.RoundToEven(nationalRecovPercentil * raw.Sembuh.KelompokUmur.ListData[0].DocCount)),
				Died:      int(math.RoundToEven(nationalDiedPercentil * raw.Meninggal.KelompokUmur.ListData[0].DocCount)),
				InCare:    int(math.RoundToEven(nationalInCarePercentil * raw.Perawatan.KelompokUmur.ListData[0].DocCount)),
			},
		},

		ByAge{
			From: 6,
			To:   17,
			Item: Item{
				Positive:  int(math.RoundToEven(nationalPositivePercentil * raw.Kasus.KelompokUmur.ListData[1].DocCount)),
				Recovered: int(math.RoundToEven(nationalRecovPercentil * raw.Sembuh.KelompokUmur.ListData[1].DocCount)),
				Died:      int(math.RoundToEven(nationalDiedPercentil * raw.Meninggal.KelompokUmur.ListData[1].DocCount)),
				InCare:    int(math.RoundToEven(nationalInCarePercentil * raw.Perawatan.KelompokUmur.ListData[1].DocCount)),
			},
		},

		ByAge{
			From: 18,
			To:   30,
			Item: Item{
				Positive:  int(math.RoundToEven(nationalPositivePercentil * raw.Kasus.KelompokUmur.ListData[2].DocCount)),
				Recovered: int(math.RoundToEven(nationalRecovPercentil * raw.Sembuh.KelompokUmur.ListData[2].DocCount)),
				Died:      int(math.RoundToEven(nationalDiedPercentil * raw.Meninggal.KelompokUmur.ListData[2].DocCount)),
				InCare:    int(math.RoundToEven(nationalInCarePercentil * raw.Perawatan.KelompokUmur.ListData[2].DocCount)),
			},
		},

		ByAge{
			From: 31,
			To:   45,
			Item: Item{
				Positive:  int(math.RoundToEven(nationalPositivePercentil * raw.Kasus.KelompokUmur.ListData[3].DocCount)),
				Recovered: int(math.RoundToEven(nationalRecovPercentil * raw.Sembuh.KelompokUmur.ListData[3].DocCount)),
				Died:      int(math.RoundToEven(nationalDiedPercentil * raw.Meninggal.KelompokUmur.ListData[3].DocCount)),
				InCare:    int(math.RoundToEven(nationalInCarePercentil * raw.Perawatan.KelompokUmur.ListData[3].DocCount)),
			},
		},

		ByAge{
			From: 46,
			To:   59,
			Item: Item{
				Positive:  int(math.RoundToEven(nationalPositivePercentil * raw.Kasus.KelompokUmur.ListData[4].DocCount)),
				Recovered: int(math.RoundToEven(nationalRecovPercentil * raw.Sembuh.KelompokUmur.ListData[4].DocCount)),
				Died:      int(math.RoundToEven(nationalDiedPercentil * raw.Meninggal.KelompokUmur.ListData[4].DocCount)),
				InCare:    int(math.RoundToEven(nationalInCarePercentil * raw.Perawatan.KelompokUmur.ListData[4].DocCount)),
			},
		},

		ByAge{
			From: 60,
			To:   0,
			Item: Item{
				Positive:  int(math.RoundToEven(nationalPositivePercentil * raw.Kasus.KelompokUmur.ListData[5].DocCount)),
				Recovered: int(math.RoundToEven(nationalRecovPercentil * raw.Sembuh.KelompokUmur.ListData[5].DocCount)),
				Died:      int(math.RoundToEven(nationalDiedPercentil * raw.Meninggal.KelompokUmur.ListData[5].DocCount)),
				InCare:    int(math.RoundToEven(nationalInCarePercentil * raw.Perawatan.KelompokUmur.ListData[5].DocCount)),
			},
		},
	}...)

	//extract data by simptom
	bySimptom := make(map[string]*BySimptom)
	//iterate confirmed data
	for _, data := range raw.Kasus.Gejala.ListData {
		bySimptom[data.Key] = &BySimptom{
			Simptom: data.Key,
			Item: Item{
				Positive: int(math.RoundToEven(nationalPositivePercentil * data.DocCount)),
			},
		}
	}

	//iterate recovered data
	for _, data := range raw.Sembuh.Gejala.ListData {
		if val, ok := bySimptom[data.Key]; ok {
			val.Item.Recovered = int(math.RoundToEven(nationalRecovPercentil * data.DocCount))
			continue
		}

		bySimptom[data.Key] = &BySimptom{
			Simptom: data.Key,
			Item: Item{
				Recovered: int(math.RoundToEven(nationalRecovPercentil * data.DocCount)),
			},
		}
	}

	//iterate died data
	for _, data := range raw.Meninggal.Gejala.ListData {
		if val, ok := bySimptom[data.Key]; ok {
			val.Item.Died = int(math.RoundToEven(nationalDiedPercentil * data.DocCount))
			continue
		}

		bySimptom[data.Key] = &BySimptom{
			Simptom: data.Key,
			Item: Item{
				Died: int(math.RoundToEven(nationalDiedPercentil * data.DocCount)),
			},
		}
	}

	//iterate in care data
	for _, data := range raw.Perawatan.Gejala.ListData {
		if val, ok := bySimptom[data.Key]; ok {
			val.Item.InCare = int(math.RoundToEven(nationalInCarePercentil * data.DocCount))
			continue
		}

		bySimptom[data.Key] = &BySimptom{
			Simptom: data.Key,
			Item: Item{
				InCare: int(math.RoundToEven(nationalInCarePercentil * data.DocCount)),
			},
		}
	}

	for _, val := range bySimptom {
		result.BySimptom = append(result.BySimptom, *val)
	}

	//extract data by accompanying condition
	byCondition := make(map[string]*ByAccompanyingCondition)

	//iterate confirmed data
	for _, data := range raw.Kasus.KondisiPenyerta.ListData {
		byCondition[data.Key] = &ByAccompanyingCondition{
			Condition: data.Key,
			Item: Item{
				Positive: int(math.RoundToEven(nationalPositivePercentil * data.DocCount)),
			},
		}
	}

	//iterate recovered data
	for _, data := range raw.Sembuh.KondisiPenyerta.ListData {
		if val, ok := byCondition[data.Key]; ok {
			val.Item.Recovered = int(math.RoundToEven(nationalRecovPercentil * data.DocCount))
			continue
		}

		byCondition[data.Key] = &ByAccompanyingCondition{
			Condition: data.Key,
			Item: Item{
				Recovered: int(math.RoundToEven(nationalRecovPercentil * data.DocCount)),
			},
		}
	}

	//iterate died data
	for _, data := range raw.Meninggal.KondisiPenyerta.ListData {
		if val, ok := byCondition[data.Key]; ok {
			val.Item.Died = int(math.RoundToEven(nationalDiedPercentil * data.DocCount))
			continue
		}

		byCondition[data.Key] = &ByAccompanyingCondition{
			Condition: data.Key,
			Item: Item{
				Died: int(math.RoundToEven(nationalDiedPercentil * data.DocCount)),
			},
		}
	}

	//iterate in care data
	for _, data := range raw.Perawatan.KondisiPenyerta.ListData {
		if val, ok := byCondition[data.Key]; ok {
			val.Item.InCare = int(math.RoundToEven(nationalInCarePercentil * data.DocCount))
			continue
		}

		byCondition[data.Key] = &ByAccompanyingCondition{
			Condition: data.Key,
			Item: Item{
				InCare: int(math.RoundToEven(nationalInCarePercentil * data.DocCount)),
			},
		}
	}

	for _, val := range byCondition {
		result.ByAccompanyingCondition = append(result.ByAccompanyingCondition, *val)
	}
}
