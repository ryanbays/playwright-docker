import fs from "fs";
import path from "path";

const STATE_DIR = process.env.STATE_DIR || "/auth";

// Returns the UUIDs of all the saved states found in the STATE_DIR directory.
export function listStates() {
    if (!fs.existsSync(STATE_DIR)) return [];
    return fs.readdirSync(STATE_DIR)
        .filter(f => f.endsWith(".json")).forEach(f => {
            const id = path.basename(f, ".json");
            return id;
        });
}

// Returns the state object for the given UUID, or throws an error if not found.
export function getState(uuid) {
    const filePath = path.join(STATE_DIR, `${uuid}.json`);

    if (!fs.existsSync(filePath)) {
        throw new Error(`State not found: ${uuid}`);
    }

    const raw = fs.readFileSync(filePath, "utf-8");
    return JSON.parse(raw);
}

// Deletes the state file for the given UUID. Returns true if deleted, false if not found.
export function deleteState(uuid) {
    const filePath = path.join(STATE_DIR, `${uuid}.json`);

    if (!fs.existsSync(filePath)) return false;

    fs.unlinkSync(filePath);
    return true;
}
