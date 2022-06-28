<template>
  <span :title="title">
    <svg t="1650536276583" class="va-icon-up" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1711" width="200" height="200"><path d="M520.3 258.2c-5.7 0-11.1 2.3-15 6.4l-431.9 454c-7.9 8.3-7.5 21.4 0.7 29.2 8.3 7.9 21.4 7.5 29.2-0.7l417-438.2L922 731.2c7.9 8.3 21 8.6 29.2 0.7s8.6-21 0.7-29.2l-416.7-438c-3.8-4.1-9.2-6.5-14.9-6.5z" p-id="1712"></path></svg> 
  </span>
</template>

<script>
export default {
  name: 'VaIconDelete',
  props: {
    title: {
      type: String,
      default: 'up',
    },
  },
}
</script>

<style lang="scss">
  .va-icon-up {
    display: inline-block;
    width: 16px;
    height: 16px;
    fill: black;
  }
  .va-icon-up:hover {
    fill: #0f0;
  }
</style>
