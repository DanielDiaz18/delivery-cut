import websocketStore from "svelte-websocket-store";

const initialValue = {};

export const myStore = websocketStore(
  "ws://localhost:8000/ws",
  initialValue,
  []
);
