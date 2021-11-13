from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

driver = webdriver.Firefox()
driver.get("https://www.uslchampionship.com/league-teams")
rosterLinks = driver.find_elements(By.XPATH, "//*[text()='Roster']")
rosterUrls = [x.get_attribute("href") for x in rosterLinks]
print(rosterUrls)
for url in rosterUrls:
    driver.get(url)
    teamInfo = driver.find_element(By.CLASS_NAME, "teamInfo")
    print(teamInfo.find_element(By.TAG_NAME, "h1").text)
    element = WebDriverWait(driver, 10).until(
        EC.presence_of_element_located((By.CLASS_NAME, "Opta-Team"))
    )
    rosterTable = driver.find_element(By.CLASS_NAME, "Opta-Team")
    rosterRows = rosterTable.find_elements(By.TAG_NAME, "tr")
    currentPosition = ""
    for row in rosterRows:
        classList = row.get_attribute("class")
        if "Opta-Position" in classList:
            if(row.text == "Goalkeeper"):
                currentPosition = "GK"
            if(row.text == "Defender"):
                currentPosition = "D"
            if(row.text == "Midfielder"):
                currentPosition = "M"
            if(row.text == "Forward"):
                currentPosition = "F"
        if "Opta-Player" in classList:
            shirtNumber = row.find_element(By.CLASS_NAME, "Opta-Shirt").text
            print(shirtNumber)
            name = row.find_element(By.CLASS_NAME, "Opta-Name").text
            print(name)
driver.close()