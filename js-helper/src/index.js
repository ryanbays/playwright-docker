import { listStates, getState, deleteState } from "./read.js";
import { saveState } from "./write.js";
import { findStates } from "./query.js";

// Named exports — `import { listStates } from "..."`
export {
    listStates,
    getState,
    deleteState,
    saveState,
    findStates,
};

// Default export — `import storageState from "..."` then `storageState.listStates()`
const storageState = {
    listStates,
    getState,
    deleteState,
    saveState,
    findStates,
};

export default storageState;
