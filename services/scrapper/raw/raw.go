package raw

type DataFromDataJSON struct {
	LastUpdate string    `json:"last_update"`
	Kasus      Kasus     `json:"kasus"`
	Sembuh     Meninggal `json:"sembuh"`
	Meninggal  Meninggal `json:"meninggal"`
	Perawatan  Meninggal `json:"perawatan"`
}

type Kasus struct {
	KondisiPenyerta KasusGejala       `json:"kondisi_penyerta"`
	JenisKelamin    KasusGejala       `json:"jenis_kelamin"`
	KelompokUmur    KasusKelompokUmur `json:"kelompok_umur"`
	Gejala          KasusGejala       `json:"gejala"`
}

type KasusGejala struct {
	CurrentData float64           `json:"current_data"`
	MissingData float64           `json:"missing_data"`
	ListData    []GejalaListDatum `json:"list_data"`
}

type GejalaListDatum struct {
	Key      string  `json:"key"`
	DocCount float64 `json:"doc_count"`
}

type KasusKelompokUmur struct {
	CurrentData float64                 `json:"current_data"`
	MissingData float64                 `json:"missing_data"`
	ListData    []KelompokUmurListDatum `json:"list_data"`
}

type KelompokUmurListDatum struct {
	Key      string  `json:"key"`
	DocCount float64 `json:"doc_count"`
	Usia     Usia    `json:"usia"`
}

type Meninggal struct {
	KondisiPenyerta MeninggalGejala       `json:"kondisi_penyerta"`
	JenisKelamin    MeninggalGejala       `json:"jenis_kelamin"`
	KelompokUmur    MeninggalKelompokUmur `json:"kelompok_umur"`
	Gejala          MeninggalGejala       `json:"gejala"`
}

type MeninggalGejala struct {
	ListData []GejalaListDatum `json:"list_data"`
}

type MeninggalKelompokUmur struct {
	ListData []KelompokUmurListDatum `json:"list_data"`
}

type Penambahan struct {
	JumlahPositif   int    `json:"jumlah_positif"`
	JumlahMeninggal int    `json:"jumlah_meninggal"`
	JumlahSembuh    int    `json:"jumlah_sembuh"`
	JumlahDirawat   int    `json:"jumlah_dirawat"`
	Tanggal         string `json:"tanggal"`
	Created         string `json:"created"`
	Positif         int64  `json:"positif"`
	Sembuh          int64  `json:"sembuh"`
	Meninggal       int64  `json:"meninggal"`
}

type Usia struct {
	Value float64 `json:"value"`
}

type DataFromProvJSON struct {
	LastDate      string      `json:"last_date"`
	CurrentData   float64     `json:"current_data"`
	MissingData   float64     `json:"missing_data"`
	TanpaProvinsi int         `json:"tanpa_provinsi"`
	ListData      []ListDatum `json:"list_data"`
}

type ListDatum struct {
	Key             string         `json:"key"`
	DocCount        float64        `json:"doc_count"`
	JumlahKasus     int            `json:"jumlah_kasus"`
	JumlahSembuh    int            `json:"jumlah_sembuh"`
	JumlahMeninggal int            `json:"jumlah_meninggal"`
	JumlahDirawat   int            `json:"jumlah_dirawat"`
	JenisKelamin    []JenisKelamin `json:"jenis_kelamin"`
	KelompokUmur    []KelompokUmur `json:"kelompok_umur"`
	Lokasi          Lokasi         `json:"lokasi"`
	Penambahan      Penambahan     `json:"penambahan"`
}

type JenisKelamin struct {
	Key      JenisKelaminKey `json:"key"`
	DocCount int             `json:"doc_count"`
}

type KelompokUmur struct {
	Key      KelompokUmurKey `json:"key"`
	DocCount int             `json:"doc_count"`
	Usia     Usia            `json:"usia"`
}

type Lokasi struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type JenisKelaminKey string

const (
	LakiLaki  JenisKelaminKey = "LAKI-LAKI"
	Perempuan JenisKelaminKey = "PEREMPUAN"
)

type KelompokUmurKey string

const (
	The05   KelompokUmurKey = "0-5"
	The1830 KelompokUmurKey = "18-30"
	The3145 KelompokUmurKey = "31-45"
	The4659 KelompokUmurKey = "46-59"
	The60   KelompokUmurKey = "≥ 60"
	The617  KelompokUmurKey = "6-17"
)

type DataFromUpdateJSON struct {
	Data   Data   `json:"data"`
	Update Update `json:"update"`
}

type Data struct {
	ID                   int `json:"id"`
	JumlahOdp            int `json:"jumlah_odp"`
	JumlahPDP            int `json:"jumlah_pdp"`
	TotalSpesimen        int `json:"total_spesimen"`
	TotalSpesimenNegatif int `json:"total_spesimen_negatif"`
}

type Update struct {
	Penambahan Penambahan `json:"penambahan"`
	Harian     []Harian   `json:"harian"`
	Total      Total      `json:"total"`
}

type Harian struct {
	KeyAsString        string `json:"key_as_string"`
	Key                int64  `json:"key"`
	DocCount           int    `json:"doc_count"`
	JumlahMeninggal    Jumlah `json:"jumlah_meninggal"`
	JumlahSembuh       Jumlah `json:"jumlah_sembuh"`
	JumlahPositif      Jumlah `json:"jumlah_positif"`
	JumlahDirawat      Jumlah `json:"jumlah_dirawat"`
	JumlahPositifKum   Jumlah `json:"jumlah_positif_kum"`
	JumlahSembuhKum    Jumlah `json:"jumlah_sembuh_kum"`
	JumlahMeninggalKum Jumlah `json:"jumlah_meninggal_kum"`
	JumlahDirawatKum   Jumlah `json:"jumlah_dirawat_kum"`
}

type Jumlah struct {
	Value int `json:"value"`
}

type Total struct {
	JumlahPositif   int `json:"jumlah_positif"`
	JumlahDirawat   int `json:"jumlah_dirawat"`
	JumlahSembuh    int `json:"jumlah_sembuh"`
	JumlahMeninggal int `json:"jumlah_meninggal"`
}
