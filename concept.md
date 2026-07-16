It is going to be the project in GoLang language, which will make following:

1. Make the the local chi server on the port 8080 
2. On the page "/weather user enter in the simple form a city and confirms it with the button "Submit"
3. The handler will request for the specified city the data from the [wttr.in - Git Repository](https://github.com/chubin/wttr.in/). It will be http request with the similar result as the following curl call: 
  curl 'wttr.in/Wetzikon?format=j2' | jq '.weather[] | {
  date: .date,
  avgtempC: .avgtempC,
  sunrise: .astronomy[0].sunrise,
  sunset: .astronomy[0].sunset
  }'
4. The recieved data are saved to the json file weather_forecast.json and printed on the web page below the request form