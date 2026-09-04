import { createTool } from "@mastra/core/tools";
import { z } from "zod";

function weatherCodeToCondition(code: number): string {
  if (code === 0) return "Clear";
  if (code <= 3) return "Partly cloudy";
  if (code <= 48) return "Foggy";
  if (code <= 67) return "Rain";
  if (code <= 77) return "Snow";
  if (code <= 82) return "Rain showers";
  if (code <= 86) return "Snow showers";
  if (code <= 99) return "Thunderstorm";
  return "Unknown";
}

export const weatherTool = createTool({
  id: "get-weather",
  description:
    "Get current weather for a city including temperature, conditions, and rain probability",
  inputSchema: z.object({
    city: z.string().describe("City name, e.g. Bangalore, London, Tokyo"),
  }),
  outputSchema: z.object({
    city: z.string(),
    temperature: z.number(),
    condition: z.string(),
    error: z.string().optional(),
  }),

  execute: async ({ city }) => {
    try {
      const geoResponse = await fetch(
        `https://geocoding-api.open-meteo.com/v1/search?name=${encodeURIComponent(city)}&count=1`,
      );

      if (!geoResponse.ok) {
        return {
          city,
          temperature: 0,
          condition: "Unavailable",
          error: `count not look the city: ${city}`,
        };
      }

      const geoData = (await geoResponse.json()) as {
        results?: Array<{ name: string; latitude: number; longitude: number }>;
      };

      const location = geoData.results?.[0];

      if (!location) {
        return {
          city,
          temperature: 0,
          condition: "Unknown city",
          error: `count not look the city: ${city}`,
        };
      }

      const weatherResponse = await fetch(
        `https://api.open-meteo.com/v1/forecast?latitude=${location.latitude}&longitude=${location.longitude}&current=temperature_2m,weather_code,precipitation_probability`,
      );

      if (!weatherResponse.ok) {
        return {
          city: location.name,
          temperature: 0,
          condition: "Unavailable",
          error: "Weather api request failed",
        };
      }

      const weatherData = (await weatherResponse.json()) as {
        current: {
          temperature_2m: number;
          weather_code: number;
        };
      };

      const curent = weatherData.current;

      return {
        city: location.name,
        temperature: curent.temperature_2m,
        condition: weatherCodeToCondition(curent.weather_code),
      };
    } catch {
      return {
        city,
        temperature: 0,
        condition: "Unavailable",
        error: `Failed to fetch weather info for this city: ${city}`,
      };
    }
  },
});
