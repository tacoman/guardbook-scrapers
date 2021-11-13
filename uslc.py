from selenium import webdriver
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By

driver = webdriver.Firefox()
driver.get("https://www.uslchampionship.com/league-teams")
rosterLinks = driver.find_elements(By.XPATH, "//*[text()='Roster']")
rosterUrls = [x.get_attribute("href") for x in rosterLinks]
print(rosterUrls)
for url in rosterUrls:
    driver.get(url)
    teamInfo = driver.find_element(By.CLASS_NAME, "teamInfo")
    print(teamInfo.find_element(By.TAG_NAME, "h1").text)
driver.close()