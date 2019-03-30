<template lang="pug">
  .login-container
    transition(name="slide-fade" mode="out-in" appear)
      el-form.login-form(v-if="mode==='login'" ref="loginForm" :model="loginForm" :rules="loginRules" auto-complete="on"
        label-position="left"
        key="login"
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
        span
          | don't have account?
          a.alink(@click="mode='register'")
            svg-icon(name="register")
            | &nbsp; Register
      el-form.login-form(v-else ref="loginForm" :model="loginForm" :rules="registerRules"
        label-position="left"
        key="register"
      )
        h3.title Create Account
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
        el-form-item(prop="email")
          span.svg-container
            svg-icon(name="email")
          el-input(v-model="loginForm.email" name="email" type="text" placeholder="email")
        el-form-item(prop="fullname")
          span.svg-container
            svg-icon(name="name")
          el-input(v-model="loginForm.fullname" name="fullname" type="text" placeholder="fullname")
        el-form-item(prop="delivery_center")
          span.svg-container
            svg-icon(name="delivery")
          el-select(v-model="loginForm.delivery_center" multiple filterable allow-create default-first-option
            placeholder="delivery center"
          )
            el-option(v-for="item in loginForm.delivery_center"
              :key="item"
              :label="item"
              :value="item"
            )
        el-form-item
          el-button(:loading="loading" type="primary" style="width:100%" @click.native.prevent="handleRegister") Create
        span
          | already have account?
          a.alink(@click="mode='login'")
            svg-icon(name="login")
            | &nbsp; Login
</template>

<script>
import { isvalidUsername, isvalidEmail } from '@/utils/validate'

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
      if (!value) {
        callback(new Error('Password must not be empty'))
      } else if (value.length < 3) {
        callback(new Error('Password length should greater than 3'))
      } else {
        callback()
      }
    }

    const validateEmail = (rule, value, callback) => {
      if (value) {
        if (!isvalidEmail(value)) {
          callback(new Error('Invalid Email'))
        } else {
          callback()
        }
      } else {
        // for now email is optional
        callback()
      }
    }

    return {
      loginForm: {
        // username: 'admin',
        // password: 'admin'
      },
      loginRules: {
        username: [{ required: true, trigger: 'blur', validator: validateUsername }],
        password: [{ required: true, trigger: 'blur', validator: validatePass }]
      },
      registerRules: {
        username: [{ required: true, trigger: 'blur', validator: validateUsername }],
        password: [{ required: true, trigger: 'blur', validator: validatePass }],
        email: [{ trigger: 'blur', validator: validateEmail }]
      },
      mode: 'login',
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
          this.$store.dispatch('Login', this.loginForm).then(() => {
            this.loading = false
            this.$router.replace({ path: this.redirect || '/' })
          }).catch((err) => {
            if (err.response && err.response.data) {
              const message = err.response.data.message
              this.$notify({
                title: 'Error',
                message,
                type: 'error'
              })
            }
          }).finally(() => {
            this.loading = false
          })
        } else {
          return false
        }
      })
    },
    handleRegister() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('Register', this.loginForm).then(() => {
            this.$notify({
              title: 'Account created',
              message: 'You have successfully created your account',
              type: 'success'
            })
          }).catch((err) => {
            if (err.response && err.response.data) {
              const message = err.response.data.message
              this.$notify({
                title: 'Error',
                message,
                type: 'error'
              })
            }
          }).finally(() => {
            this.loading = false
          })
        } else {
          return false
        }
      })
    }
  }
}
</script>

<style scoped>
/* Enter and leave animations can use different */
/* durations and timing functions.              */
.slide-fade-enter-active {
  transition: all .3s ease;
}
.slide-fade-leave-active {
  transition: all .3s ease;
}
.slide-fade-enter, .slide-fade-leave-to
/* .slide-fade-leave-active below version 2.1.8 */ {
  transform: translateY(-20px);
  opacity: 0;
}

.bounce-enter-active {
  animation: bounce-in .5s;
}
.bounce-leave-active {
  animation: bounce-in .5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
</style>

<style lang="stylus">
$bg=#2d3a4b
$light_gray=#eee

.login-container
  .el-select
    display inline-block
    height 47px
    width 85%
    .el-input.el-input--suffix
      position relative
    .el-tag
      background transparent
      font-size inherit
      color $light_gray
    .el-select__input
      background transparent
      border 0px
      -webkit-appearance none
      border-radius 0px
      padding 12px 15px 12px 15px
      margin-left 0
      color $light_gray
      height 47px
      &:-webkit-autofill
        -webkit-box-shadow 0 0 0px 1000px $bg inset !important
        -webkit-text-fill-color #fff !important
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
      margin-left 0
      color $light_gray
      height 47px
      &:-webkit-autofill
        -webkit-box-shadow 0 0 0px 1000px $bg inset !important
        -webkit-text-fill-color #fff !important
  .el-form-item
    border 1px solid rgba(255,255,255, 0.1)
    background rgba(0, 0, 0, 0.1)
    border-radius 5px
    color #eff
  span
    margin-left 5px
    color $light_gray
    .alink
      margin-left 5px
      color #cda
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
    margin: auto auto
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
