package test_ms_beer_entity

type Currencylayer struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		AED float64 `json:"USDAED"`
		AFN float64 `json:"USDAFN"`
		ALL float64 `json:"USDALL"`
		AMD float64 `json:"USDAMD"`
		ANG float64 `json:"USDANG"`
		AOA float64 `json:"USDAOA"`
		ARS float64 `json:"USDARS"`
		AUD float64 `json:"USDAUD"`
		AWG float64 `json:"USDAWG"`
		AZN float64 `json:"USDAZN"`
		BAM float64 `json:"USDBAM"`
		BBD float64 `json:"USDBBD"`
		BDT float64 `json:"USDBDT"`
		BGN float64 `json:"USDBGN"`
		BHD float64 `json:"USDBHD"`
		BIF float64 `json:"USDBIF"`
		BMD float64 `json:"USDBMD"`
		BND float64 `json:"USDBND"`
		BOB float64 `json:"USDBOB"`
		BRL float64 `json:"USDBRL"`
		BSD float64 `json:"USDBSD"`
		BTC float64 `json:"USDBTC"`
		BTN float64 `json:"USDBTN"`
		BWP float64 `json:"USDBWP"`
		BYN float64 `json:"USDBYN"`
		BYR float64 `json:"USDBYR"`
		BZD float64 `json:"USDBZD"`
		CAD float64 `json:"USDCAD"`
		CDF float64 `json:"USDCDF"`
		CHF float64 `json:"USDCHF"`
		CLF float64 `json:"USDCLF"`
		CLP float64 `json:"USDCLP"`
		CNY float64 `json:"USDCNY"`
		COP float64 `json:"USDCOP"`
		CRC float64 `json:"USDCRC"`
		CUC float64 `json:"USDCUC"`
		CUP float64 `json:"USDCUP"`
		CVE float64 `json:"USDCVE"`
		CZK float64 `json:"USDCZK"`
		DJF float64 `json:"USDDJF"`
		DKK float64 `json:"USDDKK"`
		DOP float64 `json:"USDDOP"`
		DZD float64 `json:"USDDZD"`
		EGP float64 `json:"USDEGP"`
		ERN float64 `json:"USDERN"`
		ETB float64 `json:"USDETB"`
		EUR float64 `json:"USDEUR"`
		FJD float64 `json:"USDFJD"`
		FKP float64 `json:"USDFKP"`
		GBP float64 `json:"USDGBP"`
		GEL float64 `json:"USDGEL"`
		GGP float64 `json:"USDGGP"`
		GHS float64 `json:"USDGHS"`
		GIP float64 `json:"USDGIP"`
		GMD float64 `json:"USDGMD"`
		GNF float64 `json:"USDGNF"`
		GTQ float64 `json:"USDGTQ"`
		GYD float64 `json:"USDGYD"`
		HKD float64 `json:"USDHKD"`
		HNL float64 `json:"USDHNL"`
		HRK float64 `json:"USDHRK"`
		HTG float64 `json:"USDHTG"`
		HUF float64 `json:"USDHUF"`
		IDR float64 `json:"USDIDR"`
		ILS float64 `json:"USDILS"`
		IMP float64 `json:"USDIMP"`
		INR float64 `json:"USDINR"`
		IQD float64 `json:"USDIQD"`
		IRR float64 `json:"USDIRR"`
		ISK float64 `json:"USDISK"`
		JEP float64 `json:"USDJEP"`
		JMD float64 `json:"USDJMD"`
		JOD float64 `json:"USDJOD"`
		JPY float64 `json:"USDJPY"`
		KES float64 `json:"USDKES"`
		KGS float64 `json:"USDKGS"`
		KHR float64 `json:"USDKHR"`
		KMF float64 `json:"USDKMF"`
		KPW float64 `json:"USDKPW"`
		KRW float64 `json:"USDKRW"`
		KWD float64 `json:"USDKWD"`
		KYD float64 `json:"USDKYD"`
		KZT float64 `json:"USDKZT"`
		LAK float64 `json:"USDLAK"`
		LBP float64 `json:"USDLBP"`
		LKR float64 `json:"USDLKR"`
		LRD float64 `json:"USDLRD"`
		LSL float64 `json:"USDLSL"`
		LTL float64 `json:"USDLTL"`
		LVL float64 `json:"USDLVL"`
		LYD float64 `json:"USDLYD"`
		MAD float64 `json:"USDMAD"`
		MDL float64 `json:"USDMDL"`
		MGA float64 `json:"USDMGA"`
		MKD float64 `json:"USDMKD"`
		MMK float64 `json:"USDMMK"`
		MNT float64 `json:"USDMNT"`
		MOP float64 `json:"USDMOP"`
		MRO float64 `json:"USDMRO"`
		MUR float64 `json:"USDMUR"`
		MVR float64 `json:"USDMVR"`
		MWK float64 `json:"USDMWK"`
		MXN float64 `json:"USDMXN"`
		MYR float64 `json:"USDMYR"`
		MZN float64 `json:"USDMZN"`
		NAD float64 `json:"USDNAD"`
		NGN float64 `json:"USDNGN"`
		NIO float64 `json:"USDNIO"`
		NOK float64 `json:"USDNOK"`
		NPR float64 `json:"USDNPR"`
		NZD float64 `json:"USDNZD"`
		OMR float64 `json:"USDOMR"`
		PAB float64 `json:"USDPAB"`
		PEN float64 `json:"USDPEN"`
		PGK float64 `json:"USDPGK"`
		PHP float64 `json:"USDPHP"`
		PKR float64 `json:"USDPKR"`
		PLN float64 `json:"USDPLN"`
		PYG float64 `json:"USDPYG"`
		QAR float64 `json:"USDQAR"`
		RON float64 `json:"USDRON"`
		RSD float64 `json:"USDRSD"`
		RUB float64 `json:"USDRUB"`
		RWF float64 `json:"USDRWF"`
		SAR float64 `json:"USDSAR"`
		SBD float64 `json:"USDSBD"`
		SCR float64 `json:"USDSCR"`
		SDG float64 `json:"USDSDG"`
		SEK float64 `json:"USDSEK"`
		SGD float64 `json:"USDSGD"`
		SHP float64 `json:"USDSHP"`
		SLL float64 `json:"USDSLL"`
		SOS float64 `json:"USDSOS"`
		SRD float64 `json:"USDSRD"`
		STD float64 `json:"USDSTD"`
		SVC float64 `json:"USDSVC"`
		SYP float64 `json:"USDSYP"`
		SZL float64 `json:"USDSZL"`
		THB float64 `json:"USDTHB"`
		TJS float64 `json:"USDTJS"`
		TMT float64 `json:"USDTMT"`
		TND float64 `json:"USDTND"`
		TOP float64 `json:"USDTOP"`
		TRY float64 `json:"USDTRY"`
		TTD float64 `json:"USDTTD"`
		TWD float64 `json:"USDTWD"`
		TZS float64 `json:"USDTZS"`
		UAH float64 `json:"USDUAH"`
		UGX float64 `json:"USDUGX"`
		USD float64 `json:"USDUSD"`
		UYU float64 `json:"USDUYU"`
		UZS float64 `json:"USDUZS"`
		VEF float64 `json:"USDVEF"`
		VND float64 `json:"USDVND"`
		VUV float64 `json:"USDVUV"`
		WST float64 `json:"USDWST"`
		XAF float64 `json:"USDXAF"`
		XAG float64 `json:"USDXAG"`
		XAU float64 `json:"USDXAU"`
		XCD float64 `json:"USDXCD"`
		XDR float64 `json:"USDXDR"`
		XOF float64 `json:"USDXOF"`
		XPF float64 `json:"USDXPF"`
		YER float64 `json:"USDYER"`
		ZAR float64 `json:"USDZAR"`
		ZMK float64 `json:"USDZMK"`
		ZMW float64 `json:"USDZMW"`
		ZWL float64 `json:"USDZWL"`
	} `json:"quotes"`
}
