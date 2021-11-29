<script>
  import { Route, navigateTo } from "svelte-router-spa";
  import Drawer from "../components/Drawer.svelte";
  import NavBar from "../components/NavBar.svelte";
  import { userStore, logout } from "../../stores/user_store";

  export let currentRoute;
  export let params = {};
  let drawerOpen = false;

  let toogleDrawer = () => {
    drawerOpen = !drawerOpen;
  };

  userStore.subscribe(({ isLoggedIn }) => {
    if (!isLoggedIn) {
      navigateTo("/login");
    }
  });

  let logoutUser = () => {
    logout();
  };
</script>

<div class="bg-gray-100 app">
  <NavBar onOpenDrawer={toogleDrawer}>
    <button on:click={logoutUser}>
      <span class="material-icons text-black">sign_out</span>
    </button>
  </NavBar>
  <Drawer isOpen={drawerOpen} />
  <section class="section">
    <Route {currentRoute} {params} />
  </section>
</div>
