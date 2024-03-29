from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.color import Color
from pymongo import MongoClient
import json
import os
import re

def atoi(text):
    return int(text) if text.isdigit() else 0

DB_URL = os.environ.get('DB_URL')

driver = webdriver.Firefox()
driver.get("https://www.uslchampionship.com/league-teams")
rosterLinks = driver.find_elements(By.XPATH, "//*[text()='Roster']")
rosterUrls = [x.get_attribute("href") for x in rosterLinks]
rosterUrls.append("https://www.uslchampionship.com/new-york-red-bulls-ii-roster")
rosterUrls.sort()
print(rosterUrls)
foes = []
for url in rosterUrls:
    driver.get(url)
    teamInfo = driver.find_element(By.CLASS_NAME, "teamInfo")
    print(teamInfo.find_element(By.TAG_NAME, "h1").text)
    element = WebDriverWait(driver, 30).until(
        EC.presence_of_element_located((By.CLASS_NAME, "Opta-Team"))
    )
    rosterTable = driver.find_element(By.CLASS_NAME, "Opta-Team")
    rosterRows = rosterTable.find_elements(By.TAG_NAME, "tr")
    currentPosition = ""
    foe = {}
    foe["players"] = []
    foe["opponent"] = teamInfo.find_element(By.TAG_NAME, "h1").text
    if(foe["opponent"]) == "DETROIT CITY FC":
        continue;
    foe["competition"] = "USL Championship"
    foe["textColor"] = "#000000"
    foe["backgroundColor"] = "#ffffff"
    banner = driver.find_element(By.CLASS_NAME, "teamInfoBanner")
    bannerColor = banner.value_of_css_property("background-color")
    foe["accentColor"] = Color.from_string(bannerColor).hex
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
            player = {}
            player["squadNumber"] = row.find_element(By.CLASS_NAME, "Opta-Shirt").text
            player["name"] = row.find_element(By.CLASS_NAME, "Opta-Name").text
            player["position"] = currentPosition
            foe["players"].append(player)
    foe["players"] = sorted(foe["players"], key=lambda x: atoi(x["squadNumber"]))
    foes.append(foe)

driver.close()

foeJSON = json.dumps(foes, ensure_ascii=False)
with open('foes.json', 'w') as f:
    f.write(foeJSON)

client = MongoClient(DB_URL)
db = client.ngsdetroit

for foe in foes:
    if foe["players"] == []: continue
    dbFoe = db.foes.find_one({"opponent": re.compile('^' + foe["opponent"] + '$', re.IGNORECASE), "competition": foe["competition"]})
    if dbFoe == None:
        foe["active"] = True
        db.foes.insert_one(foe)
    else:
        dbFoe["players"] = foe["players"]
        db.foes.save(dbFoe)

# case insensitive mongodb search on foe + competition
# if present: update player list and ONLY player list, update
# if not present: insert new record with everything (name, colors, etc)
client.close()
