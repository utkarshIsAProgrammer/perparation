import { Mastra } from '@mastra/core';
import { Agent } from '@mastra/core/agent';
import { createTool } from '@mastra/core/tools';
import { z } from 'zod';
import { readFile, mkdir, writeFile } from 'node:fs/promises';
import { join } from 'node:path';
import { Memory } from '@mastra/memory';
import { DuckDBStore } from '@mastra/duckdb';
import { MastraCompositeStore } from '@mastra/core/storage';
import { LibSQLStore } from '@mastra/libsql';
import { Observability, MastraStorageExporter } from '@mastra/observability';

"use strict";
const personalAssistantInstructions = `
## Role
You are a helpful personal assistant.

## Style
Answer clearly and concisely. Be friendly but direct.

## Memory
Use earlier messages in the conversation when answering follow-up questions.
If the user refers to something mentioned before (like "there" or "tomorrow"), use that context.

## Weather tool
When asked about current weather, use the getWeather tool to fetch live data.
Do not guess or invent weather information \u2014 always use the tool for live conditions.

## Notes tool
When asked to save a note or reminder, use the saveNote tool.
After saving, confirm what was saved in a friendly way.

## Multi-step tasks
For tasks that need multiple steps \u2014 for example, checking weather and then saving a note based on the result \u2014
use getWeather first, read the result, then call saveNote only if the condition is met.
`.trim();

"use strict";
function weatherCodeToCondition(code) {
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
const weatherTool = createTool({
  id: "get-weather",
  description: "Get current weather for a city including temperature, conditions, and rain probability",
  inputSchema: z.object({
    city: z.string().describe("City name, e.g. Bangalore, London, Tokyo")
  }),
  outputSchema: z.object({
    city: z.string(),
    temperature: z.number(),
    condition: z.string(),
    error: z.string().optional()
  }),
  execute: async ({ city }) => {
    try {
      const geoResponse = await fetch(
        `https://geocoding-api.open-meteo.com/v1/search?name=${encodeURIComponent(city)}&count=1`
      );
      if (!geoResponse.ok) {
        return {
          city,
          temperature: 0,
          condition: "Unavailable",
          error: `count not look the city: ${city}`
        };
      }
      const geoData = await geoResponse.json();
      const location = geoData.results?.[0];
      if (!location) {
        return {
          city,
          temperature: 0,
          condition: "Unknown city",
          error: `count not look the city: ${city}`
        };
      }
      const weatherResponse = await fetch(
        `https://api.open-meteo.com/v1/forecast?latitude=${location.latitude}&longitude=${location.longitude}&current=temperature_2m,weather_code,precipitation_probability`
      );
      if (!weatherResponse.ok) {
        return {
          city: location.name,
          temperature: 0,
          condition: "Unavailable",
          error: "Weather api request failed"
        };
      }
      const weatherData = await weatherResponse.json();
      const curent = weatherData.current;
      return {
        city: location.name,
        temperature: curent.temperature_2m,
        condition: weatherCodeToCondition(curent.weather_code)
      };
    } catch {
      return {
        city,
        temperature: 0,
        condition: "Unavailable",
        error: `Failed to fetch weather info for this city: ${city}`
      };
    }
  }
});

"use strict";
const NOTES_FILE = join(process.cwd(), "data/notes.json");
async function loadNotes() {
  try {
    const raw = await readFile(NOTES_FILE, "utf-8");
    return JSON.parse(raw);
  } catch {
    await mkdir(join(process.cwd(), "data"), { recursive: true });
    await writeFile(NOTES_FILE, [], "utf-8");
    return [];
  }
}
const saveNoteTool = createTool({
  id: "save-note",
  description: "Save a note or reminder. Use when the user wants something remembered or written down.",
  inputSchema: z.object({
    content: z.string().describe("The note or reminder text to save"),
    title: z.string().optional().describe("Optional short title for the note")
  }),
  outputSchema: z.object({
    success: z.boolean(),
    message: z.string(),
    noteId: z.string().optional()
  }),
  execute: async ({ content, title }) => {
    try {
      const notes = await loadNotes();
      const note = {
        id: crypto.randomUUID(),
        content,
        title,
        createdAt: (/* @__PURE__ */ new Date()).toISOString()
      };
      notes.push(note);
      await writeFile(NOTES_FILE, JSON.stringify(notes, null, 2), "utf-8");
      return {
        success: true,
        message: `Note saved: "${content}"`,
        noteId: note.id
      };
    } catch {
      return {
        success: false,
        message: "Failed to save a note"
      };
    }
  }
});

"use strict";
const personalAssistant = new Agent({
  id: "personal-assistant",
  name: "Personal Assistant",
  instructions: personalAssistantInstructions,
  memory: new Memory({
    options: {
      lastMessages: 20
    }
  }),
  model: "openai/gpt-4o-mini",
  tools: {
    getWeather: weatherTool,
    saveNote: saveNoteTool
  }
});

"use strict";
const observabilityStore = await new DuckDBStore().getStore("observability");
const mastra = new Mastra({
  storage: new MastraCompositeStore({
    id: "composite-storage",
    default: new LibSQLStore({
      id: "mastra-storage",
      url: "file:./mastra.db"
    }),
    domains: {
      observability: observabilityStore
    }
  }),
  observability: new Observability({
    configs: {
      default: {
        serviceName: "agentic-ai-crash-course-2026",
        exporters: [
          new MastraStorageExporter()
          // Persists observability events to Mastra Storage
        ]
      }
    }
  }),
  agents: {
    personalAssistant
  }
});

export { mastra };
