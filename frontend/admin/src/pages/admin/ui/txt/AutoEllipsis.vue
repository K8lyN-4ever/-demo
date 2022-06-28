<template>
  <div v-if="notification">
    <va-popover
      :message="text"
      placement="left"
      color="info"
      trigger="hover"
    >
      <p class="p">
        {{ text }}
      </p>
    </va-popover>
  </div>
  
  <div v-else>
    <p class="p">
      {{ text }}
    </p>
  </div>

</template>

<script>

export default {
    name: "AutoEllipsis",
    props: {
        text: {
            type: String,
            default: ""
        },
        notification: {
            type: Boolean,
            default: false,
        }
    },

}
</script>

<style lang="scss">
    .p {
        height: 19px;
        word-break: break-all;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 1; /* 这里是超出几行省略 */
        overflow: hidden;
    }
</style>