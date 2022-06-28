<template>
  <div class="dashboard">
    <dashboard-map />
  </div>
</template>

<script>
import DashboardMap from './DashboardMap'


export default {
  name: 'dashboard',
  components: {
    DashboardMap,

  },
  methods: {
    addAddressToMap ({ city, country }) {
      this.$refs.dashboardMap.addAddress({ city: city.text, country })
    },
  },
}
</script>

<style lang="scss">
  .row-equal .flex {
    .va-card {
      height: 100%;
    }
  }

  .dashboard {
    .va-card {
      margin-bottom: 0 !important;
      &__title {
        display: flex;
        justify-content: space-between;
      }
    }
  }
</style>
