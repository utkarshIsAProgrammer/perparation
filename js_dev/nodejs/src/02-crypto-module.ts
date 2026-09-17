// Hashing
// Authentication
// Encryption
// Random Number Generation

import crypto from "crypto";

// --------------------------------------------------------------------------
// 1. Random UUID
const uuid = crypto.randomUUID();
console.log("UUID: ", uuid);

// 2. Random hex token
const hexToken = crypto.randomBytes(32).toString("hex");
console.log("Hex Token: ", hexToken);

// 3. Random integer
const otp = crypto.randomInt(100000, 1000000);
console.log("OTP: ", otp);
// --------------------------------------------------------------------------

// --------------------------------------------------------------------------
// 4. Hashing data
const generateHash = (data: string) => {
    return crypto.createHash("sha256").update(data).digest("hex");
};

// createHash(sha256 or sha512) # hash algorithm
// update(data) # feed data (string or buffer)
// digest("hex" or "base64") # output format

const fileContent = "This data is meant to be hashed!";
const result = generateHash(fileContent);
console.log("Hashed Result: ", result);
// --------------------------------------------------------------------------

// --------------------------------------------------------------------------
// 5. Hash password
const hashPassword = (password: string) => {
    const salt = crypto.randomBytes(16).toString("hex");

    crypto.scrypt(password, salt, 64, (err, derivedKey) => {
        if (err) throw err;

        const hash = derivedKey.toString("hex");

        console.log("Password Salt:", salt);
        console.log("Password Hash:", hash);

        // Verify using the generated salt and hash
        verifyPassword("password", salt, hash);
    });
};

// 6. Verify password
const verifyPassword = (
    passwordAttempt: string,
    storedSalt: string,
    storedHash: string,
) => {
    crypto.scrypt(passwordAttempt, storedSalt, 64, (err, derivedKey) => {
        if (err) throw err;

        const storedHashBuffer = Buffer.from(storedHash, "hex");

        const passwordIsCorrect = crypto.timingSafeEqual(
            derivedKey,
            storedHashBuffer,
        );

        console.log(
            passwordIsCorrect
                ? "Password is correct!"
                : "Password is incorrect!",
        );
    });
};

// 7. Hash the password
hashPassword("password");

// --------------------------------------------------------------------------
