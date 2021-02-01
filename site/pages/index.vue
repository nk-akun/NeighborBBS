<template>
  <section class="main">
    <div class="container main-container left-main size-320">
      <div class="left-container">
        <div class="main-content no-padding no-bg topics-wrapper">
          <div class="topics-main">
            <load-more
              v-if="topicsPage"
              v-slot="{ results }"
              :init-data="topicsPage"
              url="/api/topics"
            >
              <topic-list :topics="results" :show-ad="true" />
            </load-more>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import TopicList from '~/components/TopicList'
import LoadMore from '~/components/LoadMore'
export default {
  components: {
    TopicList,
    LoadMore,
  },
  async asyncData({ $axios, params }) {
    try {
      const [topicsPage] = await Promise.all([$axios.get('/api/topics')])
      console.log(topicsPage)
      return { topicsPage }
    } catch (e) {
      console.error(e)
    }
  },
  data() {},
  methods: {},
}
</script>

<style lang="scss" scoped></style>
