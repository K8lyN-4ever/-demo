<template>
  <tr>
    <td class="color">
      <va-color-indicator :color="theme[colorName]" />
    </td>
    <td>
      {{ $t(`colorDropdown.${colorName}`) }}
    </td>
  </tr>
</template>

<script>
import { getColors } from 'vuestic-ui'

export default {
  props: {
    colorName: { type: String, required: true }
  },
  computed: {
    theme() {
      return getColors()
    }
  }
}
</script>

<style lang="scss" scoped>
  .color {
    width: 16px;
    height: 16px;
    &-bordered {
      border: 1px solid var(--va-gray);
      border-radius: 50%;
    }
  }
  td {
    color: var(--va-gray);
    padding: 0.25rem 0.5rem;
    box-sizing: border-box;
    vertical-align: middle;
    &:first-child {
      padding-right: 0;
    }
  }
</style>