<script>
  import { Route, Navigate } from "svelte-router-spa";
  import Drawer from "../components/Drawer.svelte";
  import ListTile from "../components/ListTile.svelte";
  import NavBar from "../components/NavBar.svelte";

  export let currentRoute;
  export let params = {};

  let drawerOpen = false;

  let toogleDrawer = () => {
    drawerOpen = !drawerOpen;
  };
</script>

<div class="bg-gray-100">
  <NavBar onOpenDrawer={toogleDrawer}>
    <button>
      <span class="material-icons">sign_out</span>
    </button>
  </NavBar>
  <Drawer isOpen={drawerOpen}>
    <Navigate to="/admin">
      <ListTile title="home">
        <span slot="leading" class="material-icons">home</span>
      </ListTile>
    </Navigate>
    <Navigate to="/admin/signup">
      <ListTile title="Registro">
        <span slot="leading" class="material-icons">person</span>
      </ListTile>
    </Navigate>
  </Drawer>
  <Route {currentRoute} {params} />
</div>
