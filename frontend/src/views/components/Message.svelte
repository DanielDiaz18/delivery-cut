<script>
  import { onMount, createEventDispatcher } from "svelte";

  export let notification = {};
  export let index = 0;
  const dispatch = createEventDispatcher();
  function closeElement() {
    dispatch("closeNotification", index);
  }
  onMount(function () {
    setTimeout(function () {
      dispatch("closeNotification", index);
    }, 5000);
  });
</script>

<div
  class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative my-2"
  role="alert"
>
  <strong class="font-bold">{notification.title ?? ""}</strong>
  <span class="block sm:inline">{notification.body}</span>
  <span
    on:click={closeElement}
    class="block absolute top-0 bottom-0 right-0 px-4 py-3 material-icons"
  >
    close
  </span>
</div>
