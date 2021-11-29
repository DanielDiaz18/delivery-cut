<script>
  import { getContext, createEventDispatcher } from "svelte";
  import { mapbox, key } from "../utils/mapbox";

  const { getMap } = getContext(key);
  const map = getMap();

  export let lngLat;
  export let color;
  export let id = `(${lngLat})`;

  const marker = new mapbox.Marker({ color }).setLngLat(lngLat).addTo(map);
  const e = marker.getElement();
  if (id) {
    e.id = id;
  }
  e.style.cursor = "pointer";
  const dispatch = createEventDispatcher();

  e.addEventListener("click", () => dispatch("markerClick", e.id));
</script>
