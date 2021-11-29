<script>
  import { Router } from "svelte-router-spa";
  import { routes } from "./routes";
  import Notifications from "./views/components/Notifications.svelte";
  import { myStore } from "./stores/websocket_store";

  myStore.subscribe(console.log);
</script>

<div class="bg-gray-100 ">
  <Notifications />
  <Router {routes} />
</div>
