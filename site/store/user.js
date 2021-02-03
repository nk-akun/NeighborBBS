import qs from 'qs'

export const state = () => ({
  current: null,
  userToken: null,
})

export const mutations = {
  setCurrent(state, user) {
    state.current = user
  },
  setUserToken(state, userToken) {
    state.userToken = userToken
  },
}

export const actions = {
  // 登录成功
  loginSuccess(context, { token, user }) {
    // console.log(token, user)
    const config = context.rootState.config.config
    const cookieMaxAge = 86400 * config.tokenExpireDays
    this.$cookies.set('userToken', token, { maxAge: cookieMaxAge, path: '/' })
    context.commit('setUserToken', token)
    context.commit('setCurrent', user)
  },

  // 获取当前登录用户
  async getCurrentUser(context) {
    const user = await this.$axios.get('/api/user/current')
    context.commit('setCurrent', user)
    return user
  },

  // 登录
  async signin(context, { username, password }) {
    const data = { username: username, password: password }
    const ret = await this.$axios.post('/api/user/login', data)
    context.dispatch('loginSuccess', ret)
    return ret.user
  },

  async signup(context, { nickname, username, email, password, rePassword }) {
    const ret = await this.$axios.post('/api/user/register', {
      nickname,
      username,
      email,
      password,
      rePassword,
    })
    context.dispatch('loginSuccess', ret)
    return ret.user
  },

  // 退出登录
  async signout(context) {
    const userToken = this.$cookies.get('userToken')
    await this.$axios.get('/api/user/logout', {
      params: {
        userToken,
      },
    })
    context.commit('setUserToken', null)
    context.commit('setCurrent', null)
    this.$cookies.remove('userToken')
  },
}
