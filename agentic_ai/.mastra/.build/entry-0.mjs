import { Mastra } from '@mastra/core';
import { Agent } from '@mastra/core/agent';

"use strict";
const personalAssistant = new Agent({
  id: "personal-assistant",
  name: "Personal Assistant",
  instructions: "You are a personal assistant. Answer clearly and concisely. When you don't have access to real-time information, do not pretend that you do. If asked about live data such as current weather, explain you cannot access it yet. ",
  model: "google/gemini-3.7-flash"
});

"use strict";
const mastra = new Mastra({
  agents: {
    personalAssistant
  }
});

export { mastra };
