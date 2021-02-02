<template>
  <section class="main">
    <div class="container">
      <div class="main-body no-bg">
        <div class="widget signin">
          <div class="widget-header">登录</div>
          <div class="widget-content">
            <div class="field">
              <label class="label">用户名/邮箱</label>
              <div class="control has-icons-left">
                <input
                  v-model="username"
                  class="input is-success"
                  type="text"
                  placeholder="请输入用户名或邮箱"
                  @keyup.enter="submitLogin"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-username" />
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
                  @keyup.enter="submitLogin"
                />
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>

            <div class="field login-button">
              <button class="button is-success" @click="submitLogin">登录</button>
              <nuxt-link class="button is-text" to="/user/signup">没有账号？点击这里去注册&gt;&gt;</nuxt-link>
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
    // console.log('这是登录页面')
    // console.log(query.ref)
    return {
      ref: query.ref,
    }
  },
  data() {
    return {
      username: '',
      password: '',
    }
  },
  computed: {
    currentUser() {
      return this.$store.state.user.current
    },
    isLogin() {
      return !!this.currentUser
    },
  },
  mounted() {
    if (this.redirectIfLogined()) {
      return
    }
  },
  methods: {
    async submitLogin() {
      try {
        if (!this.username) {
          this.$message.error('请输入用户名或邮箱')
          return
        }
        if (!this.password) {
          this.$message.error('请输入密码')
          return
        }
        const user = await this.$store.dispatch('user/signin', {
          username: this.username,
          password: this.password,
          ref: this.ref,
        })
        if (this.ref) {
          // 跳到登录前
          this.$linkTo(this.ref)
        } else {
          // 跳到个人主页
          this.$linkTo('/')
        }
      } catch (e) {
        this.$message.error(e.message || e)
      }
    },
    /**
     * 如果已经登录了，那么直接跳转
     * @returns {boolean}
     */
    redirectIfLogined() {
      if (this.isLogin) {
        const me = this
        this.$msg({
          message: '登录成功',
          onClose() {
            if (me.ref && !me.$isSigninUrl(me.ref)) {
              me.$linkTo(me.ref)
            } else {
              me.$linkTo('/')
            }
          },
        })
        return true
      }
      return false
    },
  },
  head() {
    return {
      title: this.$siteTitle('登录'),
    }
  },
}
</script>
<style scoped lang="scss">
.signin {
  max-width: 480px;
  margin: auto;
  padding: 0 20px;

  .login-captcha-input {
    width: 100%;
    margin-right: 20px;
    .input {
      width: 100% !important;
    }
  }

  .login-captcha-img {
    img {
      height: 40px;
    }
  }

  .login-button {
    .button {
      width: 100%;
    }
  }

  .third-party-line {
    border-bottom: 1px solid #dedede;
    margin-bottom: 24px;
    .third-party-title {
      margin-bottom: -12px;
      text-align: center;

      span {
        background-color: #fff;
        padding: 0 10px;
        font-size: 13px;
      }
    }
  }

  .third-parties {
    text-align: center;
    margin-bottom: 10px;
  }
}
</style>
