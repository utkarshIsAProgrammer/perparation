// Environment
// CLI Argument
// Lifecycle
// Directory
// Process Info
// Platform
// Architecture
// Resource Usage
// Async Queue
// Standard I/O

// --------------------------------------------------------------------------
const PORT = process.env.PORT;
console.log(PORT);

const NODE_ENV = process.env.NODE_ENV;
console.log(NODE_ENV);

/* 
// SETTING ENV VARIABLES IN TERMINAL
export PORT=5450
export NODE_ENV=development 
*/
// --------------------------------------------------------------------------

// --------------------------------------------------------------------------
const args = process.argv.slice(2);
console.log(args);

const userArg = args.find((arg) => arg.startsWith("--user="));
const isDryRun = args.includes("--dry-run");
const username = userArg ? userArg.split("=")[1] : "guest";

console.log(`Username: ${username}`);
console.log(`Dry Run: ${isDryRun}`);

// node src/01-process-object.ts --user=indiedev --dry-run
// --------------------------------------------------------------------------

// --------------------------------------------------------------------------
const executionDir = process.cwd();
console.log(executionDir);
// --------------------------------------------------------------------------

// --------------------------------------------------------------------------
const connectDB = async (dbUrl: string) => {
    if (!dbUrl) {
        console.log("No 'dbUrl' Provided!");
        process.exit(1);
    } else {
        console.log("DB connected successfully!");
        process.exit(0);
    }
};
connectDB("connectDb://dbIsConnecting.com");
// --------------------------------------------------------------------------
