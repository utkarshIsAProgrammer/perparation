import { Mastra } from "@mastra/core";
import { personalAssistant } from "./agents/personal-assistant";
import { DuckDBStore } from "@mastra/duckdb";
import { MastraCompositeStore } from "@mastra/core/storage";
import { LibSQLStore } from "@mastra/libsql";
import { MastraStorageExporter, Observability } from "@mastra/observability";

const observabilityStore = await new DuckDBStore().getStore("observability");

export const mastra = new Mastra({
  storage: new MastraCompositeStore({
    id: "composite-storage",
    default: new LibSQLStore({
      id: "mastra-storage",
      url: "file:./mastra.db",
    }),
    domains: {
      observability: observabilityStore,
    },
  }),
  observability: new Observability({
    configs: {
      default: {
        serviceName: "agentic-ai-crash-course-2026",
        exporters: [
          new MastraStorageExporter(), // Persists observability events to Mastra Storage
        ],
      },
    },
  }),
  agents: { personalAssistant },
});
