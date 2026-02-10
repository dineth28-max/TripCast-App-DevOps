# 🌦️ TripCast – Travel Weather & Location App

TripCast is a travel-focused web application that helps users search locations, view weather forecasts, and receive location suggestions. The system is built using a **Microservices Architecture** with a **React Frontend** and a **GraphQL API Gateway**.

---

## 🚀 Project Overview

TripCast aims to make travel planning easier and smarter by offering:  
- Location search and navigation  
- Weather forecasts (current and 5-day)  
- AI/rule-based location suggestions  
- User authentication for extended features  

---

## 🏗️ Core Features

### 1. Location Search
- Users can search for a city or tourist place.  
- Guest users see **only 1 suggested location**.  

### 2. Location Result
- Displays location name and a short description  
- Shows a map navigation option  

### 3. Weather Forecast
- Current temperature  
- Weather conditions (Rain, Sunny, Cloudy, etc.)  
- 5-day forecast  

### 4. Location Suggestions
- AI/Rule-based recommendation system  
- Guest Users: See only 1 suggested location  
- Logged-in Users: Can view all suggested locations  

### 5. Authentication
- Sign Up / Login with JWT  
- Enables full access to features  

### 6. Navigation
- After selecting a location, users can:  
  - Open a map  
  - See route directions  
  - View weather on the destination page  

---

## 🏗️ System Architecture

**Architecture Style:** Microservices + API Gateway  

### Components

1. **Frontend (React JS)**
   - Handles user input, routing, and display  
   - Communicates only with the API Gateway  

2. **API Gateway (GraphQL)**
   - Central communication layer  
   - Aggregates data from all microservices  
   - Handles authentication and request routing  

3. **Weather Service**
   - Fetches weather data from external APIs  
   - Provides current and forecast weather  

4. **Location & Recommendation Service**
   - Manages search results  
   - Suggests locations based on user behavior or rules  

5. **Authentication Service**
   - Handles user registration and login  
   - Generates JWT tokens  

---

## 🔹 Microservices Details

### 1. Weather Service
- Retrieve live weather data  
- Cache results for performance  
- Provide forecast data  
- Example APIs: OpenWeatherMap, WeatherAPI  

### 2. Location & Recommendation Service
- Location search  
- Recommendation logic  
- Guest vs Logged-in visibility rules  

### 3. Authentication Service
- User sign-up and login validation  
- Password hashing  
- Token generation (JWT)  
- Role/Access control (optional)  

---

## 🖥️ Frontend (React JS)

### Pages
- Home Page  
- Search Page  
- Location Details Page  
- Login / Register Page  
- Profile (Optional)  

### Key Components
- Search Bar  
- Weather Card  
- Map View  
- Login Modal  

---

## 📡 API Gateway (GraphQL)

### Responsibilities
- Single entry point for frontend  
- Query aggregation  
- Authorization check  
- Rate limiting  

### Example Queries
```graphql
getWeather(location)
searchLocation(keyword)
getRecommendations(userId)
login(email, password)
