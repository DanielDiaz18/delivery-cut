<script>
  import { getContext, onMount } from "svelte";
  import { mapbox, key } from "../../utils/mapbox";
  import axios from "axios";

  export let origin;
  export let destination;
  export let color = "#3887be";

  const { getMap } = getContext(key);
  const map = getMap();

  let getDirection = async () => {
    if (origin && destination) {
      const response = await axios.get(
        `https://api.mapbox.com/directions/v5/mapbox/driving/${origin.join(
          ","
        )};${destination.join(",")}`,
        {
          params: {
            geometries: "geojson",
            access_token:
              "pk.eyJ1IjoiZGFuaWVsZGlhejE4IiwiYSI6ImNrd2Z0ZnVydjBpMGwzMXB3Z3M5aDRubnEifQ.RyssATf1GE2mf6lKBnGNqA",
          },
        }
      );
      const route = response.data.routes[0].geometry.coordinates;
      const geojson = {
        type: "Feature",
        properties: {},
        geometry: {
          type: "LineString",
          coordinates: route,
        },
      };
      let id = `route-[${destination.join(",")}]`;
      if (map.getSource(id)) {
        map.getSource(id).setData(geojson);
      } else {
        map.addLayer({
          id: id,
          type: "line",
          source: {
            type: "geojson",
            data: geojson,
          },
          layout: {
            "line-join": "round",
            "line-cap": "round",
          },
          paint: {
            "line-color": color,
            "line-width": 5,
            "line-opacity": 0.75,
          },
        });
      }
    }
  };

  map.on("load", () => {
    getDirection();

    // map.addLayer({
    //   id: "point",
    //   type: "circle",
    //   source: {
    //     type: "geojson",
    //     data: {
    //       type: "FeatureCollection",
    //       features: [
    //         {
    //           type: "Feature",
    //           properties: {},
    //           geometry: {
    //             type: "Point",
    //             coordinates: origin,
    //           },
    //         },
    //       ],
    //     },
    //   },
    //   paint: {
    //     "circle-radius": 10,
    //     "circle-color": color,
    //   },
    // });
  });
</script>
