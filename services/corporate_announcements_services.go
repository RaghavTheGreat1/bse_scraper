package services

import (
	"fmt"

	"github.com/RaghavTheGreat1/bse_scraper/models"
	"github.com/RaghavTheGreat1/bse_scraper/utils"

	"github.com/playwright-community/playwright-go"
)

func ExtractCorporateAnnouncements(ctx playwright.BrowserContext, pages int) (data []models.CorporateAnnouncement, extractedPage int, err error) {

	page, err := ctx.NewPage()

	if err != nil {
		return nil, 0, fmt.Errorf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://www.bseindia.com/corporates/ann.html", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		return nil, 0, fmt.Errorf("could not go to url: %v", err)
	}

	err = page.WaitForURL("https://www.bseindia.com/corporates/ann.html")

	if err != nil {
		return nil, 0, fmt.Errorf("could not wait for url: %v", err)
	}

	currentPage := 1

	for currentPage <= pages {
		fmt.Println("Extracting data from page: ", currentPage)
		extractedData, err := extractFromTable(page)
		if err != nil {
			return nil, 0, fmt.Errorf("could not extract data: %v", err)
		}
		data = append(data, extractedData...)
		err = page.Locator(`xpath=//*[@id="idnext"]`).First().Click()
		if err != nil {
			fmt.Printf("could not click: %v\n", err)
			break
		}
		currentPage++
	}

	fmt.Println(len(data))
	return data, currentPage - 1, nil

}

func extractFromTable(page playwright.Page) (data []models.CorporateAnnouncement, err error) {
	page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State: playwright.LoadStateNetworkidle,
	})
	page.WaitForSelector("xpath=//html/body/div[1]/div[5]/div[2]/div/div[3]/div/div/table/tbody/tr/td[2]/table/tbody/tr[2]/td/table/tbody/tr[4]/td")
	announcementsTable, err := page.Locator("xpath=//html/body/div[1]/div[5]/div[2]/div/div[3]/div/div/table/tbody/tr/td[2]/table/tbody/tr[2]/td/table/tbody/tr[4]/td").All()
	if err != nil {
		return nil, fmt.Errorf("could not get table: %v", err)
	}

	if len(announcementsTable) == 0 {
		return nil, fmt.Errorf("could not find table")
	}

	subTables, err := announcementsTable[0].Locator("table").All()

	if err != nil {
		return nil, fmt.Errorf("could not get sub tables: %v", err)
	}

	for _, subTableData := range subTables {
		c := models.CorporateAnnouncement{}

		// Get Heading Data
		heading, err := subTableData.Locator("xpath=./tbody/tr[1]/td[1]/span").All()

		if err != nil {
			return nil, fmt.Errorf("could not get heading: %v", err)
		}

		headingText, err := heading[0].TextContent()

		if err != nil {
			return nil, fmt.Errorf("could not get text content: %v", err)
		}

		name, symbol, title, err := utils.ExtractCompanyInfoFromHeading(headingText)

		if err != nil {
			return nil, fmt.Errorf("could not extract company info: %v", err)
		}

		c.CompanyName = name
		c.CompanySymbol = symbol
		c.Title = title

		// Get Category Data
		category, err := subTableData.Locator("xpath=./tbody/tr[1]/td[2]").All()

		if err != nil {
			return nil, fmt.Errorf("could not get category: %v", err)
		}

		categoryText, err := category[0].TextContent()

		if err != nil {
			return nil, fmt.Errorf("could not get text content: %v", err)
		}

		c.Category = categoryText

		// Get PDF Link Data
		pdfLink, err := subTableData.Locator("xpath=./tbody/tr[1]/td[4]/a").All()

		if len(pdfLink) != 0 {

			if err != nil {
				return nil, fmt.Errorf("could not get pdf link: %v", err)
			}

			pdfLinkText, err := pdfLink[0].GetAttribute("href")

			if err != nil {
				return nil, fmt.Errorf("could not get pdf link: %v", err)
			}

			c.PDFLink = pdfLinkText
		}

		// Get Timing Data
		timing, err := subTableData.Locator("xpath=./tbody/tr[3]/td").All()

		if err != nil {
			return nil, fmt.Errorf("could not get timing: %v", err)
		}

		timingText, err := timing[0].TextContent()

		if err != nil {
			return nil, fmt.Errorf("could not get text content: %v", err)
		}

		exchangeReceivedTime, exchangeDisseminatedTime, timeTaken, err := utils.ExtractTimings(timingText)

		if err != nil {
			return nil, fmt.Errorf("could not extract timings: %v", err)
		}

		c.ExchangeReceivedTime = exchangeReceivedTime
		c.ExchangeDisseminatedTime = exchangeDisseminatedTime
		c.TimeTaken = timeTaken

		// Get Body Data

		body, err := subTableData.Locator("xpath=./tbody/tr[2]/td/div[1]/span").All()

		if err != nil {
			return nil, fmt.Errorf("could not get body: %v", err)
		}

		bodyText, err := body[0].TextContent()

		if err != nil {
			return nil, fmt.Errorf("could not get text content: %v", err)
		}

		c.Body = bodyText

		data = append(data, c)
	}

	return data, nil

}
