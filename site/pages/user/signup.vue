<template>
  <section class="main">
    <div class="container">
      <div class="main-body no-bg">
        <div class="widget signup">
          <div class="widget-header">注册</div>
          <div class="widget-content">
            <div class="field">
              <label class="label">昵称</label>
              <div class="control has-icons-left">
                <input
                  v-model="username"
                  class="input is-success"
                  type="text"
                  placeholder="请输入昵称"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-username" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">邮箱</label>
              <div class="control has-icons-left">
                <input
                  v-model="email"
                  class="input is-success"
                  type="text"
                  placeholder="请输入邮箱"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-email" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">密码</label>
              <div class="control has-icons-left">
                <input
                  v-model="password"
                  class="input"
                  type="password"
                  placeholder="请输入密码"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">确认密码</label>
              <div class="control has-icons-left">
                <input
                  v-model="rePassword"
                  class="input"
                  type="password"
                  placeholder="请再次输入密码"
                  @keyup.enter="signup"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>

            <div class="field">
              <div class="control">
                <button class="button is-success" @click="signup">注册</button>
                <github-login :ref-url="ref" />
                <qq-login :ref-url="ref" />
              </div>
            </div>

            <div class="field">
              <nuxt-link class="button is-text" to="/user/signin">已有账号，前往登录&gt;&gt;</nuxt-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  components: {},
  asyncData({ params, query }) {
    return {
      ref: query.ref,
    }
  },
  data() {
    return {
      username: '',
      email: '',
      password: '',
      rePassword: '',
    }
  },
  mounted() {},
  methods: {
    async signup() {
      try {
        await this.$store.dispatch('user/signup', {
          username: this.username,
          email: this.email,
          password: this.password,
          rePassword: this.rePassword,
          ref: this.ref,
        })
        if (this.ref) {
          // 跳到登录前
          this.$linkTo(this.ref)
        } else {
          // 跳到个人主页
          this.$linkTo('/')
        }
      } catch (err) {
        this.$message.error(err.message || err)
      }
    },
  },
  head() {
    return {
      title: this.$siteTitle('注册'),
    }
  },
}
</script>

<style lang="scss" scoped>
.signup {
  max-width: 480px;
  margin: auto;
  padding: 0 20px;
}
</style>
