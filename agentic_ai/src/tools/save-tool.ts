import { createTool } from "@mastra/core/tools";
import { readFile, writeFile, mkdir } from "node:fs/promises";
import { join } from "node:path";
import { z } from "zod";

const NOTES_FILE = join(process.cwd(), "data/notes.json");

type Note = {
  id: string;
  content: string;
  title?: string;
  createdAt: string;
};

async function loadNotes(): Promise<Note[]> {
  try {
    const raw = await readFile(NOTES_FILE, "utf-8");
    return JSON.parse(raw) as Note[];
  } catch {
    await mkdir(join(process.cwd(), "data"), { recursive: true });
    await writeFile(NOTES_FILE, [], "utf-8");
    return [];
  }
}

export const saveNoteTool = createTool({
  id: "save-note",
  description:
    "Save a note or reminder. Use when the user wants something remembered or written down.",
  inputSchema: z.object({
    content: z.string().describe("The note or reminder text to save"),
    title: z.string().optional().describe("Optional short title for the note"),
  }),
  outputSchema: z.object({
    success: z.boolean(),
    message: z.string(),
    noteId: z.string().optional(),
  }),

  execute: async ({ content, title }) => {
    try {
      const notes = await loadNotes();

      const note: Note = {
        id: crypto.randomUUID(),
        content,
        title,
        createdAt: new Date().toISOString(),
      };

      notes.push(note);
      await writeFile(NOTES_FILE, JSON.stringify(notes, null, 2), "utf-8");

      return {
        success: true,
        message: `Note saved: "${content}"`,
        noteId: note.id,
      };
    } catch {
      return {
        success: false,
        message: "Failed to save a note",
      };
    }
  },
});
