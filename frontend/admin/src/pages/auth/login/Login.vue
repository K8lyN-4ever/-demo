<template>
  <form @submit.prevent="onsubmit">
    <va-input
      class="mb-3"
      v-model="account"
      :placeholder="$t('auth.account')"
      :error="!!accountErrors.length"
      :error-messages="accountErrors"
    />

    <va-input
      class="mb-3"
      v-model="password"
      type="password"
      :placeholder="$t('auth.password')"
      :error="!!passwordErrors.length"
      :error-messages="passwordErrors"
    />

    <div class="d-flex justify--center mt-3">
      <va-button @click="onsubmit" class="my-0">{{ $t('auth.login') }}</va-button>
    </div>
  </form>
</template>

<script>
import { get } from '@/axios/html.js'
import { WS } from '@/util/WebSocket.js'

export default {
  name: 'login',
  data () {
    return {
      account: '',
      password: '',
      accountErrors: [],
      passwordErrors: [],
      toastColor: '#87CEEB',
      toastText: 'nullText',
      toastPosition: 'top-right',
      toastDuration: 2500,
    }
  },
  computed: {
    formReady () {
      return !this.accountErrors.length && !this.passwordErrors.length
    },
  },
  methods: {
    launchToast() {
      this.$vaToast.init({
        message: this.toastText,
        position: this.toastPosition,
        color: this.toastColor,
        duration: Number(this.toastDuration)
      })
    },
    onsubmit () {
      this.accountErrors = this.account ? [] : ['account is required']
      this.passwordErrors = this.password ? [] : ['Password is required']
      if (!this.formReady) {
        return
      }
      get("/api/user/logout")
      get('/api/user/login', {
        account: this.account,
        password: this.password,
        longitude: 0,
        latitude: 0
      }).then(res => {
        if(res.data.code === '0') {
          this.$store.commit('changeUserName', this.account)
          WS.connect()
          this.$router.push({ name: 'dashboard' })
        }else {
          this.toastText = res.data.msg
          this.launchToast()
        }
      }).catch(err => {
        console.log(err)
      })
        
    },
  },
}
</script>
