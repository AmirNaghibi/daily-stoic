import pandas as pd
from bs4 import BeautifulSoup
from selenium import webdriver

driver = webdriver.Chrome(executable_path ='C:/Users/Charles Arday/Documents/chromedriver/chromedriver.exe')
driver.get("https://www.goalcast.com/2019/03/15/seneca-quotes/")
results =[]
content = driver.page_source 
soup = BeautifulSoup(content)
driver.quit()


for element in soup.findAll(attrs="wp-block-pullquote"):
    SenecaQuote = element.find("blockquote")
    if SenecaQuote not in results:
        results.append(SenecaQuote.text)
df = pd.DataFrame({'Seneca Quotes':results})
df.to_csv('Seneca Quotes.csv', index =False , encoding = 'utf-8')
