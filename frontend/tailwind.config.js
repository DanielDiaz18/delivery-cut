const production = !process.env.ROLLUP_WATCH;

module.exports = {
  mode: "jit",
  future: {
    purgeLayersByDefault: true,
    removeDeprecatedGapUtilities: true,
  },
  theme: {
    container: {
      center: true,
    },
    minWidth: {
      0: "0",
      "1/4": "25%",
      "1/3": "33.333%",
      "1/2": "50%",
      "3/4": "75%",
      full: "100%",
    },
  },
  plugins: [],
  purge: {
    content: ["./src/App.svelte", "./src/**/*.svelte"],
    enabled: production, // disable purge in dev
  },
};
