import qs from 'qs'
import JSON from 'JSON'

export default function ({ $axios, app }) {
  $axios.onRequest((config) => {
    config.headers.common['X-Client'] = 'neighbor-bbs-site'
    config.headers.post['Content-Type'] = 'application/json; charset=utf-8'
    const userToken = app.$cookies.get('userToken')
    if (userToken) {
      config.headers.common['X-User-Token'] = userToken
    }
    config.transformRequest = [
      function (data) {
        // if (process.client && data instanceof FormData) {
        //   // 如果是FormData就不转换
        //   return data
        // }
        data = JSON.stringify(data)
        return data
      },
    ]
  })

  $axios.onResponse((response) => {
    if (response.status !== 200) {
      return Promise.reject(response)
    }
    const jsonResult = response.data
    // console.log(response.data.value)
    if (jsonResult.success) {
      return Promise.resolve(jsonResult.value)
    } else {
      return Promise.reject(jsonResult)
    }
  })
}
