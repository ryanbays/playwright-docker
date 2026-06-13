import { listStates, getState } from "./read.js";

export function findStates(filter = {}) {
    const files = listStates();

    return files.filter(file => {
        const state = getState(file);

        if (filter.tag) {
            return state.tags?.includes(filter.tag);
        }

        if (filter.expiresBefore) {
            return (state.expires || Infinity) < filter.expiresBefore;
        }

        return true;
    });
}
