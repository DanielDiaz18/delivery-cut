<script>
  export let onClick = () => {};
  export let title = "";
  export let subtitle = "";
</script>

<div
  on:click={onClick}
  class="text-black hover:bg-gray-200 flex flex-row justify-center items-center p-2"
>
  <div v class="flex-none m-auto">
    <slot name="leading" />
  </div>
  <div class=" flex-grow">
    <h3 class="text-bold">{title}</h3>
    <p class="bold">{subtitle}</p>
  </div>
  <slot class="flex-none" name="trailing" />
</div>
