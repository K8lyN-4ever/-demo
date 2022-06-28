<template>
  <form @submit.prevent="onsubmit()">

    <va-input
      class="mb-3"
      v-model="params.nickname"
      type="nickname"
      :placeholder="$t('auth.nickname')"
      :error="!!nicknameErrors.length"
      :error-messages="nicknameErrors"
    />

    <va-input
      class="mb-3"
      v-model="params.account"
      type="account"
      :placeholder="$t('auth.account')"
      :error="!!accountErrors.length"
      :error-messages="accountErrors"
    />

    <va-input
      class="mb-3"
      v-model="params.password"
      type="password"
      :placeholder="$t('auth.password')"
      :error="!!passwordErrors.length"
      :error-messages="passwordErrors"
    />

    <div class="d-flex justify--center mt-3">
      <va-button @click="signup" class="my-0">{{ $t('auth.sign_up') }}</va-button>
    </div>
  </form>
</template>

<script>
import { get } from '@/axios/html.js'

export default {
  name: 'signup',
  data () {
    return {
      url: {
        signup: '/api/admin/register'
      },
      params: {
        nickname: '',
        account: '',
        password: ''
      },
      emailErrors: [],
      nicknameErrors: [],
      accountErrors: [],
      passwordErrors: [],
      toastColor: '#87CEEB',
      toastText: 'nullText',
      toastPosition: 'top-right',
      toastDuration: 2500,
    }
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
    signup() {
      this.nicknameErrors = this.params.nickname ? [] : ['Nickname is required']
      this.accountErrors = this.params.account ? [] : ['Account is required']
      this.passwordErrors = this.params.password ? [] : ['Password is required']
      if (!this.formReady) {
        return
      }
      get(this.url.signup, this.params)
        .then(res => {
          this.toastText = res.data.msg
          this.launchToast()
        })
        .catch(err => {
          console.log(err)
        })
    },
  },
  computed: {
    formReady () {
      return !(this.nicknameErrors.length || this.passwordErrors.length || this.accountErrors.length)
    },
  },
}
</script>

<style lang="scss">
</style>
