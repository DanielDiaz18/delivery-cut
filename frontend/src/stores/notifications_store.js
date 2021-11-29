import { writable } from "svelte/store";

function createStore() {
  const { subscribe, set, update } = writable([]);
  return {
    subscribe,
    set,
    add: (e) => update((n) => [...n, e]),
    error: (e) => update((n) => [...n, { body: e }]),
    remove: (i) => update((n) => n.filter((e, j) => j !== i)),
  };
}

const notificationsStore = createStore();

export default notificationsStore;
