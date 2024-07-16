package services

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func InitializePlaywright() playwright.BrowserContext {
	pw, err := playwright.Run(
		&playwright.RunOptions{},
	)
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		// Headless: playwright.Bool(false), // Run in headful mode
		Args: []string{"--disable-blink-features=AutomationControlled"},
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: playwright.String("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"), // Fake user agent
		Viewport:  &playwright.Size{Width: 1920, Height: 1080},
	})

	if err != nil {
		log.Fatalf("could not create context: %v", err)
	}
	defer context.Close()

	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

	return context
}
