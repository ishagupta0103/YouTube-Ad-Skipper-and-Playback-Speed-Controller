# YouTube Ad Skipper and Playback Speed Controller

This project is a Selenium-based Go application designed to automate ad-skipping and playback speed control on YouTube. It detects ads on YouTube videos and either skips them when possible or increases the playback speed to help minimize the interruption.

## Features

- **Automatic Ad Skipping**: Detects YouTube ads and clicks the skip button when available.
- **Playback Speed Control**: Increases playback speed to 12x during ads if the skip button is not available, and resets it to normal after the ad.
- **Simulated Human-Like Interactions**: Mimics human actions to avoid bot detection by YouTube.

## Requirements

- **Go**: [Download Go](https://golang.org/dl/) and install it.
- **Selenium Server**: Requires Selenium Server JAR file. [Download Selenium Server](https://www.selenium.dev/downloads/).
- **ChromeDriver**: Matches the version of Chrome installed on your system. [Download ChromeDriver](https://sites.google.com/a/chromium.org/chromedriver/downloads).

### Directory Structure

Place the following files in a directory structure:

project-directory/   
├── selenium-server-standalone-<version>.jar   
├── chromedriver  
└── main.go    

### Setting Paths

1. **Selenium Server Path**: Update the path to your Selenium Server JAR file in the code.
2. **ChromeDriver Path**: Update the path to your `chromedriver` executable.
3. **Chrome User Profile Path**: Update the profile path for Chrome to allow the automation to work on an authenticated session.

## Setup

1. **Install Go Dependencies**: Run the following command to install the Selenium package for Go:
   ```bash
   go get -u github.com/tebeka/selenium
2. **Start Selenium Server**: Before running the program, ensure that the Selenium Server is accessible and running. You can start it with:
   ```bash
   java -jar path/to/selenium-server-<version>.jar standalone --port 8080
3. **Configure Chrome Profile Path**: Set your Chrome profile path in main.go (this enables you to keep a logged-in session).

## Running the Application

1. **Run the Go Program**: In the project directory, run:
   ```bash
   go run main.go
2. **Application Behavior**:
  - The program will open YouTube in a Chrome browser and detect ads.
  - It attempts to skip ads when a "Skip" button is available.
  - If a skippable ad isn’t detected, it increases playback speed to 12x during the ad and returns it to normal after.

## Notes
- **Browser Compatibility**: This project is configured to work with Google Chrome and requires ChromeDriver.
- **Permissions**: Ensure that Selenium Server, ChromeDriver, and Chrome all have the necessary permissions and are compatible with your OS and system configurations.

## Troubleshooting
- **ChromeDriver Version Mismatch**: Make sure your ChromeDriver version matches your Chrome version.
- **Profile Path**: Incorrect profile paths can lead to authentication errors; update as necessary.
- **Port Conflicts**: Ensure the specified port (8080 by default) is not used by other processes.
