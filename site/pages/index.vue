<template>
  <section class="main">
    <div class="container main-container left-main size-320">
      <div class="left-container">
        <div class="main-content no-padding no-bg topics-wrapper">
          <div class="topics-main">
            <load-more-articles
              v-if="topicsPage"
              v-slot="{ results }"
              :init-data="topicsPage"
              url="/api/topics"
            >
              <topic-list :articles="results" :show-ad="true" />
            </load-more-articles>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import TopicList from '~/components/TopicList'
import LoadMoreArticles from '../components/LoadMoreArticles.vue'
export default {
  components: {
    TopicList,
    LoadMoreArticles,
  },
  async asyncData({ $axios, params }) {
    try {
      const [topicsPage] = await Promise.all([$axios.get('/api/topics')])
      return { topicsPage }
    } catch (e) {
      console.error(e)
    }
  },
  data() {},
  methods: {
    twitterCreated(data) {
      if (this.topicsPage) {
        if (this.topicsPage.article_list) {
          this.topicsPage.article_list.unshift(data)
        } else {
          this.topicsPage.article_list = [data]
        }
      }
    },
  },
}
</script>

<style lang="scss" scoped></style>
