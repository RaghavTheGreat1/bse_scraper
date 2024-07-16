package services

import (
	"log"

	"github.com/RaghavTheGreat1/bse_scraper/models"
	"github.com/RaghavTheGreat1/bse_scraper/utils"

	"github.com/playwright-community/playwright-go"
)

func ExtractCorporateAnnouncements(ctx playwright.BrowserContext) (data []models.CorporateAnnouncement) {

	page, err := ctx.NewPage()

	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://www.bseindia.com/corporates/ann.html", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	err = page.WaitForURL("https://www.bseindia.com/corporates/ann.html")

	if err != nil {
		log.Fatalf("could not wait for url: %v", err)
	}

	announcementsTable, err := page.Locator("xpath=//html/body/div[1]/div[5]/div[2]/div/div[3]/div/div/table/tbody/tr/td[2]/table/tbody/tr[2]/td/table/tbody/tr[4]/td").All()
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}

	if len(announcementsTable) == 0 {
		log.Fatalf("no entries found")
	}

	subTables, err := announcementsTable[0].Locator("table").All()

	if err != nil {
		log.Fatalf("could not get sub tables: %v", err)
	}

	for _, subTableData := range subTables {
		c := models.CorporateAnnouncement{}

		// Get Heading Data
		heading, err := subTableData.Locator("xpath=./tbody/tr[1]/td[1]/span").All()

		if err != nil {
			log.Fatalf("could not get heading: %v", err)
		}

		headingText, err := heading[0].TextContent()

		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}

		name, symbol, title, err := utils.ExtractCompanyInfoFromHeading(headingText)

		if err != nil {
			log.Fatalf("could not extract company info from heading: %v", err)
		}

		c.CompanyName = name
		c.CompanySymbol = symbol
		c.Title = title

		// Get Category Data
		category, err := subTableData.Locator("xpath=./tbody/tr[1]/td[2]").All()

		if err != nil {
			log.Fatalf("could not get category: %v", err)
		}

		categoryText, err := category[0].TextContent()

		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}

		c.Category = categoryText

		// Get PDF Link Data
		pdfLink, err := subTableData.Locator("xpath=./tbody/tr[1]/td[4]/a").All()

		if err != nil {
			log.Fatalf("could not get pdf link: %v", err)
		}

		pdfLinkText, err := pdfLink[0].GetAttribute("href")

		if err != nil {
			log.Fatalf("could not get pdf link attribute: %v", err)
		}

		c.PDFLink = pdfLinkText

		// Get Timing Data
		timing, err := subTableData.Locator("xpath=./tbody/tr[3]/td").All()

		if err != nil {
			log.Fatalf("could not get timing: %v", err)
		}

		timingText, err := timing[0].TextContent()

		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}

		exchangeReceivedTime, exchangeDisseminatedTime, timeTaken, err := utils.ExtractTimings(timingText)

		if err != nil {
			log.Fatalf("could not extract timing info: %v", err)
		}

		c.ExchangeReceivedTime = exchangeReceivedTime
		c.ExchangeDisseminatedTime = exchangeDisseminatedTime
		c.TimeTaken = timeTaken

		// Get Body Data

		body, err := subTableData.Locator("xpath=./tbody/tr[2]/td/div[1]/span").All()

		if err != nil {
			log.Fatalf("could not get body: %v", err)
		}

		bodyText, err := body[0].TextContent()

		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}

		c.Body = bodyText

		data = append(data, c)
	}

	return data

}
