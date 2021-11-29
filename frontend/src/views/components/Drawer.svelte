<script>
  import { slide } from "svelte/transition";
  import { quintOut } from "svelte/easing";

  export let isOpen;

  const closeDrawer = () => {
    isOpen = false;
  };
</script>

{#if isOpen}
  <div
    on:click={closeDrawer}
    class="fixed h-full w-screen bg-black bg-opacity-25 z-10"
  >
    <aside
      transition:slide={{ delay: 250, duration: 300, easing: quintOut }}
      class="bg-white fixed top-12 left-0 w-1/4 h-full overflow-auto"
    >
      <slot />
    </aside>
  </div>
{/if}
