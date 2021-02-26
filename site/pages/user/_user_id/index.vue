<template>
  <section class="main">
    <div class="container main-container left-main size-320">
      <div class="left-container">
        <user-profile :user="user" />

        <div class="tabs-warp">
          <div class="tabs">
            <ul>
              <li :class="{ 'is-active': activeTab === 'topics' }">
                <a :href="'/user/' + user.id + '?tab=topics'">
                  <span class="icon is-small">
                    <i class="iconfont icon-topic" aria-hidden="true" />
                  </span>
                  <span>动态</span>
                </a>
              </li>
            </ul>
          </div>
          <div v-if="
                topicsPage && topicsPage.total_num
              ">
            <load-more-articles
              v-if="topicsPage"
              v-slot="{ results }"
              :init-data="topicsPage"
              :url="'/api/topics?user_id=' + user.id"
            >
              <topic-list :articles="results" :show-avatar="false" />
            </load-more-articles>
          </div>
          <div v-else class="notification is-primary">暂无动态</div>
        </div>
      </div>
      <user-center-sidebar :user="user" />
    </div>
  </section>
</template>

<script>
import TopicList from '~/components/TopicList'
import UserProfile from '~/components/UserProfile'
import UserCenterSidebar from '~/components/UserCenterSidebar'
import LoadMoreArticles from '~/components/LoadMoreArticles'

const defaultTab = 'topics'

export default {
  components: {
    TopicList,
    UserProfile,
    UserCenterSidebar,
    LoadMoreArticles,
  },
  async asyncData({ $axios, params, query, error }) {
    let user
    try {
      user = await $axios.get('/api/user/info/' + params.user_id)
    } catch (err) {
      error({
        statusCode: 404,
        message: err.message || '系统错误',
      })
      return
    }

    const activeTab = query.tab || defaultTab
    let topicsPage = null
    topicsPage = await $axios.get('/api/topics', {
      params: { user_id: params.user_id },
    })

    return {
      activeTab,
      user,
      topicsPage,
    }
  },
  data() {
    return {}
  },
  computed: {
    currentUser() {
      return this.$store.state.user.current
    },
    isOwner() {
      const current = this.$store.state.user.current
      return this.user && current && this.user.id === current.id
    },
  },
  head() {
    return {
      title: this.$siteTitle(this.user.nickname),
    }
  },
}
</script>

<style lang="scss" scoped>
.tabs-warp {
  background: #fff;
  padding: 0 10px 10px;

  .tabs {
    margin-bottom: 5px;
  }

  .more {
    text-align: right;
  }
}
</style>
