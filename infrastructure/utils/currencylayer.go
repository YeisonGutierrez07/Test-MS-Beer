package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"test_ms_beer/domain/entity/test_ms_beer_entity"
)

func GetCurrencyData() (test_ms_beer_entity.Currencylayer, error) {
	apiLayer := os.Getenv("API_LAYER")
	url := fmt.Sprintf("http://apilayer.net/api/live?access_key=" + apiLayer)

	// Llene el registro con los datos del JSON
	var record test_ms_beer_entity.Currencylayer

	// Construye la solicitud
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return record, err
	}

	client := &http.Client{}

	// Envía la solicitud a través de un cliente
	// Envía una solicitud HTTP y devuelve una respuesta HTTP
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return record, err
	}

	// Las personas que llaman deben cerrar el cuerpo
	// cuando termine de leerlo  aplazar el cierre del cuerpo
	defer resp.Body.Close()

	// Utilice json.Decode para leer flujos de datos JSON
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		return record, err
	}

	return record, nil
}

func GetCurrencyValueData(currency string, currencylayer test_ms_beer_entity.Currencylayer) float64 {
	response := 0.0

	switch currency {
	case "AED":
		response = currencylayer.Quotes.AED
		break
	case "AFN":
		response = currencylayer.Quotes.AFN
		break
	case "ALL":
		response = currencylayer.Quotes.ALL
		break
	case "AMD":
		response = currencylayer.Quotes.AMD
		break
	case "ANG":
		response = currencylayer.Quotes.ANG
		break
	case "AOA":
		response = currencylayer.Quotes.AOA
		break
	case "ARS":
		response = currencylayer.Quotes.ARS
		break
	case "AUD":
		response = currencylayer.Quotes.AUD
		break
	case "AWG":
		response = currencylayer.Quotes.AWG
		break
	case "AZN":
		response = currencylayer.Quotes.AZN
		break
	case "BAM":
		response = currencylayer.Quotes.BAM
		break
	case "BBD":
		response = currencylayer.Quotes.BBD
		break
	case "BDT":
		response = currencylayer.Quotes.BDT
		break
	case "BGN":
		response = currencylayer.Quotes.BGN
		break
	case "BHD":
		response = currencylayer.Quotes.BHD
		break
	case "BIF":
		response = currencylayer.Quotes.BIF
		break
	case "BMD":
		response = currencylayer.Quotes.BMD
		break
	case "BND":
		response = currencylayer.Quotes.BND
		break
	case "BOB":
		response = currencylayer.Quotes.BOB
		break
	case "BRL":
		response = currencylayer.Quotes.BRL
		break
	case "BSD":
		response = currencylayer.Quotes.BSD
		break
	case "BTC":
		response = currencylayer.Quotes.BTC
		break
	case "BTN":
		response = currencylayer.Quotes.BTN
		break
	case "BWP":
		response = currencylayer.Quotes.BWP
		break
	case "BYN":
		response = currencylayer.Quotes.BYN
		break
	case "BYR":
		response = currencylayer.Quotes.BYR
		break
	case "BZD":
		response = currencylayer.Quotes.BZD
		break
	case "CAD":
		response = currencylayer.Quotes.CAD
		break
	case "CDF":
		response = currencylayer.Quotes.CDF
		break
	case "CHF":
		response = currencylayer.Quotes.CHF
		break
	case "CLF":
		response = currencylayer.Quotes.CLF
		break
	case "CLP":
		response = currencylayer.Quotes.CLP
		break
	case "CNY":
		response = currencylayer.Quotes.CNY
		break
	case "COP":
		response = currencylayer.Quotes.COP
		break
	case "CRC":
		response = currencylayer.Quotes.CRC
		break
	case "CUC":
		response = currencylayer.Quotes.CUC
		break
	case "CUP":
		response = currencylayer.Quotes.CUP
		break
	case "CVE":
		response = currencylayer.Quotes.CVE
		break
	case "CZK":
		response = currencylayer.Quotes.CZK
		break
	case "DJF":
		response = currencylayer.Quotes.DJF
		break
	case "DKK":
		response = currencylayer.Quotes.DKK
		break
	case "DOP":
		response = currencylayer.Quotes.DOP
		break
	case "DZD":
		response = currencylayer.Quotes.DZD
		break
	case "EGP":
		response = currencylayer.Quotes.EGP
		break
	case "ERN":
		response = currencylayer.Quotes.ERN
		break
	case "ETB":
		response = currencylayer.Quotes.ETB
		break
	case "EUR":
		response = currencylayer.Quotes.EUR
		break
	case "FJD":
		response = currencylayer.Quotes.FJD
		break
	case "FKP":
		response = currencylayer.Quotes.FKP
		break
	case "GBP":
		response = currencylayer.Quotes.GBP
		break
	case "GEL":
		response = currencylayer.Quotes.GEL
		break
	case "GGP":
		response = currencylayer.Quotes.GGP
		break
	case "GHS":
		response = currencylayer.Quotes.GHS
		break
	case "GIP":
		response = currencylayer.Quotes.GIP
		break
	case "GMD":
		response = currencylayer.Quotes.GMD
		break
	case "GNF":
		response = currencylayer.Quotes.GNF
		break
	case "GTQ":
		response = currencylayer.Quotes.GTQ
		break
	case "GYD":
		response = currencylayer.Quotes.GYD
		break
	case "HKD":
		response = currencylayer.Quotes.HKD
		break
	case "HNL":
		response = currencylayer.Quotes.HNL
		break
	case "HRK":
		response = currencylayer.Quotes.HRK
		break
	case "HTG":
		response = currencylayer.Quotes.HTG
		break
	case "HUF":
		response = currencylayer.Quotes.HUF
		break
	case "IDR":
		response = currencylayer.Quotes.IDR
		break
	case "ILS":
		response = currencylayer.Quotes.ILS
		break
	case "IMP":
		response = currencylayer.Quotes.IMP
		break
	case "INR":
		response = currencylayer.Quotes.INR
		break
	case "IQD":
		response = currencylayer.Quotes.IQD
		break
	case "IRR":
		response = currencylayer.Quotes.IRR
		break
	case "ISK":
		response = currencylayer.Quotes.ISK
		break
	case "JEP":
		response = currencylayer.Quotes.JEP
		break
	case "JMD":
		response = currencylayer.Quotes.JMD
		break
	case "JOD":
		response = currencylayer.Quotes.JOD
		break
	case "JPY":
		response = currencylayer.Quotes.JPY
		break
	case "KES":
		response = currencylayer.Quotes.KES
		break
	case "KGS":
		response = currencylayer.Quotes.KGS
		break
	case "KHR":
		response = currencylayer.Quotes.KHR
		break
	case "KMF":
		response = currencylayer.Quotes.KMF
		break
	case "KPW":
		response = currencylayer.Quotes.KPW
		break
	case "KRW":
		response = currencylayer.Quotes.KRW
		break
	case "KWD":
		response = currencylayer.Quotes.KWD
		break
	case "KYD":
		response = currencylayer.Quotes.KYD
		break
	case "KZT":
		response = currencylayer.Quotes.KZT
		break
	case "LAK":
		response = currencylayer.Quotes.LAK
		break
	case "LBP":
		response = currencylayer.Quotes.LBP
		break
	case "LKR":
		response = currencylayer.Quotes.LKR
		break
	case "LRD":
		response = currencylayer.Quotes.LRD
		break
	case "LSL":
		response = currencylayer.Quotes.LSL
		break
	case "LTL":
		response = currencylayer.Quotes.LTL
		break
	case "LVL":
		response = currencylayer.Quotes.LVL
		break
	case "LYD":
		response = currencylayer.Quotes.LYD
		break
	case "MAD":
		response = currencylayer.Quotes.MAD
		break
	case "MDL":
		response = currencylayer.Quotes.MDL
		break
	case "MGA":
		response = currencylayer.Quotes.MGA
		break
	case "MKD":
		response = currencylayer.Quotes.MKD
		break
	case "MMK":
		response = currencylayer.Quotes.MMK
		break
	case "MNT":
		response = currencylayer.Quotes.MNT
		break
	case "MOP":
		response = currencylayer.Quotes.MOP
		break
	case "MRO":
		response = currencylayer.Quotes.MRO
		break
	case "MUR":
		response = currencylayer.Quotes.MUR
		break
	case "MVR":
		response = currencylayer.Quotes.MVR
		break
	case "MWK":
		response = currencylayer.Quotes.MWK
		break
	case "MXN":
		response = currencylayer.Quotes.MXN
		break
	case "MYR":
		response = currencylayer.Quotes.MYR
		break
	case "MZN":
		response = currencylayer.Quotes.MZN
		break
	case "NAD":
		response = currencylayer.Quotes.NAD
		break
	case "NGN":
		response = currencylayer.Quotes.NGN
		break
	case "NIO":
		response = currencylayer.Quotes.NIO
		break
	case "NOK":
		response = currencylayer.Quotes.NOK
		break
	case "NPR":
		response = currencylayer.Quotes.NPR
		break
	case "NZD":
		response = currencylayer.Quotes.NZD
		break
	case "OMR":
		response = currencylayer.Quotes.OMR
		break
	case "PAB":
		response = currencylayer.Quotes.PAB
		break
	case "PEN":
		response = currencylayer.Quotes.PEN
		break
	case "PGK":
		response = currencylayer.Quotes.PGK
		break
	case "PHP":
		response = currencylayer.Quotes.PHP
		break
	case "PKR":
		response = currencylayer.Quotes.PKR
		break
	case "PLN":
		response = currencylayer.Quotes.PLN
		break
	case "PYG":
		response = currencylayer.Quotes.PYG
		break
	case "QAR":
		response = currencylayer.Quotes.QAR
		break
	case "RON":
		response = currencylayer.Quotes.RON
		break
	case "RSD":
		response = currencylayer.Quotes.RSD
		break
	case "RUB":
		response = currencylayer.Quotes.RUB
		break
	case "RWF":
		response = currencylayer.Quotes.RWF
		break
	case "SAR":
		response = currencylayer.Quotes.SAR
		break
	case "SBD":
		response = currencylayer.Quotes.SBD
		break
	case "SCR":
		response = currencylayer.Quotes.SCR
		break
	case "SDG":
		response = currencylayer.Quotes.SDG
		break
	case "SEK":
		response = currencylayer.Quotes.SEK
		break
	case "SGD":
		response = currencylayer.Quotes.SGD
		break
	case "SHP":
		response = currencylayer.Quotes.SHP
		break
	case "SLL":
		response = currencylayer.Quotes.SLL
		break
	case "SOS":
		response = currencylayer.Quotes.SOS
		break
	case "SRD":
		response = currencylayer.Quotes.SRD
		break
	case "STD":
		response = currencylayer.Quotes.STD
		break
	case "SVC":
		response = currencylayer.Quotes.SVC
		break
	case "SYP":
		response = currencylayer.Quotes.SYP
		break
	case "SZL":
		response = currencylayer.Quotes.SZL
		break
	case "THB":
		response = currencylayer.Quotes.THB
		break
	case "TJS":
		response = currencylayer.Quotes.TJS
		break
	case "TMT":
		response = currencylayer.Quotes.TMT
		break
	case "TND":
		response = currencylayer.Quotes.TND
		break
	case "TOP":
		response = currencylayer.Quotes.TOP
		break
	case "TRY":
		response = currencylayer.Quotes.TRY
		break
	case "TTD":
		response = currencylayer.Quotes.TTD
		break
	case "TWD":
		response = currencylayer.Quotes.TWD
		break
	case "TZS":
		response = currencylayer.Quotes.TZS
		break
	case "UAH":
		response = currencylayer.Quotes.UAH
		break
	case "UGX":
		response = currencylayer.Quotes.UGX
		break
	case "USD":
		response = currencylayer.Quotes.USD
		break
	case "UYU":
		response = currencylayer.Quotes.UYU
		break
	case "UZS":
		response = currencylayer.Quotes.UZS
		break
	case "VEF":
		response = currencylayer.Quotes.VEF
		break
	case "VND":
		response = currencylayer.Quotes.VND
		break
	case "VUV":
		response = currencylayer.Quotes.VUV
		break
	case "WST":
		response = currencylayer.Quotes.WST
		break
	case "XAF":
		response = currencylayer.Quotes.XAF
		break
	case "XAG":
		response = currencylayer.Quotes.XAG
		break
	case "XAU":
		response = currencylayer.Quotes.XAU
		break
	case "XCD":
		response = currencylayer.Quotes.XCD
		break
	case "XDR":
		response = currencylayer.Quotes.XDR
		break
	case "XOF":
		response = currencylayer.Quotes.XOF
		break
	case "XPF":
		response = currencylayer.Quotes.XPF
		break
	case "YER":
		response = currencylayer.Quotes.YER
		break
	case "ZAR":
		response = currencylayer.Quotes.ZAR
		break
	case "ZMK":
		response = currencylayer.Quotes.ZMK
		break
	case "ZMW":
		response = currencylayer.Quotes.ZMW
		break
	case "ZWL":
		response = currencylayer.Quotes.ZWL
		break
	}

	return response
}
