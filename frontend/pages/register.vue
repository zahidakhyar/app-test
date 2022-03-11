<template>
  <div class="container">
    <div class="row align-items-center g-0" style="height: 80vh">
      <div class="col-12">
        <div class="card m-auto" style="max-width: 300px">
          <div class="card-header">
            <h4 class="card-title text-center">Register</h4>
          </div>
          <div class="card-body">
            <b-alert v-if="error_message" show variant="danger">{{ error_message }}</b-alert>
            <form ref="form" @submit.prevent="submit">
              <b-form-group label="Name" label-for="name">
                <b-form-input
                  id="name"
                  v-model="params.name"
                  type="text"
                  trim
                  required
                  :class="{
                    'is-invalid': errors.name,
                  }"
                  @keyup.enter="submit"
                />
                <b-form-invalid-feedback
                  v-if="errors.name"
                  id="input-live-feedback"
                >
                  {{ errors.name[0] }}
                </b-form-invalid-feedback>
              </b-form-group>
              <b-form-group label="Email" label-for="email">
                <b-form-input
                  id="email"
                  v-model="params.email"
                  type="email"
                  trim
                  required
                  :class="{
                    'is-invalid': errors.email,
                  }"
                  @keyup.enter="submit"
                />
                <b-form-invalid-feedback
                  v-if="errors.email"
                  id="input-live-feedback"
                >
                  {{ errors.email[0] }}
                </b-form-invalid-feedback>
              </b-form-group>
              <b-form-group label="Password" label-for="password">
                <b-form-input
                  id="password"
                  v-model="params.password"
                  type="password"
                  trim
                  required
                  :class="{
                    'is-invalid': errors.password,
                  }"
                  @keyup.enter="submit"
                />
                <b-form-invalid-feedback
                  v-if="errors.password"
                  id="input-live-feedback"
                >
                  {{ errors.password[0] }}
                </b-form-invalid-feedback>
              </b-form-group>
            </form>
          </div>
          <div class="card-footer">
            <b-button
              class="w-100"
              variant="primary"
              :disabled="loading"
              @click="submit()"
            >
              <span class="indicator-label" :hidden="loading"> Register </span>

              <span class="indicator-progress" :hidden="!loading">
                Please wait...
                <span
                  class="spinner-border spinner-border-sm align-middle ml-2"
                ></span>
              </span>
            </b-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  middleware: 'auth',
  auth: 'guest',
  
  data() {
    return {
      params: {
        name: '',
        email: '',
        password: '',
      },
      loading: false,
      errors: [],
      error_message: '',
    }
  },

  methods: {
    submit() {
      this.errors = []
      this.loading = true

      this.$axios
        .post('auth/register', this.params)
        .then(() => {
          this.$auth
            .loginWith('local', {
              data: this.params
            })
            .then(() => {
              this.$router.push('/')
            })
            .catch(err => {
              this.errors = err.response.data.errors
              this.loading = false
            })
        })
        .catch((error) => {
          this.loading = false
          this.errors = error.response.data.errors
          this.error_message = error.response.data.message
        })
    },
  },
}
</script>
