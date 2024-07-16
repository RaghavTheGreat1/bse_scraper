package models

type CorporateAnnouncement struct {
	CompanySymbol            string
	CompanyName              string
	Title                    string
	Body                     string
	Category                 string
	PDFLink                  string
	ExchangeReceivedTime     string
	ExchangeDisseminatedTime string
	TimeTaken                string
}
