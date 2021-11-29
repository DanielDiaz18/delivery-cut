<script>
  import { onDestroy, setContext } from "svelte";
  import { mapbox, key } from "../utils/mapbox.js";
  import axios from "axios";

  setContext(key, {
    getMap: () => map,
  });

  export let lngLat;
  export let zoom;

  let container;
  let map;

  function load() {
    map = new mapbox.Map({
      container,
      style: "mapbox://styles/mapbox/streets-v11",
      center: lngLat,
      zoom,
    });
    const bounds = [lngLat.map((n) => n - 0.5), lngLat.map((n) => n + 0.5)];
    map.setMaxBounds(bounds);
  }

  onDestroy(() => map?.remove());
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://unpkg.com/mapbox-gl/dist/mapbox-gl.css"
    on:load={load}
  />
</svelte:head>

<div class="h-full w-full" bind:this={container}>
  {#if map}
    <slot />
  {/if}
</div>
