<script>
  import { Navigate, navigateTo } from "svelte-router-spa";
  import { onDestroy } from "svelte";

  import { api } from "../dependencies";
  import { userStore, getUser } from "../stores/user_store";
  import notifications from "../stores/notifications_store";

  let form = {
    name: "",
    lastName: "",
    phoneNumber: "",
    email: "",
    password: "",
    kind: "user",
  };
  let userKind;

  export let currentRoute;
  export let params = {};

  let submit = async () => {
    try {
      if (Object.values(form).some((value) => value === "")) {
        notifications.error("Todos los campos son obligatorios");
        return;
      }
      const res = await api.post(currentRoute.path, form);
      if (currentRoute.path === "/signup") {
        navigateTo("/login");
      } else {
        navigateTo("/admin");
      }
    } catch (err) {
      if (err?.response) {
        notifications.error(err.response.data.message);
      }
    }
  };
  console.log(currentRoute.path);

  const unsubscribe = userStore.subscribe(({ kind }) => (userKind = kind));

  onDestroy(() => {
    unsubscribe();
  });
</script>

<div class="flex items-center justify-center h-screen">
  <form class="bg-white shadow-lg rounded-lg p-8 mb-4 min-w-1/3 flex flex-col">
    <h3 class="mb-8 font-bold">Registro</h3>
    <div
      class="flex items-center shadow rounded w-full px-3 my-3 text-black mb-3 leading-tight focus:shadow-outline bg-gray-200"
    >
      <label for="name" class="material-icons">person</label>
      <input
        key="name"
        class="pl-1 appearance-none focus:outline-none py-2 bg-transparent w-full"
        type="text"
        placeholder="Nombre"
        bind:value={form.name}
      />
    </div>
    <div
      class="flex items-center shadow rounded w-full px-3 my-3 text-black mb-3 leading-tight focus:shadow-outline bg-gray-200"
    >
      <label for="lastName" class="material-icons">person</label>
      <input
        key="lastName"
        class="pl-1 appearance-none focus:outline-none py-2 bg-transparent w-full"
        type="text"
        placeholder="Apellidos"
        bind:value={form.lastName}
      />
    </div>
    <div
      class="flex items-center shadow rounded w-full px-3 my-3 text-black mb-3 leading-tight focus:shadow-outline bg-gray-200"
    >
      <label for="email" class="material-icons">email</label>
      <input
        key="email"
        class="pl-1 appearance-none focus:outline-none py-2 bg-transparent w-full"
        id="username"
        type="email"
        placeholder="Email"
        bind:value={form.email}
      />
    </div>
    <div
      class="flex items-center shadow rounded w-full px-3 my-3 text-black mb-3 leading-tight focus:shadow-outline bg-gray-200"
    >
      <label for="password" class="material-icons">lock</label>
      <input
        key="password"
        class="pl-1 appearance-none focus:outline-none py-2 bg-transparent w-full"
        type="password"
        placeholder="Contraseña"
        bind:value={form.password}
      />
    </div>
    <div
      class="flex items-center shadow rounded w-full px-3 my-3 text-black mb-3 leading-tight focus:shadow-outline bg-gray-200"
    >
      <label for="phoneNumber" class="material-icons">phone</label>
      <input
        key="phoneNumber"
        class="pl-1 appearance-none focus:outline-none py-2 bg-transparent w-full"
        type="tel"
        placeholder="Celular"
        bind:value={form.phoneNumber}
      />
    </div>
    {#if currentRoute.path === "/admin/signup"}
      <div
        class="flex items-center shadow rounded w-full px-3 my-3 text-black mb-3 leading-tight focus:shadow-outline bg-gray-200"
      >
        <label class="material-icons" for="kind">account_circle</label>
        <select
          key="kind"
          name="kind"
          class="w-full py-2 bg-transparent focus:outline-none"
          bind:value={form.kind}
        >
          <option value="parner">socio</option>
          <option value="admin">admin</option>
        </select>
      </div>
    {:else}
      <div class="mb-6">
        <Navigate to="/login"
          >¿Ya tienes cuenta? <b class="underline">Inicia sesión</b></Navigate
        >
      </div>
    {/if}
    <button
      class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline self-end"
      type="button"
      on:click={submit}
    >
      Registrarse
    </button>
  </form>
</div>
