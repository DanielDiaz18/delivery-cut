<script>
  import { onDestroy, onMount } from "svelte";
  import { Navigate, navigateTo } from "svelte-router-spa";
  import { login, logout, userStore } from "../stores/user_store";
  import notifications from "../stores/notifications_store";

  let email;
  let password;

  let submit = async () => {
    try {
      if (!email?.trim() || !password?.trim()) {
        return notifications.error("Please fill in all fields");
      }
      await login(email, password);
    } catch (err) {
      if (err?.response) {
        return notifications.error(err.response.data.message);
      } else {
        return notifications.error(err.toString());
      }
    }
  };
  let unsubscribe;
  onMount(() => {
    logout();
    unsubscribe = userStore.subscribe(({ kind, isLoggedIn }) => {
      if (isLoggedIn && !!kind) {
        navigateTo(`/${kind}`);
      }
    });
  });

  onDestroy(() => unsubscribe?.call());
</script>

<div class="flex items-center justify-center h-screen w-screen">
  <form class="bg-white shadow-lg rounded-lg p-8 mb-4 min-w-1/3 flex flex-col">
    <h2 class="mb-8 font-bold">Login</h2>
    <div
      class="flex items-center shadow rounded w-full px-3 my-3 text-black mb-3 leading-tight focus:shadow-outline bg-gray-200"
    >
      <span class="material-icons">email</span>
      <input
        class="pl-1 appearance-none focus:outline-none py-2 bg-transparent w-full"
        id="username"
        type="email"
        placeholder="Email"
        bind:value={email}
      />
    </div>
    <div
      class="flex items-center shadow rounded w-full px-3 my-3 text-black mb-3 leading-tight focus:shadow-outline bg-gray-200"
    >
      <span class="material-icons">lock</span>
      <input
        class="pl-1 appearance-none focus:outline-none py-2 bg-transparent w-full"
        id="password"
        type="password"
        placeholder="Contraseña"
        bind:value={password}
        on:keyup={(event) => {
          if (event.keyCode == 13) {
            submit();
          }
        }}
      />
    </div>
    <div class="my-2">
      <Navigate to="/signup"
        >¿Aun no tienes cuenta? <b class="underline">Registro</b></Navigate
      >
    </div>

    <button
      class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline self-end"
      type="button"
      on:click={submit}
    >
      Iniciar sesión
    </button>
  </form>
</div>
