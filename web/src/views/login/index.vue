<template lang="pug">
  .login-container
    el-form.login-form(ref="loginForm" :model="loginForm" :rules="loginRules" auto-complete="on"
      label-position="left"
    )
      h3.title OFriends Login
      el-form-item(prop="username")
        span.svg-container
          svg-icon(name="user")
        el-input(v-model="loginForm.username" name="username" type="text" auto-complete="on" placeholder="username")
      el-form-item(prop="password")
        span.svg-container
          svg-icon(name="password")
        el-input(
          :type="pwdType"
          v-model="loginForm.password"
          name="password"
          auto-complete="on"
          placeholder="password"
          @keyup.enter.native="handleLogin"
        )
        span.show-pwd(@click="showPwd")
          svg-icon(:name="pwdType === 'password' ? 'eye' : 'eye-open'")
      el-form-item
        el-button(:loading="loading" type="primary" style="width:100%" @click.native.prevent="handleLogin") Sign in
      //- .tips
      //-   span(style="margin-right: 20px") username: admin
      //-   span password: admin
</template>

<script>
import { isvalidUsername } from '@/utils/validate'

export default {
  name: 'Login',
  data() {
    const validateUsername = (rule, value, callback) => {
      if (!isvalidUsername(value)) {
        callback(new Error('Invalid Username'))
      } else {
        callback()
      }
    }
    const validatePass = (rule, value, callback) => {
      if (value.length < 3) {
        callback(new Error('Password length should greater than 3'))
      } else {
        callback()
      }
    }

    return {
      loginForm: {
        username: 'admin',
        password: 'admin'
      },
      loginRules: {
        username: [{ required: true, trigger: 'blur', validator: validateUsername }],
        password: [{ required: true, trigger: 'blur', validator: validatePass }]
      },
      loading: false,
      pwdType: 'password',
      redirect: undefined
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
    showPwd() {
      if (this.pwdType === 'password') {
        this.pwdType = ''
      } else {
        this.pwdType = 'password'
      }
    },
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('Login', this.loginForm).then(() => {
            this.loading = false
            this.$router.replace({ path: this.redirect || '/' })
          }).catch((err) => {
            this.$notify({
              title: 'Error',
              message: err.message,
              type: 'error'
            })
          }).finally(() => {
            this.loading = false
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style lang="stylus">
$bg #2d3a4b
$light_gray #eee

.login-container
  .el-input
    display inline-block
    height 47px
    width 85%
    input
      background transparent
      border 0px
      -webkit-appearance none
      border-radius 0px
      padding 12px 15px 12px 15px
      color $light_gray
      height 47px
      &:-webkit-autofill
        -webkit-box-shadow 0 0 0px 1000px $bg inset !important
        -webkit-text-fill-color #fff !important
  .el-form-item
      border 1px solid rgba(255,255,255, 0.1)
      background rgba(0, 0, 0, 0.1)
      border-radius 5px
      color #454545
</style>

<style lang="stylus" scoped>
$bg = #2d3a4b
$dark_gray = #889aa4
$light_gray = #eeeeee

.login-container
  position: fixed
  height: 100%
  width: 100%
  background-color: $bg
  .login-form
    position: absolute
    left: 0
    right: 0
    width: 520px
    max-width: 100%
    padding: 35px 35px 15px 35px
    margin: 120px auto
  .tips
    font-size: 14px
    color: #fff
    margin-bottom: 10px
    span
      &:first-of-type
        margin-right: 16px
  .svg-container
    padding: 6px 5px 6px 15px
    color: $dark_gray
    vertical-align: middle
    width: 30px
    display: inline-block
  .title
    font-size: 26px
    font-weight: 400
    color: $light_gray
    margin: 0px auto 40px auto
    text-align: center
    font-weight: bold
  .show-pwd
    position: absolute
    right: 10px
    top: 7px
    font-size: 16px
    color: $dark_gray
    cursor: pointer
    user-select: none
</style>
