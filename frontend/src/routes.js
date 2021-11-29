import Login from "./views/Login.svelte";
import Home from "./views/Home.svelte";
import Signup from "./views/Signup.svelte";
import NotFound from "./views/NotFound.svelte";
import PublicLayout from "./views/layouts/PublicLayout.svelte";
import AdminLayout from "./views/layouts/AdminLayout.svelte";
import { getUser } from "./stores/user_store";

const userIsLogged = () => !!getUser().isLoggedIn;

const userIsAdmin = () => {
  const u = getUser();
  return u.kind === "admin" && !!u.isLoggedIn;
};

export const routes = [
  {
    name: "/",
    redirectTo: "/login",
  },
  {
    name: "login",
    component: Login,
  },
  {
    name: "signup",
    component: Signup,
  },
  {
    name: "user",
    component: Home,
    onlyIf: { guard: userIsLogged, redirectTo: "/login" },
    layout: PublicLayout,
  },
  {
    name: "parner",
    component: Home,
    onlyIf: { guard: userIsLogged, redirectTo: "/login" },
    layout: PublicLayout,
  },
  {
    name: "admin",
    layout: AdminLayout,
    onlyIf: { guard: userIsAdmin, redirectTo: "/login" },
    nestedRoutes: [
      {
        name: "index",
        component: Home,
      },
      {
        name: "signup",
        component: Signup,
      },
    ],
  },
  {
    name: "*",
    component: NotFound,
  },
];
