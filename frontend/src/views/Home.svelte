<script>
  import Map from "./Map.svelte";
  import MapMarker from "./MapMarker.svelte";
  import Direction from "./components/Direction.svelte";

  export let currentRoute;
  export let params;

  console.log(currentRoute, params);
</script>

<div class="flex flex-row h-screen w-screen">
  <Map lngLat={[-103.39182, 20.66682]} zoom={16}>
    <Direction
      origin={[-103.39182, 20.66682]}
      destination={[-103.40182, 20.67682]}
    />
    <Direction
      origin={[-103.39282, 20.6682]}
      destination={[-103.40282, 20.6682]}
    />
    <Direction
      origin={[-103.39182, 20.66782]}
      destination={[-103.41182, 20.68682]}
    />
    <MapMarker lngLat={[-103.39182, 20.66682]} color="#FF0000" />
    <MapMarker lngLat={[-103.39282, 20.66782]} color="#00FF00" />
    <MapMarker lngLat={[-103.39382, 20.66882]} color="#0000FF" />
  </Map>
</div>
