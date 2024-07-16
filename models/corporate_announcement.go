package models

type CorporateAnnouncement struct {
	CompanySymbol            string `json:"company_symbol"`
	CompanyName              string `json:"company_name"`
	Title                    string `json:"title"`
	Body                     string `json:"body"`
	Category                 string `json:"category"`
	PDFLink                  string `json:"pdf_link"`
	ExchangeReceivedTime     string `json:"exchange_received_time"`
	ExchangeDisseminatedTime string `json:"exchange_disseminated_time"`
	TimeTaken                string `json:"time_taken"`
}
