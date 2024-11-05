package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	fmt.Println("Simulated human-like click on skip button.")

	// Set up Selenium WebDriver with Chrome
	const (
		seleniumPath     = "C:\\Users\\ishag\\selenium_\\selenium-server-4.26.0.jar"                // path to your Selenium server JAR file
		chromeDriverPath = "C:\\Users\\ishag\\chromeDriver\\chromedriver-win64"                     // path to your ChromeDriver
		port             = 8080                                                                     // Port for Selenium
		profilePath      = "C:\\Users\\ishag\\AppData\\Local\\Google\\Chrome\\User Data\\Profile 1" // Replace with your actual profile path
	)

	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
	}

	// Start a Selenium WebDriver service
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		log.Fatalf("Error starting Selenium service: %v", err)
	}
	defer service.Stop()

	// Define Chrome capabilities
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := map[string]interface{}{
		"args": []string{fmt.Sprintf("--user-data-dir=%s", profilePath)},
	}
	caps["goog:chromeOptions"] = chromeCaps

	// Start WebDriver
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Error connecting to WebDriver: %v", err)
	}
	defer driver.Quit()

	// Open YouTube
	err = driver.Get("https://www.youtube.com")
	if err != nil {
		log.Fatalf("Error opening YouTube: %v", err)
	}

	// Monitor for ads and attempt to skip
	for {
		adDetected, err := checkAndSkipAd(driver)
		if err != nil {
			fmt.Println("Error detecting or skipping ad:", err)
		}

		if adDetected {
			fmt.Println("Ad detected - attempting to skip or adjust playback speed.")
		} else {
			fmt.Println("Playback speed set to normal.")
		}

		// Wait before checking again
		time.Sleep(1 * time.Second)
	}
}

func checkAndSkipAd(driver selenium.WebDriver) (bool, error) {
	// Check for ad badge using updated CSS selector
	adBadge, err := driver.FindElements(selenium.ByCSSSelector, ".ytp-ad-badge--clean-player.ytp-ad-badge--stark-clean-player")
	if err != nil {
		return false, fmt.Errorf("error finding ad badge: %w", err)
	}

	if len(adBadge) > 0 {
		// Try to find and click the skip button
		skipButton, err := driver.FindElement(selenium.ByCSSSelector, ".ytp-skip-ad-button")
		if err == nil {
			return clickSkipButton(driver, skipButton)
		}

		// If no skip button found, set playback speed to 12x for ads
		setPlaybackSpeed(driver, 12)
		return true, nil
	}

	// Reset playback speed to normal if no ad is playing
	setPlaybackSpeed(driver, 1)
	return false, nil
}

func clickSkipButton(driver selenium.WebDriver, button selenium.WebElement) (bool, error) {
	video, err := driver.FindElement(selenium.ByTagName, "video")
	if err != nil {
		return false, fmt.Errorf("error finding video element: %w", err)
	}

	// Speed up video temporarily before clicking skip
	_, err = driver.ExecuteScript("arguments[0].playbackRate = 12;", []interface{}{video})
	if err != nil {
		return false, fmt.Errorf("error setting playback speed for ad: %w", err)
	}
	err = button.Click()
	if err != nil {
		return false, fmt.Errorf("error clicking skip button: %w", err)
	}

	fmt.Println("Ad skipped successfully.")
	return true, nil
}

func setPlaybackSpeed(driver selenium.WebDriver, speed float64) {
	video, err := driver.FindElement(selenium.ByTagName, "video")
	if err == nil {
		_, err = driver.ExecuteScript(fmt.Sprintf("arguments[0].playbackRate = %f;", speed), []interface{}{video})
		if err == nil {
			fmt.Printf("Playback speed set to %vx.\n", speed)
		} else {
			fmt.Println("Error setting playback speed:", err)
		}
	} else {
		fmt.Println("Error finding video element:", err)
	}
}
