import fs from "fs";
import path from "path";

const STATE_DIR = process.env.STATE_DIR || "/auth";

export function saveState(name, data) {
    const filePath = path.join(STATE_DIR, name);

    fs.mkdirSync(STATE_DIR, { recursive: true });

    fs.writeFileSync(
        filePath,
        JSON.stringify(data, null, 2),
        "utf-8"
    );

    return true;
}
