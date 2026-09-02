import {Mastra} from "@mastra/core"
import {personalAssistant} from "./agents/personal-assistant"

export const mastra = new Mastra({
    agents: { personalAssistant},
})