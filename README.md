# Web API Client

## Overview

You will build a systems console application where you will request weather data
from `https://api.weather.gov` API and then display it in a human readable format.

To do this, you will:

- Interact with `geocoding.geo.census.gov` via a web API to geo-encode an address.
- Interact with `https://api.weather.gov` via a web API to get the forecasted weather based on the data from geo-coding.
- Create a human readable text output.

## Requirements

- Must use Go.
- This is a console application only.
- Must run on either Linux, Darwin or Windows.

## Steps

### Using the Geo-encoding and Weather API

#### Summary

The National Weather Service (NWS) API allows developers access to critical forecasts, alerts, and observations, along with other weather data. The API was designed with a cache-friendly approach that expires content based upon the information life cycle. The API is based upon of JSON-LD to promote machine data discovery.

The API is located at: [https://api.weather.gov](https://api.weather.gov)

#### Usage

You’ll need to know the latitude and longitude of the location in decimal degrees. (If you want to get really geo-spatially technical, your location should be a WGS 84 or EPSG 4326 coordinate.)

For our example here, we’ll use the Washington Monument in Washington, D.C. It’s located at 38.8894 latitude, -77.0352 longitude. Please note that, for efficiency purposes, the API doesn’t support more than four decimal places of precision in coordinates. If you send a more precise coordinate, you’ll receive an error giving you the closest proper coordinate. Four decimal places is about 30 feet (10 meters) over most of the United States, so that should still be close enough!

Once you know the latitude and longitude, it’s an easy three-step process from there. You can follow along in your browser with the links below:

1. Retrieve the metadata for your location from `https://api.weather.gov/points/{lat},{lon}`.
    - For our example the URL will be `https://api.weather.gov/points/38.8894,-77.0352`
2. You’ll get back a JSON document. Inside the document, find the properties object, and inside that, find the forecast property. You’ll find another URL there.
    - For our example this gives us `https://api.weather.gov/gridpoints/LWX/96,70/forecast`
    - You can also get the hour-by-hour forecast from the forecastHourly property. For our example it’s `https://api.weather.gov/gridpoints/LWX/96,70/forecast/hourly`
3. Retrieve that URL. You’ll get a JSON document containing the forecast for that location. There you go!

Please see example Web API calls in the document [wttr.http](docs/wttr.http).

### Weather Console Application

- Capture the domain in terms of the domain experts.  Domain is the business of the application.
- Embed the domain terminology and actions in the code.
- Protect domain knowledge from other domains.  (No spaghetti code, single responsibility principle).
- Questions can be asked during lab times in class.

## Submit

When you have completed your application, double-check that it runs as expected by the specifications above. Ensure your unit-tests also all pass.  When you are ready, submit the code via the following git commands:

```bash
git add -u
git tag "Submission"
git commit -m "Submission"
git push
```

You can verify your code was successfully submitted by viewing your git repo on your github account: [https://github.com](https://github.com).
