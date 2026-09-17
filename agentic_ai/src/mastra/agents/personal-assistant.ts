import { Agent } from "@mastra/core/agent";
import { personalAssistantInstructions } from "./instructions";
import { weatherTool } from "../../tools/weather-tool";
import { saveNoteTool } from "../../tools/save-tool";
import { Memory } from "@mastra/memory";

export const personalAssistant = new Agent({
  id: "personal-assistant",
  name: "Personal Assistant",
  instructions: personalAssistantInstructions,
  memory: new Memory({
    options: {
      lastMessages: 20,
    },
  }),
  model: "google/gemini-2.5-flash",
  tools: {
    getWeather: weatherTool,
    saveNote: saveNoteTool,
  },
});
