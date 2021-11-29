import { writable, get } from "svelte/store";
import { api } from "../dependencies";

let _user = JSON.parse(localStorage.getItem("user") ?? "{}");

export const userStore = writable(
  _user ?? {
    name: "",
    email: "",
    password: "",
    isLoggedIn: false,
  }
);

export async function login(email, password) {
  const res = await api.post("/login", {
    email,
    password,
  });
  const { token, user } = res.data;
  api.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  console.log("data:", user);
  userStore.set({ ...user, isLoggedIn: true });
}

export async function logout() {
  delete api.defaults.headers.common["Authorization"];
  userStore.set({ isLoggedIn: false });
}

export function getUser() {
  return get(userStore);
}

userStore.subscribe((user) => {
  console.log(user);
  if (user?.isLoggedIn) {
    localStorage.setItem("user", JSON.stringify(user));
  } else {
    api.headers?.remove("Authorization");
    localStorage.removeItem("user");
  }
});
