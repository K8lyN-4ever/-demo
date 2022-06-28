<template>
  <va-dropdown
    class="profile-dropdown"
    v-model="isShown"
    boundary-body
    position="bottom"
    :offset="[0, 13]"
  >
    <template #anchor>
      <span class="profile-dropdown__anchor">
        <slot/>
        <va-icon
          class="px-2"
          :name="isShown ? 'angle_up' :'angle_down'"
          :color="theme.primary"
        />
      </span>
    </template>
    <va-dropdown-content class="profile-dropdown__content">
      <va-list-item
        v-for="option in options"
        :key="option.name"
      >
        <router-link
          :to="{name: option.redirectTo}"
          class="profile-dropdown__item"
          @click="logout()"
        >
          {{ $t(`user.${option.name}`) }}
        </router-link>          
      </va-list-item>
    </va-dropdown-content>
  </va-dropdown>
</template>

<script>
import { get } from '@/axios/html.js'
import { WS } from '@/util/WebSocket.js'
import { useGlobalConfig } from 'vuestic-ui'

export default {
  name: 'profile-section',
  data () {
    return {
      isShown: false,
    }
  },
  props: {
    options: {
      type: Array,
      default: () => [
        {
          name: 'profile',
          redirectTo: '',
        },
        {
          name: 'logout',
          redirectTo: 'login',
        },
      ],
    },
  },
  computed: {
    theme() { return useGlobalConfig().getGlobalConfig() },
  },
  methods: {
    logout() {
      get("/api/user/logout", '')
        .then(res => {
          if(res.data.code === "0") {
            this.$store.commit("changeUserName", '')
            WS.close()
          }

        }).catch(err => {
          console.log(err)
        })
    }
  }
}
</script>

<style lang="scss">

.profile-dropdown {
  cursor: pointer;

  .va-dropdown-popper__anchor {
    display: flex;
    justify-content: flex-end;
  }

  &__content {
    width: 8rem;
  }

  &__item {
    display: block;
    color: var(--va-gray);

    &:hover,
    &:active {
      color: var(--va-primary);
    }
  }

  .va-dropdown__anchor {
    display: inline-block;
  }
}
</style>
