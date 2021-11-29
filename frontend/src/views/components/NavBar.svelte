<script>
  export let onOpenDrawer;
</script>

<nav class="flex w-screen h-12 shadow fixed z-20 justify-between bg-white">
  <button on:click={onOpenDrawer}>
    <span class="material-icons p-4">view_headline</span>
  </button>
  <slot />
</nav>
