<template>
  <div class="container">
    <div class="row align-items-center g-0" style="height: 80vh">
      <div class="col-12">
        <div class="card m-auto" style="max-width: 300px">
          <div class="card-header">
            <h4 class="card-title text-center">
              Login
            </h4>
          </div>
          <div class="card-body">
            <template v-if="errors.length">
              <error-forms :errors="errors"></error-forms>
            </template>
            <form ref="form" @submit.prevent="submit">
              <b-form-group label="Email" label-for="name">
                <b-form-input
                  id="name"
                  v-model="params.email"
                  type="email"
                  trim
                  required
                />
              </b-form-group>
              <b-form-group label="Password" label-for="password">
                <b-form-input
                  id="password"
                  v-model="params.password"
                  type="password"
                  trim
                  required
                />
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
              <span class="indicator-label" :hidden="loading"> Login </span>

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
  data() {
    return {
      params: {
        email: '',
        password: '',
      },
      loading: false,
      errors: [],
    }
  },

  methods: {
    submit() {
      this.errors = []
      this.loading = true

      this.$auth.loginWith('local', {
        data: this.params,
      })
      .then(() => this.$router.push('/'))
      .catch(error => {
        this.loading = false
        this.errors = error.response.data.errors
      })
    },
  },
}
</script>
