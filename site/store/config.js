export const state = () => ({
  config: {}
});

export const mutations = {
  setConfig(state, config) {
    state.config = config;
  }
};

export const actions = {
  async loadConfig(context) {
    console.log(this.$axios.get);
    let get = this.$axios.get.bind(this);
    const ret = await get("/api/configs");
    context.commit("setConfig", ret);
    return ret;
  }
};

export const getters = {
  siteTitle(state) {
    return state.config.siteTitle || "";
  },
  siteDescription(state) {
    return state.config.siteDescription || "";
  },
  siteKeywords(state) {
    return state.config.siteKeywords || "";
  }
};
