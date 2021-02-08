<template>
  <div class="comments">
    <load-more-comments
      v-if="commentsPage"
      ref="commentsLoadMore"
      v-slot="{ results }"
      :init-data="commentsPage"
      :params="{ entityType: entityType, article_id: entityId }"
      url="/api/comments"
    >
      <ul>
        <li
          v-for="(comment, index) in results"
          :key="comment.comment_id"
          class="comment"
          itemprop="comment"
          itemscope
          itemtype="http://schema.org/Comment"
        >
          <adsbygoogle
            v-if="showAd && (index + 1) % 3 === 0 && index !== 0"
            ad-slot="4980294904"
            ad-format="fluid"
            ad-layout-key="-ht-19-1m-3j+mu"
          />
          <div class="comment-avatar">
            <avatar
              :user="{nickname:comment.user_nickname,username:comment.user_username,id:comment.user_id,avatar_url:comment.avatar_url}"
              size="35"
            />
          </div>
          <div class="comment-meta">
            <span
              class="comment-nickname"
              itemprop="creator"
              itemscope
              itemtype="http://schema.org/Person"
            >
              <a :href="'/user/' + comment.user_id" itemprop="name">{{ comment.user_nickname }}</a>
            </span>
            <span class="comment-time">
              <time
                :datetime="
                  comment.create_time | formatDate('yyyy-MM-ddTHH:mm:ss')
                "
                itemprop="datePublished"
              >{{ comment.create_time | prettyDate }}</time>
            </span>
            <span class="comment-reply">
              <a @click="reply(comment)">回复</a>
            </span>
          </div>
          <div class="comment-content content">
            <blockquote v-if="comment.parent_comment" class="comment-quote">
              <div class="comment-quote-user">
                <avatar
                  :user="{nickname:comment.parent_comment.user_nickname,username:comment.parent_comment.user_username,id:comment.parent_comment.user_id,avatar_url:comment.parent_comment.avatar_url}"
                  size="20"
                />
                <a class="quote-nickname">{{ comment.parent_comment.user_nickname }}</a>
                <span class="quote-time">{{ comment.parent_comment.create_time | prettyDate }}</span>
              </div>
              <div
                v-lazy-container="{ selector: 'img' }"
                itemprop="text"
                v-html="comment.parent_comment.content"
              />
            </blockquote>
            <p v-lazy-container="{ selector: 'img' }" v-html="comment.content" />
          </div>
        </li>
      </ul>
    </load-more-comments>
  </div>
</template>

<script>
import Avatar from '~/components/Avatar'
import LoadMoreComments from './LoadMoreComments.vue'

export default {
  components: {
    Avatar,
    LoadMoreComments,
  },
  props: {
    entityType: {
      type: String,
      default: '',
      required: true,
    },
    entityId: {
      type: Number,
      default: 0,
      required: true,
    },
    commentsPage: {
      type: Object,
      default() {
        return {}
      },
    },
    showAd: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    user() {
      return this.$store.state.user.current
    },
    isLogin() {
      return this.$store.state.user.current != null
    },
  },
  methods: {
    append(data) {
      if (!data) return

      this.$refs.commentsLoadMore.unshiftResults(data)
    },
    reply(quote) {
      if (!this.isLogin) {
        this.$toSignin()
      }
      this.$emit('reply', quote)
    },
    cancelReply() {
      this.quote = null
    },
  },
}
</script>

<style scoped lang="scss"></style>
